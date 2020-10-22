package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type GuildAPI interface {
	// GetGuild fetches and returns the guild object for the given id.
	//
	// If `with_counts` is set to true, this endpoint will also return
	// `approximate_member_count` and `approximate_presence_count` for the guild.
	GetGuild(id dlib.Snowflake) (discord.Guild, error)

	// GetGuildPreview fetches and returns the guild preview object for the given
	// id.
	//
	// If the user is not in the guild, then the guild must be Discoverable.
	GetGuildPreview(id dlib.Snowflake) (discord.GuildPreview, error)
}
