/*
 * capio/format.go - Programmatical description of the cpio format
 *
 * Copyright (C) 2025: Pindorama
 *		Luiz Antônio Rangel (takusuman)
 *
 * SPDX-Licence-Identifier: BSD-3-Clause
 *
 */

package capio

import "errors"

var (
	ErrHeader          = errors.New("capio: invalid header; perhaps corrupt")
	ErrWriteTooLong    = errors.New("capio: write too long")
	ErrFieldTooLong    = errors.New("capio: header field too long")
	ErrWriteAfterClose = errors.New("capio: write after close")
	ErrInsecurePath    = errors.New("capio: insecure file path")
	ErrChecksum        = errors.New("capio: error computing checksum")
)
