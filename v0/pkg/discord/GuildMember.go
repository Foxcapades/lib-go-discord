package discord

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
)

type GuildMember interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// the user this guild member represents
	// The field user won't be included in the member object attached to MESSAGE_CREATE and MESSAGE_UPDATE gateway events.
	User() User
	UserIsSet() bool
	SetUser(User) GuildMember
	UnsetUser() GuildMember

	// this users guild nickname
	Nick() UserNickname
	NickIsNull() bool
	SetNick(UserNickname) GuildMember
	SetNullNick() GuildMember

	// array of role object ids
	Roles() []Snowflake
	SetRoles(roleIds []Snowflake) GuildMember

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

func NewGuildMember(validate bool) GuildMember {
	panic("implement me")
}
