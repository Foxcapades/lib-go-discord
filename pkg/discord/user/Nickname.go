package user

import "errors"

var (
	ErrBadNickLength = errors.New("nickname lengths must be 1-32 characters in length")
)

type Nickname string

func (n Nickname) IsValid() bool {
	return nil == n.Validate()
}

func (n Nickname) Validate() error {
	ln := len(scrubName(string(n)))

	if ln < 1 || ln > 32 {
		return ErrBadNickLength
	}

	return nil
}

