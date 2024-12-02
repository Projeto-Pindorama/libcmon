/*
 * dsks/mbr_or_gpt.go - Check if disk is MBR or GPT.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package dsks

import (
	"errors"
	"os"
	"pindorama.net.br/libcmon/bass"
)

const (
	MBR     = 442520
	GPT     = 475054
	Unknown = 000000
)

func IsMBRorGPT(devpath string) (int, error) {
	mbr := map[string][]byte{
		"general":    {85, 170},
		"protective": {238},
		"efi_part":   {69, 70, 73, 32, 80, 65, 82, 84},
	}
	var err1, err2, err3 error
	fi, err := os.Open(devpath)
	defer fi.Close()
	if err != nil {
		return Unknown, err
	}

	/*
	 * The protective MBR octet --- EE in hex and 238
	 * as a Go 'byte' --- is found two places after the
	 * offset 01c0, which is 448 in decimal. That's the
	 * reason why we're going at 450 and reading 1 octet
	 * from there.
	 */
	found, _, err1 :=
		bass.WalkLookinFor(mbr["protective"], fi, 450, 1)
	switch found {
	case false:
		/*
		 * Do not have a protective MBR header,
		 * search for "EFI PART".
		 * Read from the LBA 0 (also know as
		 * protective MBR) (512 bytes) plus
		 * LBA 1 (8 bytes).
		 */
		found, _, err2 =
			bass.WalkLookinFor(mbr["efi_part"], fi, 0, (512 + 8))
		switch found {
		case false:
			found, _, err3 =
				bass.WalkLookinFor(mbr["general"], fi, 510, 2)
			switch found {
			case true:
				return MBR, nil
			case false:
				return Unknown, nil
			}
		case true:
			break
		}
		fallthrough
	case true:
		return GPT, nil
	}

	err = errors.Join(err1,
		err2, err3)
	if err != nil {
		return Unknown, err
	}

	return Unknown, nil
}

/* Bogus; to be implemented. */
func CanItBoot(devpath string) (bool, error) {
	/*
		 * fi, err := os.Open(devpath)
		 * if err != nil {
	 	 *	return false, err
	 	 * }
	 	 * defer fi.Close()
		 *
		 * bytes, _, err := bass.Walk(fi, 512)
	*/
	return false, nil
}
