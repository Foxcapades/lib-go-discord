package guild

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

func NewEmoji() discord.Emoji {
	return new(emoji)
}

type emoji struct {}

func (e *emoji) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (e *emoji) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (e *emoji) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (e *emoji) IsNil() bool {
	panic("implement me")
}

func (e *emoji) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (e *emoji) NKeys() int {
	panic("implement me")
}

func (e *emoji) JSONSize() int {
	panic("implement me")
}

func (e *emoji) IsValid() bool {
	panic("implement me")
}

func (e *emoji) Validate() error {
	panic("implement me")
}

func (e *emoji) ID() discord.Snowflake {
	panic("implement me")
}

func (e *emoji) IDIsNull() bool {
	panic("implement me")
}

func (e *emoji) SetID(id discord.Snowflake) discord.Emoji {
	panic("implement me")
}

func (e *emoji) SetNullID() discord.Emoji {
	panic("implement me")
}

func (e *emoji) Name() string {
	panic("implement me")
}

func (e *emoji) NameIsNull() bool {
	panic("implement me")
}

func (e *emoji) SetName(name string) discord.Emoji {
	panic("implement me")
}

func (e *emoji) SetNullName() discord.Emoji {
	panic("implement me")
}

func (e *emoji) Roles() []discord.Role {
	panic("implement me")
}

func (e *emoji) RolesIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetRoles(roles []discord.Role) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetRoles() discord.Emoji {
	panic("implement me")
}

func (e *emoji) User() discord.User {
	panic("implement me")
}

func (e *emoji) UserIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetUser(user discord.User) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetUser() discord.Emoji {
	panic("implement me")
}

func (e *emoji) RequireColons() bool {
	panic("implement me")
}

func (e *emoji) RequireColonsIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetRequireColons(b bool) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetRequireColons() discord.Emoji {
	panic("implement me")
}

func (e *emoji) Managed() bool {
	panic("implement me")
}

func (e *emoji) ManagedIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetManaged(b bool) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetManaged() discord.Emoji {
	panic("implement me")
}

func (e *emoji) Animated() bool {
	panic("implement me")
}

func (e *emoji) AnimatedIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetAnimated(b bool) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetAnimated() discord.Emoji {
	panic("implement me")
}

func (e *emoji) Available() bool {
	panic("implement me")
}

func (e *emoji) AvailableIsSet() bool {
	panic("implement me")
}

func (e *emoji) SetAvailable(b bool) discord.Emoji {
	panic("implement me")
}

func (e *emoji) UnsetAvailable() discord.Emoji {
	panic("implement me")
}
