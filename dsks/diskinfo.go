/*
 * dsks/diskinfo.go - Collect disk information (Linux-specific for now).
 * This is (currently) a mess.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package dsks

import (
	"fmt"
	"os"
	"path/filepath"
	"pindorama.net.br/libcmon/bass"
	"regexp"
)

type DiskInfo struct {
	DevPath    string
	NSectors   uint
	NBytes     uint
	ModelName  string
	LabelType  string
	Identifier string
	Blocks     []BlockInfo
}

type BlockInfo struct {
	Device     string
	IsBootable bool
	Length     blocklen
	NSectors   uint
	Size       uint
	Id         int
	UUID       string
	FSType     string
	Dev_T
}

type blocklen struct {
	start uint64
	end   uint64
}

var partitions []PartitionInfo

func init() {
	var err error

	/*
	 * If we could not read /proc/partitions,
	 * I think we would be a little bit more
	 * in trouble to recover.
	 */
	partitions, err = GetPartitionList()
	if err != nil {
		panic(err)
	}
}

func GetAllDisksInfo() ([]DiskInfo, error) {
	var disks []DiskInfo
	diskpaths := GetSysDisksPath()

	for d := 0; d < len(diskpaths); d++ {
		diskinfo, err := GetDiskInfo(diskpaths[d])
		if err != nil {
			return []DiskInfo{}, err
		}
		disks = append(disks,
			DiskInfo{
				diskinfo.DevPath,
				diskinfo.NSectors,
				diskinfo.NBytes,
				diskinfo.ModelName,
				diskinfo.LabelType,
				diskinfo.Identifier,
				diskinfo.Blocks})
	}
	return disks, nil
}

func GetSysDisksPath() []string {
	var disks []string

	for p := 0; p < len(partitions); p++ {
		devpath := ("/dev/" + partitions[p].Name)
		if IsEntireDisk(devpath) {
			disks = append(disks, devpath)
		}
	}

	return disks
}

func GetDiskInfo(devpath string) (*DiskInfo, error) {
	//	stat, err := os.Stat(devpath)
	var err error
	if err != nil {
		return nil, err
	}
	vfspath := MakeVFSBlockPaths(devpath)
	sys_block := vfspath["sysblock"]

	/* Get disk's model name. */
	modelfi, _ := os.Open((sys_block + "/device/model"))
	_modelname, err := bass.WalkTil('\n', modelfi)
	modelname := string(_modelname)
	modelfi.Close()

	/* Get disk's BlockInfo{} slice. */
	blocks, _ := GetDiskSubBlocks(devpath)

	return &DiskInfo{
		devpath,
		0,
		0,
		modelname,
		"",
		"",
		blocks,
	}, nil
}

func GetDiskSubBlocks(devpath string) ([]BlockInfo, error) {
	var dskparts []BlockInfo

	/* Open /sys/block/<dev>. */
	vfspath := MakeVFSBlockPaths(devpath)
	syspath := vfspath["sysdevblock"]
	entries, err := os.ReadDir(syspath)
	if err != nil {
		return []BlockInfo{}, err
	}

	/*
	 * Search for blocks. Just as in IsEntireDisk(), this is sort
	 * of empirical, but there's no better way of doing it.
	 */
	for e := 0; e < len(entries); e++ {
		fname := filepath.Base(entries[e].Name())
		/*
		 * Check for a name that follows the block name convention of
		 * <name><partition number>, with the loop device edge-case
		 * of "loop<disk number>p<partition number>".
		 */
		itmatches, _ := regexp.MatchString(".*[0-9](|p[0-9])", fname)
		if itmatches {
			blkinfo, err := GetBlockInfo(("/dev/" + fname))
			if err != nil {
				return []BlockInfo{}, err
			}
			dskparts = append(dskparts, *blkinfo)
		}
	}

	return dskparts, nil
}

func GetBlockInfo(blkpath string) (*BlockInfo, error) {
	devno := GetDev_TForBlock(blkpath)

	/*
	 * Check if the block is actually listed at /proc/partitions.
	 */
	if *devno == (Dev_T{}) {
		err := fmt.Errorf("could not find dev_t numbers for %s at %s",
			blkpath, "/proc/partitions")
		return &BlockInfo{}, err
	}

	return &BlockInfo{
		blkpath,
		false,
		blocklen{},
		0,
		0,
		0,
		"",
		"",
		Dev_T{
			devno.Major,
			devno.Minor}}, nil
}

func IsEntireDisk(devpath string) bool {
	vfspath := MakeVFSBlockPaths(devpath)

	/*
	 * That's sort of empirical, but a entire disk
	 * located at /sys/block won't have a
	 * "partition" file on it, ergo it could tell
	 * us if the block is an entire disk or not.
	 */
	_, err := os.Stat((vfspath["sysdevblock"] + "/partition"))
	if os.IsNotExist(err) {
		return true
	}
	return false
}

func GetDev_TForBlock(devpath string) *Dev_T {
	devblk := filepath.Base(devpath)

	for b := 0; b < len(partitions); b++ {
		if devblk == partitions[b].Name {
			return &partitions[b].Dev_T
		}
	}
	return &Dev_T{}
}

func MakeVFSBlockPaths(devpath string) map[string]string {
	vfspath := make(map[string]string)
	devno := GetDev_TForBlock(devpath)
	devblk := filepath.Base(devpath)

	/*
	 * Make "/sys/block/<name>" and
	 * "/sys/dev/block/<m>:<n>" strings.
	 */
	vfspath["sysblock"] = ("/sys/block/" + devblk)
	vfspath["sysdevblock"] = fmt.Sprintf("/sys/dev/block/%d:%d",
		devno.Major, devno.Minor)

	return vfspath
}
