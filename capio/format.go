/*
 * capio/format.go - Programmatical description of the cpio format
 *
 * Copyright (C) 2025: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 * Constants borrowed from Heirloom Toolchest's cpio program code.
 * From cpio/cpio.h and cpio/cpio.c copyright header:
 *
 * Gunnar Ritter, Freiburg i. Br., Germany, April 2003.
 * Copyright (c) 2003 Gunnar Ritter
 *
 * SPDX-Licence-Identifier: Zlib
 *
 * The 'Header' struct and Type-* constants were partially
 * borrowed from Go's archive/tar.
 * From archive/tar/common.go copyright header:
 *
 * Copyright 2009 The Go Authors. All rights reserved.
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package capio

import (
	"os"
	"time"
)

type Format uint

/* Magic numbers and such. */
const (
	MAGIC_BINARY uint16 = 070707
	//mag_sco = 0x7ffffe00
)

var (
	MAGIC_ASCII = []byte("070701") // ... or 0x303730373031
	MAGIC_CRC   = []byte("070702") // ... or 0x303730373032
	MAGIC_ODC   = []byte("070707") // ... or 0x303730373037
)

const (
	_ Format = 0 /* Set all constants as 'Format'. */

	/* Type of header. */
	ZILTCH = 00000000 /* no format chosen yet */

	HEADER_ODC    = 00002001 /* POSIX ASCII cpio format */
	HEADER_DEC    = 00002002 /* DEC extended cpio format */
	HEADER_BINLE  = 00003001 /* binary (default) cpio format LE */
	HEADER_BINBE  = 00003101 /* binary (default) cpio format BE */
	HEADER_SGILE  = 00003201 /* IRIX-style -K binary format LE */
	HEADER_SGIBE  = 00003301 /* IRIX-style -K binary format BE */
	HEADER_ASC    = 00004001 /* SVR4 ASCII cpio format */
	HEADER_SCOASC = 00004201 /* UnixWare 7.1 ASCII cpio format */
	HEADER_CRC    = 00004401 /* SVR4 ASCII cpio format w/checksum */
	HEADER_SCOCRC = 00004601 /* UnixWare 7.1 ASCII cpio w/checksum */
	HEADER_CRAY   = 00010001 /* Cray cpio, UNICOS 6 and later */
	HEADER_CRAY5  = 00010002 /* Cray cpio, UNICOS 5 and earlier */
	HEADER_BAR    = 00400001 /* bar format type */

	/* Characteristics of the archive ('masked' into integer). */
	TYPE_BE     = 00000100 /* this binary archive is big-endian */
	TYPE_SGI    = 00000200 /* SGI cpio -K flag binary archive */
	TYPE_SCO    = 00000200 /* SCO UnixWare 7.1 extended archive */
	TYPE_CRC    = 00000400 /* this has a SVR4 'crc' checksum */
	TYPE_BINARY = 00001000 /* this is a binary cpio type */
	TYPE_OCPIO  = 00002000 /* this is an old cpio type */
	TYPE_NCPIO  = 00004000 /* this is a SVR4 cpio type */
	TYPE_CRAY   = 00010000 /* this is a Cray cpio archive */
	TYPE_CPIO   = 00077000 /* this is a cpio type */
	TYPE_BAR    = 00400000 /* this is a bar type */
)

var formatNames = map[Format]string{
	HEADER_BINLE:  "BINLE",
	HEADER_BINBE:  "BINBE",
	HEADER_ASC:    "NEWC",
	HEADER_SCOASC: "SCO",
	HEADER_CRC:    "CRC",
	HEADER_SCOCRC: "SCOCRC",
	HEADER_ODC:    "ODC",
	HEADER_DEC:    "DEC",
	HEADER_SGIBE:  "SGI",
	HEADER_CRAY:   "CRAY",
	HEADER_CRAY5:  "CRAY5",
	HEADER_BAR:    "BAR",
}

// Type flags for Header.Typeflag, indicating the
// type of header's content.
const (
	/* Type '0' indicates a regular file. */
	TypeReg byte = '0'

	/* Type '1' to '6' are header-only flags and may not have a data body. */
	TypeLink    byte = '1' /* Hard link */
	TypeSymlink byte = '2' /* Symbolic link */
	TypeChar    byte = '3' /* Character device node */
	TypeBlock   byte = '4' /* Block device node */
	TypeDir     byte = '5' /* Directory */
	TypeFifo    byte = '6' /* FIFO node */

	/* Type '7' is reserved. */
	TypeCont byte = '7'
)

// A Header represents a single header in a cpio archive.
// Some fields may not be populated.
type Header struct {
	/*
	 * Typeflag is the type of header entry.
	 * The zero value is automatically promoted to either TypeReg or TypeDir
	 * depending on the presence of a trailing slash in Name.
	 */
	Typeflag byte

	Name     string      /* Name of file entry */
	Linkname string      /* Target name of link (valid for TypeLink or TypeSymlink) */
	Namelen  uint16      /* Length of the file name */
	Size     uint64      /* Logical file size in bytes */
	Mode     os.FileMode /* Permission and mode bits */
	Uid      int         /* User ID of owner */
	Gid      int         /* Group ID of owner */

	/*
	 * If the Format is unspecified, then Writer.WriteHeader rounds ModTime
	 * to the nearest second and ignores the AccessTime and ChangeTime fields.
	 */
	ModTime    time.Time /* Modification time */
	AccessTime time.Time /* Access time */
	ChangeTime time.Time /* Change time */

	Devmajor    uint32 /* Major device number (valid for TypeLink) */
	Devminor    uint32 /* Minor device number (valid for TypeLink) */
	RawDevmajor uint32 /* Major for raw device number (valid for TypeBlock or TypeChar) */
	RawDevminor uint32 /* Minor for raw device number (valid for TypeBlock or TypeChar) */

	/*
	 * Magic contains a []byte that identifies the archive format
	 * (see 'Format' below).
	 */
	Magic []byte

	/*
	 * Format specifies the format of the tar header.
	 *
	 * If the format is unspecified when Writer.WriteHeader is called,
	 * then it uses the first format (in the order of USTAR, PAX, GNU)
	 * capable of encoding this Header (see Format).
	 */
	Format Format
}

func (f Format) String() string {
	/*
	 * TODO: Perhaps move this to another file since it
	 * isn't exactly part of describing the cpio format.
	 */
	identifier := f
	_, ok := formatNames[identifier]
	if ok {
		return formatNames[identifier]
	} else {
		return "<unknown>"
	}
}

func (f Format) HeaderLen() int64 {
	headerlen := 0
	switch {
	case ((f & TYPE_BINARY) != 0):
		headerlen = 26
	case ((f & TYPE_OCPIO) != 0):
		headerlen = 76
	case ((f & TYPE_NCPIO) != 0 || (f & TYPE_CRC) != 0):
		headerlen = 110
	}
	return int64(headerlen)
}
