package discord

import "github.com/foxcapades/lib-go-discord/pkg/dlib"

type Channel interface {
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) Channel

	Type() channel.Type
	SetType(channel.Type) Channel

	GuildID() dlib.Snowflake
	GuildIDIsSet() bool
	SetGuildID(id dlib.Snowflake) Channel
	UnsetGuildID() Channel

	Position() uint16
	PositionIsSet() bool
	SetPosition(pos uint16) Channel
	UnsetPosition() Channel

	PermissionOverwrites() []PermissionOverwrite
}

type Message interface {

}