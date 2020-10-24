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
