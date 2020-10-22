package discord

import "errors"

var (
	ErrBadExpConFilterLvl = errors.New("unrecognized explicit content filter level value")
)

type ExplicitContentFilterLevel uint8

const (
	ExpConFilterLvlDisabled ExplicitContentFilterLevel = iota
	ExpConFilterLvlMembersWithoutRoles
	ExpConFilterLvlAllMembers
)

func (e ExplicitContentFilterLevel) IsValid() bool {
	return nil == e.Validate()
}

func (e ExplicitContentFilterLevel) Validate() error {
	if e > ExpConFilterLvlAllMembers {
		return ErrBadExpConFilterLvl
	}

	return nil
}

