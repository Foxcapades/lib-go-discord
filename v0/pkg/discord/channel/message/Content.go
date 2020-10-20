package message

import "errors"

var (
	ErrBadContentLength = errors.New("content length must not exceed 2000 characters")
)

type Content string

func (c Content) IsValid() bool {
	return nil == c.Validate()
}

func (c Content) Validate() error {
	if len(c) > 2000 {
		return ErrBadContentLength
	}

	return nil
}

