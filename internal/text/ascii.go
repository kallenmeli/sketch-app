package text

import (
	"sketch/internal/errors"
	"unicode"
)

type ASCIIChar string

var (
	ErrInvalidASCIIChar = errors.Error("invalid character, you must use a valid ASCII character")
)

func (a ASCIIChar) Validate() error {
	if len(a) == 0 {
		return nil
	}

	if len(a) > 1 {
		return ErrInvalidASCIIChar
	}

	if a[0] > unicode.MaxASCII {
		return ErrInvalidASCIIChar
	}

	return nil
}
