package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/api/dio"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type ChannelAPI interface {
	// Get a channel by ID.
	//
	// Returns a channel object.
	Get(id dlib.Snowflake) (discord.Channel, error)

	// Patch updates a channel's settings.
	//
	// Requires the MANAGE_CHANNELS permission for the guild.
	//
	// Fires a Channel Update Gateway event. If modifying a category, individual
	// Channel Update events will fire for each child channel that also changes.
	//
	// Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.
	Patch(dlib.Snowflake, dio.ChannelPatch) (discord.Channel, error)

	// Delete a channel, or close a private message.
	//
	// Requires the MANAGE_CHANNELS permission for the guild.
	//
	// Deleting a category does not delete its child channels; they will have
	// their parent_id removed and a Channel Update Gateway event will fire for
	// each of them.
	//
	// Fires a Channel Delete Gateway event.
	//
	// Deleting a guild channel cannot be undone. Use this with caution, as it is
	// impossible to undo this action when performed on a guild channel. In
	// contrast, when used with a private message, it is possible to undo the
	// action by opening a private message with the recipient again.
	//
	// For Community guilds, the Rules or Guidelines channel and the Community
	// Updates channel cannot be deleted.
	Delete(id dlib.Snowflake) error

	// EditPermissions updates the channel permission overwrites for a user or
	// role in a channel.
	//
	// Only usable for guild channels.
	//
	// Requires the MANAGE_ROLES permission.
	EditPermissions(id dlib.Snowflake, patch dio.ChannelPermissionPatch) error

	// GetInvites returns a list of invite objects (with invite metadata) for the
	// channel.
	//
	// Only usable for guild channels.
	//
	// Requires the MANAGE_CHANNELS permission.
	GetInvites(id dlib.Snowflake) ([]discord.Invite, error)

	Messages() MessageAPI
}