package dlib

import (
	"encoding/json"
	"time"
)

type TriStateField interface {
	json.Marshaler
	json.Unmarshaler

	OptionContainer
	NullContainer

	// SetNull sets the current field value to null.
	SetNull() TriStateField

	// Unset marks the current field as absent.
	Unset() TriStateField
}

////////////////////////////////////////////////////////////////////////////////

type I8TriStateField struct {
	value int8
	flags flag
}

func (i I8TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i I8TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i I8TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i I8TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *I8TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *I8TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *I8TriStateField) Set(val int8) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i I8TriStateField) Value() int8 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i I8TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i I8TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *int8

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type I16TriStateField struct {
	value int16
	flags flag
}

func (i I16TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i I16TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i I16TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i I16TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *I16TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *I16TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *I16TriStateField) Set(val int16) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i I16TriStateField) Value() int16 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i I16TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i I16TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *int16

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type I32TriStateField struct {
	value int32
	flags flag
}

func (i I32TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i I32TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i I32TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i I32TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *I32TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *I32TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *I32TriStateField) Set(val int32) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i I32TriStateField) Value() int32 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i I32TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i I32TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *int32

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type I64TriStateField struct {
	value int64
	flags flag
}

func (i I64TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i I64TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i I64TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i I64TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *I64TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *I64TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *I64TriStateField) Set(val int64) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i I64TriStateField) Value() int64 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i I64TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i I64TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *int64

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type U8TriStateField struct {
	value uint8
	flags flag
}

func (i U8TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i U8TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i U8TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i U8TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *U8TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *U8TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *U8TriStateField) Set(val uint8) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i U8TriStateField) Value() uint8 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i U8TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i U8TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint8

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type U16TriStateField struct {
	value uint16
	flags flag
}

func (i U16TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i U16TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i U16TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i U16TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *U16TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *U16TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *U16TriStateField) Set(val uint16) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i U16TriStateField) Value() uint16 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i U16TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i U16TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint16

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type U32TriStateField struct {
	value uint32
	flags flag
}

func (i U32TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i U32TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i U32TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i U32TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *U32TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *U32TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *U32TriStateField) Set(val uint32) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i U32TriStateField) Value() uint32 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i U32TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i U32TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint32

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type U64TriStateField struct {
	value uint64
	flags flag
}

func (i U64TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i U64TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i U64TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i U64TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *U64TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *U64TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *U64TriStateField) Set(val uint64) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i U64TriStateField) Value() uint64 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i U64TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i U64TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint64

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type F32TriStateField struct {
	value float32
	flags flag
}

func (i F32TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i F32TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i F32TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i F32TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *F32TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *F32TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *F32TriStateField) Set(val float32) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i F32TriStateField) Value() float32 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i F32TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i F32TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *float32

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type F64TriStateField struct {
	value float64
	flags flag
}

func (i F64TriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i F64TriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i F64TriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i F64TriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *F64TriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *F64TriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *F64TriStateField) Set(val float64) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i F64TriStateField) Value() float64 {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i F64TriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i F64TriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *float64

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type BoolTriStateField struct {
	value bool
	flags flag
}

func (i BoolTriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i BoolTriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i BoolTriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i BoolTriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *BoolTriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *BoolTriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *BoolTriStateField) Set(val bool) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i BoolTriStateField) Value() bool {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i BoolTriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i BoolTriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *bool

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type StrTriStateField struct {
	value string
	flags flag
}

func (i StrTriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i StrTriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i StrTriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i StrTriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *StrTriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *StrTriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *StrTriStateField) Set(val string) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i StrTriStateField) Value() string {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i StrTriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i StrTriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *string

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type SnowflakeTriStateField struct {
	value Snowflake
	flags flag
}

func (i SnowflakeTriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i SnowflakeTriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i SnowflakeTriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i SnowflakeTriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *SnowflakeTriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *SnowflakeTriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *SnowflakeTriStateField) Set(val Snowflake) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i SnowflakeTriStateField) Value() Snowflake {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i SnowflakeTriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i SnowflakeTriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *Snowflake

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type TimeTriStateField struct {
	value time.Time
	flags flag
}

func (i TimeTriStateField) IsNull() bool {
	return !isSet(i.flags, flagNotNull)
}

func (i TimeTriStateField) IsNotNull() bool {
	return isSet(i.flags, flagNotNull)
}

func (i TimeTriStateField) IsUnset() bool {
	return !isSet(i.flags, flagPresent)
}

func (i TimeTriStateField) IsSet() bool {
	return isSet(i.flags, flagPresent)
}

func (i *TimeTriStateField) SetNull() TriStateField {
	i.flags = add(sub(i.flags, flagNotNull), flagPresent)

	return i
}

func (i *TimeTriStateField) Unset() TriStateField {
	i.flags = sub(sub(i.flags, flagPresent), flagNotNull)

	return i
}

func (i *TimeTriStateField) Set(val time.Time) TriStateField {
	i.flags = add(add(i.flags, flagNotNull), flagPresent)
	i.value = val

	return i
}

func (i TimeTriStateField) Value() time.Time {
	return i.value
}

// MarshalJSON serializes the current field to JSON.
//
// If this field is unset, this method will error with
func (i TimeTriStateField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

func (i TimeTriStateField) UnmarshalJSON(bytes []byte) error {
	var tmp *time.Time

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.flags = add(i.flags, flagNotNull)
		i.value = *tmp
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type flag uint8

const (
	flagPresent flag = 1 << iota
	flagNotNull
)

func isSet(val, test flag) bool {
	return val & test == test
}

func add(val, add flag) flag {
	return val | add
}

func sub(val, sub flag) flag {
	return val & ^sub
}
