package discord

import "errors"

var (
	ErrBadConnVisibility = errors.New("unrecognized connection visibility value")
)

type ConnectionVisibility uint8

const (
	ConnVisibilityNone ConnectionVisibility = iota
	ConnVisibilityEveryone
)

func (c ConnectionVisibility) IsValid() bool {
	return nil == c.Validate()
}

func (c ConnectionVisibility) Validate() error {
	if c > ConnVisibilityEveryone {
		return ErrBadConnVisibility
	}

	return nil
}
