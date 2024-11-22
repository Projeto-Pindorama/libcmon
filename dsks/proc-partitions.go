package dsks

import (
	"bufio"
	"fmt"
	"os"
)

type Dev_T struct {
	major  uint
	minor  uint
}

type PartitionInfo struct {
	Dev_T
	blocks uint64
	name   string
}

func GetPartitionList() ([]PartitionInfo, error) {
	partitions := []PartitionInfo{}
	procfi, err := os.Open("/proc/partitions")
	if err != nil {
		return nil, err
	}

	rd := bufio.NewScanner(procfi)
	for rd.Scan() {
		line := rd.Text()
		switch rd.Err() {
		case nil:
			partinfo := doTheParse(line)
			/*
			 * Check if struct isn't empty before appending
			 * it into []partitions, of course.
			 */
			if *partinfo != (PartitionInfo{}) {
				partitions = append(partitions,
					PartitionInfo{
						Dev_T{
							partinfo.major,
							partinfo.minor},
						partinfo.blocks,
						partinfo.name})
			}
		default:
			return nil, err
		}
	}
	procfi.Close()

	return partitions, nil
}

func doTheParse(s string) *PartitionInfo {
	/* To receive from fmt.Fscanf(). */
	var (
		major  uint
		minor  uint
		blocks uint64
		dev    string
	)

	nparsed, _ := fmt.Sscanf(s, "%d %d %d %s",
		&major, &minor, &blocks, &dev)
	/*
	 * Since /proc/partitions' format is of
	 * "uint uint uint64 string", we can use
	 * the maxim of disposing any string if
	 * it doesn't parses right (which we can
	 * view per the number of elements parsed
	 * on 'nparsed').
	 */
	if nparsed < 4 {
		return &PartitionInfo{}
	}

	info := &PartitionInfo{
		Dev_T{
			major,
			minor},
		blocks,
		dev}

	return info
}
