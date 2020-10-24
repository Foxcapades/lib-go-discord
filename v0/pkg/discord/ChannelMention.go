package discord

type ChannelMention interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the channel.
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) ChannelMention

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the id of the guild containing the channel.
	GuildID() Snowflake

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(Snowflake) ChannelMention

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of the channel.
	Type() ChannelType

	// SetType overwrites the current value of this record's `type` field.
	SetType(ChannelType) ChannelMention

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the channel.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) ChannelMention
}
