package discord

import (
	"errors"
)

var (
	ErrBadTargetUserType = errors.New("unrecognized target user type value")
)

type InviteTargetUserType uint8

const (
	TargetUserTypeStream InviteTargetUserType = 1
)

func (t InviteTargetUserType) IsValid() bool {
	return nil == t.Validate()
}

func (t InviteTargetUserType) Validate() error {
	if t != TargetUserTypeStream {
		return ErrBadTargetUserType
	}

	return nil
}
