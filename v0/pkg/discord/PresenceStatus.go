package discord

import "errors"

var (
	ErrBadPresenceStatus = errors.New("unrecognized presence status value")
)

type PresenceStatus string

const (
	PresenceStatusIdle         PresenceStatus = "idle"
	PresenceStatusDoNotDisturb PresenceStatus = "dnd"
	PresenceStatusOnline       PresenceStatus = "online"
	PresenceStatusOffline      PresenceStatus = "offline"
)

func (p PresenceStatus) IsValid() bool {
	return nil == p.Validate()
}

func (p PresenceStatus) Validate() error {
	switch p {
	case PresenceStatusIdle:
		fallthrough
	case PresenceStatusDoNotDisturb:
		fallthrough
	case PresenceStatusOnline:
		fallthrough
	case PresenceStatusOffline:
		return nil
	default:
		return ErrBadPresenceStatus
	}
}

