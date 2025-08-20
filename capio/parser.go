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

	/*
	 * TODO: Determine the block size for the
	 * device in question for more efficience.
	 */
	buffer, _, err := bass.Walk(file, 512)
	var teste time.Time
	if err != nil {
		return nil, err
	}
	header_format := whatHeaderIsIt(buffer)
	header_len := uint(0)

	for i := 0; i < len(buffer); i += 2 {
		if i == 0 {
			continue
		}
		/* TODO: Fix for afio and odc/original ASCII format. */
		if (buffer[(i-1)] != nula) &&
			bytes.Equal(buffer[i:(i+2)], []byte{nula, nula}) {
			header = buffer[:(i + 2)]
		}
	}

	/* These will populate the fields of the Header struct. */
	typeflag := nula
	file_name := ""
	link_name := ""
	file_size := int64(0)
	file_mode := int64(0)
	file_uid := 0
	file_gid := 0
	devmajor := int64(0)
	devminor := int64(0)
	magic := []byte("")

	/*
	 * The header format goes around this:
	 * 	    M  D  I  Md U  G  nL Mn T  Ns Fsz
	 * Binary: [2][2][2][2][2][2][2][2][4][2][4]
	 * ASCII:  [6][6][6][6][6][6][6][6][11][6][11]
	 *
	 * The struct for the New ASCII format differs a little.
	 */
	switch header_format & TYPE_BINARY {
	case 0: /* ASCII, ODC, CRC, etc. */
		if (header_format & TYPE_OCPIO) != 0 {
			header_len = 76 /* Broken for now. */
		} else if (header_format&TYPE_NCPIO) != 0 ||
			(header_format&TYPE_CRC) != 0 {
			header_len = 110
		}
		magic = header[:6]
	default: /* Binary, either little or big endian. */
		magic = header[:2]
		switch header_format & TYPE_BE {
		case 0: /* Little endian. */
			println("LE")
		default: /* Big endian. */
			println("BE")
		}
	}

	file_name = string(header[header_len:(len(header) - 2)])
	if file_name[(len(file_name) - 1)] == '/' {
		typeflag = TypeDir
	}

	entry = &Header{
		Typeflag:   typeflag,
		Name:       file_name,
		Linkname:   link_name,
		Size:       file_size,
		Mode:       file_mode,
		Uid:        file_uid,
		Gid:        file_gid,
		Uname:      "",
		Gname:      "",
		ModTime:    teste,
		AccessTime: teste,
		ChangeTime: teste,
		Devmajor:   devmajor,
		Devminor:   devminor,
		Magic:      magic,
		Format:     header_format,
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
