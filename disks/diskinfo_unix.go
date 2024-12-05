/*
 * disks/diskinfo_unix.go - UNIX®-specific functions for
 * collecting disk information.
 * I mean, you wouldn't use major and minor numbers at Windows.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */
package disks

import "path/filepath"

func GetDev_TForBlock(devpath string) *Dev_T {
	devblk := filepath.Base(devpath)

	for b := 0; b < len(partitions); b++ {
		if devblk == partitions[b].Name {
			return &partitions[b].Dev_T
		}
	}
	return &Dev_T{}
}
