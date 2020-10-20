package message

import "github.com/foxcapades/lib-go-discord/v0/pkg/dlib"

type Reference interface {
	// MessageID returns the current value of this record's `message_id` field.
	//
	// The `message_id` field contains the id of the originating message.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MessageIDIsSet to check if the field is present before use.
	MessageID() dlib.Snowflake

	// MessageIDIsSet returns whether this record's `message_id` field is
	// currently present.
	MessageIDIsSet() bool

	// SetMessageID overwrites the current value of this record's `message_id`
	// field.
	SetMessageID(dlib.Snowflake) Reference

	// UnsetMessageID removes this record's `message_id` field.
	UnsetMessageID() Reference

	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the id of the originating message's
	// channel.
	ChannelID() dlib.Snowflake

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(dlib.Snowflake) Reference

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the id of the originating message's guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use GuildIDIsSet to check if the field is present before use.
	GuildID() dlib.Snowflake

	// GuildIDIsSet returns whether this record's `guild_id` field is currently
	// present.
	GuildIDIsSet() bool

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(dlib.Snowflake) Reference

	// UnsetGuildID removes this record's `guild_id` field.
	UnsetGuildID() Reference
}
