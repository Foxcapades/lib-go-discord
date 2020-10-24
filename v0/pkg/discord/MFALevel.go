package discord

import "errors"

var (
	ErrBadMFALevel = errors.New("unrecognized MFA level value")
)

type MFALevel uint8

const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

func (m MFALevel) IsValid() bool {
	return nil == m.Validate()
}

func (m MFALevel) Validate() error {
	if m > MFALevelElevated {
		return ErrBadMFALevel
	}

	return nil
}

