package comm

import "github.com/foxcapades/lib-go-discord/pkg/dlib"

// FollowedChannel
// TODO: Document me
type FollowedChannel interface {
	// source channel id
	ChannelID() dlib.Snowflake
	SetChannelID(dlib.Snowflake) FollowedChannel

	// created target webhook id
	WebhookID() dlib.Snowflake
	SetWebhookID(dlib.Snowflake) FollowedChannel
}
