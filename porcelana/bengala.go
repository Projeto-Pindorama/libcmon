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

func IntWidth(reflectk reflect.Kind) int {
	switch reflectk {
	case reflect.Int:
		var x int
		return int(reflect.TypeOf(x).Size())
	case reflect.Int8:
		return 8
	case reflect.Int16:
		return 16
	case reflect.Int32:
		return 32
	case reflect.Int64:
		return 64
	default:
		return -1
	}
}
