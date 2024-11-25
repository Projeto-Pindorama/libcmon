package main
import (
	"pindorama.net.br/libcmon/dsks"
	"fmt"	
	)

/* major:min (int:int): /sys/block/<basename /dev/xxx>/dev */
/* sysfs_blkdev: /sys/dev/block/major(devno):minor(devno) */
func main() {
	disks, err := dsks.GetAllDisksInfo()
	fmt.Printf("%+v\n%v\n", disks, err)
}
