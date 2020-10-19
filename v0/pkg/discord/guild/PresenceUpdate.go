package guild

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/gateway"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)


// A user's presence is their current state on a guild. This event is sent when
// a user's presence or info, such as name or avatar, is updated.
//
// If you are using Gateway Intents, you must specify the GUILD_PRESENCES intent
// in order to receive Presence Update events
//
// The user object within this event can be partial, the only field which must
// be sent is the id field, everything else is optional. Along with this
// limitation, no fields are required, and the types of the fields are not
// validated. Your client should expect any combination of fields and types
// within this event.
type PresenceUpdate interface {
	// User returns the current value of this record's `user` field.
	//
	// The `user` field contains the
	//
	// TODO: This should probably be it's own type since the API may return
	//       incompatible values.
	User() user.User

	// SetUser overwrites the current value of this record's `user` field.
	SetUser(user.User) PresenceUpdate

	// GuildID returns the current value of this record's `guild_id` field.
	//
	// The `guild_id` field contains the
	GuildID() dlib.Snowflake

	// SetGuildID overwrites the current value of this record's `guild_id` field.
	SetGuildID(dlib.Snowflake) PresenceUpdate

	// Status returns the current value of this record's `status` field.
	//
	// The `status` field contains the
	Status() PresenceStatus

	// SetStatus overwrites the current value of this record's `status` field.
	SetStatus(PresenceStatus) PresenceUpdate

	// Activities returns the current value of this record's `activities` field.
	//
	// The `activities` field contains the
	Activities() []gateway.Activity

	// SetActivities overwrites the current value of this record's `activities` field.
	SetActivities([]gateway.Activity) PresenceUpdate
}
