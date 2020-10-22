package dlib

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrUnsetField = errors.New("cannot get a value from an absent field; " +
		"use the *IsSet method before attempting to unwrap a nullable field")
	ErrSerializeUnset = errors.New("cannot serialize an absent field")
)

type OptionContainer interface {
	// IsSet returns whether the current field value is present.
	IsSet() bool

	// IsUnset returns whether the current field value is absent.
	IsUnset() bool
}

type OptionalField interface {
	json.Marshaler
	json.Unmarshaler

	OptionContainer

	// Unset marks the field as absent.
	Unset() OptionalField
}

////////////////////////////////////////////////////////////////////////////////

type I8OptionalField struct {
	value int8
	isSet bool
}

func (i I8OptionalField) IsSet() bool {
	return i.isSet
}

func (i I8OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *I8OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *I8OptionalField) Set(val int8) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i I8OptionalField) Get() int8 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *I8OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp int8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *I8OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I16OptionalField struct {
	value int16
	isSet bool
}

func (i I16OptionalField) IsSet() bool {
	return i.isSet
}

func (i I16OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *I16OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *I16OptionalField) Set(val int16) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i I16OptionalField) Get() int16 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *I16OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp int16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *I16OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I32OptionalField struct {
	value int32
	isSet bool
}

func (i I32OptionalField) IsSet() bool {
	return i.isSet
}

func (i I32OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *I32OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *I32OptionalField) Set(val int32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i I32OptionalField) Get() int32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *I32OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp int32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *I32OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I64OptionalField struct {
	value int64
	isSet bool
}

func (i I64OptionalField) IsSet() bool {
	return i.isSet
}

func (i I64OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *I64OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *I64OptionalField) Set(val int64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i I64OptionalField) Get() int64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *I64OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp int64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *I64OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U8OptionalField struct {
	value uint8
	isSet bool
}

func (i U8OptionalField) IsSet() bool {
	return i.isSet
}

func (i U8OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *U8OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *U8OptionalField) Set(val uint8) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i U8OptionalField) Get() uint8 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *U8OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp uint8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *U8OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U16OptionalField struct {
	value uint16
	isSet bool
}

func (i U16OptionalField) IsSet() bool {
	return i.isSet
}

func (i U16OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *U16OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *U16OptionalField) Set(val uint16) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i U16OptionalField) Get() uint16 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *U16OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp uint16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *U16OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U32OptionalField struct {
	value uint32
	isSet bool
}

func (i U32OptionalField) IsSet() bool {
	return i.isSet
}

func (i U32OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *U32OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *U32OptionalField) Set(val uint32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i U32OptionalField) Get() uint32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *U32OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp uint32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *U32OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U64OptionalField struct {
	value uint64
	isSet bool
}

func (i U64OptionalField) IsSet() bool {
	return i.isSet
}

func (i U64OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *U64OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *U64OptionalField) Set(val uint64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i U64OptionalField) Get() uint64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *U64OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp uint64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *U64OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type F32OptionalField struct {
	value float32
	isSet bool
}

func (i F32OptionalField) IsSet() bool {
	return i.isSet
}

func (i F32OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *F32OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *F32OptionalField) Set(val float32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i F32OptionalField) Get() float32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *F32OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp float32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *F32OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type F64OptionalField struct {
	value float64
	isSet bool
}

func (i F64OptionalField) IsSet() bool {
	return i.isSet
}

func (i F64OptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *F64OptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *F64OptionalField) Set(val float64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i F64OptionalField) Get() float64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *F64OptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp float64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *F64OptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type BoolOptionalField struct {
	value bool
	isSet bool
}

func (i BoolOptionalField) IsSet() bool {
	return i.isSet
}

func (i BoolOptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *BoolOptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *BoolOptionalField) Set(val bool) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i BoolOptionalField) Get() bool {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *BoolOptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp bool

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *BoolOptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type StrOptionalField struct {
	value string
	isSet bool
}

func (i StrOptionalField) IsSet() bool {
	return i.isSet
}

func (i StrOptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *StrOptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *StrOptionalField) Set(val string) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i StrOptionalField) Get() string {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *StrOptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *StrOptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type SnowflakeOptionalField struct {
	value Snowflake
	isSet bool
}

func (i SnowflakeOptionalField) IsSet() bool {
	return i.isSet
}

func (i SnowflakeOptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *SnowflakeOptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *SnowflakeOptionalField) Set(val Snowflake) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i SnowflakeOptionalField) Get() Snowflake {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *SnowflakeOptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp Snowflake

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *SnowflakeOptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type TimeOptionalField struct {
	value time.Time
	isSet bool
}

func (i TimeOptionalField) IsSet() bool {
	return i.isSet
}

func (i TimeOptionalField) IsUnset() bool {
	return !i.isSet
}

func (i *TimeOptionalField) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *TimeOptionalField) Set(val time.Time) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i TimeOptionalField) Get() time.Time {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *TimeOptionalField) UnmarshalJSON(bytes []byte) error {
	var tmp time.Time

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *TimeOptionalField) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}
