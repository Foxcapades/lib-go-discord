package discord

import "errors"

var (
	ErrBadMsgType = errors.New("unrecognized message type value")
)

// MessageType
// TODO: Document me
type MessageType uint8

const (
	MsgTypeDefault MessageType = iota
	MsgTypeRecipientAdd
	MsgTypeRecipientRemove
	MsgTypeCall
	MsgTypeChannelNameChange
	MsgTypeChannelIconChange
	MsgTypeChannelPinnedMessage
	MsgTypeGuildMemberJoin
	MsgTypeUserPremiumGuildSubscription
	MsgTypeUserPremiumGuildSubscriptionTier1
	MsgTypeUserPremiumGuildSubscriptionTier2
	MsgTypeUserPremiumGuildSubscriptionTier3
	MsgTypeChannelFollowAdd
	MsgTypeGuildDiscoveryDisqualified MessageType = 1 + iota
	MsgTypeGuildDiscoveryRequalified
)

func (m MessageType) IsValid() bool {
	return nil == m.Validate()
}

func (m MessageType) Validate() error {
	if m > 15 || m == 13 {
		return ErrBadMsgType
	}

	return nil
}

