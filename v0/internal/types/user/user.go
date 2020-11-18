package user

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"io"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewUser() User {
	return new(user)
}

const userBaseSize = js.BracketSize +
	js.FirstKeySize + uint32(len(serial.KeyID)) +
	js.NextKeySize + uint32(len(serial.KeyUsername)) +
	js.NextKeySize + uint32(len(serial.KeyDiscriminator)) +
	js.NextKeySize + uint32(len(serial.KeyAvatar)) +
	js.NextKeySize + uint32(len(serial.KeyBot)) +
	js.NextKeySize + uint32(len(serial.KeySystem)) +
	js.NextKeySize + uint32(len(serial.KeyMFAEnabled)) +
	js.NextKeySize + uint32(len(serial.KeyLocale)) +
	js.NextKeySize + uint32(len(serial.KeyVerified)) +
	js.NextKeySize + uint32(len(serial.KeyEmail)) +
	js.NextKeySize + uint32(len(serial.KeyFlags)) +
	js.NextKeySize + uint32(len(serial.KeyPremiumType)) +
	js.NextKeySize + uint32(len(serial.KeyPublicFlags))

type user struct {
	nullEmail bool

	id            Snowflake
	username      Username
	discriminator Discriminator
	avatar        *ImageHash
	bot           *bool
	system        *bool
	mfaEnabled    *bool
	locale        *string
	verified      *bool
	email         *string
	flags         *UserFlag
	premiumType   *UserPremiumType
	publicFlags   *UserFlag
}

func (u *user) JSONSize() uint32 {
	return userBaseSize +
		utils.NullableSize(u.id) +
		u.username.JSONSize() +
		u.discriminator.JSONSize() +
		utils.OptionalSize(u.avatar) +
		utils.OptionalBoolSize(u.bot) +
		utils.OptionalBoolSize(u.system) +
		utils.OptionalBoolSize(u.mfaEnabled) +
		utils.OptionalStringSize(u.locale) +
		utils.OptionalBoolSize(u.verified) +
		utils.TriStateStringSize(u.email, u.nullEmail) +
		utils.OptionalSize(u.flags) +
		utils.OptionalSize(u.premiumType) +
		utils.OptionalSize(u.publicFlags)
}

func (u *user) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(serial.KeyID, u.id).
		AddString(serial.KeyUsername, string(u.username)).
		AddString(serial.KeyDiscriminator, u.discriminator.String()).
		AddNullableString(serial.KeyAvatar, (*string)(u.avatar)).
		AddOptionalBool(serial.KeyBot, u.bot).
		AddOptionalBool(serial.KeySystem, u.system).
		AddOptionalBool(serial.KeyMFAEnabled, u.mfaEnabled).
		AddOptionalString(serial.KeyLocale, u.locale).
		AddOptionalBool(serial.KeyVerified, u.verified).
		AddTriStateString(serial.KeyEmail, u.email, u.nullEmail).
		AddOptionalUint32(serial.KeyFlags, (*uint32)(u.flags)).
		AddOptionalUint8(serial.KeyPremiumType, (*uint8)(u.premiumType)).
		AddOptionalUint32(serial.KeyPublicFlags, (*uint32)(u.publicFlags))
}

func (u *user) IsNil() bool {
	return u == nil
}

func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {
	switch serial.Key(key) {
	case serial.KeyID:
		u.id, err = com.DecodeSnowflake(dec)
	case serial.KeyUsername:
		err = dec.DecodeString((*string)(&u.username))
	case serial.KeyDiscriminator:
		u.discriminator, err = DecodeDiscriminator(dec)
	case serial.KeyAvatar:
		u.avatar, err = DecodeImageHash(dec)
	case serial.KeyBot:
		err = dec.BoolNull(&u.bot)
	case serial.KeySystem:
		err = dec.BoolNull(&u.system)
	case serial.KeyMFAEnabled:
		err = dec.BoolNull(&u.mfaEnabled)
	case serial.KeyLocale:
		err = dec.StringNull(&u.locale)
	case serial.KeyVerified:
		err = dec.BoolNull(&u.verified)
	case serial.KeyEmail:
		err = dec.StringNull(&u.email)
		u.nullEmail = u.email == nil
	case serial.KeyFlags:
		u.flags, err = DecodeFlags(dec)
	case serial.KeyPremiumType:
		u.premiumType, err = DecodePremiumType(dec)
	case serial.KeyPublicFlags:
		u.publicFlags, err = DecodeFlags(dec)
	}

	return
}

func (u *user) NKeys() int {
	return 0
}

func (u *user) IsValid() bool {
	return nil == u.Validate()
}

func (u *user) Validate() error {
	out := lib.NewValidationErrorSet()

	out.AppendRawKeyedError(serial.KeyID, u.id.Validate())
	out.AppendRawKeyedError(serial.KeyUsername, u.username.Validate())
	out.AppendRawKeyedError(serial.KeyDiscriminator, u.discriminator.Validate())

	if u.flags != nil {
		out.AppendRawKeyedError(serial.KeyFlags, u.flags.Validate())
	}

	if u.premiumType != nil {
		out.AppendRawKeyedError(serial.KeyPremiumType, u.premiumType.Validate())
	}

	if u.publicFlags != nil {
		out.AppendRawKeyedError(serial.KeyPublicFlags, u.publicFlags.Validate())
	}

	if out.Len() > 0 {
		return out
	}

	return nil
}

func (u *user) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, u.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeObject(u); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (u *user) UnmarshalJSON(in []byte) (err error) {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(u)
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

func (u *user) AvatarHash() ImageHash {
	if u.avatar == nil {
		panic(e.ErrGetNull)
	}

	return *u.avatar
}

func (u *user) AvatarHashIsNull() bool {
	return u.avatar == nil
}

func (u *user) SetAvatarHash(s ImageHash) User {
	u.avatar = &s
	return u
}

func (u *user) SetNullAvatarHash() User {
	u.avatar = nil
	return u
}

func (u *user) BotFlag() bool {
	if u.bot == nil {
		panic(e.ErrGetUnset)
	}

	return *u.bot
}

func (u *user) BotFlagIsSet() bool {
	return u.bot != nil
}

func (u *user) SetBotFlag(b bool) User {
	u.bot = &b
	return u
}

func (u *user) UnsetBotFlag() User {
	u.bot = nil
	return u
}

func (u *user) SystemFlag() bool {
	if u.system == nil {
		panic(e.ErrGetUnset)
	}

	return *u.system
}

func (u *user) SystemFlagIsSet() bool {
	return u.system != nil
}

func (u *user) SetSystemFlag(b bool) User {
	u.system = &b
	return u
}

func (u *user) UnsetSystemFlag() User {
	u.system = nil
	return u
}

func (u *user) MFAEnabled() bool {
	if u.mfaEnabled == nil {
		panic(e.ErrGetUnset)
	}

	return *u.mfaEnabled
}

func (u *user) MFAEnabledIsSet() bool {
	return u.mfaEnabled != nil
}

func (u *user) SetMFAEnabled(b bool) User {
	u.mfaEnabled = &b
	return u
}

func (u *user) UnsetMFAEnabled() User {
	u.mfaEnabled = nil
	return u
}

func (u *user) Locale() string {
	if u.locale == nil {
		panic(e.ErrGetUnset)
	}

	return *u.locale
}

func (u *user) LocaleIsSet() bool {
	return u.locale != nil
}

func (u *user) SetLocale(s string) User {
	u.locale = &s
	return u
}

func (u *user) UnsetLocale() User {
	u.locale = nil
	return u
}

func (u *user) VerifiedFlag() bool {
	if u.verified == nil {
		panic(e.ErrGetUnset)
	}

	return *u.verified
}

func (u *user) VerifiedFlagIsSet() bool {
	return u.verified != nil
}

func (u *user) SetVerifiedFlag(b bool) User {
	u.verified = &b
	return u
}

func (u *user) UnsetVerifiedFlag() User {
	u.verified = nil

	return u
}

func (u *user) Email() string {
	if u.email == nil {
		if u.nullEmail {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *u.email
}

func (u *user) EmailIsSet() bool {
	return u.email != nil
}

func (u *user) EmailIsNull() bool {
	return u.email == nil && u.nullEmail
}

func (u *user) SetEmail(s string) User {
	u.email = &s
	u.nullEmail = false

	return u
}

func (u *user) UnsetEmail() User {
	u.email = nil
	u.nullEmail = true

	return u
}

func (u *user) SetNullEmail() User {
	u.email = nil
	u.nullEmail = true

	return u
}

func (u *user) Flags() UserFlag {
	return *u.flags
}

func (u *user) FlagsIsSet() bool {
	return u.flags != nil
}

func (u *user) SetFlags(flag UserFlag) User {
	u.flags = &flag
	return u
}

func (u *user) UnsetFlags() User {
	u.flags = nil
	return u
}

func (u *user) AddFlag(flag UserFlag) User {
	if u.flags != nil {
		*u.flags |= flag
	} else {
		u.flags = &flag
	}

	return u
}

func (u *user) RemoveFlag(flag UserFlag) User {
	if u.flags != nil {
		*u.flags &= ^flag
	}

	return u
}

func (u *user) FlagsContains(flag UserFlag) bool {
	if u.flags != nil {
		return *u.flags&flag == flag
	}

	return false
}

func (u *user) PremiumType() UserPremiumType {
	return *u.premiumType
}

func (u *user) PremiumTypeIsSet() bool {
	return u.premiumType != nil
}

func (u *user) SetPremiumType(val UserPremiumType) User {
	u.premiumType = &val
	return u
}

func (u *user) UnsetPremiumType() User {
	u.premiumType = nil
	return u
}

func (u *user) PublicFlags() UserFlag {
	return *u.publicFlags
}

func (u *user) PublicFlagsIsSet() bool {
	return u.publicFlags != nil
}

func (u *user) SetPublicFlags(flag UserFlag) User {
	u.publicFlags = &flag
	return u
}

func (u *user) UnsetPublicFlags() User {
	u.publicFlags = nil
	return u
}

func (u *user) AddPublicFlag(flag UserFlag) User {
	if u.publicFlags != nil {
		*u.publicFlags |= flag
	} else {
		u.publicFlags = &flag
	}

	return u
}

func (u *user) RemovePublicFlag(flag UserFlag) User {
	if u.publicFlags != nil {
		*u.publicFlags &= ^flag
	}

	return u
}

func (u *user) PublicFlagsContains(flag UserFlag) bool {
	if u.publicFlags == nil {
		return false
	}

	return *u.publicFlags&flag == flag
}

func (u *user) ToJSONBytes() []byte {
	out := new(bytes.Buffer)
	out.Grow(int(u.JSONSize()))
	out.Reset()

	out.WriteByte(js.OpenObject)

	_ = serial.KeyID.AppendJSONBytes(out)
	out.WriteByte(js.PairSeparator)
	_ = u.id.AppendJSONBytes(out)
	out.WriteByte(js.FieldDivider)

	_ = serial.KeyUsername.AppendJSONBytes(out)
	out.WriteByte(js.PairSeparator)
	_ = u.username.AppendJSONBytes(out)
	out.WriteByte(js.FieldDivider)

	_ = serial.KeyDiscriminator.AppendJSONBytes(out)
	out.WriteByte(js.PairSeparator)
	_ = u.discriminator.AppendJSONBytes(out)
	out.WriteByte(js.FieldDivider)

	if u.bot != nil {
		_ = serial.KeyBot.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalBool(*u.bot))
		out.WriteByte(js.FieldDivider)
	}

	if u.system != nil {
		_ = serial.KeySystem.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalBool(*u.system))
		out.WriteByte(js.FieldDivider)
	}

	if u.mfaEnabled != nil {
		_ = serial.KeyMFAEnabled.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalBool(*u.mfaEnabled))
		out.WriteByte(js.FieldDivider)
	}

	if u.locale != nil {
		_ = serial.KeyLocale.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalString(*u.locale))
		out.WriteByte(js.FieldDivider)
	}

	if u.verified != nil {
		_ = serial.KeyVerified.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalBool(*u.verified))
		out.WriteByte(js.FieldDivider)
	}

	if u.email != nil {
		_ = serial.KeyEmail.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(utils.MarshalString(*u.email))
		out.WriteByte(js.FieldDivider)
	} else if u.nullEmail {
		_ = serial.KeyEmail.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		out.Write(js.NullBytesBuf)
		out.WriteByte(js.FieldDivider)
	}

	if u.flags != nil {
		_ = serial.KeyFlags.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		_ = u.flags.AppendJSONBytes(out)
		out.WriteByte(js.FieldDivider)
	}

	if u.premiumType != nil {
		_ = serial.KeyPremiumType.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		_ = u.premiumType.AppendJSONBytes(out)
		out.WriteByte(js.FieldDivider)
	}

	if u.publicFlags != nil {
		_ = serial.KeyPublicFlags.AppendJSONBytes(out)
		out.WriteByte(js.PairSeparator)
		_ = u.publicFlags.AppendJSONBytes(out)
		out.WriteByte(js.FieldDivider)
	}

	_ = serial.KeyAvatar.AppendJSONBytes(out)
	out.WriteByte(js.PairSeparator)
	_ = u.avatar.AppendJSONBytes(out)

	out.WriteByte(js.CloseObject)

	return out.Bytes()
}

func (u *user) AppendJSONBytes(writer io.Writer) error {
	_, err := writer.Write(u.ToJSONBytes())
	return err
}
