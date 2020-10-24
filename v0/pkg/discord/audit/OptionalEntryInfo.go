package audit

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/permission"
)

// TODO: The fields on this type seems pretty sus.  Are they all strings?  Are
//       they all required?  Its name is Optional, so perhaps all fields are
//       optional?
type OptionalEntryInfo interface {
	// DeleteMemberDays returns the current value of this record's
	// `delete_member_days` field.
	//
	// The `delete_member_days` field contains the number of days after which
	// inactive members were kicked.
	//
	// Applies to action type(s): MEMBER_PRUNE
	DeleteMemberDays() string

	// SetDeleteMemberDays overwrites the current value of this record's
	// `delete_member_days` field.
	SetDeleteMemberDays(string) OptionalEntryInfo

	// MembersRemoved returns the current value of this record's `members_removed`
	// field.
	//
	// The `members_removed` field contains the number of members removed by the
	// prune.
	//
	// Applies to action type(s): MEMBER_PRUNE
	MembersRemoved() string

	// SetMembersRemoved overwrites the current value of this record's
	// `members_removed` field.
	SetMembersRemoved(string) OptionalEntryInfo

	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the channel in which the entities were
	// targeted.
	//
	// Applies to action type(s): MEMBER_MOVE, MESSAGE_PIN, MESSAGE_UNPIN,
	// MESSAGE_DELETE
	ChannelID() discord.Snowflake

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(discord.Snowflake) OptionalEntryInfo

	// MessageID returns the current value of this record's `message_id` field.
	//
	// The `message_id` field contains the id of the message that was targeted.
	//
	// Applies to action type(s): MESSAGE_PIN, MESSAGE_UNPIN
	MessageID() discord.Snowflake

	// SetMessageID overwrites the current value of this record's `message_id` field.
	SetMessageID(discord.Snowflake) OptionalEntryInfo

	// Count returns the current value of this record's `count` field.
	//
	// The `count` field contains the number of entities that were targeted.
	//
	// Applies to action type(s): MESSAGE_DELETE, MESSAGE_BULK_DELETE,
	// MEMBER_DISCONNECT, MEMBER_MOVE
	Count() string

	// SetCount overwrites the current value of this record's `count` field.
	SetCount(string) OptionalEntryInfo

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the overwritten entity.
	//
	// Applies to action type(s): CHANNEL_OVERWRITE_CREATE,
	// CHANNEL_OVERWRITE_UPDATE, CHANNEL_OVERWRITE_DELETE
	ID() discord.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(discord.Snowflake) OptionalEntryInfo

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of overwritten entity - "0" for "role"
	// or "1" for "member".
	//
	// Applies to action type(s): CHANNEL_OVERWRITE_CREATE,
	// CHANNEL_OVERWRITE_UPDATE, CHANNEL_OVERWRITE_DELETE
	Type() permission.Type

	// SetType overwrites the current value of this record's `type` field.
	SetType(permission.Type) OptionalEntryInfo

	// RoleName returns the current value of this record's `role_name` field.
	//
	// The `role_name` field contains the name of the role if type is "0" (not
	// present if type is "1").
	//
	// Applies to action type(s): CHANNEL_OVERWRITE_CREATE,
	// CHANNEL_OVERWRITE_UPDATE, CHANNEL_OVERWRITE_DELETE
	RoleName() string

	// SetRoleName overwrites the current value of this record's `role_name` field.
	SetRoleName(string) OptionalEntryInfo
}
