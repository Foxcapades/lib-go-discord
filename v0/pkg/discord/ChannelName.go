package discord

import (
	"errors"
	"github.com/francoispqt/gojay"
)

var (
	ErrBadChannelNameLength = errors.New("channel names must be 2-100" +
		" characters in length")
)

type ChannelName string

func (n ChannelName) IsValid() bool {
	return nil == n.Validate()
}

func (n ChannelName) Validate() error {
	l := len(n)
	if l < 2 || l > 100 {
		return ErrBadChannelNameLength
	}

	return nil
}

func DecodeChannelName(dec *gojay.Decoder) (*ChannelName, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	return (*ChannelName)(tmp), nil
}
