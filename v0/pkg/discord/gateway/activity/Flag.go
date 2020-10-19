package activity

import "errors"

var (
	ErrBadActivityFlag = errors.New("unrecognized activity flag value")
)

type Flag uint8

const (
	FlagInstance Flag = 1 << iota
	FlagJoin
	FlagSpectate
	FlagJoinRequest
	FlagSync
	FlagPlay
)

const flagMax Flag = 0b0011_1111

func (f Flag) IsValid() bool {
	return nil == f.Validate()
}

func (f Flag) Validate() error {
	if f > flagMax {
		return ErrBadActivityFlag
	}

	return nil
}

