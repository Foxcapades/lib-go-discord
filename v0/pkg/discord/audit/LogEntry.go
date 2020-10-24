package audit

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type LogEntry interface {
	// TargetID returns the current value of this record's `target_id` field.
	//
	// The `target_id` field contains the id of the affected entity (webhook,
	// user, role, etc.).
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use TargetIDIsNull to check if the field is null before use.
	TargetID() string

	// TargetIDIsNull returns whether this record's `target_id` field is currently
	// null.
	TargetIDIsNull() bool

	// SetTargetID overwrites the current value of this record's `target_id`
	// field.
	SetTargetID(string) LogEntry

	// SetNullTargetID overwrites the current value of this record's `target_id`
	// field with `null`.
	SetNullTargetID() LogEntry

	// Changes returns the current value of this record's `changes` field.
	//
	// The `changes` field contains an array of changes made to the target_id.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ChangesIsSet to check if the field is present before use.
	Changes() []LogChange

	// ChangesIsSet returns whether this record's `changes` field is currently
	// present.
	ChangesIsSet() bool

	// SetChanges overwrites the current value of this record's `changes` field.
	SetChanges([]LogChange) LogEntry

	// UnsetChanges removes this record's `changes` field.
	UnsetChanges() LogEntry

	// UserID returns the current value of this record's `user_id` field.
	//
	// The `user_id` field contains the id of the user who made the changes.
	UserID() discord.Snowflake

	// SetUserID overwrites the current value of this record's `user_id` field.
	SetUserID(discord.Snowflake) LogEntry

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the entry.
	ID() discord.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(discord.Snowflake) LogEntry

	// ActionType returns the current value of this record's `action_type` field.
	//
	// The `action_type` field contains the type of action that occurred.
	ActionType() LogEvent

	// SetActionType overwrites the current value of this record's `action_type`
	// field.
	SetActionType(LogEvent) LogEntry

	// Options returns the current value of this record's `options` field.
	//
	// The `options` field contains additional info for certain action types.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use OptionsIsSet to check if the field is present before use.
	Options() OptionalEntryInfo

	// OptionsIsSet returns whether this record's `options` field is currently
	// present.
	OptionsIsSet() bool

	// SetOptions overwrites the current value of this record's `options` field.
	SetOptions(OptionalEntryInfo) LogEntry

	// UnsetOptions removes this record's `options` field.
	UnsetOptions() LogEntry

	// Reason returns the current value of this record's `reason` field.
	//
	// The `reason` field contains the reason for the change (0-512 characters).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ReasonIsSet to check if the field is present before use.
	Reason() LogReason

	// ReasonIsSet returns whether this record's `reason` field is currently
	// present.
	ReasonIsSet() bool

	// SetReason overwrites the current value of this record's `reason` field.
	SetReason(LogReason) LogEntry

	// UnsetReason removes this record's `reason` field.
	UnsetReason() LogEntry
}
