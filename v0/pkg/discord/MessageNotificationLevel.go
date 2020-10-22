package discord

import "errors"

var (
	ErrBadMsgNoteLevel = errors.New("unrecognized message notification level value")
)

type MessageNotificationLevel uint8

const (
	MsgNoteLvlAllMessages MessageNotificationLevel = iota
	MsgNoteLvlOnlyMentions
)

func (m MessageNotificationLevel) IsValid() bool {
	return nil == m.Validate()
}

func (m MessageNotificationLevel) Validate() error {
	if m > MsgNoteLvlOnlyMentions {
		return ErrBadMsgNoteLevel
	}

	return nil
}
