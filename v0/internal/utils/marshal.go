package utils

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
)

var (
	ErrTooShort = errors.New("value is too short to be a valid JSON string")
	ErrBadFmt   = errors.New("badly formatted JSON string")
)

// deprecated
func UnmarshalString(val []byte) (string, error) {
	sz := len(val)

	if sz < 2 {
		return "", ErrTooShort
	}

	if val[0] != '"' || val[sz-1] != '"' {
		return "", ErrBadFmt
	}

	return string(val[1 : sz-1]), nil
}

func UnmarshalStringInto(val []byte, to *string) error {
	sz := len(val)

	if sz < 2 {
		return ErrTooShort
	}

	if val[0] != '"' || val[sz-1] != '"' {
		return ErrBadFmt
	}

	*to = string(val[1 : sz-1])

	return nil
}

func MarshalString(str string) (buf []byte) {
	sz := len(str)
	buf = make([]byte, sz+js.QuoteSize)
	buf[0] = '"'
	buf[sz-1] = '"'
	copy(buf[1:], str)

	return
}

var (
	tBuf = []byte("true")
	fBuf = []byte("false")
)

func MarshalBool(v bool) (buf []byte) {
	if v {
		return tBuf
	}

	return fBuf
}
