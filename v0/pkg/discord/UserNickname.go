package discord

import (
	"errors"
)

var (
	ErrBadNickLength = errors.New("nickname lengths must be 1-32 characters in length")
)

type UserNickname string

func (n UserNickname) IsValid() bool {
	return nil == n.Validate()
}

func (n UserNickname) Validate() error {
	ln := len(scrubName(string(n)))

	if ln < 1 || ln > 32 {
		return ErrBadNickLength
	}

	return nil
}

