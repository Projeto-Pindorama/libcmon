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
)

func Walk(f *os.File, to int) ([]byte, error) {
	var b []byte
	buf := make([]byte, 1)

coda:
	for ; to != 0; to-- {
		_, err := f.Read(buf)
		switch err {
		case nil:
			b = append(b, buf[0])
		case io.EOF:
			break coda
		default:
			return nil, err
		}
	}
	return b, nil
}

func WalkTil(here byte, f *os.File) ([]byte, error) {
	var b []byte
	bb, err := Walk(f, -1)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(bb); i++ {
		if bb[i] != here {
			b = append(b, bb[i])
		} else {
			break
		}
	}

	return b, nil
}
