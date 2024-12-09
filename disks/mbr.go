/*
 * disks/mbr.go - MBR-specific functions.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 */

package disks

import (
	"errors"
	"os"
	"pindorama.net.br/libcmon/bass"
	"strconv"
)

func GetMBREntryForPart(blkpath string) (int64, error) {
	vfspath := MakeVFSBlockPaths(blkpath)
	fi, err1 := os.Open((vfspath["sysdevblock"] + "/partition"))
	defer fi.Close()

	/*
	 * Get only one character since we can't get any number
	 * larger than 4 for a legacy MBR disk in any case; ergo,
	 * there's no necessity for walking until the line break.
	 */
	bnpart, _, err2 := bass.Walk(fi, 1)
	npart, err3 := strconv.Atoi(string(bnpart))

	err := errors.Join(err1, err2, err3)
	if err != nil {
		/*
		 * Since we're expecting to one use bass.Walk(),
		 * we shall return an negative number, so
		 * os.Seek() will error out and the function
		 * won't be doing nothing; otherwise it would
		 * be triggering an unwanted payload if one
		 * decides to plainly ignore the error.
		 */
		return -1, err
	}

	/*
	 * npart × 16 (0x10) plus 430 (0x1ae), which
	 * is 16 places before the start of the MBR
	 * partition table.
	 * The first partition starts at 0x1be (446),
	 * and there's four until 0x1ee (494) divided
	 * in intervals of 16 octets for each one.
	 */
	a := int64((npart * 0x10) + 0x1ae)

	return a, nil
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

/* Meant for MBR block devices only. */
func CanItBoot(blkpath string) (bool, error) {
	boot_octet := []byte{0x80}

	devpath, err1 := GetBlockMainDisk(blkpath)
	fi, err2 := os.Open(devpath)
	defer fi.Close()

	entry, err3 := GetMBREntryForPart(blkpath)
	found, _, err4 := bass.WalkLookinFor(boot_octet, fi, 1, entry)

	err := errors.Join(err1, err2, err3, err4)
	if err != nil {
		return false, err
	}

	return found, err
}
