package discord

import (
	"errors"
	"github.com/francoispqt/gojay"
)

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

func DecodeChannelTopic(dec *gojay.Decoder) (*ChannelTopic, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	return (*ChannelTopic)(tmp), nil
}
