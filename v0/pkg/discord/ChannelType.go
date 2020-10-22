package discord

import "errors"

var (
	ErrBadChannelType = errors.New("unrecognized channel type value")
)

type ChannelType uint8

const (
	// A text channel within a server
	ChannelTypeGuildText ChannelType = iota

	// A direct message between users
	ChannelTypeDM

	// A voice channel within a server
	ChannelTypeGuildVoice

	// A direct message between multiple users
	ChannelTypeGroupDM

	// An organizational category that contains up to 50 channels
	ChannelTypeGuildCategory

	// A channel that users can follow and crosspost into their own server
	ChannelTypeGuildNews

	// A channel in which game developers can sell their game on Discord
	ChannelTypeGuildStore
)

func (t ChannelType) IsValid() bool {
	return nil == t.Validate()
}

func (t ChannelType) Validate() error {
	if t > ChannelTypeGuildStore {
		return ErrBadChannelType
	}

	return nil
}

