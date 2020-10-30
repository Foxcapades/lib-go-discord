package types

import (
	"encoding/json"
	"errors"
)

var (
	ErrSetNilTriStateVal = errors.New("attempted to set a nil value using a" +
		" TriStateField's Set method")
)

// TriStateField represents a field which may be absent, null, or containing a
// value.
type TriStateField interface {
	json.Marshaler
	json.Unmarshaler

	OptionContainer
	NullContainer

	// SetNull sets the current field value to null.
	SetNull() TriStateField

	// Unset marks the current field as absent.
	Unset() TriStateField

	// IsReadable returns true if the current field is set and non-null.
	IsReadable() bool
}

////////////////////////////////////////////////////////////////////////////////

// TriStateUint8 represents a(n) uint8 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateUint8 struct {
	value *uint8
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateUint8) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateUint8) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateUint8) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateUint8) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateUint8) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateUint8) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateUint8) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateUint8) Set(val uint8) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateUint8) Get() uint8 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateUint8) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateUint8) UnmarshalJSON(bytes []byte) error {
	var tmp *uint8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = tmp
		i.null = false
	} else {
		i.value = nil
		i.null = true
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// TriStateString represents a(n) string field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateString struct {
	value *string
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateString) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateString) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateString) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateString) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateString) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateString) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateString) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateString) Set(val string) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateString) Get() string {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateString) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateString) UnmarshalJSON(bytes []byte) error {
	var tmp *string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = tmp
		i.null = false
	} else {
		i.value = nil
		i.null = true
	}

	return nil
}
