package user

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewUser() User {
	return new(user)
}

type user struct {
	id            Snowflake
	username      Username
	discriminator Discriminator
	avatar        types.NullableString
	bot           types.OptionalBool
	system        types.OptionalBool
	mfaEnabled    types.OptionalBool
	locale        types.OptionalString
	verified      types.OptionalBool
	email         types.TriStateString
	flags         types.OptionalUint32
	premiumType   types.OptionalUint8
	publicFlags   types.OptionalUint32
}

func (u *user) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (u *user) IsNil() bool {
	panic("implement me")
}

func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (u *user) NKeys() int {
	panic("implement me")
}

func (u *user) IsValid() bool {
	panic("implement me")
}

func (u *user) Validate() error {
	panic("implement me")
}

func (u *user) MarshalJSON() ([]byte, error) {
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

func (u *user) UnmarshalJSON(bytes []byte) (err error) {
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

func (u *user) ID() Snowflake {
	return u.id
}

func (u *user) SetID(id Snowflake) User {
	u.id = id

	return u
}

func (u *user) Username() string {
	return string(u.username)
}

func (u *user) SetUsername(name Username) User {
	u.username = name

	return u
}

func (u *user) Discriminator() Discriminator {
	return u.discriminator
}

func (u *user) SetDiscriminator(code Discriminator) User {
	u.discriminator = code

	return u
}

func (u *user) AvatarHash() string {
	return u.avatar.Get()
}

func (u *user) AvatarHashIsNull() bool {
	return u.avatar.IsNull()
}

func (u *user) SetAvatarHash(s string) User {
	u.avatar.Set(s)

	return u
}

func (u *user) SetNullAvatarHash() User {
	u.avatar.SetNull()

	return u
}

func (u *user) BotFlag() bool {
	return u.bot.Get()
}

func (u *user) BotFlagIsSet() bool {
	return u.bot.IsSet()
}

func (u *user) SetBotFlag(b bool) User {
	u.bot.Set(b)

	return u
}

func (u *user) UnsetBotFlag() User {
	u.bot.Unset()

	return u
}

func (u *user) SystemFlag() bool {
	return u.system.Get()
}

func (u *user) SystemFlagIsSet() bool {
	return u.system.IsSet()
}

func (u *user) SetSystemFlag(b bool) User {
	u.system.Set(b)

	return u
}

func (u *user) UnsetSystemFlag() User {
	u.system.Unset()

	return u
}

func (u *user) MFAEnabled() bool {
	return u.mfaEnabled.Get()
}

func (u *user) MFAEnabledIsSet() bool {
	return u.mfaEnabled.IsSet()
}

func (u *user) SetMFAEnabled(b bool) User {
	u.mfaEnabled.Set(b)

	return u
}

func (u *user) UnsetMFAEnabled() User {
	u.mfaEnabled.Unset()

	return u
}

func (u *user) Locale() string {
	return u.locale.Get()
}

func (u *user) LocaleIsSet() bool {
	return u.locale.IsSet()
}

func (u *user) SetLocale(s string) User {
	u.locale.Set(s)

	return u
}

func (u *user) UnsetLocale() User {
	u.locale.Unset()

	return u
}

func (u *user) VerifiedFlag() bool {
	return u.verified.Get()
}

func (u *user) VerifiedFlagIsSet() bool {
	return u.verified.IsSet()
}

func (u *user) SetVerifiedFlag(b bool) User {
	u.verified.Set(b)

	return u
}

func (u *user) UnsetVerifiedFlag() User {
	u.verified.Unset()

	return u
}

func (u *user) Email() string {
	return u.email.Get()
}

func (u *user) EmailIsSet() bool {
	return u.email.IsSet()
}

func (u *user) EmailIsNull() bool {
	return u.email.IsNull()
}

func (u *user) SetEmail(s string) User {
	u.email.Set(s)

	return u
}

func (u *user) UnsetEmail() User {
	u.email.Unset()

	return u
}

func (u *user) SetNullEmail() User {
	u.email.SetNull()

	return u
}

func (u *user) Flags() UserFlag {
	return UserFlag(u.flags.Get())
}

func (u *user) FlagsIsSet() bool {
	return u.flags.IsSet()
}

func (u *user) SetFlags(flag UserFlag) User {
	u.flags.Set(uint32(flag))

	return u
}

func (u *user) UnsetFlags() User {
	u.flags.Unset()

	return u
}

func (u *user) AddFlag(flag UserFlag) User {
	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() | uint32(flag))
	} else {
		u.flags.Set(uint32(flag))
	}

	return u
}

func (u *user) RemoveFlag(flag UserFlag) User {
	if u.flags.IsSet() {
		u.flags.Set(u.flags.Get() & ^uint32(flag))
	}

	return u
}

func (u *user) FlagsContains(flag UserFlag) bool {
	if u.flags.IsSet() {
		return u.flags.Get()|uint32(flag) == uint32(flag)
	}

	return false
}

func (u *user) PremiumType() UserPremiumType {
	return UserPremiumType(u.premiumType.Get())
}

func (u *user) PremiumTypeIsSet() bool {
	return u.premiumType.IsSet()
}

func (u *user) SetPremiumType(val UserPremiumType) User {
	u.premiumType.Set(uint8(val))

	return u
}

func (u *user) UnsetPremiumType() User {
	u.premiumType.Unset()

	return u
}

func (u *user) PublicFlags() UserFlag {
	return UserFlag(u.publicFlags.Get())
}

func (u *user) PublicFlagsIsSet() bool {
	return u.publicFlags.IsSet()
}

func (u *user) SetPublicFlags(flag UserFlag) User {
	u.publicFlags.Set(uint32(flag))
	return u
}

func (u *user) UnsetPublicFlags() User {
	u.publicFlags.Unset()
	return u
}

func (u *user) AddPublicFlag(flag UserFlag) User {
	if u.publicFlags.IsSet() {
		u.publicFlags.Set(uint32(flag) | u.publicFlags.Get())
	} else {
		u.publicFlags.Set(uint32(flag))
	}

	return u
}

func (u *user) RemovePublicFlag(flag UserFlag) User {
	if u.publicFlags.IsSet() {
		u.publicFlags.Set(u.publicFlags.Get() & ^uint32(flag))
	}

	return u
}

func (u *user) PublicFlagsContains(flag UserFlag) bool {
	c := uint32(flag)

	return u.publicFlags.Get()&c == c
}

func (*user) appendIfSet1(mp map[serial.Key]interface{}, key serial.Key, field types.OptionalField) {
	if field.IsSet() {
		mp[key] = field
	}
}

func (*user) appendIfSet2(mp map[serial.Key]interface{}, key serial.Key, field types.TriStateField) {
	if field.IsSet() {
		mp[key] = field
	}
}
