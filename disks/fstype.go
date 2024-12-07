/*
 * disks/fstype.go - Map for associating MBR partition types with
 * its human-readable names.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package disks

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
