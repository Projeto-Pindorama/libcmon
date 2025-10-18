/*
 * capio/parser.go - (Not so) Underground functions for parsing cpio files
 *
 * Copyright (C) 2025: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

// if on_disk {
// for ;; { 
// b, _, err := bass.Walk(f, hdrsize)
// parseHeader(b)
// start = <current location> - hdrsize
// end = <current location> + file size + jump "TRAILER!!!"
// entries[] = { Name: fileName, start: start, end: end }
// }
// bass.Walk(f, 0, 0) /* Rewind */
// }

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package capio 

import (
	"fmt"
	"os"
	"pindorama.net.br/libcmon/bass"
)

func readEntry(archive *os.File) (*Header, error) {
	magicbuf, _, err := bass.Walk(archive, 6)
	if err != nil {
		return nil, err
	}
	headerfmt := whatHeaderIsIt(magicbuf)


	headerbuf, _, err := bass.Walk(archive, (headerfmt.HeaderLen() - int64(6)))
	header := append(magicbuf, headerbuf...)
	fname, _, err := bass.WalkTil(nula, archive)
	if err != nil {
		return nil, err
	}

	/* Rebuild the header string. */
	entrydata := append(header, fname...)
	entrydata = append(entrydata, nula)
	fmt.Printf("%s\n", entrydata)
	fmt.Printf("%o\n", entrydata)
	return doTheParse(entrydata, headerfmt), nil
}

func CallFromTest(file *os.File) {
	entry, err := readEntry(file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%+v\n", entry)
}
