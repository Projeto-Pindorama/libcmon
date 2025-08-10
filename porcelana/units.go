/*
 * porcelana/units.go - General unit conversion functions
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package prcl

const (
	/* Metric. */
	K = 1000
	M = (K * 1000)
	G = (M * 1000)
	T = (G * 1000)
	P = (T * 1000)
	/* ISO/IEC 80000. */
	Ki = 1024
	Mi = (Ki * 1024)
	Gi = (Mi * 1024)
	Ti = (Gi * 1024)
	Pi = (Ti * 1024)
)

/*
 * DiskSectorsTo calculates how much a number of sectors ('have') is
 * in human-comprehensible units, with 'want' being a constant from
 * this file that relates the mentioned unit with its size in octets.
 * 'sectsize' is usually 512, but it is possible that it changes, so please
 * attempt to obtain it in your specific environment --- on Linux, this can
 * be obtained for a specific disk from /sys/block.
 */
func DiskSectorsTo(want uint64, sectsiz uint16, have uint64) float64 {
	/* Octets per sector divided per the unit. */
	return (float64(have*uint64(sectsiz)) / float64(want))
}

/*
 * DiskSectorsToHuman is a boilerplate to DiskSectorsTo. A really nice one.
 * It also returns the unit name as a string.
 */
func DiskSectorsToHuman(sectsiz uint16, nsectors uint64) (float64, string) {
	v := float64((nsectors * uint64(sectsiz)))
	unit := "B"

	if v >= Ki && v < Mi {
		v = DiskSectorsTo(Ki, sectsiz, nsectors)
		unit = "KiB"
	} else if v >= Mi && v < Gi {
		v = DiskSectorsTo(Mi, sectsiz, nsectors)
		unit = "MiB"
	} else if v >= Gi && v < Ti {
		v = DiskSectorsTo(Gi, sectsiz, nsectors)
		unit = "GiB"
	} else if v >= Ti && v < Pi {
		v = DiskSectorsTo(Ti, sectsiz, nsectors)
		unit = "TiB"
	} else if v >= Pi {
		v = DiskSectorsTo(Pi, sectsiz, nsectors)
		unit = "PiB"
	}

	return v, unit
}
