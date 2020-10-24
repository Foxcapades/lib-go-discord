package discord

import (
	"encoding/json"
	"errors"
	"time"
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

// TriStateInt8 represents a(n) int8 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateInt8 struct {
	value *int8
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateInt8) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateInt8) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateInt8) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateInt8) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateInt8) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateInt8) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateInt8) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateInt8) Set(val int8) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateInt8) Get() int8 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateInt8) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateInt8) UnmarshalJSON(bytes []byte) error {
	var tmp *int8

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

// TriStateInt16 represents a(n) int16 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateInt16 struct {
	value *int16
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateInt16) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateInt16) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateInt16) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateInt16) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateInt16) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateInt16) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateInt16) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateInt16) Set(val int16) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateInt16) Get() int16 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateInt16) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateInt16) UnmarshalJSON(bytes []byte) error {
	var tmp *int16

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

// TriStateInt32 represents a(n) int32 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateInt32 struct {
	value *int32
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateInt32) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateInt32) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateInt32) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateInt32) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateInt32) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateInt32) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateInt32) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateInt32) Set(val int32) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateInt32) Get() int32 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateInt32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateInt32) UnmarshalJSON(bytes []byte) error {
	var tmp *int32

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

// TriStateInt64 represents a(n) int64 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateInt64 struct {
	value *int64
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateInt64) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateInt64) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateInt64) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateInt64) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateInt64) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateInt64) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateInt64) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateInt64) Set(val int64) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateInt64) Get() int64 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateInt64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateInt64) UnmarshalJSON(bytes []byte) error {
	var tmp *int64

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

// TriStateUint16 represents a(n) uint16 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateUint16 struct {
	value *uint16
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateUint16) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateUint16) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateUint16) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateUint16) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateUint16) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateUint16) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateUint16) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateUint16) Set(val uint16) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateUint16) Get() uint16 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateUint16) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateUint16) UnmarshalJSON(bytes []byte) error {
	var tmp *uint16

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

// TriStateUint32 represents a(n) uint32 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateUint32 struct {
	value *uint32
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateUint32) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateUint32) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateUint32) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateUint32) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateUint32) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateUint32) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateUint32) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateUint32) Set(val uint32) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateUint32) Get() uint32 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateUint32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateUint32) UnmarshalJSON(bytes []byte) error {
	var tmp *uint32

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

// TriStateUint64 represents a(n) uint64 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateUint64 struct {
	value *uint64
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateUint64) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateUint64) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateUint64) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateUint64) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateUint64) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateUint64) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateUint64) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateUint64) Set(val uint64) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateUint64) Get() uint64 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateUint64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateUint64) UnmarshalJSON(bytes []byte) error {
	var tmp *uint64

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

// TriStateFloat32 represents a(n) float32 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateFloat32 struct {
	value *float32
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateFloat32) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateFloat32) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateFloat32) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateFloat32) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateFloat32) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateFloat32) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateFloat32) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateFloat32) Set(val float32) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateFloat32) Get() float32 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateFloat32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateFloat32) UnmarshalJSON(bytes []byte) error {
	var tmp *float32

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

// TriStateFloat64 represents a(n) float64 field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateFloat64 struct {
	value *float64
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateFloat64) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateFloat64) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateFloat64) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateFloat64) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateFloat64) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateFloat64) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateFloat64) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateFloat64) Set(val float64) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateFloat64) Get() float64 {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateFloat64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateFloat64) UnmarshalJSON(bytes []byte) error {
	var tmp *float64

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

// TriStateBool represents a(n) bool field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateBool struct {
	value *bool
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateBool) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateBool) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateBool) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateBool) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateBool) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateBool) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateBool) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateBool) Set(val bool) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateBool) Get() bool {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateBool) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateBool) UnmarshalJSON(bytes []byte) error {
	var tmp *bool

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

////////////////////////////////////////////////////////////////////////////////

// TriStateSnowflake represents a(n) Snowflake field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateSnowflake struct {
	value *Snowflake
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateSnowflake) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateSnowflake) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateSnowflake) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateSnowflake) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateSnowflake) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateSnowflake) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateSnowflake) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateSnowflake) Set(val Snowflake) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateSnowflake) Get() Snowflake {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateSnowflake) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateSnowflake) UnmarshalJSON(bytes []byte) error {
	var tmp *Snowflake

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

// TriStateTime represents a(n) time.Time field which may
// be one of 3 states: filled, nil, or absent entirely.
type TriStateTime struct {
	value *time.Time
	null  bool
}

// IsNull returns whether the value of the current field is nil (JSON null).
//
// This will return false if the field contains a value or is unset.
func (i TriStateTime) IsNull() bool {
	return i.value == nil && i.null
}

// IsNotNull returns whether the value of the current field is not nil (JSON
// null).
//
// This will return false if the field is null or is absent.
func (i TriStateTime) IsNotNull() bool {
	return i.value != nil || !i.null
}

// IsUnset returns whether the value of the current field is absent/unset.
//
// Returns false if the field contains a value or is null.
func (i TriStateTime) IsUnset() bool {
	return i.value == nil && !i.null
}

// IsSet returns whether the value of the current field is present.
//
// Returns false if the field is absent or is null.
func (i TriStateTime) IsSet() bool {
	return i.value != nil
}

// IsReadable returns whether the field is present and readable.
//
// Returns false if the field is unset or is null.
func (i TriStateTime) IsReadable() bool {
	return i.value != nil
}

// SetNull overwrites the current field value with JSON null.
func (i *TriStateTime) SetNull() TriStateField {
	i.value = nil
	i.null = true

	return i
}

// Unset removes the current field value entirely.
func (i *TriStateTime) Unset() TriStateField {
	i.value = nil
	i.null = false

	return i
}

// Set overwrites the current field value with the given value.
//
// If the type is nillable and a nil value is passed here, this method will
// panic.
func (i *TriStateTime) Set(val time.Time) TriStateField {
	i.value = &val
	i.null = false

	return i
}

// Get returns the current field value.
//
// If the current field is unset or is null, this method will panic.  Check if
// this method is safe to call by using the IsReadable method.
func (i TriStateTime) Get() time.Time {
	return *i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will return an ErrSerializeUnset error.
func (i TriStateTime) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TriStateTime) UnmarshalJSON(bytes []byte) error {
	var tmp *time.Time

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
