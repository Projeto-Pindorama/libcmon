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

/* Matches the integer width number for a
 * integer passed per reflect.Kind. */
func IntWidth(reflectk reflect.Kind) int {
	switch reflectk {
	case reflect.Uint, reflect.Int:
		var x int
		return int(reflect.TypeOf(x).Size())
	case reflect.Uint8, reflect.Int8:
		return 8
	case reflect.Uint16, reflect.Int16:
		return 16
	case reflect.Uint32, reflect.Int32:
		return 32
	case reflect.Uint64, reflect.Int64:
		return 64
	default:
		break
	}
	return -1 /* Not an integer. */
}
