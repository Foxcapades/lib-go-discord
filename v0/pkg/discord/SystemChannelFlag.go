package discord

import "errors"

var (
	ErrBadSystemChannelFlag = errors.New("unrecognized system_channel_flags value")
)

type SystemChannelFlag uint8

const (
	// Suppress member join notifications
	SystemChannelFlagSuppressJoinNotifications SystemChannelFlag = 1 << iota

	// Suppress server boost notifications
	SystemChannelFlagSuppressPremiumSubscriptions
)

func (s SystemChannelFlag) IsValid() bool {
	return nil == s.Validate()
}

func (s SystemChannelFlag) Validate() error {
	if s > 3 {
		return ErrBadSystemChannelFlag
	}

	return nil
}

