package message

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

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

func (f Flag) Add(flag Flag) Flag {
	return f | flag
}

func (f Flag) Remove(flag Flag) Flag {
	return f &^ flag
}

func (f Flag) Mask(flag Flag) Flag {
	return f & flag
}

type TriStateFlagField struct {
	value *Flag
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

func (t TriStateFlagField) Get() Flag {
	if t.value == nil {
		if t.null {
			panic(dlib.ErrNullField)
		} else {
			panic(dlib.ErrUnsetField)
		}
	}

	return *t.value
}

func (t TriStateFlagField) Set(f Flag) {
	t.value = &f
	t.null = false
}