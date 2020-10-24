package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/derr"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Errors                                                              ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

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

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Types                                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║    User Type                                                           ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// User TODO: Document me
type User interface {
	json.Marshaler
	json.Unmarshaler

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
	SetUsername(name user.Username) User

	// Discriminator returns the value of the `discriminator` field currently set
	// on this user record.
	//
	// The `discriminator` field contains the user's 4-digit discord tag.
	//
	// Getting this field from the Discord API requires the `identify` OAuth2
	// scope.
	Discriminator() user.Discriminator

	// SetDiscriminator overwrites the current value of this user record's
	// `discriminator` field with the given param value.
	SetDiscriminator(code user.Discriminator) User

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
	Flags() user.Flag

	// FlagsIsSet returns whether the current user record's flags field exists.
	FlagsIsSet() bool

	// SetFlags overwrites the current user record's flag(s) with the given value.
	SetFlags(user.Flag) User

	// UnsetFlags removes this user record's flags field.
	UnsetFlags() User

	// AddFlag appends the given user flag to this User record's current flag
	// value(s).
	//
	// If the field is unset when this method is called this method, the given
	// parameter value will become the new field value.
	AddFlag(user.Flag) User

	// RemoveFlag removes the given user flag from this User record's current flag
	// value(s).
	//
	// If the field is unset when this method is called, this method does nothing.
	// become the new field value.
	RemoveFlag(user.Flag) User

	// FlagsContains returns whether or not the public flag value(s) currently set
	// on this user record contain or equal the given value.
	FlagsContains(user.Flag) bool

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
	PremiumType() user.PremiumType
	PremiumTypeIsSet() bool
	SetPremiumType(user.PremiumType) User
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
	PublicFlags() user.Flag

	// PublicFlagsIsSet returns whether the current user record's public flags
	// field exists.
	PublicFlagsIsSet() bool

	// SetPublicFlags overwrites the current user record's public flag(s) with the
	// given value.
	SetPublicFlags(user.Flag) User

	// UnsetPublicFlags removes this user record's flags field.
	UnsetPublicFlags() User

	// AddPublicFlag appends the given user flag to this User record's current
	// public flag value(s).
	//
	// If the field is unset when this method is called this method, the given
	// parameter value will become the new field value.
	AddPublicFlag(user.Flag) User

	// RemovePublicFlag removes the given user flag from this User record's
	// current public flag value(s).
	//
	// If the field is unset when this method is called, this method does nothing.
	// become the new field value.
	RemovePublicFlag(user.Flag) User

	// PublicFlagsContains returns whether or not the public flag value(s)
	// currently set on this user record contain or equal the given value.
	PublicFlagsContains(user.Flag) bool
}

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Functions                                                    ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// Returns a new, empty User instance.
//
// Data Validation
//
// The "validate" flag enables/disables internal input validation.  This will
// check that values being set via the type's setter methods are valid according
// to the documented Discord restrictions.
//
// Note: internal input validation does not guarantee that Discord will accept
// the value, there are undocumented internal checks that the Discord API does
// in addition to the documented checks.  This feature is primarily intended to
// help catch internal bugs and/or avoid making requests with obviously bad
// data.
func NewUser(validate bool) User {
	return &userImpl{validate: validate}
}

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Private Types                                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

type userImpl struct {
	validate bool

	id            Snowflake
	username      user.Username
	discriminator user.Discriminator
	avatar        NullableString
	bot           OptionalBool
	system        OptionalBool
	mfaEnabled    OptionalBool
	locale        OptionalString
	verified      OptionalBool
	email         TriStateString
	flags         OptionalUint32
	premiumType   OptionalUint8
	publicFlags   OptionalUint32
}

func (u *userImpl) MarshalJSON() ([]byte, error) {
	out := make(map[FieldKey]interface{}, 10)

	if u.id.RawValue() == 0 {
		return nil, ErrNoId
	} else {
		out[FieldKeyID] = u.id
	}

	if u.username == "" {
		return nil, ErrNoUsername
	} else {
		out[FieldKeyUsername] = u.username
	}

	if u.discriminator == 0 {
		return nil, ErrNoDiscriminator
	} else {
		out[FieldKeyDiscriminator] = u.discriminator
	}

	out[FieldKeyAvatar] = u.avatar
	u.appendIfSet1(out, FieldKeyBot, &u.bot)
	u.appendIfSet1(out, FieldKeySystem, &u.system)
	u.appendIfSet1(out, FieldKeyMFAEnabled, &u.mfaEnabled)
	u.appendIfSet1(out, FieldKeyLocale, &u.locale)
	u.appendIfSet1(out, FieldKeyVerified, &u.verified)
	u.appendIfSet2(out, FieldKeyEmail, &u.email)
	u.appendIfSet1(out, FieldKeyFlags, &u.flags)
	u.appendIfSet1(out, FieldKeyPremiumType, &u.premiumType)
	u.appendIfSet1(out, FieldKeyPublicFlags, &u.publicFlags)

	return json.Marshal(out)
}

func (u *userImpl) UnmarshalJSON(bytes []byte) (err error) {
	in := make(map[string]json.RawMessage, 10)

	if err = json.Unmarshal(bytes, &in); err != nil {
		return
	}

	multis := map[FieldKey]json.Unmarshaler{
		FieldKeyAvatar:      &u.avatar,
		FieldKeyBot:         &u.bot,
		FieldKeySystem:      &u.system,
		FieldKeyMFAEnabled:  &u.mfaEnabled,
		FieldKeyLocale:      &u.locale,
		FieldKeyVerified:    &u.verified,
		FieldKeyEmail:       &u.email,
		FieldKeyFlags:       &u.flags,
		FieldKeyPremiumType: &u.premiumType,
		FieldKeyPublicFlags: &u.publicFlags,
	}

	singles := map[FieldKey]interface{}{
		FieldKeyID:            &u.id,
		FieldKeyUsername:      &u.username,
		FieldKeyDiscriminator: &u.discriminator,
	}

	for k, v := range in {
		key := FieldKey(k)

		if f, ok := singles[key]; ok {
			err = json.Unmarshal(v, f)
		} else if f, ok := multis[key]; ok {
			err = f.UnmarshalJSON(v)
		}

		if err != nil {
			return
		}
	}

	return nil
}

func (u *userImpl) ID() Snowflake {
	return u.id
}

func (u *userImpl) SetID(id Snowflake) User {
	u.id = id

	return u
}

func (u *userImpl) Username() string {
	return string(u.username)
}

func (u *userImpl) SetUsername(name user.Username) User {
	if u.validate {
		if err := name.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyUsername), name, err))
		}
	}

	u.username = name

	return u
}

func (u *userImpl) Discriminator() user.Discriminator {
	return u.discriminator
}

func (u *userImpl) SetDiscriminator(code user.Discriminator) User {
	if u.validate {
		if err := code.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyDiscriminator), code, err))
		}
	}

	u.discriminator = code

	return u
}

func (u *userImpl) AvatarHash() string {
	return u.avatar.Get()
}

func (u *userImpl) AvatarHashIsNull() bool {
	return u.avatar.IsNull()
}

func (u *userImpl) SetAvatarHash(s string) User {
	u.avatar.Set(s)

	return u
}

func (u *userImpl) SetNullAvatarHash() User {
	u.avatar.SetNull()

	return u
}

func (u *userImpl) BotFlag() bool {
	return u.bot.Get()
}

func (u *userImpl) BotFlagIsSet() bool {
	return u.bot.IsSet()
}

func (u *userImpl) SetBotFlag(b bool) User {
	u.bot.Set(b)

	return u
}

func (u *userImpl) UnsetBotFlag() User {
	u.bot.Unset()

	return u
}

func (u *userImpl) SystemFlag() bool {
	return u.system.Get()
}

func (u *userImpl) SystemFlagIsSet() bool {
	return u.system.IsSet()
}

func (u *userImpl) SetSystemFlag(b bool) User {
	u.system.Set(b)

	return u
}

func (u *userImpl) UnsetSystemFlag() User {
	u.system.Unset()

	return u
}

func (u *userImpl) MFAEnabled() bool {
	return u.mfaEnabled.Get()
}

func (u *userImpl) MFAEnabledIsSet() bool {
	return u.mfaEnabled.IsSet()
}

func (u *userImpl) SetMFAEnabled(b bool) User {
	u.mfaEnabled.Set(b)

	return u
}

func (u *userImpl) UnsetMFAEnabled() User {
	u.mfaEnabled.Unset()

	return u
}

func (u *userImpl) Locale() string {
	return u.locale.Get()
}

func (u *userImpl) LocaleIsSet() bool {
	return u.locale.IsSet()
}

func (u *userImpl) SetLocale(s string) User {
	u.locale.Set(s)

	return u
}

func (u *userImpl) UnsetLocale() User {
	u.locale.Unset()

	return u
}

func (u *userImpl) VerifiedFlag() bool {
	return u.verified.Get()
}

func (u *userImpl) VerifiedFlagIsSet() bool {
	return u.verified.IsSet()
}

func (u *userImpl) SetVerifiedFlag(b bool) User {
	u.verified.Set(b)

	return u
}

func (u *userImpl) UnsetVerifiedFlag() User {
	u.verified.Unset()

	return u
}

func (u *userImpl) Email() string {
	return u.email.Get()
}

func (u *userImpl) EmailIsSet() bool {
	return u.email.IsSet()
}

func (u *userImpl) EmailIsNull() bool {
	return u.email.IsNull()
}

func (u *userImpl) SetEmail(s string) User {
	u.email.Set(s)

	return u
}

func (u *userImpl) UnsetEmail() User {
	u.email.Unset()

	return u
}

func (u *userImpl) SetNullEmail() User {
	u.email.SetNull()

	return u
}

func (u *userImpl) Flags() user.Flag {
	return user.Flag(u.flags.Get())
}

func (u *userImpl) FlagsIsSet() bool {
	return u.flags.IsSet()
}

func (u *userImpl) SetFlags(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyFlags), flag, err))
		}
	}

	u.flags.Set(uint32(flag))

	return u
}

func (u *userImpl) UnsetFlags() User {
	u.flags.Unset()

	return u
}

func (u *userImpl) AddFlag(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() | uint32(flag))
	} else {
		u.flags.Set(uint32(flag))
	}

	return u
}

func (u *userImpl) RemoveFlag(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() & ^uint32(flag))
	}

	return u
}

func (u *userImpl) FlagsContains(flag user.Flag) bool {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		return u.flags.Get()|uint32(flag) == uint32(flag)
	}

	return false
}

func (u *userImpl) PremiumType() user.PremiumType {
	return user.PremiumType(u.premiumType.Get())
}

func (u *userImpl) PremiumTypeIsSet() bool {
	return u.premiumType.IsSet()
}

func (u *userImpl) SetPremiumType(val user.PremiumType) User {
	if u.validate {
		if err := val.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyPremiumType), val, err))
		}
	}

	u.premiumType.Set(uint8(val))

	return u
}

func (u *userImpl) UnsetPremiumType() User {
	u.premiumType.Unset()

	return u
}

func (u *userImpl) PublicFlags() user.Flag {
	return user.Flag(u.publicFlags.Get())
}

func (u *userImpl) PublicFlagsIsSet() bool {
	return u.publicFlags.IsSet()
}

func (u *userImpl) SetPublicFlags(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyPublicFlags), flag, err))
		}
	}

	u.publicFlags.Set(uint32(flag))
	return u
}

func (u *userImpl) UnsetPublicFlags() User {
	u.publicFlags.Unset()
	return u
}

func (u *userImpl) AddPublicFlag(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyPublicFlags), flag, err))
		}
	}

	if u.publicFlags.IsSet() {
		u.publicFlags.Set(uint32(flag) | u.publicFlags.Get())
	} else {
		u.publicFlags.Set(uint32(flag))
	}

	return u
}

func (u *userImpl) RemovePublicFlag(flag user.Flag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyPublicFlags), flag, err))
		}
	}

	if u.publicFlags.IsSet() {
		u.publicFlags.Set(u.publicFlags.Get() & ^uint32(flag))
	}

	return u
}

func (u *userImpl) PublicFlagsContains(flag user.Flag) bool {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(FieldKeyPublicFlags), flag, err))
		}
	}

	c := uint32(flag)

	return u.publicFlags.Get()&c == c
}

func (*userImpl) appendIfSet1(mp map[FieldKey]interface{}, key FieldKey, field OptionalField) {
	if field.IsSet() {
		mp[key] = field
	}
}

func (*userImpl) appendIfSet2(mp map[FieldKey]interface{}, key FieldKey, field TriStateField) {
	if field.IsSet() {
		mp[key] = field
	}
}
