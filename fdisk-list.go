/*
 * fdisk-list.go - Test for disks/.
 * Mimics util-linux fdisk's '-l' option.
 */

package main

import (
	"fmt"
	"pindorama.net.br/libcmon/disks"
	"pindorama.net.br/libcmon/porcelana"
)

/* major:min (int:int): /sys/block/<basename /dev/xxx>/dev */
/* sysfs_blkdev: /sys/dev/block/major(devno):minor(devno) */
func main() {
	disks, _ := disks.GetAllDisksInfo()

	for d := 0; d < len(disks); d++ {
		disk := disks[d]
		disksiz, unit := prcl.DiskSectorsToHuman(disk.QueueLimits.Logical_Block_Size, disk.NSectors)
		fmt.Printf(("Disk %s: %.1f %s, %d bytes, %d sectors\nDisk model: %s\n" +
			"Units: sectors of 1 * %d = %d\nSector size (logical/physical): %d / %d\n" +
			"I/O size (minimum/optimal): %d bytes / %d bytes\n" +
			"Disklabel type: %s\nDisk identifier: %s\n"),
			disk.DevPath, disksiz, unit,
			disk.NBytes, disk.NSectors, disk.ModelName,
			disk.QueueLimits.Logical_Block_Size, disk.QueueLimits.Logical_Block_Size,
			disk.QueueLimits.Logical_Block_Size, disk.QueueLimits.Physical_Block_Size,
			0, 0,
			disk.LabelType, disk.Identifier)
		blknamsiz := len(disk.DevPath)
		fmt.Printf("%*s Boot Start End   Sectors%10s  %10s\n", (blknamsiz + 1), "Device", "Size", "Type")
		for b := 0; b < len(disk.Blocks); b++ {
			block := disk.Blocks[b]
			fmt.Printf("%*s %v %d%5d%15d%15d  %s\n",
				blknamsiz, block.Device, block.IsBootable,
				block.Range.Start, block.Range.End,
				block.NSectors, block.Size, block.FSType)
		}
	}
}
