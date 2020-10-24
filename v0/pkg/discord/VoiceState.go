package discord

import (
	"encoding/json"
)

// VoiceState
// TODO: Document me
type VoiceState interface {
	json.Marshaler
	json.Unmarshaler

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the guild id this voice state is for.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use GuildIDIsSet to check if the field is present before use.
	GuildID() Snowflake

	// GuildIDIsSet returns whether this record's `guild_id` field is currently
	// present.
	GuildIDIsSet() bool

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(Snowflake) VoiceState

	// UnsetGuildID removes this record's `guild_id` field.
	UnsetGuildID() VoiceState

	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the channel id this user is connected to.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use ChannelIDIsNull to check if the field is null before use.
	ChannelID() Snowflake

	// ChannelIDIsNull returns whether this record's `channel_id` field is
	// currently null.
	ChannelIDIsNull() bool

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(Snowflake) VoiceState

	// SetNullChannelID overwrites the current value of this record's `channel_id`
	// field with `null`.
	SetNullChannelID() VoiceState

	// UserID returns the current value of this record's `user_id` field.
	//
	// The `user_id` field contains the user id this voice state is for.
	UserID() Snowflake

	// SetUserID overwrites the current value of this record's `user_id` field.
	SetUserID(Snowflake) VoiceState

	// Member returns the current value of this record's `member` field.
	//
	// The `member` field contains the guild member this voice state is for.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MemberIsSet to check if the field is present before use.
	Member() GuildMember

	// MemberIsSet returns whether this record's `member` field is currently
	// present.
	MemberIsSet() bool

	// SetMember overwrites the current value of this record's `member` field.
	SetMember(GuildMember) VoiceState

	// UnsetMember removes this record's `member` field.
	UnsetMember() VoiceState

	// SessionID returns the current value of this record's `session_id` field.
	//
	// The `session_id` field contains the session id for this voice state.
	SessionID() string

	// SetSessionID overwrites the current value of this record's `session_id`
	// field.
	SetSessionID(string) VoiceState

	// Deaf returns the current value of this record's `deaf` field.
	//
	// The `deaf` field indicates whether this user is deafened by the server.
	Deaf() bool

	// SetDeaf overwrites the current value of this record's `deaf` field.
	SetDeaf(bool) VoiceState

	// Mute returns the current value of this record's `mute` field.
	//
	// The `mute` field indicates whether this user is muted by the server.
	Mute() bool

	// SetMute overwrites the current value of this record's `mute` field.
	SetMute(bool) VoiceState

	// SelfDeaf returns the current value of this record's `self_deaf` field.
	//
	// The `self_deaf` field indicates whether this user is locally deafened.
	SelfDeaf() bool

	// SetSelfDeaf overwrites the current value of this record's `self_deaf`
	// field.
	SetSelfDeaf(bool) VoiceState

	// SelfMute returns the current value of this record's `self_mute` field.
	//
	// The `self_mute` field indicates whether this user is locally muted.
	SelfMute() bool

	// SetSelfMute overwrites the current value of this record's `self_mute`
	// field.
	SetSelfMute(bool) VoiceState

	// SelfStream returns the current value of this record's `self_stream` field.
	//
	// The `self_stream` field indicates whether this user is streaming using
	// "Go Live".
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SelfStreamIsSet to check if the field is present before use.
	SelfStream() bool

	// SelfStreamIsSet returns whether this record's `self_stream` field is
	// currently present.
	SelfStreamIsSet() bool

	// SetSelfStream overwrites the current value of this record's `self_stream`
	// field.
	SetSelfStream(bool) VoiceState

	// UnsetSelfStream removes this record's `self_stream` field.
	UnsetSelfStream() VoiceState

	// SelfVideo returns the current value of this record's `self_video` field.
	//
	// The `self_video` field indicates whether this user's camera is enabled.
	SelfVideo() bool

	// SetSelfVideo overwrites the current value of this record's `self_video`
	// field.
	SetSelfVideo(bool) VoiceState

	// Suppress returns the current value of this record's `suppress` field.
	//
	// The `suppress` field indicates whether this user is muted by the current
	// user.
	Suppress() bool

	// SetSuppress overwrites the current value of this record's `suppress` field.
	SetSuppress(bool) VoiceState
}

func NewVoiceState(validate bool) VoiceState {
	return &voiceState{validate: validate}
}

type voiceState struct {
	validate   bool

	guildID    OptionalSnowflake
	channelID  NullableSnowflake
	userID     Snowflake
	member     OptionalGuildMember
	sessionId  string
	deaf       bool
	mute       bool
	selfDeaf   bool
	selfMute   bool
	selfStream OptionalBool
	selfVideo  bool
	suppress   bool
}

func (v *voiceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(outBuilder, 12).
		AppendOptional(FieldKeyGuildID, &v.guildID).
		AppendNullable(FieldKeyChannelID, &v.channelID).
		AppendRaw(FieldKeyUserID, &v.userID).
		AppendOptional(FieldKeyMember, &v.member).
		AppendRaw(FieldKeySessionID, &v.sessionId).
		AppendRaw(FieldKeyDeaf, &v.deaf).
		AppendRaw(FieldKeyMute, &v.mute).
		AppendRaw(FieldKeySelfDeaf, &v.selfDeaf).
		AppendRaw(FieldKeySelfMute, &v.selfMute).
		AppendOptional(FieldKeySelfStream, &v.selfStream).
		AppendRaw(FieldKeySelfVideo, &v.selfVideo).
		AppendRaw(FieldKeySuppress, &v.suppress))
}

func (v *voiceState) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]json.RawMessage, 12)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	unm := map[FieldKey]json.Unmarshaler{
		FieldKeyGuildID:    &v.guildID,
		FieldKeyChannelID:  &v.channelID,
		FieldKeyUserID:     &v.userID,
		FieldKeyMember:     &v.member,
		FieldKeySelfStream: &v.selfStream,
	}

	for k, v := range unm {
		if _, ok := tmp[k]; ok {
			if err := v.UnmarshalJSON(tmp[k]); err != nil {
				return err
			}
		}
	}

	oth := map[FieldKey]interface{}{
		FieldKeySessionID: &v.sessionId,
		FieldKeyDeaf:      &v.deaf,
		FieldKeyMute:      &v.mute,
		FieldKeySelfDeaf:  &v.selfDeaf,
		FieldKeySelfMute:  &v.selfMute,
		FieldKeySelfVideo: &v.selfVideo,
		FieldKeySuppress:  &v.suppress,
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

func (v *voiceState) GuildID() Snowflake {
	return v.guildID.Get()
}

func (v *voiceState) GuildIDIsSet() bool {
	return v.guildID.IsSet()
}

func (v *voiceState) SetGuildID(id Snowflake) VoiceState {
	v.guildID.Set(id)
	return v
}

func (v *voiceState) UnsetGuildID() VoiceState {
	v.guildID.Unset()
	return v
}

func (v *voiceState) ChannelID() Snowflake {
	return v.channelID.Get()
}

func (v *voiceState) ChannelIDIsNull() bool {
	return v.channelID.IsNull()
}

func (v *voiceState) SetChannelID(id Snowflake) VoiceState {
	v.channelID.Set(id)
	return v
}

func (v *voiceState) SetNullChannelID() VoiceState {
	v.channelID.SetNull()
	return v
}

func (v *voiceState) UserID() Snowflake {
	return v.userID
}

func (v *voiceState) SetUserID(id Snowflake) VoiceState {
	v.userID = id
	return v
}

func (v *voiceState) Member() GuildMember {
	return v.member.Get()
}

func (v *voiceState) MemberIsSet() bool {
	return v.member.IsSet()
}

func (v *voiceState) SetMember(member GuildMember) VoiceState {
	v.member.Set(member)
	return v
}

func (v *voiceState) UnsetMember() VoiceState {
	v.member.Unset()
	return v
}

func (v *voiceState) SessionID() string {
	return v.sessionId
}

func (v *voiceState) SetSessionID(s string) VoiceState {
	v.sessionId = s
	return v
}

func (v *voiceState) Deaf() bool {
	return v.deaf
}

func (v *voiceState) SetDeaf(b bool) VoiceState {
	v.deaf = b
	return v
}

func (v *voiceState) Mute() bool {
	return v.mute
}

func (v *voiceState) SetMute(b bool) VoiceState {
	v.mute = b
	return v
}

func (v *voiceState) SelfDeaf() bool {
	return v.selfDeaf
}

func (v *voiceState) SetSelfDeaf(b bool) VoiceState {
	v.selfDeaf = b
	return v
}

func (v *voiceState) SelfMute() bool {
	return v.selfMute
}

func (v *voiceState) SetSelfMute(b bool) VoiceState {
	v.selfMute = b
	return v
}

func (v *voiceState) SelfStream() bool {
	return v.selfStream.Get()
}

func (v *voiceState) SelfStreamIsSet() bool {
	return v.selfStream.IsSet()
}

func (v *voiceState) SetSelfStream(b bool) VoiceState {
	v.selfStream.Set(b)
	return v
}

func (v *voiceState) UnsetSelfStream() VoiceState {
	v.selfStream.Unset()
	return v
}

func (v *voiceState) SelfVideo() bool {
	return v.selfVideo
}

func (v *voiceState) SetSelfVideo(b bool) VoiceState {
	v.selfVideo = b
	return v
}

func (v *voiceState) Suppress() bool {
	return v.suppress
}

func (v *voiceState) SetSuppress(b bool) VoiceState {
	v.suppress = b
	return v
}
