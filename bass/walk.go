/*
 * bass/walk.go - Some hip functions for walking through
 * file pointers and making 'em into []byte(s).
 * Ain't you ever seen a walking bass?
 *
 * Copyright (C) 2024: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package bass

import (
	"bytes"
	"errors"
	"io"
	"os"
)

/*
 * Walk 'walks' a file descriptor until a certain point ('place'),
 * or from one point to another, with the end point and the initial
 * one being informed, respectively. If the end point is informed as
 * '-1', it walks until EOF.
 */
func Walk(f *os.File, place ...int64) ([]byte, int64, error) {
	var r int64
	var b []byte
	buf := make([]byte, 1)

	if len(place) < 1 {
		return nil, 0,
			errors.New("Walk() requires at least one point to go.")
	} else if len(place) >= 2 {
		/* Place to walk from. */
		from := place[1]
		p, err := f.Seek(from, 0)
		if err != nil {
			return nil, 0, err
		}

		/*
		 * Since 'r' is the number of bytes
		 * walked through in total, sum
		 * where we're walking from.
		 */
		r += p
	}
	to := place[0]

coda:
	for r = 0; to != 0; to, r = (to - 1), (r + 1) {
		_, err := f.Read(buf)
		switch err {
		case nil:
			b = append(b, buf[0])
		case io.EOF:
			break coda
		default:
			return nil, 0, err
		}
	}
	return b, r, nil
}

/*
 * WalkTil walks a file descriptor until a certain octet ('here') is
 * encountered and returns a 'byte' array containing everything before
 * it. Can be really useful when parsing a binary header, or reading lines
 * terminated in '\n'.
 */
func WalkTil(here byte, f *os.File) ([]byte, int, error) {
	var b []byte
	var i int
	bb, _, err := Walk(f, -1)
	if err != nil {
		return nil, 0, err
	}

	for i = 0; i < len(bb); i++ {
		if bb[i] != here {
			b = append(b, bb[i])
		} else {
			break
		}
	}

	return b, i, nil
}

/*
 * WalkLookinFor works similarly to WalkTil, but it looks for an entire
 * 'byte' array and returns a boolean for whether the said array is
 * encountered in the file (or not). Can be useful for detecting certain
 * characteristics in a binary.
 */
func WalkLookinFor(this []byte, at *os.File, place ...int64) (bool, int64, error) {
	var n, stop, nstep int64
	nstep = int64(len(this))
	found := false
	n = 0
	stop = -1

	/*
	 * Just like bass.Walk() itself, we can set from
	 * whence we want to start Walk()ing, saving some
	 * CPU cycles and I/O accesses.
	 */
	places := len(place)
	if places > 0 {
		stop = place[0]
		/*
		 * We can also set until where we want to go,
		 * for avoiding reading a entire file if we
		 * know that the said information won't be
		 * present after the said place.
		 */
		if places >= 2 {
			n += place[1]
		}
	}

exit:
	for ; !found; n++ {
		/* Check if it arrived to the limit, if any. */
		if stop == 0 {
			break exit
		}
		stop--

		b, _, err := Walk(at, nstep, n)
		if err != nil {
			return false, n, err
		}

		if (bytes.Equal(this, b)) {
			found = true
			/*
			 * Some funky arithmetic I discovered for returning
			 * the correct number of "binary places" Walk()ed
			 * pass-by.
			 */
			n--
			n += nstep
		}
	}

	return found, n, nil
}
