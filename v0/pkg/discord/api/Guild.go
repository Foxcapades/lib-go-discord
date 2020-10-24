package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/api/dio"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/audit"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type GuildAPI interface {
	// CreateGuild creates a new guild.
	//
	// Returns a guild object on success.
	//
	// Fires a Guild Create Gateway event.
	CreateGuild(dio.CreateGuildRequest) (Guild, error)

	// GetGuild fetches and returns the guild object for the given id.
	//
	// If `with_counts` is set to true, this endpoint will also return
	// `approximate_member_count` and `approximate_presence_count` for the guild.
	GetGuild(id Snowflake) (Guild, error)

	// GetGuildPreview fetches and returns the guild preview object for the given
	// id.
	//
	// If the user is not in the guild, then the guild must be Discoverable.
	GetGuildPreview(id Snowflake) (GuildPreview, error)

	// GetAuditLogs returns an audit log object for the guild.
	//
	// Requires the 'VIEW_AUDIT_LOG' permission.
	GetAuditLogs(id Snowflake) (audit.Log, error)
}
