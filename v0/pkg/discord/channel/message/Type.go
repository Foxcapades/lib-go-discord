package message

import "errors"

var (
	ErrBadMsgType = errors.New("unrecognized message type value")
)

// Type
// TODO: Document me
type Type uint8

const (
	MsgTypeDefault Type = iota
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
	MsgTypeGuildDiscoveryDisqualified Type = 1 + iota
	MsgTypeGuildDiscoveryRequalified
)

func (m Type) IsValid() bool {
	return nil == m.Validate()
}

func (m Type) Validate() error {
	if m > 15 || m == 13 {
		return ErrBadMsgType
	}

	return nil
}

