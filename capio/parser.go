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
	"os"
	"time"

	"golang.org/x/sys/unix"
	"pindorama.net.br/libcmon/porcelana"
)

// nula is a null (\0) character.
var nula = byte(0)

// rawBINHeader and rawASCIIHeader are non-exposed structs
// for making it easier to parse binary/Programmer's Workbench
// and Old/New ASCII formats respectively.
type rawBINHeader struct {
	h_dev      []byte /* 2 bytes */
	h_inode    []byte /* 2 bytes */
	h_mode     []byte /* 2 bytes */
	h_uid      []byte /* 2 bytes */
	h_gid      []byte /* 2 bytes */
	h_nlink    []byte /* 2 bytes */
	h_majmin   []byte /* 2 bytes */
	h_mtime    []byte /* 4 bytes */
	h_namesize []byte /* 2 bytes */
	h_filesize []byte /* 4 bytes */
}

type rawASCIIHeader struct {
	c_dev       []byte /* 6 bytes, ODC specific. */
	c_inode     []byte /* 6 bytes for ODC, 8 for NEWC. */
	c_mode      []byte /* 6 bytes for ODC, 8 for NEWC. */
	c_uid       []byte /* 6 bytes for ODC, 8 for NEWC. */
	c_gid       []byte /* 6 bytes for ODC, 8 for NEWC. */
	c_nlink     []byte /* 6 bytes for ODC, 8 for NEWC.*/
	c_rdev      []byte /* 6 bytes, ODC specific. */
	c_mtime     []byte /* 11 bytes for ODC, 8 for NEWC. */
	c_namesize  []byte /* 6 bytes for ODC, 8 for NEWC. */
	c_filesize  []byte /* 11 bytes for ODC, 8 for NEWC. */
	c_devmajor  []byte /* 8 bytes, NEWC specific. */
	c_devminor  []byte /* 8 bytes, NEWC specific. */
	c_rdevmajor []byte /* 8 bytes, NEWC specific. */
	c_rdevminor []byte /* 8 bytes, NEWC specific. */
	c_check     []byte /* 8 bytes, NEWC specific. */
}

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

func doTheParse(header []byte, headerfmt Format) *Header {
	entry := &Header{}

	/* Sane way to store raw data. */
	ascii_header := rawASCIIHeader{}
	bin_header := rawBINHeader{}

	/* These will populate the fields of the Header struct. */
	name_len := uint16(0)
	file_size := uint64(0)
	file_mode := os.FileMode(0)
	file_uid := 0
	file_gid := 0
	m_time := time.Unix(0, 0)
	devmajor := uint32(0)
	devminor := uint32(0)
	rdevmajor := uint32(0)
	rdevminor := uint32(0)
	magic := []byte("")

	switch headerfmt & TYPE_BINARY {
	case 0: /* ASCII, ODC, CRC, etc. */
		magic = header[:6]

		/* TODO: Special treatment for HEADER_DEC.
		 * "The DEC format is rubbish." - Gunnar Ritter
		 *
		 * if (itsaDEC) {
		 *	return HEADER_DEC, nil
		 * }
		 */

		/*
		 * Just used as medium for being passed
		 * into the Header struct.
		 */
		devmaj := uint64(0)
		devmin := uint64(0)
		rdevmaj := uint64(0)
		rdevmin := uint64(0)

		switch headerfmt & TYPE_OCPIO {
		case 0: /* New ASCII/CRC. */
			ascii_header = rawASCIIHeader{
				c_inode:     header[6:14],
				c_mode:      header[14:22],
				c_uid:       header[22:30],
				c_gid:       header[30:38],
				c_nlink:     header[38:46],
				c_mtime:     header[46:54],
				c_filesize:  header[54:62],
				c_devmajor:  header[62:70],
				c_devminor:  header[70:78],
				c_rdevmajor: header[78:86],
				c_rdevminor: header[86:94],
				c_namesize:  header[94:102],
				c_check:     header[102:110],
			}

			name_len = uint16(prcl.HexaToInt(ascii_header.c_namesize))
			file_mode = os.FileMode(prcl.HexaToInt(ascii_header.c_mode))
			file_uid = int(prcl.HexaToInt(ascii_header.c_uid))
			file_gid = int(prcl.HexaToInt(ascii_header.c_gid))
			file_size = uint64(prcl.HexaToInt(ascii_header.c_filesize))
			m_time = time.Unix(prcl.HexaToInt(ascii_header.c_mtime), 0)
			devmaj = uint64(prcl.HexaToInt(ascii_header.c_devmajor))
			devmin = uint64(prcl.HexaToInt(ascii_header.c_devminor))
			rdevmaj = uint64(prcl.HexaToInt(ascii_header.c_rdevmajor))
			rdevmin = uint64(prcl.HexaToInt(ascii_header.c_rdevminor))
		default: /* ODC. */
			ascii_header = rawASCIIHeader{
				c_dev:      header[6:12],
				c_inode:    header[12:18],
				c_mode:     header[18:24],
				c_uid:      header[24:30],
				c_gid:      header[30:36],
				c_nlink:    header[36:42],
				c_rdev:     header[42:48],
				c_mtime:    header[48:59],
				c_namesize: header[59:65],
				c_filesize: header[65:76],
			}
			name_len = uint16(prcl.OctalToInt(ascii_header.c_namesize))
			file_mode = os.FileMode(prcl.OctalToInt(ascii_header.c_mode))
			file_uid = int(prcl.OctalToInt(ascii_header.c_uid))
			file_gid = int(prcl.OctalToInt(ascii_header.c_gid))
			file_size = uint64(prcl.OctalToInt(ascii_header.c_filesize))
			m_time = time.Unix(prcl.OctalToInt(ascii_header.c_mtime), 0)
			devmaj = uint64(prcl.OctalToInt(ascii_header.c_dev))
			devmin = devmaj
			rdevmaj = uint64(prcl.OctalToInt(ascii_header.c_rdev))
			rdevmin = rdevmaj
		}

		devmajor = unix.Major(devmaj)
		devminor = unix.Minor(devmin)
		rdevmajor = unix.Major(rdevmaj)
		rdevminor = unix.Minor(rdevmin)
	default:
		magic = header[:2]
		bin_header = rawBINHeader{
			h_dev:      header[2:4],
			h_inode:    header[4:6],
			h_mode:     header[6:8],
			h_uid:      header[8:10],
			h_gid:      header[10:12],
			h_nlink:    header[12:14],
			h_majmin:   header[14:16],
			h_mtime:    header[16:20],
			h_namesize: header[20:22],
			h_filesize: header[22:26],
		}

		/*
		 * These will be used as medium before being
		 * finally converted for the Header struct.
		 * Perhaps this could be simpler, but who knows?
		 */
		mtime := uint32(0)
		majmin := uint16(0)

		switch headerfmt & TYPE_BE {
		case 0: /* Binary, little endian. */
			name_len = binary.LittleEndian.Uint16(bin_header.h_namesize)
			file_mode = os.FileMode(binary.LittleEndian.Uint16(bin_header.h_mode))
			file_uid = int(binary.LittleEndian.Uint16(bin_header.h_uid))
			file_gid = int(binary.LittleEndian.Uint16(bin_header.h_gid))
			/*
			 * For both h_filesize and h_mtime, we will have to use
			 * a "middle-endian" format. For more detail, check code
			 * for the 'prcl' package.
			 */
			file_size = uint64(prcl.MixedEndian.Uint32(bin_header.h_filesize))
			mtime = prcl.MixedEndian.Uint32(bin_header.h_mtime)

			/* Major and minor numbers. */
			majmin = binary.LittleEndian.Uint16(bin_header.h_majmin)
		default: /* Binary, big endian. */
			name_len = binary.BigEndian.Uint16(bin_header.h_namesize)
			file_mode = os.FileMode(binary.BigEndian.Uint16(bin_header.h_mode))
			file_uid = int(binary.BigEndian.Uint16(bin_header.h_uid))
			file_gid = int(binary.BigEndian.Uint16(bin_header.h_gid))
			file_size = uint64(binary.BigEndian.Uint32(bin_header.h_filesize))
			mtime = binary.BigEndian.Uint32(bin_header.h_mtime)
			majmin = binary.BigEndian.Uint16(bin_header.h_majmin)
		}

		m_time = time.Unix(int64(mtime), 0)
		rdevmajor = unix.Major(uint64(majmin))
		rdevminor = unix.Minor(uint64(majmin))
	}

	entry = &Header{
		Typeflag:    nula,
		Name:        "",
		Linkname:    "",
		Namelen:     name_len,
		Size:        file_size,
		Mode:        file_mode,
		Uid:         file_uid,
		Gid:         file_gid,
		ModTime:     m_time,
		Devmajor:    devmajor,
		Devminor:    devminor,
		RawDevmajor: rdevmajor,
		RawDevminor: rdevminor,
		Magic:       magic,
		Format:      headerfmt,
	}
	return entry
}
