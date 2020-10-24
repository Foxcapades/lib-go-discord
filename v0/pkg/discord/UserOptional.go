package discord

import (
	"encoding/json"
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

func (i *OptionalUser) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUser) Set(val User) OptionalField {
	if val == nil {
		panic(ErrSetNilUser)
	}

	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUser) Get() User {
	if !i.isSet {
		panic(ErrUnsetField)
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
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

