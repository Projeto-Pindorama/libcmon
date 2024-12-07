/*
 * disks/gpt.go - EFI/GPT-specific functions.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 */

package disks

import (
	"errors"
)

func GetGPTPartType(blkname string) (string, error) {
	return "", errors.New("Not implemented (yet).")
}
