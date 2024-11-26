/*
 * bass/walk.go - Some hip functions for walking through
 * file pointers and making 'em into []byte(s).
 * Ain't you ever seen a walking bass?
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package bass

import (
	"io"
	"os"
	"strings"
)

func Walk(f *os.File, to int) ([]byte, int, error) {
	var r int
	var b []byte
	buf := make([]byte, 1)

coda:
	for r = 0; to != 0; to, r = (to - 1), (r + 1) {
		_, err := f.Read(buf)
		switch err {
		case nil:
			b = append(b, buf[0])
		case io.EOF:
			break coda
		default:
			return nil, 0, err
		}
	}
	return b, r, nil
}

func WalkTil(here byte, f *os.File) ([]byte, int, error) {
	var b []byte
	var i int
	bb, _, err := Walk(f, -1)
	if err != nil {
		return nil, 0, err
	}

	for i = 0; i < len(bb); i++ {
		if bb[i] != here {
			b = append(b, bb[i])
		} else {
			break
		}
	}

	return b, i, nil
}

func Strncmp(s1, s2 string, upto uint) bool {
	var i uint

	if (uint(len(s1)) < upto) || (uint(len(s2)) < upto) {
		return false
	}

	for i = 0; i < upto; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func Strcmp(s1, s2 string) bool {
	/* 
	 * Using strings.Compare() instead of implementing
	 * manually because of the optmizations made per
	 * the Go crew.
	 */
	 r := strings.Compare(s1, s2)

	 if r != 0 {
		 return false
	 }
	 return true
}
