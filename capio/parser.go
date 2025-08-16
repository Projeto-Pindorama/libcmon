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
)

// whatHeaderIsIt does what its name implies: verifies what
// the cpio header is. It correlates the magic number with
// a internal constant for it, making it easier to deal with
// each header type.
func whatHeaderIsIt(file *os.File) (uint, error) {
	buffer, _, err := bass.Walk(file, 512)
	if err != nil {
		return ZILTCH, err
	}

	/* First two bytes for the binary magic. */
	switch MAGIC_BINARY {
	case (binary.LittleEndian.Uint16(buffer[:2])):
		return HEADER_BINLE, nil
	case (binary.BigEndian.Uint16(buffer[:2])):
		return HEADER_BINBE, nil
	}

	/* Modern ascii format and variations. */
	if bytes.Equal(buffer[:6], MAGIC_ODC) {
		fmt.Println("ODC")
		return HEADER_ODC, nil
	} else if bytes.Equal(buffer[:6], MAGIC_CRC) {
		fmt.Println("CRC")
		return HEADER_CRC, nil
	} else if bytes.Equal(buffer[:6], MAGIC_ASCII) {
		/* TODO: Special treatment for HEADER_DEC.
		 *
			 * if (itsaDEC) {
		 *	return HEADER_DEC, nil
		 * }
		 */
		return HEADER_ASC, nil
	}

	return ZILTCH, err
}

func CallFromTest(file *os.File) {
	header, _ := whatHeaderIsIt(file)
	println((header & TYPE_BINARY))
	fmt.Println(formatNames[header])
}
