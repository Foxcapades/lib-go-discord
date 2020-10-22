package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type Webhook interface {
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
	Avatar() string

	// AvatarIsNull returns whether this record's `avatar` field is currently
	// null.
	AvatarIsNull() bool

	// SetAvatar overwrites the current value of this record's `avatar` field.
	SetAvatar(string) Webhook

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
