package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
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
	value *int8
}

func (i NullableInt8) IsNull() bool {
	return i.value == nil
}

func (i NullableInt8) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableInt8) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableInt8) Set(val int8) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt8) Get() int8 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableInt8) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *int16
}

func (i NullableInt16) IsNull() bool {
	return i.value == nil
}

func (i NullableInt16) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableInt16) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableInt16) Set(val int16) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt16) Get() int16 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableInt16) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *int32
}

func (i NullableInt32) IsNull() bool {
	return i.value == nil
}

func (i NullableInt32) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableInt32) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableInt32) Set(val int32) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt32) Get() int32 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableInt32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *int64
}

func (i NullableInt64) IsNull() bool {
	return i.value == nil
}

func (i NullableInt64) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableInt64) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableInt64) Set(val int64) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableInt64) Get() int64 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableInt64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *uint8
}

func (i NullableUint8) IsNull() bool {
	return i.value == nil
}

func (i NullableUint8) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableUint8) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableUint8) Set(val uint8) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint8) Get() uint8 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableUint8) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *uint16
}

func (i NullableUint16) IsNull() bool {
	return i.value == nil
}

func (i NullableUint16) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableUint16) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableUint16) Set(val uint16) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint16) Get() uint16 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableUint16) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *uint32
}

func (i NullableUint32) IsNull() bool {
	return i.value == nil
}

func (i NullableUint32) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableUint32) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableUint32) Set(val uint32) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint32) Get() uint32 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableUint32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *uint64
}

func (i NullableUint64) IsNull() bool {
	return i.value == nil
}

func (i NullableUint64) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableUint64) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableUint64) Set(val uint64) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableUint64) Get() uint64 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableUint64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *float32
}

func (i NullableFloat32) IsNull() bool {
	return i.value == nil
}

func (i NullableFloat32) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableFloat32) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableFloat32) Set(val float32) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableFloat32) Get() float32 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableFloat32) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *float64
}

func (i NullableFloat64) IsNull() bool {
	return i.value == nil
}

func (i NullableFloat64) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableFloat64) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableFloat64) Set(val float64) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableFloat64) Get() float64 {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableFloat64) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *bool
}

func (i NullableBool) IsNull() bool {
	return i.value == nil
}

func (i NullableBool) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableBool) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableBool) Set(val bool) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableBool) Get() bool {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableBool) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *string
}

func (i NullableString) IsNull() bool {
	return i.value == nil
}

func (i NullableString) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableString) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableString) Set(val string) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableString) Get() string {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableString) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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
	value *discord.Snowflake
}

func (i NullableSnowflake) IsNull() bool {
	return i.value == nil
}

func (i NullableSnowflake) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableSnowflake) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableSnowflake) Set(val discord.Snowflake) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableSnowflake) Get() discord.Snowflake {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableSnowflake) UnmarshalJSON(bytes []byte) error {
	tmp := NewSnowflakeImpl(false)
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	if tmp != nil {
		t2 := (discord.Snowflake)(tmp)
		i.value = &t2
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
	value *time.Time
}

func (i NullableTime) IsNull() bool {
	return i.value == nil
}

func (i NullableTime) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableTime) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableTime) Set(val time.Time) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableTime) Get() time.Time {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableTime) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
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

type NullableVerificationLevel struct {
	value *discord.VerificationLevel
}

func (i NullableVerificationLevel) IsNull() bool {
	return i.value == nil
}

func (i NullableVerificationLevel) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableVerificationLevel) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableVerificationLevel) Set(val discord.VerificationLevel) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableVerificationLevel) Get() discord.VerificationLevel {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableVerificationLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableVerificationLevel) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableMessageNotificationLevel struct {
	value *discord.MessageNotificationLevel
}

func (i NullableMessageNotificationLevel) IsNull() bool {
	return i.value == nil
}

func (i NullableMessageNotificationLevel) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableMessageNotificationLevel) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableMessageNotificationLevel) Set(val discord.MessageNotificationLevel) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableMessageNotificationLevel) Get() discord.MessageNotificationLevel {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableMessageNotificationLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableMessageNotificationLevel) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableExplicitContentFilterLevel struct {
	value *discord.ExplicitContentFilterLevel
}

func (i NullableExplicitContentFilterLevel) IsNull() bool {
	return i.value == nil
}

func (i NullableExplicitContentFilterLevel) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableExplicitContentFilterLevel) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableExplicitContentFilterLevel) Set(val discord.ExplicitContentFilterLevel) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableExplicitContentFilterLevel) Get() discord.ExplicitContentFilterLevel {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableExplicitContentFilterLevel) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableExplicitContentFilterLevel) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableChannelTopic struct {
	value *discord.ChannelTopic
}

func (i NullableChannelTopic) IsNull() bool {
	return i.value == nil
}

func (i NullableChannelTopic) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableChannelTopic) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableChannelTopic) Set(val discord.ChannelTopic) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableChannelTopic) Get() discord.ChannelTopic {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableChannelTopic) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableChannelTopic) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableAny struct {
	value *interface{}
}

func (i NullableAny) IsNull() bool {
	return i.value == nil
}

func (i NullableAny) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableAny) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableAny) Set(val interface{}) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableAny) Get() interface{} {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableAny) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableAny) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

var nullValue = []byte("null");
