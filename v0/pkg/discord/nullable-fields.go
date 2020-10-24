package discord

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

type NullableInt8 struct {
	value int8
	isSet bool
}

func (i NullableInt8) IsNull() bool {
	return !i.isSet
}

func (i NullableInt8) IsNotNull() bool {
	return i.isSet
}

func (i *NullableInt8) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableInt8) Set(val int8) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt8) Get() int8 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableInt8) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableInt8) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableInt16 struct {
	value int16
	isSet bool
}

func (i NullableInt16) IsNull() bool {
	return !i.isSet
}

func (i NullableInt16) IsNotNull() bool {
	return i.isSet
}

func (i *NullableInt16) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableInt16) Set(val int16) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt16) Get() int16 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableInt16) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableInt16) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableInt32 struct {
	value int32
	isSet bool
}

func (i NullableInt32) IsNull() bool {
	return !i.isSet
}

func (i NullableInt32) IsNotNull() bool {
	return i.isSet
}

func (i *NullableInt32) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableInt32) Set(val int32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt32) Get() int32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableInt32) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableInt32) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableInt64 struct {
	value int64
	isSet bool
}

func (i NullableInt64) IsNull() bool {
	return !i.isSet
}

func (i NullableInt64) IsNotNull() bool {
	return i.isSet
}

func (i *NullableInt64) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableInt64) Set(val int64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt64) Get() int64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableInt64) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableInt64) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableUint8 struct {
	value uint8
	isSet bool
}

func (i NullableUint8) IsNull() bool {
	return !i.isSet
}

func (i NullableUint8) IsNotNull() bool {
	return i.isSet
}

func (i *NullableUint8) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableUint8) Set(val uint8) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint8) Get() uint8 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableUint8) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableUint8) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableUint16 struct {
	value uint16
	isSet bool
}

func (i NullableUint16) IsNull() bool {
	return !i.isSet
}

func (i NullableUint16) IsNotNull() bool {
	return i.isSet
}

func (i *NullableUint16) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableUint16) Set(val uint16) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint16) Get() uint16 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableUint16) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableUint16) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableUint32 struct {
	value uint32
	isSet bool
}

func (i NullableUint32) IsNull() bool {
	return !i.isSet
}

func (i NullableUint32) IsNotNull() bool {
	return i.isSet
}

func (i *NullableUint32) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableUint32) Set(val uint32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint32) Get() uint32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableUint32) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableUint32) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableUint64 struct {
	value uint64
	isSet bool
}

func (i NullableUint64) IsNull() bool {
	return !i.isSet
}

func (i NullableUint64) IsNotNull() bool {
	return i.isSet
}

func (i *NullableUint64) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableUint64) Set(val uint64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint64) Get() uint64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableUint64) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableUint64) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableFloat32 struct {
	value float32
	isSet bool
}

func (i NullableFloat32) IsNull() bool {
	return !i.isSet
}

func (i NullableFloat32) IsNotNull() bool {
	return i.isSet
}

func (i *NullableFloat32) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableFloat32) Set(val float32) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableFloat32) Get() float32 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableFloat32) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableFloat32) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableFloat64 struct {
	value float64
	isSet bool
}

func (i NullableFloat64) IsNull() bool {
	return !i.isSet
}

func (i NullableFloat64) IsNotNull() bool {
	return i.isSet
}

func (i *NullableFloat64) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableFloat64) Set(val float64) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableFloat64) Get() float64 {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableFloat64) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableFloat64) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableBool struct {
	value bool
	isSet bool
}

func (i NullableBool) IsNull() bool {
	return !i.isSet
}

func (i NullableBool) IsNotNull() bool {
	return i.isSet
}

func (i *NullableBool) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableBool) Set(val bool) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableBool) Get() bool {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableBool) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableBool) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableString struct {
	value string
	isSet bool
}

func (i NullableString) IsNull() bool {
	return !i.isSet
}

func (i NullableString) IsNotNull() bool {
	return i.isSet
}

func (i *NullableString) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableString) Set(val string) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableString) Get() string {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableString) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableString) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableSnowflake struct {
	value Snowflake
	isSet bool
}

func (i NullableSnowflake) IsNull() bool {
	return !i.isSet
}

func (i NullableSnowflake) IsNotNull() bool {
	return i.isSet
}

func (i *NullableSnowflake) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableSnowflake) Set(val Snowflake) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableSnowflake) Get() Snowflake {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableSnowflake) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableSnowflake) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableTime struct {
	value time.Time
	isSet bool
}

func (i NullableTime) IsNull() bool {
	return !i.isSet
}

func (i NullableTime) IsNotNull() bool {
	return i.isSet
}

func (i *NullableTime) SetNull() NullableField {
	i.isSet = false

	return i
}

func (i *NullableTime) Set(val time.Time) NullableField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableTime) Get() time.Time {
	if !i.isSet {
		panic(ErrNullField)
	}

	return i.value
}

func (i *NullableTime) UnmarshalJSON(bytes []byte) error {
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

func (i *NullableTime) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

var nullValue = []byte("null");
