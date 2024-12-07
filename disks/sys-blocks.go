/*
 * disks/sys-blocks.go - Make VFS paths for collecting disk information.
 * Linux-specific --- at least until some hip system decides to do it too.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package disks

import (
	"fmt"
	"errors"
	"os"
	"path/filepath"
	"pindorama.net.br/libcmon/bass"
	"pindorama.net.br/libcmon/porcelana"
	"reflect"
	"strconv"
	"strings"
)

/*
 * On other systems, this may still exists, but
 * it will be populated in another manner and some
 * information will possibly be zeroed out for being
 * Linux-specific.
 */
type QueueLimits struct {
	Physical_Block_Size uint16
	Logical_Block_Size  uint16
}

func MakeVFSBlockPaths(devpath string) map[string]string {
	vfspath := make(map[string]string)
	devno := GetDev_TForBlock(devpath)
	devblk := filepath.Base(devpath)

	/* 
	 * Make "/sys/block/<name>" string
	 * only if it is an entire block (or
	 * if /sys/block/</dev block> actually
	 * exists).
	 */
	if _, err := os.Stat(devpath); err == nil {
		vfspath["sysblock"] = ("/sys/block/" + devblk)
	}
	vfspath["sysdevblock"] = fmt.Sprintf("/sys/dev/block/%d:%d",
		devno.Major, devno.Minor)

	return vfspath
}

func GetDiskQueueLimits(devpath string) (*QueueLimits, error) {
	vfspath := MakeVFSBlockPaths(devpath)
	queuedir := (vfspath["sysblock"] + "/queue/")
	lim := &QueueLimits{}
	v := reflect.Indirect(reflect.ValueOf(lim))

	for q := 0; q < v.NumField(); q++ {
		e := v.Field(q)
		intw := prcl.IntWidth(e.Kind())
		switch intw {
		default:
			name := strings.ToLower(v.Type().Field(q).Name)
			qpath := (queuedir + name)
			fi, err1 := os.Open(qpath)
			s, _, err2 := bass.WalkTil('\n', fi)
			x, err3 := strconv.ParseUint(string(s), 0, intw)
			err := errors.Join(err1, err2, err3)
			if err != nil {
				return &QueueLimits{}, err
			}
			v.Field(q).SetUint(x)
			fi.Close()
		case -1:
			break
		}
	}

	return lim, nil
}

func GetBlockMainDisk(blkpath string) (string, error) {
	vfspath := MakeVFSBlockPaths(blkpath)
	real_sysdevblock, err1 := os.Readlink(vfspath["sysdevblock"])
	main_disk_sysdev, err2 := filepath.Abs((real_sysdevblock + "/.."))
	err := errors.Join(err1, err2)
	if err != nil {
		return "", err
	}
	main_disk := filepath.Base(main_disk_sysdev)

	return ("/dev/" + main_disk), nil
}
