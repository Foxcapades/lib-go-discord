package types

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/derr"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewUserImpl(val bool) *UserImpl {
	return &UserImpl{
		validate: val,
		id:       NewSnowflakeImpl(val),
	}
}

type UserImpl struct {
	validate bool

	id            Snowflake
	username      Username
	discriminator Discriminator
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

func (u *UserImpl) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{}, 10)

	if u.id.RawValue() == 0 {
		return nil, ErrNoId
	} else {
		out[serial.KeyID] = u.id
	}

	if u.username == "" {
		return nil, ErrNoUsername
	} else {
		out[serial.KeyUsername] = u.username
	}

	if u.discriminator == 0 {
		return nil, ErrNoDiscriminator
	} else {
		out[serial.KeyDiscriminator] = u.discriminator
	}

	out[serial.KeyAvatar] = u.avatar
	u.appendIfSet1(out, serial.KeyBot, &u.bot)
	u.appendIfSet1(out, serial.KeySystem, &u.system)
	u.appendIfSet1(out, serial.KeyMFAEnabled, &u.mfaEnabled)
	u.appendIfSet1(out, serial.KeyLocale, &u.locale)
	u.appendIfSet1(out, serial.KeyVerified, &u.verified)
	u.appendIfSet2(out, serial.KeyEmail, &u.email)
	u.appendIfSet1(out, serial.KeyFlags, &u.flags)
	u.appendIfSet1(out, serial.KeyPremiumType, &u.premiumType)
	u.appendIfSet1(out, serial.KeyPublicFlags, &u.publicFlags)

	return json.Marshal(out)
}

func (u *UserImpl) UnmarshalJSON(bytes []byte) (err error) {
	in := make(map[string]json.RawMessage, 10)

	if err = json.Unmarshal(bytes, &in); err != nil {
		return
	}

	multis := map[serial.Key]json.Unmarshaler{
		serial.KeyAvatar:      &u.avatar,
		serial.KeyBot:         &u.bot,
		serial.KeySystem:      &u.system,
		serial.KeyMFAEnabled:  &u.mfaEnabled,
		serial.KeyLocale:      &u.locale,
		serial.KeyVerified:    &u.verified,
		serial.KeyEmail:       &u.email,
		serial.KeyFlags:       &u.flags,
		serial.KeyPremiumType: &u.premiumType,
		serial.KeyPublicFlags: &u.publicFlags,
	}

	singles := map[serial.Key]interface{}{
		serial.KeyID:            &u.id,
		serial.KeyUsername:      &u.username,
		serial.KeyDiscriminator: &u.discriminator,
	}

	for k, v := range in {
		key := serial.Key(k)

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

func (u *UserImpl) ID() Snowflake {
	return u.id
}

func (u *UserImpl) SetID(id Snowflake) User {
	u.id = id

	return u
}

func (u *UserImpl) Username() string {
	return string(u.username)
}

func (u *UserImpl) SetUsername(name Username) User {
	if u.validate {
		if err := name.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyUsername), name, err))
		}
	}

	u.username = name

	return u
}

func (u *UserImpl) Discriminator() Discriminator {
	return u.discriminator
}

func (u *UserImpl) SetDiscriminator(code Discriminator) User {
	if u.validate {
		if err := code.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyDiscriminator), code, err))
		}
	}

	u.discriminator = code

	return u
}

func (u *UserImpl) AvatarHash() string {
	return u.avatar.Get()
}

func (u *UserImpl) AvatarHashIsNull() bool {
	return u.avatar.IsNull()
}

func (u *UserImpl) SetAvatarHash(s string) User {
	u.avatar.Set(s)

	return u
}

func (u *UserImpl) SetNullAvatarHash() User {
	u.avatar.SetNull()

	return u
}

func (u *UserImpl) BotFlag() bool {
	return u.bot.Get()
}

func (u *UserImpl) BotFlagIsSet() bool {
	return u.bot.IsSet()
}

func (u *UserImpl) SetBotFlag(b bool) User {
	u.bot.Set(b)

	return u
}

func (u *UserImpl) UnsetBotFlag() User {
	u.bot.Unset()

	return u
}

func (u *UserImpl) SystemFlag() bool {
	return u.system.Get()
}

func (u *UserImpl) SystemFlagIsSet() bool {
	return u.system.IsSet()
}

func (u *UserImpl) SetSystemFlag(b bool) User {
	u.system.Set(b)

	return u
}

func (u *UserImpl) UnsetSystemFlag() User {
	u.system.Unset()

	return u
}

func (u *UserImpl) MFAEnabled() bool {
	return u.mfaEnabled.Get()
}

func (u *UserImpl) MFAEnabledIsSet() bool {
	return u.mfaEnabled.IsSet()
}

func (u *UserImpl) SetMFAEnabled(b bool) User {
	u.mfaEnabled.Set(b)

	return u
}

func (u *UserImpl) UnsetMFAEnabled() User {
	u.mfaEnabled.Unset()

	return u
}

func (u *UserImpl) Locale() string {
	return u.locale.Get()
}

func (u *UserImpl) LocaleIsSet() bool {
	return u.locale.IsSet()
}

func (u *UserImpl) SetLocale(s string) User {
	u.locale.Set(s)

	return u
}

func (u *UserImpl) UnsetLocale() User {
	u.locale.Unset()

	return u
}

func (u *UserImpl) VerifiedFlag() bool {
	return u.verified.Get()
}

func (u *UserImpl) VerifiedFlagIsSet() bool {
	return u.verified.IsSet()
}

func (u *UserImpl) SetVerifiedFlag(b bool) User {
	u.verified.Set(b)

	return u
}

func (u *UserImpl) UnsetVerifiedFlag() User {
	u.verified.Unset()

	return u
}

func (u *UserImpl) Email() string {
	return u.email.Get()
}

func (u *UserImpl) EmailIsSet() bool {
	return u.email.IsSet()
}

func (u *UserImpl) EmailIsNull() bool {
	return u.email.IsNull()
}

func (u *UserImpl) SetEmail(s string) User {
	u.email.Set(s)

	return u
}

func (u *UserImpl) UnsetEmail() User {
	u.email.Unset()

	return u
}

func (u *UserImpl) SetNullEmail() User {
	u.email.SetNull()

	return u
}

func (u *UserImpl) Flags() UserFlag {
	return UserFlag(u.flags.Get())
}

func (u *UserImpl) FlagsIsSet() bool {
	return u.flags.IsSet()
}

func (u *UserImpl) SetFlags(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyFlags), flag, err))
		}
	}

	u.flags.Set(uint32(flag))

	return u
}

func (u *UserImpl) UnsetFlags() User {
	u.flags.Unset()

	return u
}

func (u *UserImpl) AddFlag(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() | uint32(flag))
	} else {
		u.flags.Set(uint32(flag))
	}

	return u
}

func (u *UserImpl) RemoveFlag(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() & ^uint32(flag))
	}

	return u
}

func (u *UserImpl) FlagsContains(flag UserFlag) bool {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyFlags), flag, err))
		}
	}

	if u.flags.IsSet() {
		return u.flags.Get()|uint32(flag) == uint32(flag)
	}

	return false
}

func (u *UserImpl) PremiumType() UserPremiumType {
	return UserPremiumType(u.premiumType.Get())
}

func (u *UserImpl) PremiumTypeIsSet() bool {
	return u.premiumType.IsSet()
}

func (u *UserImpl) SetPremiumType(val UserPremiumType) User {
	if u.validate {
		if err := val.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyPremiumType), val, err))
		}
	}

	u.premiumType.Set(uint8(val))

	return u
}

func (u *UserImpl) UnsetPremiumType() User {
	u.premiumType.Unset()

	return u
}

func (u *UserImpl) PublicFlags() UserFlag {
	return UserFlag(u.publicFlags.Get())
}

func (u *UserImpl) PublicFlagsIsSet() bool {
	return u.publicFlags.IsSet()
}

func (u *UserImpl) SetPublicFlags(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyPublicFlags), flag, err))
		}
	}

	u.publicFlags.Set(uint32(flag))
	return u
}

func (u *UserImpl) UnsetPublicFlags() User {
	u.publicFlags.Unset()
	return u
}

func (u *UserImpl) AddPublicFlag(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyPublicFlags), flag, err))
		}
	}

	if u.publicFlags.IsSet() {
		u.publicFlags.Set(uint32(flag) | u.publicFlags.Get())
	} else {
		u.publicFlags.Set(uint32(flag))
	}

	return u
}

func (u *UserImpl) RemovePublicFlag(flag UserFlag) User {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyPublicFlags), flag, err))
		}
	}

	if u.publicFlags.IsSet() {
		u.publicFlags.Set(u.publicFlags.Get() & ^uint32(flag))
	}

	return u
}

func (u *UserImpl) PublicFlagsContains(flag UserFlag) bool {
	if u.validate {
		if err := flag.Validate(); err != nil {
			panic(derr.NewInternalValidationError(string(serial.KeyPublicFlags), flag, err))
		}
	}

	c := uint32(flag)

	return u.publicFlags.Get()&c == c
}

func (*UserImpl) appendIfSet1(mp map[serial.Key]interface{}, key serial.Key, field OptionalField) {
	if field.IsSet() {
		mp[key] = field
	}
}

func (*UserImpl) appendIfSet2(mp map[serial.Key]interface{}, key serial.Key, field TriStateField) {
	if field.IsSet() {
		mp[key] = field
	}
}
