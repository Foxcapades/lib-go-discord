package discord

import (
	"encoding/json"
	"errors"
	"strconv"
)

var (
	ErrBadPermission = errors.New("unrecognized permission flag")
)

// TODO: Should this be in the user package?

// TODO: Document this
type Permission uint32

const (
	// Allows creation of instant invites
	//
	// Channel Type(s): T, V
	PermissionCreateInstantInvite Permission = 1 << iota

	// Allows kicking members
	//
	// Requires 2FA
	PermissionKickMembers

	// Allows banning members
	//
	// Requires 2FA
	PermissionBanMembers

	// Allows all permissions and bypasses channel permission overwrites
	//
	// Requires 2FA
	PermissionAdministrator

	// Allows management and editing of channels
	//
	// Channel Type(s): T, V
	// Requires 2FA
	PermissionManageChannels

	// Allows management and editing of the guild
	//
	// Requires 2FA
	PermissionManageGuild

	// Allows for the addition of reactions to messages
	//
	// Channel Type(s): T
	PermissionAddReactions

	// Allows for viewing of audit logs
	//
	// Channel Type(s): 
	PermissionViewAuditLog

	// Allows for using priority speaker in a voice channel
	//
	// Channel Type(s): V
	PermissionPrioritySpeaker

	// Allows the user to go live
	//
	// Channel Type(s): V
	PermissionStream

	// Allows guild members to view a channel, which includes reading messages in text channels
	//
	// Channel Type(s): T, V
	PermissionViewChannel

	// Allows for sending messages in a channel
	//
	// Channel Type(s): T
	PermissionSendMessages

	// Allows for sending of /tts messages
	//
	// Channel Type(s): T
	PermissionSendTTSMessages

	// Allows for deletion of other users messages
	//
	// Channel Type(s): T
	// Requires 2FA
	PermissionManageMessages
	// Links sent by users with this permission will be auto-embedded
	//
	// Channel Type(s): T
	PermissionEmbedLinks

	// Allows for uploading images and files
	//
	// Channel Type(s): T
	PermissionAttachFiles

	// Allows for reading of message history
	//
	// Channel Type(s): T
	PermissionReadMessageHistory

	// Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
	//
	// Channel Type(s): T
	PermissionMentionEveryone

	// Allows the usage of custom emojis from other servers
	//
	// Channel Type(s): T
	PermissionUseExternalEmojis

	// Allows for viewing guild insights
	//
	// Channel Type(s): 
	PermissionViewGuildInsights

	// Allows for joining of a voice channel
	//
	// Channel Type(s): V
	PermissionConnect

	// Allows for speaking in a voice channel
	//
	// Channel Type(s): V
	PermissionSpeak

	// Allows for muting members in a voice channel
	//
	// Channel Type(s): V
	PermissionMuteMembers

	// Allows for deafening of members in a voice channel
	//
	// Channel Type(s): V
	PermissionDeafenMembers

	// Allows for moving of members between voice channels
	//
	// Channel Type(s): V
	PermissionMoveMembers

	// Allows for using voice-activity-detection in a voice channel
	//
	// Channel Type(s): V
	PermissionUseVAD

	// Allows for modification of own nickname
	//
	// Channel Type(s): 
	PermissionChangeNickname

	// Allows for modification of other users nicknames
	//
	// Channel Type(s): 
	PermissionManageNicknames

	// Allows management and editing of roles
	//
	// Channel Type(s): T, V
	// Requires 2FA
	PermissionManageRoles

	// Allows management and editing of webhooks
	//
	// Channel Type(s): T, V
	// Requires 2FA
	PermissionManageWebhooks

	// Allows management and editing of emojis
	//
	// Requires 2FA
	PermissionManageEmojis
)

func (p Permission) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(p), 10)), nil
}

func (p *Permission) UnmarshalJSON(b []byte) error {
	var tmp string

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if v, err := strconv.ParseUint(tmp, 10, 32); err != nil {
		return err
	} else {
		*p = Permission(v)
	}

	return nil
}

func (p Permission) IsValid() bool {
	return nil == p.Validate()
}

const maxPermission Permission = 2_147_483_647

func (p Permission) Validate() error {
	if p > maxPermission {
		return ErrBadPermission
	}

	return nil
}
