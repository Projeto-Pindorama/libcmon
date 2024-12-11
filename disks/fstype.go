/*
 * disks/fstype.go - Functions for getting partition types
 * and disk/partition identifiers.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package disks

import "fmt"

func GetPartType(blkpath string, label int) (string, error) {
	var s string
	var err error

	switch label {
	case MBR:
		s, err = GetMBRPartType(blkpath)
	case GPT:
		s, err = GetGPTPartType(blkpath)
	}

	return s, err
}

func GetDiskIdentifier(devpath string, label int) (string, error) {
	var s string
	var err error

	switch label {
	case MBR:
		h, err1 := GetMBRDiskID(devpath)
		err = err1
		s = fmt.Sprintf("0x%08x", h)
	case GPT:
		s, err = GetGPTDiskID(devpath)
	}

	return s, err
}
