package api

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/api/dio"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/channel"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/guild"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type DiscordAPI interface {
	Guilds() GuildAPI

	Channels() ChannelAPI

	// GetVoiceRegions fetches and returns an array of voice region objects that
	// can be used when creating servers.
	GetVoiceRegions() ([]guild.VoiceRegion, error)
}

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
	GetGuildPreview(id dlib.Snowflake) (guild.Preview, error)
}

type ChannelAPI interface {
	// Get a channel by ID.
	//
	// Returns a channel object.
	Get(id dlib.Snowflake) (channel.Channel, error)

	// Patch updates a channel's settings.
	//
	// Requires the MANAGE_CHANNELS permission for the guild.
	//
	// Fires a Channel Update Gateway event. If modifying a category, individual
	// Channel Update events will fire for each child channel that also changes.
	//
	// Returns a channel on success, and a 400 BAD REQUEST on invalid parameters.
	Patch(dlib.Snowflake, dio.ChannelPatch) (channel.Channel, error)

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
	//
	// Returns a channel object on success.
	Delete(id dlib.Snowflake) error

	Messages() MessageAPI
}

type MessageAPI interface {
	// GetMessages returns the messages for a channel.
	//
	// If operating on a guild channel, this endpoint requires the VIEW_CHANNEL
	// permission to be present on the current user. If the current user is
	// missing the 'READ_MESSAGE_HISTORY' permission in the channel then this will
	// return no messages (since they cannot read the message history).
	//
	// Passing a nil value to this method will result in no filters being applied.
	//
	// Returns an array of message objects on success.
	GetMessages(dio.MessageFilter) ([]channel.Message, error)

	// GetMessage Returns a specific message in the channel. If operating on a
	// guild channel, this endpoint requires the 'READ_MESSAGE_HISTORY' permission
	// to be present on the current user.
	//
	// Returns a message object on success.
	GetMessage(id dlib.Snowflake) (channel.Message, error)

	// PostJSONMessage creates a new message for a guild text or DM channel.
	//
	// Before using this endpoint, you must connect to and identify with a gateway
	// at least once.
	//
	// Discord may strip certain characters from message content, like invalid
	// unicode characters or characters which cause unexpected message formatting.
	//
	// If you are passing user-generated strings into message content, consider
	// sanitizing the data to prevent unexpected behavior and utilizing
	// `allowed_mentions` to prevent unexpected mentions.
	//
	// If operating on a guild channel, this endpoint requires the SEND_MESSAGES
	// permission to be present on the current user. If the `tts` field is set to
	// true, the SEND_TTS_MESSAGES permission is required for the message to be
	// spoken.
	//
	// Fires a Message Create Gateway event. See
	// https://discord.com/developers/docs/reference#message-formatting for more
	// information on how to properly format messages.
	//
	// The maximum request size when sending a message is 8MB.
	//
	// Note that when sending application/json you must send at least one of the
	// fields `content` or `embed`.
	//
	// Returns a message object.
	PostJSONMessage(post dio.JSONMessagePost) (channel.Message, error)

	// PostMultipartMessage creates a new message for a guild text or DM channel.
	//
	// Before using this endpoint, you must connect to and identify with a gateway
	// at least once.
	//
	// Discord may strip certain characters from message content, like invalid
	// unicode characters or characters which cause unexpected message formatting.
	//
	// If you are passing user-generated strings into message content, consider
	// sanitizing the data to prevent unexpected behavior and utilizing
	// `allowed_mentions` to prevent unexpected mentions.
	//
	// If operating on a guild channel, this endpoint requires the SEND_MESSAGES
	// permission to be present on the current user. If the `tts` field is set to
	// true, the SEND_TTS_MESSAGES permission is required for the message to be
	// spoken.
	//
	// Fires a Message Create Gateway event. See
	// https://discord.com/developers/docs/reference#message-formatting for more
	// information on how to properly format messages.
	//
	// The maximum request size when sending a message is 8MB.
	//
	// Note that when sending multipart/form-data, you must send at least one of
	// the fields `content`, `embed` or `file`.
	//
	// Returns a message object.
	PostMultipartMessage(post dio.MultipartMessagePost) (channel.Message, error)
}