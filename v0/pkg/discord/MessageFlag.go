package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
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

type TriStateFlagField struct {
	value *MessageFlag
	null  bool
}

func (t TriStateFlagField) IsSet() bool {
	return t.value != nil
}

func (t TriStateFlagField) IsUnset() bool {
	return t.value == nil && !t.null
}

func (t TriStateFlagField) IsNull() bool {
	return t.value == nil && t.null
}

func (t TriStateFlagField) IsNotNull() bool {
	return !t.null
}

func (t TriStateFlagField) IsReadable() bool {
	return t.value != nil
}

func (t TriStateFlagField) SetNull() {
	t.value = nil
	t.null = true
}

func (t TriStateFlagField) Unset() {
	t.value = nil
	t.null = false
}

func (t TriStateFlagField) Get() MessageFlag {
	if t.value == nil {
		if t.null {
			panic(dlib.ErrNullField)
		} else {
			panic(dlib.ErrUnsetField)
		}
	}

	return *t.value
}

func (t TriStateFlagField) Set(f MessageFlag) {
	t.value = &f
	t.null = false
}