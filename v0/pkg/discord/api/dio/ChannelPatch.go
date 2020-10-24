package dio

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type ChannelPatch interface {
	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the 2-100 character channel name.
	//
	// Applies to channel type(s): All
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NameIsSet to check if the field is present before use.
	Name() discord.ChannelName

	// NameIsSet returns whether this record's `name` field is currently present.
	NameIsSet() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(discord.ChannelName) ChannelPatch

	// UnsetName removes this record's `name` field.
	UnsetName() ChannelPatch

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the the type of channel; only conversion between
	// text and news is supported and only in guilds with the "NEWS" feature.
	//
	// Applies to channel type(s): Text, News
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TypeIsSet to check if the field is present before use.
	Type() discord.ChannelType

	// TypeIsSet returns whether this record's `type` field is currently present.
	TypeIsSet() bool

	// SetType overwrites the current value of this record's `type` field.
	SetType(discord.ChannelType) ChannelPatch

	// UnsetType removes this record's `type` field.
	UnsetType() ChannelPatch

	// Position returns the current value of this record's `position` field.
	//
	// The `position` field contains the position of the channel in the left-hand
	// listing.
	//
	// Applies to channel type(s): All
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use PositionIsReadable to check if the field is
	// present and non-null before use.
	Position() uint16

	// PositionIsNull returns whether this record's `position` field is currently
	// null.
	PositionIsNull() bool

	// PositionIsSet returns whether this record's `position` field is currently
	// present.
	PositionIsSet() bool

	// PositionIsReadable returns whether this record's `position` field is
	// currently set and non-null.
	PositionIsReadable() bool

	// SetPosition overwrites the current value of this record's `position` field.
	SetPosition(uint16) ChannelPatch

	// SetNullPosition overwrites the current value of this record's `position`
	// field with `null`.
	SetNullPosition() ChannelPatch

	// UnsetPosition removes this record's `position` field.
	UnsetPosition() ChannelPatch

	// Topic returns the current value of this record's `topic` field.
	//
	// The `topic` field contains the 0-1024 character channel topic.
	//
	// Applies to channel type(s): Text, News
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use TopicIsReadable to check if the field is
	// present and non-null before use.
	Topic() discord.ChannelTopic

	// TopicIsNull returns whether this record's `topic` field is currently null.
	TopicIsNull() bool

	// TopicIsSet returns whether this record's `topic` field is currently
	// present.
	TopicIsSet() bool

	// TopicIsReadable returns whether this record's `topic` field is currently
	// set and non-null.
	TopicIsReadable() bool

	// SetTopic overwrites the current value of this record's `topic` field.
	SetTopic(discord.ChannelTopic) ChannelPatch

	// SetNullTopic overwrites the current value of this record's `topic` field
	// with `null`.
	SetNullTopic() ChannelPatch

	// UnsetTopic removes this record's `topic` field.
	UnsetTopic() ChannelPatch

	// NSFW returns the current value of this record's `nsfw` field.
	//
	// The `nsfw` field indicates whether the channel is nsfw.
	//
	// Applies to channel type(s): Text, News, Store
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use NSFWIsReadable to check if the field is
	// present and non-null before use.
	NSFW() bool

	// NSFWIsNull returns whether this record's `nsfw` field is currently null.
	NSFWIsNull() bool

	// NSFWIsSet returns whether this record's `nsfw` field is currently present.
	NSFWIsSet() bool

	// NSFWIsReadable returns whether this record's `nsfw` field is currently set
	// and non-null.
	NSFWIsReadable() bool

	// SetNSFW overwrites the current value of this record's `nsfw` field.
	SetNSFW(bool) ChannelPatch

	// SetNullNSFW overwrites the current value of this record's `nsfw` field
	// with `null`.
	SetNullNSFW() ChannelPatch

	// UnsetNSFW removes this record's `nsfw` field.
	UnsetNSFW() ChannelPatch

	// RateLimitPerUser returns the current value of this record's
	// `rate_limit_per_user` field.
	//
	// The `rate_limit_per_user` field contains the amount of seconds a user has
	// to wait before sending another message (0-21600); bots, as well as users
	// with the permission manage_messages or manage_channel, are unaffected.
	//
	// Applies to channel type(s): Text
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use RateLimitPerUserIsReadable to check if the
	// field is present and non-null before use.
	RateLimitPerUser() discord.PerUserRateLimit

	// RateLimitPerUserIsNull returns whether this record's `rate_limit_per_user`
	// field is currently null.
	RateLimitPerUserIsNull() bool

	// RateLimitPerUserIsSet returns whether this record's `rate_limit_per_user`
	// field is currently present.
	RateLimitPerUserIsSet() bool

	// RateLimitPerUserIsReadable returns whether this record's
	// `rate_limit_per_user` field is currently set and non-null.
	RateLimitPerUserIsReadable() bool

	// SetRateLimitPerUser overwrites the current value of this record's
	// `rate_limit_per_user` field.
	SetRateLimitPerUser(discord.PerUserRateLimit) ChannelPatch

	// SetNullRateLimitPerUser overwrites the current value of this record's
	// `rate_limit_per_user` field with `null`.
	SetNullRateLimitPerUser() ChannelPatch

	// UnsetRateLimitPerUser removes this record's `rate_limit_per_user` field.
	UnsetRateLimitPerUser() ChannelPatch

	// PermissionOverwrites returns the current value of this record's
	// `permission_overwrites` field.
	//
	// The `permission_overwrites` field contains the channel or category-specific
	// permissions.
	//
	// Applies to channel type(s): All
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use PermissionOverwritesIsReadable to check if the
	// field is present and non-null before use.
	PermissionOverwrites() []discord.PermissionOverwrite

	// PermissionOverwritesIsNull returns whether this record's
	// `permission_overwrites` field is currently null.
	PermissionOverwritesIsNull() bool

	// PermissionOverwritesIsSet returns whether this record's
	// `permission_overwrites` field is currently present.
	PermissionOverwritesIsSet() bool

	// PermissionOverwritesIsReadable returns whether this record's
	// `permission_overwrites` field is currently set and non-null.
	PermissionOverwritesIsReadable() bool

	// SetPermissionOverwrites overwrites the current value of this record's
	// `permission_overwrites` field.
	SetPermissionOverwrites([]discord.PermissionOverwrite) ChannelPatch

	// SetNullPermissionOverwrites overwrites the current value of this record's
	// `permission_overwrites` field with `null`.
	SetNullPermissionOverwrites() ChannelPatch

	// UnsetPermissionOverwrites removes this record's `permission_overwrites`
	// field.
	UnsetPermissionOverwrites() ChannelPatch

	// ParentID returns the current value of this record's `parent_id` field.
	//
	// The `parent_id` field contains the id of the new parent category for a
	// channel.
	//
	// Applies to channel type(s): Text, News, Store, Voice
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use ParentIDIsReadable to check if the field is
	// present and non-null before use.
	ParentID() discord.Snowflake

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
	SetParentID(discord.Snowflake) ChannelPatch

	// SetNullParentID overwrites the current value of this record's `parent_id`
	// field with `null`.
	SetNullParentID() ChannelPatch

	// UnsetParentID removes this record's `parent_id` field.
	UnsetParentID() ChannelPatch
}
