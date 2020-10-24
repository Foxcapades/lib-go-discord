package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

var (
	ErrUnsetField = errors.New("cannot get a value from an absent field; " +
		"use the *IsSet method before attempting to unwrap a nullable field")
	ErrSerializeUnset = errors.New("cannot serialize an absent field")
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

type OptionalInt8 struct {
	value *int8
}

func (i OptionalInt8) IsSet() bool {
	return i.value != nil
}

func (i OptionalInt8) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalInt8) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalInt8) Set(val int8) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt8) Get() int8 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalInt8) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *int16
}

func (i OptionalInt16) IsSet() bool {
	return i.value != nil
}

func (i OptionalInt16) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalInt16) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalInt16) Set(val int16) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt16) Get() int16 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalInt16) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *int32
}

func (i OptionalInt32) IsSet() bool {
	return i.value != nil
}

func (i OptionalInt32) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalInt32) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalInt32) Set(val int32) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt32) Get() int32 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalInt32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *int64
}

func (i OptionalInt64) IsSet() bool {
	return i.value != nil
}

func (i OptionalInt64) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalInt64) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalInt64) Set(val int64) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalInt64) Get() int64 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalInt64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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

type OptionalUint32 struct {
	value *uint32
}

func (i OptionalUint32) IsSet() bool {
	return i.value != nil
}

func (i OptionalUint32) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalUint32) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalUint32) Set(val uint32) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalUint32) Get() uint32 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalUint32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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

type OptionalFloat32 struct {
	value *float32
}

func (i OptionalFloat32) IsSet() bool {
	return i.value != nil
}

func (i OptionalFloat32) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalFloat32) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalFloat32) Set(val float32) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalFloat32) Get() float32 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalFloat32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	value *float64
}

func (i OptionalFloat64) IsSet() bool {
	return i.value != nil
}

func (i OptionalFloat64) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalFloat64) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalFloat64) Set(val float64) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalFloat64) Get() float64 {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalFloat64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

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
	tmp := NewSnowflakeImpl(false)
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

type OptionalTime struct {
	value *time.Time
}

func (i OptionalTime) IsSet() bool {
	return i.value != nil
}

func (i OptionalTime) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalTime) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalTime) Set(val time.Time) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalTime) Get() time.Time {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalTime) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalTime) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalVerificationLevel struct {
	value *discord.VerificationLevel
}

func (i OptionalVerificationLevel) IsSet() bool {
	return i.value != nil
}

func (i OptionalVerificationLevel) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalVerificationLevel) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalVerificationLevel) Set(val discord.VerificationLevel) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalVerificationLevel) Get() discord.VerificationLevel {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalVerificationLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalVerificationLevel) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalMessageNotificationLevel struct {
	value *discord.MessageNotificationLevel
}

func (i OptionalMessageNotificationLevel) IsSet() bool {
	return i.value != nil
}

func (i OptionalMessageNotificationLevel) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalMessageNotificationLevel) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalMessageNotificationLevel) Set(val discord.MessageNotificationLevel) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalMessageNotificationLevel) Get() discord.MessageNotificationLevel {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalMessageNotificationLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalMessageNotificationLevel) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalExplicitContentFilterLevel struct {
	value *discord.ExplicitContentFilterLevel
}

func (i OptionalExplicitContentFilterLevel) IsSet() bool {
	return i.value != nil
}

func (i OptionalExplicitContentFilterLevel) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalExplicitContentFilterLevel) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalExplicitContentFilterLevel) Set(val discord.ExplicitContentFilterLevel) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalExplicitContentFilterLevel) Get() discord.ExplicitContentFilterLevel {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalExplicitContentFilterLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalExplicitContentFilterLevel) MarshalJSON() ([]byte, error) {
	if i.IsUnset() {
		return nil, ErrSerializeUnset
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type OptionalChannelTopic struct {
	value *discord.ChannelTopic
}

func (i OptionalChannelTopic) IsSet() bool {
	return i.value != nil
}

func (i OptionalChannelTopic) IsUnset() bool {
	return i.value == nil
}

func (i *OptionalChannelTopic) Unset() OptionalField {
	i.value = nil

	return i
}

func (i *OptionalChannelTopic) Set(val discord.ChannelTopic) OptionalField {
	i.value = &val

	return i
}

// Value returns the raw value contained by this field.
//
// If this field is unset, this method will panic.
func (i OptionalChannelTopic) Get() discord.ChannelTopic {
	if i.value == nil {
		panic(ErrUnsetField)
	}

	return *i.value
}

func (i *OptionalChannelTopic) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *OptionalChannelTopic) MarshalJSON() ([]byte, error) {
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
