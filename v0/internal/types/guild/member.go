package guild

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
	"time"
)

func NewMember() discord.GuildMember {
	return new(member)
}

type member struct{}

func (m *member) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (m *member) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (m *member) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (m *member) IsNil() bool {
	panic("implement me")
}

func (m *member) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (m *member) NKeys() int {
	panic("implement me")
}

func (m *member) IsValid() bool {
	panic("implement me")
}

func (m *member) Validate() error {
	panic("implement me")
}

func (m *member) User() discord.User {
	panic("implement me")
}

func (m *member) UserIsSet() bool {
	panic("implement me")
}

func (m *member) SetUser(user discord.User) discord.GuildMember {
	panic("implement me")
}

func (m *member) UnsetUser() discord.GuildMember {
	panic("implement me")
}

func (m *member) Nick() discord.UserNickname {
	panic("implement me")
}

func (m *member) NickIsNull() bool {
	panic("implement me")
}

func (m *member) SetNick(nick discord.UserNickname) discord.GuildMember {
	panic("implement me")
}

func (m *member) SetNullNick() discord.GuildMember {
	panic("implement me")
}

func (m *member) Roles() []discord.Snowflake {
	panic("implement me")
}

func (m *member) SetRoles(roleIds []discord.Snowflake) discord.GuildMember {
	panic("implement me")
}

func (m *member) JoinedAt() time.Time {
	panic("implement me")
}

func (m *member) SetJoinedAt(t time.Time) discord.GuildMember {
	panic("implement me")
}

func (m *member) PremiumSince() time.Time {
	panic("implement me")
}

func (m *member) PremiumSinceIsSet() bool {
	panic("implement me")
}

func (m *member) PremiumSinceIsNull() bool {
	panic("implement me")
}

func (m *member) SetPremiumSince(t time.Time) discord.GuildMember {
	panic("implement me")
}

func (m *member) SetNullPremiumSince() discord.GuildMember {
	panic("implement me")
}

func (m *member) UnsetPremiumSince() discord.GuildMember {
	panic("implement me")
}

func (m *member) Deaf() bool {
	panic("implement me")
}

func (m *member) SetDeaf(b bool) discord.GuildMember {
	panic("implement me")
}

func (m *member) Mute() bool {
	panic("implement me")
}

func (m *member) SetMute(b bool) discord.GuildMember {
	panic("implement me")
}
