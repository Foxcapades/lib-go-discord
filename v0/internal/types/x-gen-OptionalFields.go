package types

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

var (
	ErrUnsetField = errors.New("cannot get a value from an absent field; " +
		"use the *IsSet method before attempting to unwrap a nullable field")
	ErrSerializeUnset    = errors.New("cannot serialize an absent field")
	ErrSetNilOptionalVal = errors.New("attempted to set a nil value using an" +
		" OptionalField's Set method")
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

type OptionalUint8 struct {
	value *uint8
}

func (i OptionalUint8) IsSet() bool {
	return i.value != nil
}

func (i OptionalUint8) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalUint8) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalUint8) Set(val uint8) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint8) Get() uint8 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalUint8) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *uint16
}

func (i OptionalUint16) IsSet() bool {
	return i.value != nil
}

func (i OptionalUint16) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalUint16) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalUint16) Set(val uint16) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint16) Get() uint16 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalUint16) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalUint16) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalUint64 struct {
	value *uint64
}

func (i OptionalUint64) IsSet() bool {
	return i.value != nil
}

func (i OptionalUint64) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalUint64) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalUint64) Set(val uint64) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint64) Get() uint64 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalUint64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalUint64) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalBool struct {
	value *bool
}

func (i OptionalBool) IsSet() bool {
	return i.value != nil
}

func (i OptionalBool) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalBool) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalBool) Set(val bool) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalBool) Get() bool {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalBool) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *string
}

func (i OptionalString) IsSet() bool {
	return i.value != nil
}

func (i OptionalString) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalString) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalString) Set(val string) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalString) Get() string {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalString) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *discord.Snowflake
}

func (i OptionalSnowflake) IsSet() bool {
	return i.value != nil
}

func (i OptionalSnowflake) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalSnowflake) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalSnowflake) Set(val discord.Snowflake) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalSnowflake) Get() discord.Snowflake {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalSnowflake) UnmarshalJSON(bytes []byte) error {
	tmp := com.NewSnowflake()
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	t2 := (discord.Snowflake)(tmp)
	i.value = &t2

	return nil
}

func (i *OptionalSnowflake) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalAny struct {
	value *interface{}
}

func (i OptionalAny) IsSet() bool {
	return i.value != nil
}

func (i OptionalAny) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalAny) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalAny) Set(val interface{}) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalAny) Get() interface{} {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalAny) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalAny) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}
