package webhook

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type WebhookImpl struct {
	id        Snowflake
	kind      WebhookType
	guildID   types.OptionalSnowflake
	channelID Snowflake
	user      types.OptionalAny
	name      types.NullableString
	avatar    types.NullableString
	token     types.OptionalString
	appID     types.NullableSnowflake
}

func (w *WebhookImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(types.OutBuilder, 9).
		AppendRaw(serial.KeyID, w.id).
		AppendRaw(serial.KeyType, w.kind).
		AppendOptional(serial.KeyGuildID, &w.guildID).
		AppendRaw(serial.KeyChannelID, w.channelID).
		AppendOptional(serial.KeyUser, &w.user).
		AppendNullable(serial.KeyName, &w.name).
		AppendNullable(serial.KeyAvatar, &w.avatar).
		AppendOptional(serial.KeyToken, &w.token).
		AppendNullable(serial.KeyApplicationID, &w.appID))
}

func (w *WebhookImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage)
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	unm := map[serial.Key]json.Unmarshaler{
		serial.KeyID:            w.id,
		serial.KeyGuildID:       &w.guildID,
		serial.KeyChannelID:     w.channelID,
		serial.KeyUser:          &w.user,
		serial.KeyName:          &w.name,
		serial.KeyAvatar:        &w.avatar,
		serial.KeyToken:         &w.token,
		serial.KeyApplicationID: &w.appID,
	}

	for k, u := range unm {
		if _, ok := tmp[k]; ok {
			if err := u.UnmarshalJSON(tmp[k]); err != nil {
				return err
			}
		}
	}

	var kind uint8
	if err := json.Unmarshal(tmp[serial.KeyType], &kind); err != nil {
		return err
	}
	w.kind = WebhookType(kind)

	return nil
}

func (w *WebhookImpl) ID() Snowflake {
	return w.id
}

func (w *WebhookImpl) SetID(id Snowflake) Webhook {
	w.id = id
	return w
}

func (w *WebhookImpl) Type() WebhookType {
	return w.kind
}

func (w *WebhookImpl) SetType(kind WebhookType) Webhook {
	w.kind = kind
	return w
}

func (w *WebhookImpl) GuildID() Snowflake {
	return w.guildID.Get().(Snowflake)
}

func (w *WebhookImpl) GuildIDIsSet() bool {
	return w.guildID.IsSet()
}

func (w *WebhookImpl) SetGuildID(id Snowflake) Webhook {
	w.guildID.Set(id)
	return w
}

func (w *WebhookImpl) UnsetGuildID() Webhook {
	w.guildID.Unset()
	return w
}

func (w *WebhookImpl) ChannelID() Snowflake {
	return w.channelID
}

func (w *WebhookImpl) SetChannelID(id Snowflake) Webhook {
	w.channelID = id
	return w
}

func (w *WebhookImpl) User() User {
	return w.user.Get().(User)
}

func (w *WebhookImpl) UserIsSet() bool {
	return w.user.IsSet()
}

func (w *WebhookImpl) SetUser(user User) Webhook {
	w.user.Set(user)
	return w
}

func (w *WebhookImpl) UnsetUser() Webhook {
	w.user.Unset()
	return w
}

func (w *WebhookImpl) Name() string {
	return w.name.Get()
}

func (w *WebhookImpl) NameIsNull() bool {
	return w.name.IsNull()
}

func (w *WebhookImpl) SetName(name string) Webhook {
	w.name.Set(name)
	return w
}

func (w *WebhookImpl) SetNullName() Webhook {
	w.name.SetNull()
	return w
}

func (w *WebhookImpl) Avatar() ImageHash {
	return ImageHash(w.avatar.Get())
}

func (w *WebhookImpl) AvatarIsNull() bool {
	return w.avatar.IsNull()
}

func (w *WebhookImpl) SetAvatar(ava ImageHash) Webhook {
	w.avatar.Set(string(ava))
	return w
}

func (w *WebhookImpl) SetNullAvatar() Webhook {
	w.avatar.SetNull()
	return w
}

func (w *WebhookImpl) Token() string {
	return w.token.Get()
}

func (w *WebhookImpl) TokenIsSet() bool {
	return w.token.IsSet()
}

func (w *WebhookImpl) SetToken(tok string) Webhook {
	w.token.Set(tok)
	return w
}

func (w *WebhookImpl) UnsetToken() Webhook {
	w.token.Unset()
	return w
}

func (w *WebhookImpl) ApplicationID() Snowflake {
	return w.appID.Get().(Snowflake)
}

func (w *WebhookImpl) ApplicationIDIsNull() bool {
	return w.appID.IsNull()
}

func (w *WebhookImpl) SetApplicationID(id Snowflake) Webhook {
	w.appID.Set(id)
	return w
}

func (w *WebhookImpl) SetNullApplicationID() Webhook {
	w.appID.SetNull()
	return w
}
