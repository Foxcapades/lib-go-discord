package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/api/dio"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type ChannelAPI interface {
	// Get a channel by ID.
	//
	// Returns a channel object.
	Get(id Snowflake) (Channel, error)

	// Patch updates a channel's settings.
	//
	// Requires the MANAGE_CHANNELS permission for the guild.
	//
	// Fires a Channel Update Gateway event. If modifying a category, individual
	// Channel Update events will fire for each child channel that also changes.
	//
	// Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.
	Patch(Snowflake, dio.ChannelPatch) (Channel, error)

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
	Delete(id Snowflake) error

	// EditPermissions updates the channel permission overwrites for a user or
	// role in a channel.
	//
	// Only usable for guild channels.
	//
	// Requires the MANAGE_ROLES permission.
	EditPermissions(id Snowflake, patch dio.ChannelPermissionPatch) error

	// GetInvites returns a list of invite objects (with invite metadata) for the
	// channel.
	//
	// Only usable for guild channels.
	//
	// Requires the MANAGE_CHANNELS permission.
	GetInvites(id Snowflake) ([]Invite, error)

	// DeletePermission deletes a channel permission overwrite for a user or role
	// in a channel.
	//
	// Only usable for guild channels.
	//
	// Requires the MANAGE_ROLES permission.
	DeletePermission(chanId, permId Snowflake) error

	// FollowNewsChannel follows a News Channel to send messages to a target
	// channel.
	//
	// Requires the MANAGE_WEBHOOKS permission in the target channel.
	FollowNewsChannel(chanId, webhookChanId Snowflake) (FollowedChannel, error)

	// TriggerTypingIndicator starts a typing indicator for the specified channel.
	//
	// Generally bots should not implement this route. However, if a bot is
	// responding to a command and expects the computation to take a few seconds,
	// this endpoint may be called to let the user know that the bot is processing
	// their message.
	//
	// Fires a Typing Start Gateway event.
	TriggerTypingIndicator(chanId Snowflake) error

	// GetPinnedMessages returns all pinned messages in the channel as an array of
	// message objects.
	GetPinnedMessages(chanId Snowflake) ([]Message, error)

	// PinChannelMessage pins a message in a channel.
	//
	// Requires the MANAGE_MESSAGES permission.
	//
	// The max pinned messages is 50.
	PinMessage(chanId, msgId Snowflake) error

	// UnpinMessage deletes a pinned message in a channel.
	//
	// Requires the MANAGE_MESSAGES permission.
	UnpinMessage(chanId, msgId Snowflake) error

	// AddGroupRecipient adds a recipient to a Group DM using their access token.
	//
	// TODO: does this return anything? API docs unclear.
	AddGroupRecipient(request dio.AddRecipientRequest) error

	// RemoveGroupRecipient removes a recipient from a group DM.
	RemoveGroupRecipient(chanId, userId Snowflake) error

	Messages() MessageAPI
}