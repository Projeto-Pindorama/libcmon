/* 
 * fdisk-list.go - Test for dsks/.
 * Mimics util-linux fdisk's '-l' option.
 */

 package main

 import (
	"fmt"
	"pindorama.net.br/libcmon/dsks"
)

/* major:min (int:int): /sys/block/<basename /dev/xxx>/dev */
/* sysfs_blkdev: /sys/dev/block/major(devno):minor(devno) */
func main() {
	fdiskfmt := "Disk %s: %f GiB, %d bytes, %d sectors\nDisk model: %s\nUnits: sectors of 1 * %d = %d\nSector size (logical/physical): %d bytes / %d bytes\nI/O size (minimum/optimal): %d bytes / %d bytes\nDisklabel type: %s\nDisk identifier: %s\n"
	disks, _ := dsks.GetAllDisksInfo()
	for d := 0; d < len(disks); d++ {
		disk := disks[d]
		fmt.Printf(fdiskfmt, disk.DevPath, (disk.NBytes / (1024*1024*1024)),
		disk.NBytes, disk.NSectors, disk.ModelName,
		disk.QueueLimits.Logical_Block_Size, disk.QueueLimits.Logical_Block_Size,
		disk.QueueLimits.Logical_Block_Size, disk.QueueLimits.Physical_Block_Size,
		0, 0,
		disk.LabelType, disk.Identifier)
		fmt.Printf("%s %10s %2s %5s %5s %5s\n", "Device", "Start", "End", "Sectors", "Size", "Type")
		for b := 0; b < len(disk.Blocks); b++ {
			block := disk.Blocks[b]
			fmt.Printf("%s %5d %4d %5d %5d %5s\n",
			block.Device, block.Range.Start, block.Range.End,
				block.NSectors, block.Size, block.FSType)
		}
	}
}
