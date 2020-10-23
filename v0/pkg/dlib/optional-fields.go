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

type OptionalInt8 struct {
	value int8
	isSet bool
}

func (i OptionalInt8) IsSet() bool {
	return i.isSet
}

func (i OptionalInt8) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalInt8) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalInt8) Set(val int8) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt8) Get() int8 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalInt8) UnmarshalJSON(bytes []byte) error {
	var tmp int8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalInt8) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalInt16 struct {
	value int16
	isSet bool
}

func (i OptionalInt16) IsSet() bool {
	return i.isSet
}

func (i OptionalInt16) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalInt16) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalInt16) Set(val int16) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt16) Get() int16 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalInt16) UnmarshalJSON(bytes []byte) error {
	var tmp int16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalInt16) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalInt32 struct {
	value int32
	isSet bool
}

func (i OptionalInt32) IsSet() bool {
	return i.isSet
}

func (i OptionalInt32) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalInt32) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalInt32) Set(val int32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt32) Get() int32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalInt32) UnmarshalJSON(bytes []byte) error {
	var tmp int32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalInt32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalInt64 struct {
	value int64
	isSet bool
}

func (i OptionalInt64) IsSet() bool {
	return i.isSet
}

func (i OptionalInt64) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalInt64) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalInt64) Set(val int64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt64) Get() int64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalInt64) UnmarshalJSON(bytes []byte) error {
	var tmp int64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalInt64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalUint8 struct {
	value uint8
	isSet bool
}

func (i OptionalUint8) IsSet() bool {
	return i.isSet
}

func (i OptionalUint8) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalUint8) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUint8) Set(val uint8) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint8) Get() uint8 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalUint8) UnmarshalJSON(bytes []byte) error {
	var tmp uint8

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalUint8) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalUint16 struct {
	value uint16
	isSet bool
}

func (i OptionalUint16) IsSet() bool {
	return i.isSet
}

func (i OptionalUint16) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalUint16) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUint16) Set(val uint16) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint16) Get() uint16 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalUint16) UnmarshalJSON(bytes []byte) error {
	var tmp uint16

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalUint16) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalUint32 struct {
	value uint32
	isSet bool
}

func (i OptionalUint32) IsSet() bool {
	return i.isSet
}

func (i OptionalUint32) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalUint32) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUint32) Set(val uint32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint32) Get() uint32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalUint32) UnmarshalJSON(bytes []byte) error {
	var tmp uint32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalUint32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalUint64 struct {
	value uint64
	isSet bool
}

func (i OptionalUint64) IsSet() bool {
	return i.isSet
}

func (i OptionalUint64) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalUint64) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalUint64) Set(val uint64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint64) Get() uint64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalUint64) UnmarshalJSON(bytes []byte) error {
	var tmp uint64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalUint64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalFloat32 struct {
	value float32
	isSet bool
}

func (i OptionalFloat32) IsSet() bool {
	return i.isSet
}

func (i OptionalFloat32) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalFloat32) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalFloat32) Set(val float32) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalFloat32) Get() float32 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalFloat32) UnmarshalJSON(bytes []byte) error {
	var tmp float32

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalFloat32) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalFloat64 struct {
	value float64
	isSet bool
}

func (i OptionalFloat64) IsSet() bool {
	return i.isSet
}

func (i OptionalFloat64) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalFloat64) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalFloat64) Set(val float64) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalFloat64) Get() float64 {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalFloat64) UnmarshalJSON(bytes []byte) error {
	var tmp float64

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalFloat64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalBool struct {
	value bool
	isSet bool
}

func (i OptionalBool) IsSet() bool {
	return i.isSet
}

func (i OptionalBool) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalBool) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalBool) Set(val bool) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalBool) Get() bool {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalBool) UnmarshalJSON(bytes []byte) error {
	var tmp bool

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalBool) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalString struct {
	value string
	isSet bool
}

func (i OptionalString) IsSet() bool {
	return i.isSet
}

func (i OptionalString) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalString) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalString) Set(val string) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalString) Get() string {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalString) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalString) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalSnowflake struct {
	value Snowflake
	isSet bool
}

func (i OptionalSnowflake) IsSet() bool {
	return i.isSet
}

func (i OptionalSnowflake) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalSnowflake) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalSnowflake) Set(val Snowflake) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalSnowflake) Get() Snowflake {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalSnowflake) UnmarshalJSON(bytes []byte) error {
	var tmp Snowflake

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalSnowflake) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalTime struct {
	value time.Time
	isSet bool
}

func (i OptionalTime) IsSet() bool {
	return i.isSet
}

func (i OptionalTime) IsUnset() bool {
	return !i.isSet
}

func (i *OptionalTime) Unset() OptionalField {
	i.isSet = false

	return i
}

func (i *OptionalTime) Set(val time.Time) OptionalField {
	i.isSet = true
	i.value = val

	return i
}

// Get returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalTime) Get() time.Time {
	if !i.isSet {
		panic(ErrUnsetField)
	}

	return i.value
}

func (i *OptionalTime) UnmarshalJSON(bytes []byte) error {
	var tmp time.Time

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	i.value = tmp
	i.isSet = true

	return nil
}

func (i *OptionalTime) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}
