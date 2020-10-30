package types

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
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
	tmp := com.NewSnowflake()
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	if tmp != nil {
		t2 := tmp
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

type NullableActivityEmoji struct {
	value *discord.ActivityEmoji
}

func (i NullableActivityEmoji) IsNull() bool {
	return i.value == nil
}

func (i NullableActivityEmoji) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableActivityEmoji) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableActivityEmoji) Set(val discord.ActivityEmoji) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableActivityEmoji) Get() discord.ActivityEmoji {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableActivityEmoji) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableActivityEmoji) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableActivityParty struct {
	value *discord.ActivityParty
}

func (i NullableActivityParty) IsNull() bool {
	return i.value == nil
}

func (i NullableActivityParty) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableActivityParty) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableActivityParty) Set(val discord.ActivityParty) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableActivityParty) Get() discord.ActivityParty {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableActivityParty) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableActivityParty) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableActivityAssets struct {
	value *discord.ActivityAssets
}

func (i NullableActivityAssets) IsNull() bool {
	return i.value == nil
}

func (i NullableActivityAssets) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableActivityAssets) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableActivityAssets) Set(val discord.ActivityAssets) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableActivityAssets) Get() discord.ActivityAssets {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableActivityAssets) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableActivityAssets) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

type NullableActivitySecrets struct {
	value *discord.ActivitySecrets
}

func (i NullableActivitySecrets) IsNull() bool {
	return i.value == nil
}

func (i NullableActivitySecrets) IsNotNull() bool {
	return i.value != nil
}

func (i *NullableActivitySecrets) SetNull() NullableField {
	i.value = nil

	return i
}

func (i *NullableActivitySecrets) Set(val discord.ActivitySecrets) NullableField {
	i.value = &val

	return i
}

// Value returns the raw wrapped field value.
//
// This method panics if the field is marked as null.
func (i NullableActivitySecrets) Get() discord.ActivitySecrets {
	if i.value == nil {
		panic(ErrNullField)
	}

	return *i.value
}

func (i *NullableActivitySecrets) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &i.value); err != nil {
		return err
	}

	return nil
}

func (i *NullableActivitySecrets) MarshalJSON() ([]byte, error) {
	if i.IsNull() {
		return nullValue, nil
	}

	return json.Marshal(i.value)
}

////////////////////////////////////////////////////////////////////////////////

var nullValue = []byte("null")
