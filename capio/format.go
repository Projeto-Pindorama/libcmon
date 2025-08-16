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
 */

package capio

//type Format uint 

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
	/* Type of header. */
	ZILTCH uint = 00000000 /* no format chosen yet */

	HEADER_ODC    uint = 00002001 /* POSIX ASCII cpio format */
	HEADER_DEC    uint = 00002002 /* DEC extended cpio format */
	HEADER_BINLE  uint = 00003001 /* binary (default) cpio format LE */
	HEADER_BINBE  uint = 00003101 /* binary (default) cpio format BE */
	HEADER_SGILE  uint = 00003201 /* IRIX-style -K binary format LE */
	HEADER_SGIBE  uint = 00003301 /* IRIX-style -K binary format BE */
	HEADER_ASC    uint = 00004001 /* SVR4 ASCII cpio format */
	HEADER_SCOASC uint = 00004201 /* UnixWare 7.1 ASCII cpio format */
	HEADER_CRC    uint = 00004401 /* SVR4 ASCII cpio format w/checksum */
	HEADER_SCOCRC uint = 00004601 /* UnixWare 7.1 ASCII cpio w/checksum */
	HEADER_CRAY   uint = 00010001 /* Cray cpio, UNICOS 6 and later */
	HEADER_CRAY5  uint = 00010002 /* Cray cpio, UNICOS 5 and earlier */
	HEADER_BAR    uint = 00400001 /* bar format type */

	/* Characteristics of the archive ('masked' into integer). */
	TYPE_BE     uint = 00000100 /* this binary archive is big-endian */
	TYPE_SGI    uint = 00000200 /* SGI cpio -K flag binary archive */
	TYPE_SCO    uint = 00000200 /* SCO UnixWare 7.1 extended archive */
	TYPE_CRC    uint = 00000400 /* this has a SVR4 'crc' checksum */
	TYPE_BINARY uint = 00001000 /* this is a binary cpio type */
	TYPE_OCPIO  uint = 00002000 /* this is an old cpio type */
	TYPE_NCPIO  uint = 00004000 /* this is a SVR4 cpio type */
	TYPE_CRAY   uint = 00010000 /* this is a Cray cpio archive */
	TYPE_CPIO   uint = 00077000 /* this is a cpio type */
	TYPE_BAR    uint = 00400000 /* this is a bar type */
)

var formatNames = map[uint]string{
	ZILTCH:        "ZILTCH",
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

//func (f Format) String() string {
//	identifier := f
//	_, ok := formatNames[identifier]
//	if ok {
//		return formatNames[identifier]
//	} else {
//		return "<unknown>"
//	}
//}
