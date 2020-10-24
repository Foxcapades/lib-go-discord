package discord

import (
	"errors"
)

var (
	ErrBadContentLength = errors.New("content length must not exceed 2000 characters")
)

type MessageContent string

func (c MessageContent) IsValid() bool {
	return nil == c.Validate()
}

func (c MessageContent) Validate() error {
	if len(c) > 2000 {
		return ErrBadContentLength
	}

	return nil
}
