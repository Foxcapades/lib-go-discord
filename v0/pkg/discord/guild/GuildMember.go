package guild

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"time"
)

type Member interface {
	// the user this guild member represents
	// The field user won't be included in the member object attached to MESSAGE_CREATE and MESSAGE_UPDATE gateway events.
	User() user.User
	UserIsSet() bool
	SetUser(user.User) Member
	UnsetUser() Member

	// this users guild nickname
	Nick() user.Nickname
	NickIsNull() bool
	SetNick(user.Nickname) Member
	SetNullNick() Member

	// array of role object ids
	Roles() []dlib.Snowflake
	SetRoles(roleIds []dlib.Snowflake) Member

	// when the user joined the guild
	JoinedAt() time.Time
	SetJoinedAt(time.Time) Member

	// when the user started boosting the guild
	PremiumSince() time.Time
	PremiumSinceIsSet() bool
	PremiumSinceIsNull() bool
	SetPremiumSince(time.Time) Member
	SetNullPremiumSince() Member
	UnsetPremiumSince() Member

	// whether the user is deafened in voice channels
	Deaf() bool
	SetDeaf(bool) Member

	// whether the user is muted in voice channels
	Mute() bool
	SetMute(bool) Member
}
