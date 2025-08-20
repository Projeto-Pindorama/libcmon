/*
 * capio/parser.go - (Not so) Underground functions for parsing cpio files
 *
 * Copyright (C) 2025: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package capio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"pindorama.net.br/libcmon/bass"
	"time"
)

/* nula is a null (\0) character. */
var nula = byte(0)

// whatHeaderIsIt does what its name implies: verifies what
// the cpio header is. It correlates the magic number with
// a internal constant for it, making it easier to deal with
// each header type.
func whatHeaderIsIt(buffer []byte) Format {
	/* First two bytes for the binary magic. */
	switch MAGIC_BINARY {
	case (binary.LittleEndian.Uint16(buffer[:2])):
		return HEADER_BINLE
	case (binary.BigEndian.Uint16(buffer[:2])):
		return HEADER_BINBE
	}

	/* Modern ascii format and variations. */
	if bytes.Equal(buffer[:6], MAGIC_ODC) {
		/* TODO: Special treatment for HEADER_DEC.
		 * "The DEC format is rubbish." - Gunnar Ritter
		 *
		 * if (itsaDEC) {
		 *	return HEADER_DEC, nil
		 * }
		 */
		return HEADER_ODC
	} else if bytes.Equal(buffer[:6], MAGIC_CRC) {
		return HEADER_CRC
	} else if bytes.Equal(buffer[:6], MAGIC_ASCII) {
		return HEADER_ASC
	}

	return ZILTCH
}


// whatsTheHeaderType determines the specific type of the archive,
// besides whether it is binary (L.E. or B.E.) or ASCII/new cpio.
/*
 * func whatsTheHeaderType(buffer []byte) (error) {
 *
 *
 * }
 */

/*
 * func validateFile(file *os.File) {
 * }
 */

func doTheParse(file *os.File) (*Header, error) {
	header := []byte("")
	entry := &Header{}
	buffer, _, err := bass.Walk(file, 512)
	var teste time.Time
	if err != nil {
		return nil, err
	}
	header_format := whatHeaderIsIt(buffer)
	fmt.Println(header_format.String())

	for i := 0; i < len(buffer); i += 2 {
		if i == 0 {
			continue
		}
		if (buffer[(i-1)] != nula) &&
			bytes.Equal(buffer[i:(i+2)], []byte{nula, nula}) {
			header = buffer[:(i + 2)]
		}
	}

	switch (header_format & TYPE_BINARY) {
	case 0: /* ASCII, ODC, CRC, etc. */
		println("Isso é tudo, pe-pessoal!")
		println("... Por ora.")
	default: /* Binary, either little or big endian. */
		entry = &Header{
			Typeflag:   nula,
			Name:       "",
			Linkname:   "",
			Size:       0,
			Mode:       0,
			Uid:        0,
			Gid:        0,
			Uname:      "",
			Gname:      "",
			ModTime:    teste,
			AccessTime: teste,
			ChangeTime: teste,
			Devmajor:   0,
			Devminor:   0,
			Magic:      header[:2],
			Format:     header_format,
		}
	}

	fmt.Printf("%#v\n", header)
	return entry, nil
}

func CallFromTest(file *os.File) {
	entry, err := doTheParse(file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%#v\n", entry)
}
