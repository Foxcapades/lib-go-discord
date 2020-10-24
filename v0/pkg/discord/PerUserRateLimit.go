package discord

import "errors"

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



