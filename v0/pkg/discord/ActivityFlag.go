package discord

import "errors"

var (
	ErrBadActivityFlag = errors.New("unrecognized activity flag value")
)

type ActivityFlag uint8

const (
	FlagInstance ActivityFlag = 1 << iota
	FlagJoin
	FlagSpectate
	FlagJoinRequest
	FlagSync
	FlagPlay
)

const flagMax ActivityFlag = 0b0011_1111

func (f ActivityFlag) IsValid() bool {
	return nil == f.Validate()
}

func (f ActivityFlag) Validate() error {
	if f > flagMax {
		return ErrBadActivityFlag
	}

	return nil
}

