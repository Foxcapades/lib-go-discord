package channel

import "errors"

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
