package discord

import "errors"

var (
	ErrBadOverwriteType = errors.New("unrecognized overwrite type value")
)

type Type uint8

const (
	OverwriteTypeRole Type = iota
	OverwriteTypeMember
)

func (o Type) IsValid() bool {
	return nil == o.Validate()
}

func (o Type) Validate() error {
	if o > 1 {
		return ErrBadOverwriteType
	}

	return nil
}

