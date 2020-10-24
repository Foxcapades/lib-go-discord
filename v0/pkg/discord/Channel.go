package discord

import (
	"encoding/json"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
)

type Channel interface {
	json.Marshaler
	json.Unmarshaler
	lib.Validatable

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of this channel.
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) Channel

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of the channel.
	Type() ChannelType

	// SetType overwrites the current value of this record's `type` field.
	SetType(ChannelType) Channel

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the id of the guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use GuildIDIsSet to check if the field is present before use.
	GuildID() Snowflake

	// GuildIDIsSet returns whether this record's `guild_id` field is currently
	// present.
	GuildIDIsSet() bool

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(Snowflake) Channel

	// UnsetGuildID removes this record's `guild_id` field.
	UnsetGuildID() Channel

	// Position returns the current value of this record's `position` field.
	//
	// The `position` field contains the sorting position of the channel.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PositionIsSet to check if the field is present before use.
	Position() uint16

	// PositionIsSet returns whether this record's `position` field is currently
	// present.
	PositionIsSet() bool

	// SetPosition overwrites the current value of this record's `position` field.
	SetPosition(uint16) Channel

	// UnsetPosition removes this record's `position` field.
	UnsetPosition() Channel

	// PermissionOverwrites returns the current value of this record's
	// `permission_overwrites` field.
	//
	// The `permission_overwrites` field contains an array of explicit permission
	// overwrites for members and roles.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PermissionOverwritesIsSet to check if the field is present before use.
	PermissionOverwrites() []PermissionOverwrite

	// PermissionOverwritesIsSet returns whether this record's
	// `permission_overwrites` field is currently present.
	PermissionOverwritesIsSet() bool

	// SetPermissionOverwrites overwrites the current value of this record's
	// `permission_overwrites` field.
	SetPermissionOverwrites([]PermissionOverwrite) Channel

	// UnsetPermissionOverwrites removes this record's `permission_overwrites`
	// field.
	UnsetPermissionOverwrites() Channel

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the channel (2-100 characters).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NameIsSet to check if the field is present before use.
	Name() ChannelName

	// NameIsSet returns whether this record's `name` field is currently present.
	NameIsSet() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(ChannelName) Channel

	// UnsetName removes this record's `name` field.
	UnsetName() Channel

	// Topic returns the current value of this record's `topic` field.
	//
	// The `topic` field contains the channel topic (0-1024 characters).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use TopicIsReadable to check if the field is
	// present and non-null before use.
	Topic() ChannelTopic

	// TopicIsNull returns whether this record's `topic` field is currently null.
	TopicIsNull() bool

	// TopicIsSet returns whether this record's `topic` field is currently
	// present.
	TopicIsSet() bool

	// TopicIsReadable returns whether this record's `topic` field is currently
	// set and non-null.
	TopicIsReadable() bool

	// SetTopic overwrites the current value of this record's `topic` field.
	SetTopic(ChannelTopic) Channel

	// SetNullTopic overwrites the current value of this record's `topic` field
	// with `null`.
	SetNullTopic() Channel

	// UnsetTopic removes this record's `topic` field.
	UnsetTopic() Channel

	// NSFW returns the current value of this record's `nsfw` field.
	//
	// The `nsfw` field indicates whether the channel is nsfw.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NSFWIsSet to check if the field is present before use.
	NSFW() bool

	// NSFWIsSet returns whether this record's `nsfw` field is currently present.
	NSFWIsSet() bool

	// SetNSFW overwrites the current value of this record's `nsfw` field.
	SetNSFW(bool) Channel

	// UnsetNSFW removes this record's `nsfw` field.
	UnsetNSFW() Channel

	// LastMessageID returns the current value of this record's `last_message_id`
	// field.
	//
	// The `last_message_id` field contains the id of the last message sent in
	// this channel (may not point to an existing or valid message).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use LastMessageIDIsReadable to check if the field
	// is present and non-null before use.
	LastMessageID() Snowflake

	// LastMessageIDIsNull returns whether this record's `last_message_id` field
	// is currently null.
	LastMessageIDIsNull() bool

	// LastMessageIDIsSet returns whether this record's `last_message_id` field is
	// currently present.
	LastMessageIDIsSet() bool

	// LastMessageIDIsReadable returns whether this record's `last_message_id`
	// field is currently set and non-null.
	LastMessageIDIsReadable() bool

	// SetLastMessageID overwrites the current value of this record's
	// `last_message_id` field.
	SetLastMessageID(Snowflake) Channel

	// SetNullLastMessageID overwrites the current value of this record's
	// `last_message_id` field with `null`.
	SetNullLastMessageID() Channel

	// UnsetLastMessageID removes this record's `last_message_id` field.
	UnsetLastMessageID() Channel

	// Bitrate returns the current value of this record's `bitrate` field.
	//
	// The `bitrate` field contains the bitrate (in bits) of the voice channel.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use BitrateIsSet to check if the field is present before use.
	Bitrate() Bitrate

	// BitrateIsSet returns whether this record's `bitrate` field is currently
	// present.
	BitrateIsSet() bool

	// SetBitrate overwrites the current value of this record's `bitrate` field.
	SetBitrate(Bitrate) Channel

	// UnsetBitrate removes this record's `bitrate` field.
	UnsetBitrate() Channel

	// UserLimit returns the current value of this record's `user_limit` field.
	//
	// The `user_limit` field contains the user limit of the voice channel.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use UserLimitIsSet to check if the field is present before use.
	UserLimit() UserLimit

	// UserLimitIsSet returns whether this record's `user_limit` field is
	// currently present.
	UserLimitIsSet() bool

	// SetUserLimit overwrites the current value of this record's `user_limit`
	// field.
	SetUserLimit(UserLimit) Channel

	// UnsetUserLimit removes this record's `user_limit` field.
	UnsetUserLimit() Channel

	// RateLimitPerUser returns the current value of this record's
	// `rate_limit_per_user` field.
	//
	// The `rate_limit_per_user` field contains the amount of seconds a user has
	// to wait before sending another message (0-21600); bots, as well as users
	// with the permission manage_messages or manage_channel, are unaffected.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RateLimitPerUserIsSet to check if the field is present before use.
	RateLimitPerUser() PerUserRateLimit

	// RateLimitPerUserIsSet returns whether this record's `rate_limit_per_user`
	// field is currently present.
	RateLimitPerUserIsSet() bool

	// SetRateLimitPerUser overwrites the current value of this record's
	// `rate_limit_per_user` field.
	SetRateLimitPerUser(PerUserRateLimit) Channel

	// UnsetRateLimitPerUser removes this record's `rate_limit_per_user` field.
	UnsetRateLimitPerUser() Channel

	// Recipients returns the current value of this record's `recipients` field.
	//
	// The `recipients` field contains the recipients of the DM.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RecipientsIsSet to check if the field is present before use.
	Recipients() []User

	// RecipientsIsSet returns whether this record's `recipients` field is
	// currently present.
	RecipientsIsSet() bool

	// SetRecipients overwrites the current value of this record's `recipients`
	// field.
	SetRecipients([]User) Channel

	// UnsetRecipients removes this record's `recipients` field.
	UnsetRecipients() Channel

	// Icon returns the current value of this record's `icon` field.
	//
	// The `icon` field contains the icon hash.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use IconIsReadable to check if the field is
	// present and non-null before use.
	Icon() ImageHash

	// IconIsNull returns whether this record's `icon` field is currently null.
	IconIsNull() bool

	// IconIsSet returns whether this record's `icon` field is currently present.
	IconIsSet() bool

	// IconIsReadable returns whether this record's `icon` field is currently set
	// and non-null.
	IconIsReadable() bool

	// SetIcon overwrites the current value of this record's `icon` field.
	SetIcon(ImageHash) Channel

	// SetNullIcon overwrites the current value of this record's `icon` field
	// with `null`.
	SetNullIcon() Channel

	// UnsetIcon removes this record's `icon` field.
	UnsetIcon() Channel

	// OwnerID returns the current value of this record's `owner_id` field.
	//
	// The `owner_id` field contains the id of the DM creator.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use OwnerIDIsSet to check if the field is present before use.
	OwnerID() Snowflake

	// OwnerIDIsSet returns whether this record's `owner_id` field is currently
	// present.
	OwnerIDIsSet() bool

	// SetOwnerID overwrites the current value of this record's `owner_id` field.
	SetOwnerID(Snowflake) Channel

	// UnsetOwnerID removes this record's `owner_id` field.
	UnsetOwnerID() Channel

	// ApplicationID returns the current value of this record's `application_id`
	// field.
	//
	// The `application_id` field contains the application id of the group DM
	// creator if it is bot-created.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ApplicationIDIsSet to check if the field is present before use.
	ApplicationID() Snowflake

	// ApplicationIDIsSet returns whether this record's `application_id` field is
	// currently present.
	ApplicationIDIsSet() bool

	// SetApplicationID overwrites the current value of this record's
	// `application_id` field.
	SetApplicationID(Snowflake) Channel

	// UnsetApplicationID removes this record's `application_id` field.
	UnsetApplicationID() Channel

	// ParentID returns the current value of this record's `parent_id` field.
	//
	// The `parent_id` field contains the id of the parent category for a channel
	// (each parent category can contain up to 50 channels).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use ParentIDIsReadable to check if the field is
	// present and non-null before use.
	ParentID() Snowflake

	// ParentIDIsNull returns whether this record's `parent_id` field is currently
	// null.
	ParentIDIsNull() bool

	// ParentIDIsSet returns whether this record's `parent_id` field is currently
	// present.
	ParentIDIsSet() bool

	// ParentIDIsReadable returns whether this record's `parent_id` field is
	// currently set and non-null.
	ParentIDIsReadable() bool

	// SetParentID overwrites the current value of this record's `parent_id`
	// field.
	SetParentID(Snowflake) Channel

	// SetNullParentID overwrites the current value of this record's `parent_id`
	// field with `null`.
	SetNullParentID() Channel

	// UnsetParentID removes this record's `parent_id` field.
	UnsetParentID() Channel

	// LastPinTimestamp returns the current value of this record's
	// `last_pin_timestamp` field.
	//
	// The `last_pin_timestamp` field contains the timestamp of when the last
	// pinned message was pinned. This may be null in events such as GUILD_CREATE
	// when a message is not pinned.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use LastPinTimestampIsReadable to check if the
	// field is present and non-null before use.
	LastPinTimestamp() time.Time

	// LastPinTimestampIsNull returns whether this record's `last_pin_timestamp`
	// field is currently null.
	LastPinTimestampIsNull() bool

	// LastPinTimestampIsSet returns whether this record's `last_pin_timestamp`
	// field is currently present.
	LastPinTimestampIsSet() bool

	// LastPinTimestampIsReadable returns whether this record's
	// `last_pin_timestamp` field is currently set and non-null.
	LastPinTimestampIsReadable() bool

	// SetLastPinTimestamp overwrites the current value of this record's
	// `last_pin_timestamp` field.
	SetLastPinTimestamp(time.Time) Channel

	// SetNullLastPinTimestamp overwrites the current value of this record's
	// `last_pin_timestamp` field with `null`.
	SetNullLastPinTimestamp() Channel

	// UnsetLastPinTimestamp removes this record's `last_pin_timestamp` field.
	UnsetLastPinTimestamp() Channel
}
