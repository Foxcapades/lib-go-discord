package channel

import "errors"

var (
	ErrBadChannelType = errors.New("unrecognized channel type value")
)

type Type uint8

const (
	// A text channel within a server
	TypeGuildText Type = iota

	// A direct message between users
	TypeDM

	// A voice channel within a server
	TypeGuildVoice

	// A direct message between multiple users
	TypeGroupDM

	// An organizational category that contains up to 50 channels
	TypeGuildCategory

	// A channel that users can follow and crosspost into their own server
	TypeGuildNews

	// A channel in which game developers can sell their game on Discord
	TypeGuildStore
)

func (t Type) IsValid() bool {
	return nil == t.Validate()
}

func (t Type) Validate() error {
	if t > TypeGuildStore {
		return ErrBadChannelType
	}

	return nil
}

