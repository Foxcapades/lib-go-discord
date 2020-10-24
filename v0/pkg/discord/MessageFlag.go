package discord

import (
	"errors"
)

var (
	ErrBadMsgFlag = errors.New("unrecognized message flag value")
)

type MessageFlag uint8

const (
	MsgFlagCrossposted MessageFlag = 1 << iota
	MsgFlagIsCrossposted
	MsgFlagSuppressEmbeds
	MsgFlagSourceMessageDeleted
	MsgFlagUrgent
)

const maxFlags MessageFlag = 0b0001_1111

func (f MessageFlag) IsValid() bool {
	return nil == f.Validate()
}

func (f MessageFlag) Validate() error {
	if f > maxFlags {
		return ErrBadMsgFlag
	}

	return nil
}

func (f MessageFlag) Add(flag MessageFlag) MessageFlag {
	return f | flag
}

func (f MessageFlag) Remove(flag MessageFlag) MessageFlag {
	return f &^ flag
}

func (f MessageFlag) Mask(flag MessageFlag) MessageFlag {
	return f & flag
}
