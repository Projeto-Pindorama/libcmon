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
	"time"

	"golang.org/x/sys/unix"
	"pindorama.net.br/libcmon/bass"
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
	header_len := uint(0)
	header_end := uint(0)
	entry := &Header{}

	/* Sane way to store raw data. */
	ascii_header := rawASCIIHeader{}
	bin_header := rawBINHeader{}

	/*
	 * TODO: Determine the block size for the
	 * device in question for more efficience.
	 */
	buffer, _, err := bass.Walk(file, 512)
	if err != nil {
		return nil, err
	}
	header_format := whatHeaderIsIt(buffer)

	if (header_format & TYPE_BINARY) != 0 {
		header_len = 26
	} else if (header_format & TYPE_OCPIO) != 0 {
		header_len = 76
	} else if (header_format&TYPE_NCPIO) != 0 ||
		(header_format&TYPE_CRC) != 0 {
		header_len = 110
	}
	header_end = (header_len + uint(bytes.IndexByte(buffer[header_len:], nula)))

	/* These will populate the fields of the Header struct. */
	header := buffer[:header_len]
	typeflag := nula
	file_name := ""
	link_name := ""
	name_len := uint16(0)
	file_size := uint64(0)
	file_mode := os.FileMode(0)
	file_uid := 0
	file_gid := 0
	m_time := time.Unix(0, 0)
	devmajor := uint32(0)
	devminor := uint32(0)
	magic := []byte("")

	file_name = string(buffer[header_len:header_end])
	if file_name[(len(file_name)-1)] == '/' {
		typeflag = TypeDir
	}
	switch header_format & TYPE_BE {
	case 0: /* ASCII, ODC, CRC, etc and little endian binary. */
		switch header_format & TYPE_BINARY {
		case 0: /* ASCII, ODC, CRC, etc. */
			magic = header[:6]
			switch header_format & TYPE_OCPIO {
			case 0: /* New ASCII/CRC. */
				ascii_header = rawASCIIHeader{
					c_inode:     header[6:12],
					c_mode:      header[12:20],
					c_uid:       header[20:28],
					c_gid:       header[28:36],
					c_nlink:     header[36:44],
					c_mtime:     header[44:52],
					c_filesize:  header[52:60],
					c_devmajor:  header[60:68],
					c_devminor:  header[68:76],
					c_rdevmajor: header[76:84],
					c_rdevminor: header[84:92],
					c_namesize:  header[92:100],
					c_check:     header[100:108],
				}
				fmt.Println("New CPIO")
			default: /* ODC. */
				ascii_header = rawASCIIHeader{c_dev: header[6:12],
					c_inode:    header[12:18],
					c_mode:     header[18:24],
					c_uid:      header[24:30],
					c_gid:      header[30:36],
					c_nlink:    header[30:36],
					c_rdev:     header[36:42],
					c_mtime:    header[42:53],
					c_namesize: header[53:59],
					c_filesize: header[59:70],
				}
			}
			fmt.Printf("%#o\n", ascii_header)
			goto parsed
		default:
			break /* Common binary parsing code. */
		}
		fallthrough
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
		switch header_format & TYPE_BE {
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
			mtime := int64(prcl.MixedEndian.Uint32(bin_header.h_mtime))
			m_time = time.Unix(mtime, 0)

			/* Major and minor numbers. */
			majmin := uint64(binary.LittleEndian.Uint16(bin_header.h_majmin))
			devmajor = unix.Major(majmin)
			devminor = unix.Minor(majmin)
		default: /* Binary, big endian. */
			name_len = binary.BigEndian.Uint16(bin_header.h_namesize)
			file_mode = os.FileMode(binary.BigEndian.Uint16(bin_header.h_mode))
			file_uid = int(binary.BigEndian.Uint16(bin_header.h_uid))
			file_gid = int(binary.BigEndian.Uint16(bin_header.h_gid))
			file_size = uint64(binary.BigEndian.Uint32(bin_header.h_filesize))
			mtime := int64(binary.BigEndian.Uint32(bin_header.h_mtime))
			m_time = time.Unix(mtime, 0)
			majmin := uint64(binary.LittleEndian.Uint16(bin_header.h_majmin))
			devmajor = unix.Major(majmin)
			devminor = unix.Minor(majmin)
		}
	}
parsed: /* Jump falthrough. */

	entry = &Header{
		Typeflag: typeflag,
		Name:     file_name,
		Linkname: link_name,
		Namelen:  name_len,
		Size:     file_size,
		Mode:     file_mode,
		Uid:      file_uid,
		Gid:      file_gid,
		Uname:    "",
		Gname:    "",
		ModTime:  m_time,
		Devmajor: devmajor,
		Devminor: devminor,
		Magic:    magic,
		Format:   header_format,
	}
	return entry, nil
}

func CallFromTest(file *os.File) {
	entry, err := doTheParse(file)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", entry)
}
