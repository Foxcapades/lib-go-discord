package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
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

type OptionalTargetUserType struct {
	value *InviteTargetUserType
}

func (o OptionalTargetUserType) IsSet() bool {
	return o.value != nil
}

func (o OptionalTargetUserType) IsUnset() bool {
	return o.value == nil
}

func (o OptionalTargetUserType) Unset() {
	o.value = nil
}

func (o OptionalTargetUserType) Get() InviteTargetUserType {
	if o.value == nil {
		panic(dlib.ErrUnsetField)
	}

	return *o.value
}

func (o OptionalTargetUserType) Set(t InviteTargetUserType) {
	o.value = &t
}
