package gateway

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/gateway/activity"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"time"
)

// Activity
//
// TODO: document me
//
// Bots are only able to send name, type, and optionally url.
type Activity interface {
	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the activity's name.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) Activity

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the activity type.
	Type() activity.Type

	// SetType overwrites the current value of this record's `type` field.
	SetType(activity.Type) Activity

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the stream url, is validated when type is 1.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use URLIsReadable to check if the field is
	// present and non-null before use.
	URL() string

	// URLIsNull returns whether this record's `url` field is currently null.
	URLIsNull() bool

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// URLIsReadable returns whether this record's `url` field is currently set
	// and non-null.
	URLIsReadable() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) Activity

	// SetNullURL overwrites the current value of this record's `url` field
	// with `null`.
	SetNullURL() Activity

	// UnsetURL removes this record's `url` field.
	UnsetURL() Activity

	// CreatedAt returns the current value of this record's `created_at` field.
	//
	// The `created_at` field contains the unix timestamp of when the activity was
	// added to the user's session.
	CreatedAt() time.Time

	// SetCreatedAt overwrites the current value of this record's `created_at` field.
	SetCreatedAt(time.Time) Activity

	// Timestamps returns the current value of this record's `timestamps` field.
	//
	// The `timestamps` field contains the unix timestamps for start and/or end of
	// the game.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TimestampsIsSet to check if the field is present before use.
	Timestamps() activity.Timestamps

	// TimestampsIsSet returns whether this record's `timestamps` field is currently present.
	TimestampsIsSet() bool

	// SetTimestamps overwrites the current value of this record's `timestamps` field.
	SetTimestamps(activity.Timestamps) Activity

	// UnsetTimestamps removes this record's `timestamps` field.
	UnsetTimestamps() Activity

	// ApplicationID returns the current value of this record's `application_id` field.
	//
	// The `application_id` field contains the application id for the game.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ApplicationIDIsSet to check if the field is present before use.
	ApplicationID() dlib.Snowflake

	// ApplicationIDIsSet returns whether this record's `application_id` field is currently present.
	ApplicationIDIsSet() bool

	// SetApplicationID overwrites the current value of this record's `application_id` field.
	SetApplicationID(dlib.Snowflake) Activity

	// UnsetApplicationID removes this record's `application_id` field.
	UnsetApplicationID() Activity

	// Details returns the current value of this record's `details` field.
	//
	// The `details` field contains what the player is currently doing.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use DetailsIsReadable to check if the field is
	// present and non-null before use.
	Details() string

	// DetailsIsNull returns whether this record's `details` field is currently null.
	DetailsIsNull() bool

	// DetailsIsSet returns whether this record's `details` field is currently present.
	DetailsIsSet() bool

	// DetailsIsReadable returns whether this record's `details` field is currently set
	// and non-null.
	DetailsIsReadable() bool

	// SetDetails overwrites the current value of this record's `details` field.
	SetDetails(string) Activity

	// SetNullDetails overwrites the current value of this record's `details` field
	// with `null`.
	SetNullDetails() Activity

	// UnsetDetails removes this record's `details` field.
	UnsetDetails() Activity

	// State returns the current value of this record's `state` field.
	//
	// The `state` field contains the user's current party status.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use StateIsReadable to check if the field is
	// present and non-null before use.
	State() string

	// StateIsNull returns whether this record's `state` field is currently null.
	StateIsNull() bool

	// StateIsSet returns whether this record's `state` field is currently present.
	StateIsSet() bool

	// StateIsReadable returns whether this record's `state` field is currently set
	// and non-null.
	StateIsReadable() bool

	// SetState overwrites the current value of this record's `state` field.
	SetState(string) Activity

	// SetNullState overwrites the current value of this record's `state` field
	// with `null`.
	SetNullState() Activity

	// UnsetState removes this record's `state` field.
	UnsetState() Activity

	// Emoji returns the current value of this record's `emoji` field.
	//
	// The `emoji` field contains the emoji used for a custom status.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use EmojiIsReadable to check if the field is
	// present and non-null before use.
	Emoji() activity.Emoji

	// EmojiIsNull returns whether this record's `emoji` field is currently null.
	EmojiIsNull() bool

	// EmojiIsSet returns whether this record's `emoji` field is currently present.
	EmojiIsSet() bool

	// EmojiIsReadable returns whether this record's `emoji` field is currently set
	// and non-null.
	EmojiIsReadable() bool

	// SetEmoji overwrites the current value of this record's `emoji` field.
	SetEmoji(activity.Emoji) Activity

	// SetNullEmoji overwrites the current value of this record's `emoji` field
	// with `null`.
	SetNullEmoji() Activity

	// UnsetEmoji removes this record's `emoji` field.
	UnsetEmoji() Activity

	// Party returns the current value of this record's `party` field.
	//
	// The `party` field contains information for the current party of the player.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PartyIsSet to check if the field is present before use.
	Party() activity.Party

	// PartyIsSet returns whether this record's `party` field is currently present.
	PartyIsSet() bool

	// SetParty overwrites the current value of this record's `party` field.
	SetParty(activity.Party) Activity

	// UnsetParty removes this record's `party` field.
	UnsetParty() Activity

	// Assets returns the current value of this record's `assets` field.
	//
	// The `assets` field contains images for the presence and their hover texts.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AssetsIsSet to check if the field is present before use.
	Assets() activity.Assets

	// AssetsIsSet returns whether this record's `assets` field is currently present.
	AssetsIsSet() bool

	// SetAssets overwrites the current value of this record's `assets` field.
	SetAssets(activity.Assets) Activity

	// UnsetAssets removes this record's `assets` field.
	UnsetAssets() Activity

	// Secrets returns the current value of this record's `secrets` field.
	//
	// The `secrets` field contains secrets for Rich Presence joining and
	// spectating.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SecretsIsSet to check if the field is present before use.
	Secrets() activity.Secrets

	// SecretsIsSet returns whether this record's `secrets` field is currently present.
	SecretsIsSet() bool

	// SetSecrets overwrites the current value of this record's `secrets` field.
	SetSecrets(activity.Secrets) Activity

	// UnsetSecrets removes this record's `secrets` field.
	UnsetSecrets() Activity

	// Instanced returns the current value of this record's `instanced` field.
	//
	// The `instanced` field indicates whether or not the activity is an instanced
	// game session.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use InstancedIsSet to check if the field is present before use.
	Instanced() bool

	// InstancedIsSet returns whether this record's `instanced` field is currently present.
	InstancedIsSet() bool

	// SetInstanced overwrites the current value of this record's `instanced` field.
	SetInstanced(bool) Activity

	// UnsetInstanced removes this record's `instanced` field.
	UnsetInstanced() Activity

	// Flags returns the current value of this record's `flags` field.
	//
	// The `flags` field contains activity flags ORd together, describes what the
	// payload includes.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use FlagsIsSet to check if the field is present before use.
	Flags() activity.Flag

	// FlagsIsSet returns whether this record's `flags` field is currently present.
	FlagsIsSet() bool

	// SetFlags overwrites the current value of this record's `flags` field.
	SetFlags(activity.Flag) Activity

	// UnsetFlags removes this record's `flags` field.
	UnsetFlags() Activity
}
