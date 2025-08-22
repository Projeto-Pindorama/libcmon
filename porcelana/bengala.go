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
	"math"
	"reflect"
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

// OctalToInt converts a octal number contained in a []byte to int64.
func OctalToInt(data []byte) int64 {
	val := int64(0)

	for c, i := (len(data) - 1), 0; i < len(data); i++ {
		switch {
		case '0' <= data[i] && data[i] <= '7':
			cs := int64(math.Pow(float64(8), float64(c)))
			val += int64((data[i] - '0')) * cs
		default:
			c -= 1
			continue
		}
		c -= 1
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
