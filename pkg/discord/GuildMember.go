package discord

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
	"time"
)

type GuildMember interface {
	// the user this guild member represents
	// The field user won't be included in the member object attached to MESSAGE_CREATE and MESSAGE_UPDATE gateway events.
	User() user.User
	UserIsSet() bool
	SetUser(user.User) GuildMember
	UnsetUser() GuildMember

	// this users guild nickname
	Nick() user.Nickname
	NickIsNull() bool
	SetNick(user.Nickname) GuildMember
	SetNullNick() GuildMember

	// array of role object ids
	Roles() []dlib.Snowflake
	SetRoles(roleIds []dlib.Snowflake) GuildMember

	// when the user joined the guild
	JoinedAt() time.Time
	SetJoinedAt(time.Time) GuildMember

	// when the user started boosting the guild
	PremiumSince() time.Time
	PremiumSinceIsSet() bool
	PremiumSinceIsNull() bool
	SetPremiumSince(time.Time) GuildMember
	SetNullPremiumSince() GuildMember
	UnsetPremiumSince() GuildMember

	// whether the user is deafened in voice channels
	Deaf() bool
	SetDeaf(bool) GuildMember

	// whether the user is muted in voice channels
	Mute() bool
	SetMute(bool) GuildMember
}
