package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type AllowedMentions interface {
	// Parse returns the current value of this record's `parse` field.
	//
	// The `parse` field contains an array of allowed mention types to parse from
	// the content.
	Parse() []AllowedMentionType

	// SetParse overwrites the current value of this record's `parse` field.
	SetParse([]AllowedMentionType) AllowedMentions

	// Roles returns the current value of this record's `roles` field.
	//
	// The `roles` field contains an array of role_ids to mention (Max size of
	// 100).
	Roles() []discord.Snowflake

	// SetRoles overwrites the current value of this record's `roles` field.
	SetRoles([]discord.Snowflake) AllowedMentions

	// Users returns the current value of this record's `users` field.
	//
	// The `users` field contains an array of user_ids to mention (Max size of
	// 100).
	Users() []discord.Snowflake

	// SetUsers overwrites the current value of this record's `users` field.
	SetUsers([]discord.Snowflake) AllowedMentions
}
