package invite

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

var (
	ErrBadTargetUserType = errors.New("unrecognized target user type value")
)

type TargetUserType uint8

const (
	TargetUserTypeStream TargetUserType = 1
)

func (t TargetUserType) IsValid() bool {
	return nil == t.Validate()
}

func (t TargetUserType) Validate() error {
	if t != TargetUserTypeStream {
		return ErrBadTargetUserType
	}

	return nil
}

type OptionalTargetUserType struct {
	value *TargetUserType
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

func (o OptionalTargetUserType) Get() TargetUserType {
	if o.value == nil {
		panic(dlib.ErrUnsetField)
	}

	return *o.value
}

func (o OptionalTargetUserType) Set(t TargetUserType) {
	o.value = &t
}
