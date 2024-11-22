package dsks

import (
	"fmt"
	"os"
	"path/filepath"
	)

type DiskInfo struct {
	devpath string
	nsectors uint
	nbytes uint
	modelname string
	labeltype string
	identifier string
	blocks []BlockInfo
}

type BlockInfo struct {
	device string
	boot bool
	length blocklen
	nsectors uint
	size uint
	id int
	uuid string
	fstype string
	Dev_T
}

type blocklen struct {
	start uint64
	end uint64
}

func GetAllDisksInfo() ([]DiskInfo, error) {
	var disks []DiskInfo
	partitions, err := GetPartitionList()
	if err != nil {
		return nil, err
	}

	diskpaths := GetSysDisksPath(partitions)

	for d := 0; d < len(diskpaths); d++ {
		diskinfo, err := GetDiskInfo(diskpaths[d])
		if err != nil {
			return []DiskInfo{}, err
		}
		disks = append(disks,
			 DiskInfo{
				diskinfo.devpath,
				diskinfo.nsectors,
				diskinfo.nbytes,
				diskinfo.modelname,
				diskinfo.labeltype,
				diskinfo.identifier,
				diskinfo.blocks})
	}
	return disks, nil
}

func GetSysDisksPath(partlist []PartitionInfo) ([]string){
	var disks []string

	for p := 0; p < len(partlist); p++ {
		devpath := ("/dev/" + partlist[p].name)
		if IsEntireDisk(partlist, devpath) {
			disks = append(disks, devpath)
		}
	}

	return disks
}


func GetDiskInfo(devpath string) (*DiskInfo, error) {
	stat, err := os.Stat(devpath)
	if err != nil {
		return nil, err
	}

	/* Debug. */
	fmt.Printf("%+v\n", stat)

	return &DiskInfo{
		devpath: devpath,
		nsectors: 0,
		nbytes: 0,
		modelname: "", /* /sys/block/<basename /dev/xxx>/device/model */
		labeltype: "",
		identifier: "",
		blocks: []BlockInfo{},
		}, nil
}

func GetBlockInfo(blkpath string) (*BlockInfo, error) {
	var err error
	err = nil
	return &BlockInfo{}, err 
}

func IsEntireDisk(partlist []PartitionInfo, devpath string) (bool) {
	devno := GetDev_TForBlock(partlist, devpath)
	/*
	 * That's sort of empirical, but a entire disk
	 * located at /sys/block won't have a
	 * "partition" file on it, ergo it could tell
	 * us if the block is an entire disk or not.
	 */
	 _, err := os.Stat(fmt.Sprintf(
		 "/sys/dev/block/%d:%d/partition",
 			devno.major, devno.minor))
	 if os.IsNotExist(err) {
		return true
	 }
	 return false
}

func GetDev_TForBlock(partlist []PartitionInfo, devpath string) (*Dev_T) {
	devblk := filepath.Base(devpath)
	
	for b := 0; b < len(partlist); b++ {
		if devblk == partlist[b].name	{
			return &partlist[b].Dev_T
		}
	}
	return &Dev_T{}
}

/* major:min (int:int): /sys/block/<basename /dev/xxx>/dev */
/* sysfs_blkdev: /sys/dev/block/major(devno):minor(devno) */
