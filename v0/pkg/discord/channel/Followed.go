package channel

import "github.com/foxcapades/lib-go-discord/v0/pkg/dlib"

type Followed interface {
	// ChannelID returns the current value of this record's `channel_id` field.
	//
	// The `channel_id` field contains the source channel id.
	ChannelID() dlib.Snowflake

	// SetChannelID overwrites the current value of this record's `channel_id`
	// field.
	SetChannelID(dlib.Snowflake) Followed

	// WebhookID returns the current value of this record's `webhook_id` field.
	//
	// The `webhook_id` field contains the created target webhook id.
	WebhookID() dlib.Snowflake

	// SetWebhookID overwrites the current value of this record's `webhook_id`
	// field.
	SetWebhookID(dlib.Snowflake) Followed
}
