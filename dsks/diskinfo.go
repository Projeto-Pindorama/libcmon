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
	"pindorama.net.br/libcmon/bass"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

var partitions []PartitionInfo

func init() {
	var err error
	partitions, err = GetPartitionList()
	if err != nil {
		panic(err)		
	}
}

/* major:min (int:int): /sys/block/<basename /dev/xxx>/dev */
/* sysfs_blkdev: /sys/dev/block/major(devno):minor(devno) */
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

func GetSysDisksPath() ([]string){
	var disks []string

	for p := 0; p < len(partitions); p++ {
		devpath := ("/dev/" + partitions[p].name)
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

	/* 
	 * Get block device name and
	 * dev_t numbers for the disk. 
	 */
	devblk := filepath.Base(devpath)

	/* Make "/sys/block/<block name>" and "/sys/block/<m>:<n>" strings. */
	sys_block := ("/sys/block/" + devblk)
	
	/* Get disk's model name. */
	modelfi, _ := os.Open((sys_block + "/device/model"))
	_modelname, err := bass.WalkTil('\n', modelfi)
	modelname := string(_modelname)
	modelfi.Close()

	blocks, _ := GetDiskSubBlocks(devpath)
	
	return &DiskInfo{
		devpath: devpath,
		nsectors: 0,
		nbytes: 0,
		modelname: modelname,
		labeltype: "",
		identifier: "",
		blocks: blocks,
		}, nil
}

func GetDiskSubBlocks(devpath string) ([]BlockInfo, error) {
	var dskparts []BlockInfo

	devno := GetDev_TForBlock(devpath)
	syspath := fmt.Sprintf("/sys/dev/block/%d:%d",
			devno.major, devno.minor)
	entries, err := os.ReadDir(syspath)
	if err != nil {
		return nil, err
	}

	for e := 0; e < len(entries); e++ {
		fname := filepath.Base(entries[e].Name())
		itmatches, _ := regexp.MatchString(".*[0-9](|p[0-9])", fname)
		if (itmatches) {
			blkinfo, err := GetBlockInfo(("/dev/" + fname))
			if err != nil {
				return nil, err
			}
			dskparts = append(dskparts, *blkinfo)
		}
	}

	return dskparts, nil
}

func GetBlockInfo(blkpath string) (*BlockInfo, error) {
	devno := GetDev_TForBlock(blkpath)

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
		devno.major,
		devno.minor}}, nil 
}

func IsEntireDisk(devpath string) (bool) {
	devno := GetDev_TForBlock(devpath)
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

func GetDev_TForBlock(devpath string) (*Dev_T) {
	devblk := filepath.Base(devpath)
	
	for b := 0; b < len(partitions); b++ {
		if devblk == partitions[b].name	{
			return &partitions[b].Dev_T
		}
	}
	return &Dev_T{}
}

