/*
 * bass/strcmp.go - strncmp(3) replica.
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package bass

func Strncmp(s1, s2 string, upto uint) bool {
	minlen := int(upto)
	if len(s1) < minlen {
		minlen = len(s1)
	}
	if len(s2) < minlen {
		minlen = len(s2)
	}

	return s1[:minlen] == s2[:minlen]
}
