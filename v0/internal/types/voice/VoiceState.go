package voice

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type VoiceStateImpl struct {
	Validate bool

	guildID    types.OptionalSnowflake
	channelID  types.NullableSnowflake
	userID     Snowflake
	member     types.OptionalAny
	sessionId  string
	deaf       bool
	mute       bool
	selfDeaf   bool
	selfMute   bool
	selfStream types.OptionalBool
	selfVideo  bool
	suppress   bool
}

func (v *VoiceStateImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(types.OutBuilder, 12).
		AppendOptional(serial.KeyGuildID, &v.guildID).
		AppendNullable(serial.KeyChannelID, &v.channelID).
		AppendRaw(serial.KeyUserID, &v.userID).
		AppendOptional(serial.KeyMember, &v.member).
		AppendRaw(serial.KeySessionID, &v.sessionId).
		AppendRaw(serial.KeyDeaf, &v.deaf).
		AppendRaw(serial.KeyMute, &v.mute).
		AppendRaw(serial.KeySelfDeaf, &v.selfDeaf).
		AppendRaw(serial.KeySelfMute, &v.selfMute).
		AppendOptional(serial.KeySelfStream, &v.selfStream).
		AppendRaw(serial.KeySelfVideo, &v.selfVideo).
		AppendRaw(serial.KeySuppress, &v.suppress))
}

func (v *VoiceStateImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 12)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	unm := map[serial.Key]json.Unmarshaler{
		serial.KeyGuildID:    &v.guildID,
		serial.KeyChannelID:  &v.channelID,
		serial.KeyUserID:     v.userID,
		serial.KeyMember:     &v.member,
		serial.KeySelfStream: &v.selfStream,
	}

	for k, v := range unm {
		if _, ok := tmp[k]; ok {
			if err := v.UnmarshalJSON(tmp[k]); err != nil {
				return err
			}
		}
	}

	oth := map[serial.Key]interface{}{
		serial.KeySessionID: &v.sessionId,
		serial.KeyDeaf:      &v.deaf,
		serial.KeyMute:      &v.mute,
		serial.KeySelfDeaf:  &v.selfDeaf,
		serial.KeySelfMute:  &v.selfMute,
		serial.KeySelfVideo: &v.selfVideo,
		serial.KeySuppress:  &v.suppress,
	}

	for k, v := range oth {
		if _, ok := tmp[k]; ok {
			if err := json.Unmarshal(tmp[k], v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (v *VoiceStateImpl) GuildID() Snowflake {
	return v.guildID.Get().(Snowflake)
}

func (v *VoiceStateImpl) GuildIDIsSet() bool {
	return v.guildID.IsSet()
}

func (v *VoiceStateImpl) SetGuildID(id Snowflake) VoiceState {
	v.guildID.Set(id)
	return v
}

func (v *VoiceStateImpl) UnsetGuildID() VoiceState {
	v.guildID.Unset()
	return v
}

func (v *VoiceStateImpl) ChannelID() Snowflake {
	return v.channelID.Get().(Snowflake)
}

func (v *VoiceStateImpl) ChannelIDIsNull() bool {
	return v.channelID.IsNull()
}

func (v *VoiceStateImpl) SetChannelID(id Snowflake) VoiceState {
	v.channelID.Set(id)
	return v
}

func (v *VoiceStateImpl) SetNullChannelID() VoiceState {
	v.channelID.SetNull()
	return v
}

func (v *VoiceStateImpl) UserID() Snowflake {
	return v.userID
}

func (v *VoiceStateImpl) SetUserID(id Snowflake) VoiceState {
	v.userID = id
	return v
}

func (v *VoiceStateImpl) Member() GuildMember {
	return v.member.Get().(GuildMember)
}

func (v *VoiceStateImpl) MemberIsSet() bool {
	return v.member.IsSet()
}

func (v *VoiceStateImpl) SetMember(member GuildMember) VoiceState {
	v.member.Set(member)
	return v
}

func (v *VoiceStateImpl) UnsetMember() VoiceState {
	v.member.Unset()
	return v
}

func (v *VoiceStateImpl) SessionID() string {
	return v.sessionId
}

func (v *VoiceStateImpl) SetSessionID(s string) VoiceState {
	v.sessionId = s
	return v
}

func (v *VoiceStateImpl) Deaf() bool {
	return v.deaf
}

func (v *VoiceStateImpl) SetDeaf(b bool) VoiceState {
	v.deaf = b
	return v
}

func (v *VoiceStateImpl) Mute() bool {
	return v.mute
}

func (v *VoiceStateImpl) SetMute(b bool) VoiceState {
	v.mute = b
	return v
}

func (v *VoiceStateImpl) SelfDeaf() bool {
	return v.selfDeaf
}

func (v *VoiceStateImpl) SetSelfDeaf(b bool) VoiceState {
	v.selfDeaf = b
	return v
}

func (v *VoiceStateImpl) SelfMute() bool {
	return v.selfMute
}

func (v *VoiceStateImpl) SetSelfMute(b bool) VoiceState {
	v.selfMute = b
	return v
}

func (v *VoiceStateImpl) SelfStream() bool {
	return v.selfStream.Get()
}

func (v *VoiceStateImpl) SelfStreamIsSet() bool {
	return v.selfStream.IsSet()
}

func (v *VoiceStateImpl) SetSelfStream(b bool) VoiceState {
	v.selfStream.Set(b)
	return v
}

func (v *VoiceStateImpl) UnsetSelfStream() VoiceState {
	v.selfStream.Unset()
	return v
}

func (v *VoiceStateImpl) SelfVideo() bool {
	return v.selfVideo
}

func (v *VoiceStateImpl) SetSelfVideo(b bool) VoiceState {
	v.selfVideo = b
	return v
}

func (v *VoiceStateImpl) Suppress() bool {
	return v.suppress
}

func (v *VoiceStateImpl) SetSuppress(b bool) VoiceState {
	v.suppress = b
	return v
}
