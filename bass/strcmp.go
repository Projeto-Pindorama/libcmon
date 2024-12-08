/*
 * bass/strcmp.go - str(n)cmp(3) replicas. 
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package bass

func Strncmp(s1, s2 string, upto uint) bool {
	if (uint(len(s1)) < upto) || (uint(len(s2)) < upto) {
		return false
	}

	return Strcmp(s1[:upto], s2[:upto])
}

func Strcmp(s1, s2 string) bool {
	/*
	 * As per 'go doc strings.Compare':
	 * "It is usually clearer and always faster to use the
	 * built-in string comparison operators ==, <, >, and so on".
	 */

	return (s1 == s2)
}
