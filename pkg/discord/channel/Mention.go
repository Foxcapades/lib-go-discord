package channel

import "github.com/foxcapades/lib-go-discord/pkg/dlib"

type Mention interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the channel.
	ID() dlib.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(dlib.Snowflake) Mention

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the id of the guild containing the channel.
	GuildID() dlib.Snowflake

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(dlib.Snowflake) Mention

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of the channel.
	Type() Type

	// SetType overwrites the current value of this record's `type` field.
	SetType(Type) Mention

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the channel.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) Mention
}
