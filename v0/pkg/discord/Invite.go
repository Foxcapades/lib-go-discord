package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/channel"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/invite"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/user"
)

type Invite interface {
	// Code returns the current value of this record's `code` field.
	//
	// The `code` field contains the invite code (unique ID).
	Code() string

	// SetCode overwrites the current value of this record's `code` field.
	SetCode(string) Invite

	// Guild returns the current value of this record's `guild` field.
	//
	// The `guild` field contains the guild this invite is for.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use GuildIsSet to check if the field is present before use.
	Guild() Guild

	// GuildIsSet returns whether this record's `guild` field is currently
	// present.
	GuildIsSet() bool

	// SetGuild overwrites the current value of this record's `guild` field.
	SetGuild(Guild) Invite

	// UnsetGuild removes this record's `guild` field.
	UnsetGuild() Invite

	// Channel returns the current value of this record's `channel` field.
	//
	// The `channel` field contains the channel this invite is for.
	Channel() channel.Channel

	// SetChannel overwrites the current value of this record's `channel` field.
	SetChannel(channel.Channel) Invite

	// Inviter returns the current value of this record's `inviter` field.
	//
	// The `inviter` field contains the user who created the invite.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use InviterIsSet to check if the field is present before use.
	Inviter() user.User

	// InviterIsSet returns whether this record's `inviter` field is currently
	// present.
	InviterIsSet() bool

	// SetInviter overwrites the current value of this record's `inviter` field.
	SetInviter(user.User) Invite

	// UnsetInviter removes this record's `inviter` field.
	UnsetInviter() Invite

	// TargetUser returns the current value of this record's `target_user` field.
	//
	// The `target_user` field contains the target user for this invite.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TargetUserIsSet to check if the field is present before use.
	TargetUser() user.User

	// TargetUserIsSet returns whether this record's `target_user` field is
	// currently present.
	TargetUserIsSet() bool

	// SetTargetUser overwrites the current value of this record's `target_user`
	// field.
	SetTargetUser(user.User) Invite

	// UnsetTargetUser removes this record's `target_user` field.
	UnsetTargetUser() Invite

	// TargetUserType returns the current value of this record's
	// `target_user_type` field.
	//
	// The `target_user_type` field contains the type of user target for this
	// invite.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TargetUserTypeIsSet to check if the field is present before use.
	TargetUserType() invite.TargetUserType

	// TargetUserTypeIsSet returns whether this record's `target_user_type` field
	// is currently present.
	TargetUserTypeIsSet() bool

	// SetTargetUserType overwrites the current value of this record's
	// `target_user_type` field.
	SetTargetUserType(invite.TargetUserType) Invite

	// UnsetTargetUserType removes this record's `target_user_type` field.
	UnsetTargetUserType() Invite

	// ApproximatePresenceCount returns the current value of this record's
	// `approximate_presence_count` field.
	//
	// The `approximate_presence_count` field contains the approximate count of
	// online members (only present when target_user is set).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ApproximatePresenceCountIsSet to check if the field is present before
	// use.
	ApproximatePresenceCount() uint32

	// ApproximatePresenceCountIsSet returns whether this record's
	// `approximate_presence_count` field is currently present.
	ApproximatePresenceCountIsSet() bool

	// SetApproximatePresenceCount overwrites the current value of this record's
	// `approximate_presence_count` field.
	SetApproximatePresenceCount(uint32) Invite

	// UnsetApproximatePresenceCount removes this record's
	// `approximate_presence_count` field.
	UnsetApproximatePresenceCount() Invite

	// ApproximateMemberCount returns the current value of this record's
	// `approximate_member_count` field.
	//
	// The `approximate_member_count` field contains the approximate count of
	// total members.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ApproximateMemberCountIsSet to check if the field is present before
	// use.
	ApproximateMemberCount() uint32

	// ApproximateMemberCountIsSet returns whether this record's
	// `approximate_member_count` field is currently present.
	ApproximateMemberCountIsSet() bool

	// SetApproximateMemberCount overwrites the current value of this record's
	// `approximate_member_count` field.
	SetApproximateMemberCount(uint32) Invite

	// UnsetApproximateMemberCount removes this record's
	// `approximate_member_count` field.
	UnsetApproximateMemberCount() Invite
}
