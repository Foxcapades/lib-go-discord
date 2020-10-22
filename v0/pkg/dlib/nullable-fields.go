package dlib

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrNullField = errors.New("cannot get a value from a null field; use " +
		"the *IsNull method before attempting to unwrap a nullable field")
)

type NullContainer interface {
	// IsNull returns whether the current field value is null.
	IsNull() bool

	// IsNotNull returns whether the current field value is not null.
	IsNotNull() bool
}

type NullableField interface {
	json.Marshaler
	json.Unmarshaler

	NullContainer

	// SetNull sets the current field to null.
	SetNull() NullableField
}

////////////////////////////////////////////////////////////////////////////////

type I8NullableField struct {
	value int8
	isSet bool
}

func (i I8NullableField) IsNull() bool {
	return !i.isSet
}

func (i I8NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *I8NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *I8NullableField) Set(val int8) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i I8NullableField) Get() int8 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *I8NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *int8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *I8NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I16NullableField struct {
	value int16
	isSet bool
}

func (i I16NullableField) IsNull() bool {
	return !i.isSet
}

func (i I16NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *I16NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *I16NullableField) Set(val int16) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i I16NullableField) Get() int16 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *I16NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *int16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *I16NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I32NullableField struct {
	value int32
	isSet bool
}

func (i I32NullableField) IsNull() bool {
	return !i.isSet
}

func (i I32NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *I32NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *I32NullableField) Set(val int32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i I32NullableField) Get() int32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *I32NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *int32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *I32NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type I64NullableField struct {
	value int64
	isSet bool
}

func (i I64NullableField) IsNull() bool {
	return !i.isSet
}

func (i I64NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *I64NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *I64NullableField) Set(val int64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i I64NullableField) Get() int64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *I64NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *int64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *I64NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U8NullableField struct {
	value uint8
	isSet bool
}

func (i U8NullableField) IsNull() bool {
	return !i.isSet
}

func (i U8NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *U8NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *U8NullableField) Set(val uint8) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i U8NullableField) Get() uint8 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *U8NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *U8NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U16NullableField struct {
	value uint16
	isSet bool
}

func (i U16NullableField) IsNull() bool {
	return !i.isSet
}

func (i U16NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *U16NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *U16NullableField) Set(val uint16) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i U16NullableField) Get() uint16 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *U16NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *U16NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U32NullableField struct {
	value uint32
	isSet bool
}

func (i U32NullableField) IsNull() bool {
	return !i.isSet
}

func (i U32NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *U32NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *U32NullableField) Set(val uint32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i U32NullableField) Get() uint32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *U32NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *U32NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type U64NullableField struct {
	value uint64
	isSet bool
}

func (i U64NullableField) IsNull() bool {
	return !i.isSet
}

func (i U64NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *U64NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *U64NullableField) Set(val uint64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i U64NullableField) Get() uint64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *U64NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *uint64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *U64NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type F32NullableField struct {
	value float32
	isSet bool
}

func (i F32NullableField) IsNull() bool {
	return !i.isSet
}

func (i F32NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *F32NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *F32NullableField) Set(val float32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i F32NullableField) Get() float32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *F32NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *float32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *F32NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type F64NullableField struct {
	value float64
	isSet bool
}

func (i F64NullableField) IsNull() bool {
	return !i.isSet
}

func (i F64NullableField) IsNotNull() bool {
	return i.isSet
}

func (i *F64NullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *F64NullableField) Set(val float64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i F64NullableField) Get() float64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *F64NullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *float64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *F64NullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type BoolNullableField struct {
	value bool
	isSet bool
}

func (i BoolNullableField) IsNull() bool {
	return !i.isSet
}

func (i BoolNullableField) IsNotNull() bool {
	return i.isSet
}

func (i *BoolNullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *BoolNullableField) Set(val bool) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i BoolNullableField) Get() bool {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *BoolNullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *bool

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *BoolNullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type StrNullableField struct {
	value string
	isSet bool
}

func (i StrNullableField) IsNull() bool {
	return !i.isSet
}

func (i StrNullableField) IsNotNull() bool {
	return i.isSet
}

func (i *StrNullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *StrNullableField) Set(val string) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i StrNullableField) Get() string {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *StrNullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *StrNullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type SnowflakeNullableField struct {
	value Snowflake
	isSet bool
}

func (i SnowflakeNullableField) IsNull() bool {
	return !i.isSet
}

func (i SnowflakeNullableField) IsNotNull() bool {
	return i.isSet
}

func (i *SnowflakeNullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *SnowflakeNullableField) Set(val Snowflake) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i SnowflakeNullableField) Get() Snowflake {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *SnowflakeNullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *Snowflake

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *SnowflakeNullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type TimeNullableField struct {
	value time.Time
	isSet bool
}

func (i TimeNullableField) IsNull() bool {
	return !i.isSet
}

func (i TimeNullableField) IsNotNull() bool {
	return i.isSet
}

func (i *TimeNullableField) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *TimeNullableField) Set(val time.Time) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i TimeNullableField) Get() time.Time {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *TimeNullableField) UnmarshalJSON(bytes []byte) error {
	var tmp *time.Time

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if tmp != nil {
		i.value = *tmp
		i.isSet = true
	}

	return nil
}

func (i *TimeNullableField) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

var nullValue = []byte("null");
