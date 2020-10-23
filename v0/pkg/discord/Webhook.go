package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type Webhook interface {
	json.Marshaler
	json.Unmarshaler

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the webhook.
	ID() dlib.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(dlib.Snowflake) Webhook

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of the webhook.
	Type() WebhookType

	// SetType overwrites the current value of this record's `type` field.
	SetType(WebhookType) Webhook

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the guild id this webhook is for.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use GuildIDIsSet to check if the field is present before use.
	GuildID() dlib.Snowflake

	// GuildIDIsSet returns whether this record's `guild_id` field is currently
	// present.
	GuildIDIsSet() bool

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(dlib.Snowflake) Webhook

	// UnsetGuildID removes this record's `guild_id` field.
	UnsetGuildID() Webhook

	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the channel id this webhook is for.
	ChannelID() dlib.Snowflake

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(dlib.Snowflake) Webhook

	// User returns the current value of this record's `user` field.
	//
	// The `user` field contains the user this webhook was created by (not
	// returned when getting a webhook with its token).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use UserIsSet to check if the field is present before use.
	User() User

	// UserIsSet returns whether this record's `user` field is currently present.
	UserIsSet() bool

	// SetUser overwrites the current value of this record's `user` field.
	SetUser(User) Webhook

	// UnsetUser removes this record's `user` field.
	UnsetUser() Webhook

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the default name of the webhook.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use NameIsNull to check if the field is null before use.
	Name() string

	// NameIsNull returns whether this record's `name` field is currently null.
	NameIsNull() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) Webhook

	// SetNullName overwrites the current value of this record's `name` field
	// with `null`.
	SetNullName() Webhook

	// Avatar returns the current value of this record's `avatar` field.
	//
	// The `avatar` field contains the default avatar of the webhook.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use AvatarIsNull to check if the field is null before use.
	Avatar() comm.ImageHash

	// AvatarIsNull returns whether this record's `avatar` field is currently
	// null.
	AvatarIsNull() bool

	// SetAvatar overwrites the current value of this record's `avatar` field.
	SetAvatar(comm.ImageHash) Webhook

	// SetNullAvatar overwrites the current value of this record's `avatar` field
	// with `null`.
	SetNullAvatar() Webhook

	// Token returns the current value of this record's `token` field.
	//
	// The `token` field contains the secure token of the webhook (returned for
	// Incoming Webhooks).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TokenIsSet to check if the field is present before use.
	Token() string

	// TokenIsSet returns whether this record's `token` field is currently
	// present.
	TokenIsSet() bool

	// SetToken overwrites the current value of this record's `token` field.
	SetToken(string) Webhook

	// UnsetToken removes this record's `token` field.
	UnsetToken() Webhook

	// ApplicationID returns the current value of this record's `application_id`
	// field.
	//
	// The `application_id` field contains the bot/OAuth2 application that created
	// this webhook.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use ApplicationIDIsNull to check if the field is null before use.
	ApplicationID() dlib.Snowflake

	// ApplicationIDIsNull returns whether this record's `application_id` field is
	// currently null.
	ApplicationIDIsNull() bool

	// SetApplicationID overwrites the current value of this record's
	// `application_id` field.
	SetApplicationID(dlib.Snowflake) Webhook

	// SetNullApplicationID overwrites the current value of this record's
	// `application_id` field with `null`.
	SetNullApplicationID() Webhook
}

type webhookImpl struct {
	id        dlib.Snowflake
	kind      WebhookType
	guildID   dlib.OptionalSnowflake
	channelID dlib.Snowflake
	user      OptionalUser
	name      dlib.NullableString
	avatar    comm.NullableImageHash
	token     dlib.OptionalString
	appID     dlib.NullableSnowflake
}

func (w *webhookImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(outBuilder, 9).
		AppendRaw(FieldKeyID, w.id).
		AppendRaw(FieldKeyType, w.kind).
		AppendOptional(FieldKeyGuildID, &w.guildID).
		AppendRaw(FieldKeyChannelID, w.channelID).
		AppendOptional(FieldKeyUser, &w.user).
		AppendNullable(FieldKeyName, &w.name).
		AppendNullable(FieldKeyAvatar, &w.avatar).
		AppendOptional(FieldKeyToken, &w.token).
		AppendNullable(FieldKeyApplicationID, &w.appID))
}

func (w *webhookImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]json.RawMessage)
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	unm := map[FieldKey]json.Unmarshaler{
		FieldKeyID:            &w.id,
		FieldKeyGuildID:       &w.guildID,
		FieldKeyChannelID:     &w.channelID,
		FieldKeyUser:          &w.user,
		FieldKeyName:          &w.name,
		FieldKeyAvatar:        &w.avatar,
		FieldKeyToken:         &w.token,
		FieldKeyApplicationID: &w.appID,
	}

	for k, u := range unm {
		if err := u.UnmarshalJSON(tmp[k]); err != nil {
			return err
		}
	}

	var kind uint8
	if err := json.Unmarshal(tmp[FieldKeyType], &kind); err != nil {
		return err
	}
	w.kind = WebhookType(kind)

	return nil
}

func (w *webhookImpl) ID() dlib.Snowflake {
	return w.id
}

func (w *webhookImpl) SetID(id dlib.Snowflake) Webhook {
	w.id = id
	return w
}

func (w *webhookImpl) Type() WebhookType {
	return w.kind
}

func (w *webhookImpl) SetType(kind WebhookType) Webhook {
	w.kind = kind
	return w
}

func (w *webhookImpl) GuildID() dlib.Snowflake {
	return w.guildID.Get()
}

func (w *webhookImpl) GuildIDIsSet() bool {
	return w.guildID.IsSet()
}

func (w *webhookImpl) SetGuildID(id dlib.Snowflake) Webhook {
	w.guildID.Set(id)
	return w
}

func (w *webhookImpl) UnsetGuildID() Webhook {
	w.guildID.Unset()
	return w
}

func (w *webhookImpl) ChannelID() dlib.Snowflake {
	return w.channelID
}

func (w *webhookImpl) SetChannelID(id dlib.Snowflake) Webhook {
	w.channelID = id
	return w
}

func (w *webhookImpl) User() User {
	return w.user.Get()
}

func (w *webhookImpl) UserIsSet() bool {
	return w.user.IsSet()
}

func (w *webhookImpl) SetUser(user User) Webhook {
	w.user.Set(user)
	return w
}

func (w *webhookImpl) UnsetUser() Webhook {
	w.user.Unset()
	return w
}

func (w *webhookImpl) Name() string {
	return w.name.Get()
}

func (w *webhookImpl) NameIsNull() bool {
	return w.name.IsNull()
}

func (w *webhookImpl) SetName(name string) Webhook {
	w.name.Set(name)
	return w
}

func (w *webhookImpl) SetNullName() Webhook {
	w.name.SetNull()
	return w
}

func (w *webhookImpl) Avatar() comm.ImageHash {
	return w.avatar.Get()
}

func (w *webhookImpl) AvatarIsNull() bool {
	return w.avatar.IsNull()
}

func (w *webhookImpl) SetAvatar(ava comm.ImageHash) Webhook {
	w.avatar.Set(ava)
	return w
}

func (w *webhookImpl) SetNullAvatar() Webhook {
	w.avatar.SetNull()
	return w
}

func (w *webhookImpl) Token() string {
	return w.token.Get()
}

func (w *webhookImpl) TokenIsSet() bool {
	return w.token.IsSet()
}

func (w *webhookImpl) SetToken(tok string) Webhook {
	w.token.Set(tok)
	return w
}

func (w *webhookImpl) UnsetToken() Webhook {
	w.token.Unset()
	return w
}

func (w *webhookImpl) ApplicationID() dlib.Snowflake {
	return w.appID.Get()
}

func (w *webhookImpl) ApplicationIDIsNull() bool {
	return w.appID.IsNull()
}

func (w *webhookImpl) SetApplicationID(id dlib.Snowflake) Webhook {
	w.appID.Set(id)
	return w
}

func (w *webhookImpl) SetNullApplicationID() Webhook {
	w.appID.SetNull()
	return w
}
