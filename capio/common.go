package capio

import "errors"

var (
	ErrHeader          = errors.New("capio: invalid header")
	ErrWriteTooLong    = errors.New("capio: write too long")
	ErrFieldTooLong    = errors.New("capio: header field too long")
	ErrWriteAfterClose = errors.New("capio: write after close")
	ErrInsecurePath    = errors.New("capio: insecure file path")
)
