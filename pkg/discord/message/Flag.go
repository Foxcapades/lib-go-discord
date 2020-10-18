package message

import "errors"

var (
	ErrBadMsgFlag = errors.New("unrecognized message flag value")
)

type Flag uint8

const (
	FlagCrossposted Flag = 1 << iota
	FlagIsCrossposted
	FlagSuppressEmbeds
	FlagSourceMessageDeleted
	FlagUrgent
)

const maxFlags Flag = 0b0001_1111

func (f Flag) IsValid() bool {
	return nil == f.Validate()
}

func (f Flag) Validate() error {
	if f > maxFlags {
		return ErrBadMsgFlag
	}

	return nil
}

