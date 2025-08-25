/*
 * porcelana/bengala.go - Walking-stick-functions for avoiding repetition
 *
 * Copyright (C) 2025: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package prcl

import (
	"encoding/binary"
	"reflect"
	"strconv"
)

/* Just for the interface */
type mixedEndian struct{}

// MixedEndian is used on little-endian systems
// to store 32-bit integers in two 16-bit blocks.
var MixedEndian mixedEndian

func (mixedEndian) Uint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32([]byte{
		data[2],
		data[3],
		data[0],
		data[1],
	})
}

// OctalToInt converts an octal number (a.k.a. base-8)
// contained in a []byte to int64.
func OctalToInt(data []byte) int64 {
	numstr := ""

	/*
	 * Clean string so obtaining foolish-induced errors from
	 * strconv is unnecessary.
	 */
	for j := 0; j < len(data); j++ {
		switch {
		case '0' <= data[j] && data[j] <= '7':
			numstr += string(data[j])
		default:
			continue
		}
	}

	val, err := strconv.ParseInt(numstr, 8, 64)
	if err != nil {
		panic(err)
	}
	return val
}

// HexaToInt converts a hexadecimal number (a.k.a. base-16)
// contained in a []byte array into a int64.
func HexaToInt(data []byte) int64 {
	tdata := []byte("")
	numstr := ""
	j := int(0)

	/* Jump '0x', since it is considered invalid by strconv. */
	if string(data[:2]) == "0x" {
		j += 2
	}

	/*
	 * All to uppercase.
	 * We can also exclude actually undesired characters now,
	 * such as spaces/blanks, before treating it further below.
	 */
	for ; j < len(data); j++ {
		if data[j] == ' ' {
			continue
		} else if 'a' <= data[j] && data[j] <= 'z' {
			tdata = append(tdata, (data[j] - 32))
			continue
		}
		tdata = append(tdata, data[j])
	}

	/*
	 * Clean string so obtaining foolish-induced errors from
	 * strconv is unnecessary.
	 */
	for i := 0; i < len(tdata); i++ {
		switch {
		case ('0' <= tdata[i] && tdata[i] <= '9') ||
			('A' <= tdata[i] && tdata[i] <= 'F'):
			numstr += string(tdata[i])
		default:
			continue
		}
	}

	/*
	 * Just like OctalToInt, use Go stdlib's strconv function for actually
	 * converting the hexadecimal string into a integer, since it will be
	 * far more competent than something we could implement from scratch
	 * using the crude logic that we learn at our's university courses ---
	 * and as competent as the implementation that we could've borrowed and
	 * adapted from Heirloom's cpio.
	 */
	val, err := strconv.ParseInt(numstr, 16, 64)
	if err != nil {
		panic(err)
	}
	return val
}

// IntWidth matches the integer width number for
// a integer passed per reflect.Kind.
func IntWidth(reflectk reflect.Kind) int {
	/* (u)int type-to-size map. */
	int_sizes := map[reflect.Kind]int{
		reflect.Uint8:  8,
		reflect.Int8:   8,
		reflect.Uint16: 16,
		reflect.Int16:  16,
		reflect.Uint32: 32,
		reflect.Int32:  32,
		reflect.Uint64: 64,
		reflect.Int64:  64,
	}

	switch reflectk {
	/* Determine generic (u)int size per platform. */
	case reflect.Uint, reflect.Int:
		var x int
		return int(reflect.TypeOf(x).Size())
	default:
		break
	}

	/* Otherwise, verify the map. */
	if size, isitevenint := int_sizes[reflectk]; isitevenint {
		return size
	}

	return -1 /* Not an integer. */
}
