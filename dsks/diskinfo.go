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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"pindorama.net.br/libcmon/bass"
	"regexp"
	"strconv"
)

type DiskInfo struct {
	DevPath    string
	NSectors   uint64
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
	NSectors   uint64
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

	nsectors, err := GetBlockNSectors(devpath)
	modelname, err := GetDiskModelName(devpath)

	/* Get disk's BlockInfo{} slice. */
	blocks, err := GetDiskSubBlocks(devpath)
	if err != nil {
		return &DiskInfo{}, err
	}

	return &DiskInfo{
		devpath,
		nsectors,
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
		fname := entries[e].Name()
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

	nsectors, err := GetBlockNSectors(blkpath)
	if err != nil {
		return &BlockInfo{}, err
	}

	uuid := GetUUIDForBlock(blkpath)

	return &BlockInfo{
		blkpath,
		false,
		blocklen{},
		nsectors,
		0,
		0,
		uuid,
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

func GetBlockNSectors(devpath string) (uint64, error) {
	vfspath := MakeVFSBlockPaths(devpath)

	/* Open /sys/dev/block/m:n/size. */
	fi, err := os.Open((vfspath["sysdevblock"] + "/size"))
	defer fi.Close()
	if err != nil {
		return 0, err
	}

	s, _, err := bass.WalkTil('\n', fi)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(string(s), 0, 64)
}

func GetUUIDForBlock(devpath string) string {
	if IsEntireDisk(devpath) {
		/*
		 * There is no UUID for
		 * an entire disk.
		 */
		return ""
	}

	disk_by_uuid_path := "/dev/disk/by-uuid/"
	devblk := filepath.Base(devpath)
	entries, _ := os.ReadDir(disk_by_uuid_path)

	for e := 0; e < len(entries); e++ {
		/*
		 * os.ReadDir() actually returns the base name
		 * as entry's .Name(), so we must amend the
		 * "/dev/disk/by-uuid" directory before it.
		 */
		devpath_per_uuid_path, _ :=
			os.Readlink((disk_by_uuid_path + entries[e].Name()))
		devblk_per_uuid_path :=
			filepath.Base(devpath_per_uuid_path)

		/* Then check if the block device name matches. */
		if devblk == devblk_per_uuid_path {
			uuid := entries[e].Name()
			return uuid
		}
	}

	/* Not found case. */
	return ""
}

func GetDiskModelName(devpath string) (string, error) {
	devblk := filepath.Base(devpath)

	if bass.Strncmp(devblk, "loop", 4) {
		/*
		 * loop devices have no model name.
		 * Despite this fact, shall not error.
		 */
		return "", nil
	} else if !IsEntireDisk(devpath) {
		return "", errors.New("single block devices do not have a model name.")
	}

	vfspath := MakeVFSBlockPaths(devpath)
	modelfi, err := os.Open((vfspath["sysblock"] + "/device/model"))
	if err != nil {
		return "", err
	}

	modelname, _, err := bass.WalkTil('\n', modelfi)
	modelfi.Close()

	return string(modelname), nil
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
