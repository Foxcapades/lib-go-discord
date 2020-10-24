package discord

// FollowedChannel
// TODO: Document me
type FollowedChannel interface {
	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the current record's source channel id.
	ChannelID() Snowflake

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(Snowflake) FollowedChannel

	// WebhookID returns the current value of this record's `webhook_id` field.
	//
	// The `webhook_id` field contains the current record's created target webhook
	// id.
	WebhookID() Snowflake

	// SetWebhookID overwrites the current value of this record's `webhook_id`
	// field.
	SetWebhookID(Snowflake) FollowedChannel
}
