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

import (
	"errors"
	"os"
	"pindorama.net.br/libcmon/bass"
)

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

func GetMBRPartType(blkpath string) (string, error) {
	devpath, err1 := GetBlockMainDisk(blkpath)
	fi, err2 := os.Open(devpath)
	defer fi.Close()

	entry, err3 := GetMBREntryForPart(blkpath)

	/* Also known as 01?0 + 0x2, or 0x1?2. */
	b, _, err4 := bass.Walk(fi, 1, (entry + 4))
	err := errors.Join(err1, err2, err3, err4)
	if err != nil {
		return "", err
	}
	partname := MBRPartNames[b[0]]
	if partname == "" {
		partname = "Unknown"
	}
	return partname, nil
}

func GetGPTPartType(blkname string) (string, error) {
	return "", errors.New("Not implemented (yet).")
}
