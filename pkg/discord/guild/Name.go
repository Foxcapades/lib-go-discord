package guild

import (
	"errors"
	"strings"
)

var (
	ErrBadNameLength = errors.New("guild names must be 2-100 characters in " +
		"length excluding leading/trailing whitespace")
)

type Name string

func (n Name) IsValid() bool {
	return nil == n.Validate()
}

func (n Name) Validate() error {
	v := strings.TrimSpace(string(n))
	l := len(v)

	if l < 2 || l > 100 {
		return ErrBadNameLength
	}

	return nil
}

