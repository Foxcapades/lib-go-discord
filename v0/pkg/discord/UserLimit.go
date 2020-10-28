package discord

import (
	"errors"
	"github.com/francoispqt/gojay"
)

var (
	ErrBadUserLimit = errors.New("invalid user limit value, user limits must be 0-99")
)

type UserLimit uint8

func (u UserLimit) IsValid() bool {
	return nil == u.Validate()
}

func (u UserLimit) Validate() error {
	if u > 99 {
		return ErrBadUserLimit
	}

	return nil
}

func DecodeUserLimit(dec *gojay.Decoder) (*UserLimit, error) {
	var tmp *uint8

	if err := dec.Uint8Null(&tmp); err != nil {
		return nil, err
	}

	return (*UserLimit)(tmp), nil
}
