package discord

import "errors"

var (
	ErrBadTopicLength = errors.New("channel topics cannot be longer than 1024 characters")
)

type ChannelTopic string

func (t ChannelTopic) IsValid() bool {
	return nil == t.Validate()
}

func (t ChannelTopic) Validate() error {
	if len(t) > 1024 {
		return ErrBadTopicLength
	}

	return nil
}


