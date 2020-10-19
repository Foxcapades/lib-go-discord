package channel

import "errors"

var (
	ErrBadNameLength = errors.New("channel names must be 2-100 characters in length")
)

type Name string

func (n Name) IsValid() bool {
	return nil == n.Validate()
}

func (n Name) Validate() error {
	l := len(n)
	if l < 2 || l > 100 {
		return ErrBadNameLength
	}

	return nil
}

