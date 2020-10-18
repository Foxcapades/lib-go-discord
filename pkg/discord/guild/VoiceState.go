package guild

import "github.com/foxcapades/lib-go-discord/pkg/dlib"

// VoiceState
// TODO: Document me
type VoiceState interface {
	// the guild id this voice state is for
	GuildID() dlib.Snowflake
	GuildIDIsSet() bool
	SetGuildID(id dlib.Snowflake) VoiceState
	UnsetGuildID() VoiceState

	// the channel id this user is connected to
	ChannelID() dlib.Snowflake
	ChannelIDIsNull() bool
	SetChannelID(id dlib.Snowflake) VoiceState
	SetNullChannelID() VoiceState

	// the user id this voice state is for
	UserID() dlib.Snowflake
	SetUserID(id dlib.Snowflake) VoiceState

	// the guild member this voice state is for
	Member() Member
	MemberIsSet() bool
	SetMember(member Member) VoiceState
	UnsetMember() VoiceState

	// the session id for this voice state
	SessionID() string
	SetSessionID(id string) VoiceState

	// whether this user is deafened by the server
	Deaf() bool
	SetDeaf(bool) VoiceState

	// whether this user is muted by the server
	Mute() bool
	SetMute(bool) VoiceState

	// whether this user is locally deafened
	SelfDeaf() bool
	SetSelfDeaf(bool) VoiceState

	// whether this user is locally muted
	SelfMute() bool
	SetSelfMute(bool) VoiceState

	// whether this user is streaming using "Go Live"
	SelfStream() bool
	SelfStreamIsSet() bool
	SetSelfStream(bool) VoiceState
	UnsetSelfStream() VoiceState

	// whether this user's camera is enabled
	SelfVideo() bool
	SetSelfVideo(bool) VoiceState

	// whether this user is muted by the current user
	Suppress() bool
	SetSuppress(bool) VoiceState
}
