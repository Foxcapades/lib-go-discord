package user

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type OptionalUser struct {
	value User
	isSet bool
}

func (i OptionalUser) IsSet() bool {
	return i.isSet
}

func (i OptionalUser) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalUser) Unset() dlib.OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUser) Set(val User) dlib.OptionalField {
	if val == nil {
		panic(ErrSetNilUser)
	}

	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUser) Value() User {
	if !i.isSet {
		panic(dlib.ErrUnsetField)
	}

	return i.value
}

func (i *OptionalUser) UnmarshalJSON(bytes []byte) error {
	var tmp userImpl

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = &tmp
	i.isSet = true

	return nil
}

func (i *OptionalUser) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, dlib.ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

