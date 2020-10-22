package discord

import "errors"

var (
	ErrBadWebhookType = errors.New("unrecognized webhook type value")
)

type WebhookType uint8

const (
	WebhookTypeIncoming        WebhookType = 1
	WebhookTypeChannelFollower WebhookType = 2
)

func (t WebhookType) IsValid() bool {
	return nil == t.Validate()
}

func (t WebhookType) Validate() error {
	if t < 1 || t > 2 {
		return ErrBadWebhookType
	}

	return nil
}

