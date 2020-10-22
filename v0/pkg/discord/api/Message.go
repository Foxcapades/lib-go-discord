package api

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/api/dio"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/channel"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/guild"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

var (
	ErrNilEmoji = errors.New("passed a nil Emoji to an API request requiring" +
		" an emoji value")
)

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

	// CrosspostMessage crossposts the given message in a News Channel to all
	// subscriber channels.
	//
	// This endpoint requires the 'SEND_MESSAGES' permission, if the current user
	// sent the message, or additionally the 'MANAGE_MESSAGES' permission, for all
	// other messages, to be present for the current user.
	//
	// Returns a message object.
	CrosspostMessage(id dlib.Snowflake) (channel.Message, error)

	// CreateReaction adds a reaction to the message with the given ID.
	//
	// This endpoint requires the 'READ_MESSAGE_HISTORY' permission to be present
	// on the current user. Additionally, if nobody else has reacted to the
	// message using this emoji, this endpoint requires the 'ADD_REACTIONS'
	// permission to be present on the current user.
	CreateReaction(msgId dlib.Snowflake, emoji guild.Emoji) error

	// DeleteOwnReaction delete a reaction the current user has made for the
	// message.
	DeleteOwnReaction(msgId dlib.Snowflake, emoji guild.Emoji) error

	// DeleteUserReaction deletes another user's reaction.
	//
	// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on
	// the current user.
	DeleteUserReaction(msgId, userId dlib.Snowflake, emoji guild.Emoji) error

	// GetReactions returns a list of users that reacted with this emoji.
	GetReactions(query dio.GetReactionsQuery) ([]user.User, error)

	// DeleteAllReactions deletes all reactions on a message.
	//
	// This endpoint requires the 'MANAGE_MESSAGES' permission to be present on
	// the current user.
	//
	// Fires a Message Reaction Remove All Gateway event.
	DeleteAllReactions(id dlib.Snowflake) error

	// DeleteAllReactionsFor deletes all the reactions for the given emoji on a
	// message.
	//
	// This endpoint requires the MANAGE_MESSAGES permission to be present on the
	// current user. Fires a Message Reaction Remove Emoji Gateway event.
	DeleteAllReactionsFor(msgId dlib.Snowflake, emoji guild.Emoji) error

	// EditMessage patches a previously sent message.
	//
	// The fields `content`, `embed`, `allowed_mentions` and flags can be edited
	// by the original message author. Other users can only edit flags and only if
	// they have the MANAGE_MESSAGES permission in the corresponding channel.
	//
	// When specifying flags, ensure to include all previously set flags/bits in
	// addition to ones that you are modifying. Only flags documented in the table
	// below may be modified by users (unsupported flag changes are currently
	// ignored without error).
	//
	// Fires a Message Update Gateway event.
	//
	// Returns a message object.
	EditMessage(id dlib.Snowflake, patch dio.MessagePatch) (channel.Message, error)

	// DeleteMessage deletes an existing message.
	//
	// If operating on a guild channel and trying to delete a message that was not
	// sent by the current user, this endpoint requires the MANAGE_MESSAGES
	// permission.
	//
	// Fires a Message Delete Gateway event.
	DeleteMessage(id dlib.Snowflake) error

	// BulkDeleteMessage deletes multiple messages in a single request.
	//
	// This endpoint can only be used on guild channels and requires the
	// MANAGE_MESSAGES permission.
	//
	// Fires a Message Delete Bulk Gateway event.
	//
	// Any message IDs given that do not exist or are invalid will count towards
	// the minimum and maximum message count (currently 2 and 100 respectively).
	//
	// This endpoint will not delete messages older than 2 weeks, and will fail
	// with a 400 BAD REQUEST if any message provided is older than that or if any
	// duplicate message IDs are provided.
	BulkDeleteMessages([]dlib.Snowflake) error
}