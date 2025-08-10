/*
 * porcelana/bengala.go - Walking-stick-functions for avoiding repetition
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package prcl

import "reflect"

/*
 * IntWidth matches the integer width number for
 * a integer passed per reflect.Kind.
 */
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
