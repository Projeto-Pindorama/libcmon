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

package capio 

import (
	"fmt"
	//"io"
	"os"
	"pindorama.net.br/libcmon/bass"

)

//func readEntry(archive *io.Reader) (*Header, error) {
// TODO: Remove bass.Walk entirely.
func readEntry(archive *os.File) (*Header, error) {
	entry := &Header{}
	magicbuf, _, err := bass.Walk(archive, 6)
	if err != nil {
		return nil, err
	}
	headerfmt := whatHeaderIsIt(magicbuf)

	headerbuf, _, err := bass.Walk(archive, (headerfmt.HeaderLen() - int64(6)))
	entrydata := append(magicbuf, headerbuf...)
	fname, _, err := bass.WalkTil(nula, archive)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%#v\n", headerfmt.Limits())
	/* Rebuild the header string. */
	fmt.Printf("%s\n", entrydata)
	fmt.Printf("%o\n", entrydata)
	whatsTheHeaderType(headerfmt, entrydata)
	entry = doTheParse(entrydata, headerfmt)
	entry.Name = string(fname)
	return entry, nil
}

// whatsTheHeaderType determines the specific type of the archive,
// besides whether it is binary (L.E. or B.E.) or ASCII/new cpio.
func whatsTheHeaderType(base Format, buffer []byte) (error) {
	fmt.Println(base, len(buffer))
	return nil
}

/*
 * func validateFile(file *os.File) {
 * }
 */


func CallFromTest(file *os.File) {
	entry, err := readEntry(file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%+v\n", entry)
}
