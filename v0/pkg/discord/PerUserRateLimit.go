package discord

import (
	"errors"
	"github.com/francoispqt/gojay"
)

var (
	ErrBadUserRateLimit = errors.New("per_user_rate_limit values cannot be " +
		"greater than 21600")
)

type PerUserRateLimit uint16

func (p PerUserRateLimit) IsValid() bool {
	return nil == p.Validate()
}

func (p PerUserRateLimit) Validate() error {
	if p > 21600 {
		return ErrBadUserRateLimit
	}

	return nil
}

func DecodeUserRateLimit(dec *gojay.Decoder) (*PerUserRateLimit, error) {
	var tmp *uint16

	if err := dec.Uint16Null(&tmp); err != nil {
		return nil, err
	}

	return (*PerUserRateLimit)(tmp), nil
}
