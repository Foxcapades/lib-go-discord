package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"
)

// Errors thrown by methods in this file.
var (
	// Thrown if attempting to serialize a User instance that has no `id` value
	// set.
	//
	// This will only be thrown if the User instance has validation enabled.
	ErrNoId = errors.New("id field missing on User instance")

	// Thrown if attempting to serialize a User instance that has no `username`
	// value set.
	//
	// This will only be thrown if the User instance has validation enabled.
	ErrNoUsername = errors.New("username field missing for User instance")

	// Thrown if attempting to serialize a User instance that has no
	// `discriminator` value set.
	//
	// This will only be thrown if the User instance has validation enabled.
	ErrNoDiscriminator = errors.New("discriminator field missing for User instance")

	ErrSetNilUser = errors.New("cannot call OptionalUser.Set with a nil value")
)

// User TODO: Document me
type User interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// ID returns the value of the `id` field currently set on this user record.
	//
	// The `id` field contains the the user's Snowflake id.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	ID() Snowflake

	// SetID overwrites the current value of this user record's `id` field with
	// the given param value.
	SetID(snowflake Snowflake) User

	// Username returns the value of the `username` field currently set on this
	// user record.
	//
	// The `username` field contains the the user's username (not unique across
	// the platform).
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	Username() string

	// SetUsername overwrites the current value of this user record's `username`
	// field with the given param value.
	SetUsername(name Username) User

	// Discriminator returns the value of the `discriminator` field currently set
	// on this user record.
	//
	// The `discriminator` field contains the user's 4-digit discord tag.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	Discriminator() Discriminator

	// SetDiscriminator overwrites the current value of this user record's
	// `discriminator` field with the given param value.
	SetDiscriminator(code Discriminator) User

	// AvatarHash returns the value of the `avatar_hash` field currently set on
	// this user record.
	//
	// The `avatar_hash` field contains the user's the user's avatar hash
	// (https://discord.com/developers/docs/reference#image-formatting).
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `avatar_hash` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// AvatarHashIsSet method.
	AvatarHash() string
	AvatarHashIsNull() bool
	SetAvatarHash(string) User
	SetNullAvatarHash() User

	// BotFlag returns the value of the `bot` field currently set on this user
	// record.
	//
	// The `bot` flag indicates whether the user belongs to an OAuth2 application.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `bot` field is unset when this method is called, this
	// method will panic.  Check for the field's existence with the BotFlagIsSet
	// method.
	BotFlag() bool
	BotFlagIsSet() bool
	SetBotFlag(bool) User
	UnsetBotFlag() User

	// SystemFlag returns the value of the `system` field currently set on this
	// user record.
	//
	// The `system` flag indicates whether the user is an Official Discord System
	// user (part of the urgent message system).
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `system` field is unset when this method is called,
	// this method will panic.  Check for the field's existence with the
	// SystemFlagIsSet method.
	SystemFlag() bool
	SystemFlagIsSet() bool
	SetSystemFlag(bool) User
	UnsetSystemFlag() User

	// MFAEnabled returns the value of the `mfa_enabled` field currently set on
	// this user record.
	//
	// The `mfa_enabled` flag indicates whether the user has two factor
	// authentication enabled on their account.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `mfa_enabled` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// MFAEnabledIsSet method.
	MFAEnabled() bool
	MFAEnabledIsSet() bool
	SetMFAEnabled(bool) User
	UnsetMFAEnabled() User

	// Locale returns the value of the `locale` field currently set on this user
	// record.
	//
	// The `locale` field contains the user's chosen language option.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `locale` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// LocaleIsSet method.
	Locale() string
	LocaleIsSet() bool
	SetLocale(string) User
	UnsetLocale() User

	// VerifiedFlag returns the value of the `verified` field currently set on
	// this user record.
	//
	// The `verified` flag indicates whether the email on this account has been
	// verified.
	//
	// Getting this field from the Discord API requires the `email` OAuth2 scope.
	//
	// If the user record's `verified` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// VerifiedFlagIsSet method.
	VerifiedFlag() bool
	VerifiedFlagIsSet() bool
	SetVerifiedFlag(bool) User
	UnsetVerifiedFlag() User

	// Email returns the value of the `email` field currently set on this user
	// record.
	//
	// The `email` field contains the user's email address.
	//
	// Getting this field from the Discord API requires the `email` OAuth2 scope.
	//
	// If the user record's `email` field is unset or null when this method is
	// called, this method will panic.  Check for the field's existence and null
	// state with the EmailIsSet and EmailIsNull methods.
	Email() string
	EmailIsSet() bool
	EmailIsNull() bool
	SetEmail(string) User
	UnsetEmail() User
	SetNullEmail() User

	// Flags returns the value of the `flags` field currently set on this user
	// record.
	//
	// The `flags` field contains the flags set on a user's account.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `flags` field is unset when this method is called,
	// this method will panic.  Check for the field's existence with the
	// FlagsIsSet method.
	Flags() UserFlag

	// FlagsIsSet returns whether the current user record's flags field exists.
	FlagsIsSet() bool

	// SetFlags overwrites the current user record's flag(s) with the given value.
	SetFlags(UserFlag) User

	// UnsetFlags removes this user record's flags field.
	UnsetFlags() User

	// AddFlag appends the given user flag to this User record's current flag
	// value(s).
	//
	// If the field is unset when this method is called this method, the given
	// parameter value will become the new field value.
	AddFlag(UserFlag) User

	// RemoveFlag removes the given user flag from this User record's current flag
	// value(s).
	//
	// If the field is unset when this method is called, this method does nothing.
	// become the new field value.
	RemoveFlag(UserFlag) User

	// FlagsContains returns whether or not the public flag value(s) currently set
	// on this user record contain or equal the given value.
	FlagsContains(UserFlag) bool

	// PremiumType returns the value of the `premium_type` field currently set on
	// this user record.
	//
	// The `premium_type` field contains the the flags set on a user's account.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `premium_type` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// PremiumTypeIsSet method.
	PremiumType() UserPremiumType
	PremiumTypeIsSet() bool
	SetPremiumType(UserPremiumType) User
	UnsetPremiumType() User

	// PublicFlags returns the value of the `public_flags` field currently set on
	// this user record.
	//
	// The `public_flags` field contains the the public flags on a user's account.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	//
	// If the user record's `public_flags` field is unset when this method is
	// called, this method will panic.  Check for the field's existence with the
	// PublicFlagsIsSet method.
	PublicFlags() UserFlag

	// PublicFlagsIsSet returns whether the current user record's public flags
	// field exists.
	PublicFlagsIsSet() bool

	// SetPublicFlags overwrites the current user record's public flag(s) with the
	// given value.
	SetPublicFlags(UserFlag) User

	// UnsetPublicFlags removes this user record's flags field.
	UnsetPublicFlags() User

	// AddPublicFlag appends the given user flag to this User record's current
	// public flag value(s).
	//
	// If the field is unset when this method is called this method, the given
	// parameter value will become the new field value.
	AddPublicFlag(UserFlag) User

	// RemovePublicFlag removes the given user flag from this User record's
	// current public flag value(s).
	//
	// If the field is unset when this method is called, this method does nothing.
	// become the new field value.
	RemovePublicFlag(UserFlag) User

	// PublicFlagsContains returns whether or not the public flag value(s)
	// currently set on this user record contain or equal the given value.
	PublicFlagsContains(UserFlag) bool
}
