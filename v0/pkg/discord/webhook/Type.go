package webhook

import "errors"

var (
	ErrBadWebhookType = errors.New("unrecognized webhook type value")
)

type Type uint8

const (
	TypeIncoming        Type = 1
	TypeChannelFollower Type = 2
)

func (t Type) IsValid() bool {
	return nil == t.Validate()
}

func (t Type) Validate() error {
	if t < 1 || t > 2 {
		return ErrBadWebhookType
	}

	return nil
}

