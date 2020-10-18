package discord

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/pkg/discord/guild"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
	"time"
)

type Guild interface {
	// ID returns the value of the `id` field currently set on this guild record.
	//
	// The `id` field contains the the guild's Snowflake id.
	ID() dlib.Snowflake
	SetID(id dlib.Snowflake) Guild

	// ID returns the value of the `name` field currently set on this guild record.
	//
	// The `name` field contains the the guild's name.
	Name() guild.Name
	SetName(name guild.Name) Guild

	Icon() comm.ImageHash
	IconIsNull() bool
	SetIcon(hash comm.ImageHash) Guild
	SetNullIcon() Guild

	IconHash() comm.ImageHash
	IconHashIsNull() bool
	IconHashIsSet() bool
	SetIconHash(hash comm.ImageHash) Guild
	SetNullIconHash() Guild
	UnsetIconHash() Guild

	Splash() comm.ImageHash
	SplashIsNull() bool
	SetSplash(hash comm.ImageHash) Guild
	SetNullSplash() Guild

	DiscoverySplash() comm.ImageHash
	DiscoverySplashIsNull() bool
	SetDiscoverySplash(hash comm.ImageHash) Guild
	SetNullDiscoverySplash() Guild

	Owner() bool
	OwnerIsSet() bool
	SetOwner(bool) Guild
	UnsetOwner() Guild

	OwnerID() dlib.Snowflake
	SetOwnerID(id dlib.Snowflake) Guild

	Permissions() comm.Permission
	PermissionsIsSet() bool
	SetPermissions(perm comm.Permission) Guild
	AddPermission(perm comm.Permission) Guild
	RemovePermission(perm comm.Permission) Guild
	PermissionsContain(perm comm.Permission) bool

	Region() string
	SetRegion(region string) Guild

	AFKChannelID() dlib.Snowflake
	AFKChannelIDIsNull() bool
	SetAFKChannelID(id dlib.Snowflake) Guild
	SetNullAFKChannelID() Guild

	AFKTimeout() uint64
	SetAFKTimeout(seconds uint64) Guild

	WidgetEnabled() bool
	WidgetEnabledIsSet() bool
	SetWidgetEnabled(bool) Guild
	UnsetWidgetEnabled() Guild

	WidgetChannelID() dlib.Snowflake
	WidgetChannelIDIsSet() bool
	WidgetChannelIDIsNull() bool
	SetWidgetChannelID(id dlib.Snowflake) Guild
	SetNullWidgetChannelID() Guild
	UnsetWidgetChannelID() Guild

	VerificationLevel() guild.VerificationLevel
	SetVerificationLevel(lvl guild.VerificationLevel) Guild

	DefaultMessageNotifications() guild.MessageNotificationLevel
	SetDefaultMessageNotifications(lvl guild.MessageNotificationLevel) Guild

	ExplicitContentFilter() guild.ExplicitContentFilterLevel
	SetExplicitContentFilter(lvl guild.ExplicitContentFilterLevel) Guild

	Roles() []comm.Role
	SetRoles(roles []comm.Role) Guild

	Emojis() []Emoji
	SetEmojis(emoji []Emoji) Guild

	Features() []guild.Feature
	SetFeatures(feats []guild.Feature) Guild

	MFALevel() guild.MFALevel
	SetMFALevel(lvl guild.MFALevel) Guild

	ApplicationID() dlib.Snowflake
	ApplicationIDIsNull() bool
	SetApplicationID(id dlib.Snowflake) Guild
	SetNullApplicationID() Guild

	SystemChannelID() dlib.Snowflake
	SystemChannelIDIsNull() bool
	SetSystemChannelID(id dlib.Snowflake) Guild
	SetNullSystemChannelID() Guild

	SystemChannelFlags() guild.SystemChannelFlag
	SetSystemChannelFlags(flag guild.SystemChannelFlag) Guild
	AddSystemChannelFlag(flag guild.SystemChannelFlag) Guild
	RemoveSystemChannelFlag(flag guild.SystemChannelFlag) Guild
	SystemChannelFlagsContains(flag guild.SystemChannelFlag) bool

	RulesChannelID() dlib.Snowflake
	RulesChannelIDIsNull() bool
	SetRulesChannelID(id dlib.Snowflake) Guild
	SetNullRulesChannelID() Guild

	JoinedAt() time.Time
	JoinedAtIsSet() bool
	SetJoinedAt(time.Time) Guild
	UnsetJoinedAt() Guild

	Large() bool
	LargeIsSet() bool
	SetLarge(bool) Guild
	UnsetLarge() Guild

	Unavailable() bool
	UnavailableIsSet() bool
	SetUnavailable(bool) Guild
	UnsetUnavailable() Guild

	MemberCount() uint32
	MemberCountIsSet() bool
	SetMemberCount(count uint32) Guild
	UnsetMemberCount() Guild

	VoiceStates() []guild.VoiceState
	VoiceStatesIsSet() bool
	SetVoiceStates([]guild.VoiceState) Guild
	UnsetVoiceStates() Guild

	Members() []guild.Member
	MembersIsSet() bool
	SetMembers([]guild.Member) Guild
	UnsetMembers() Guild

	Channels() []Channel
	ChannelsIsSet() bool
	SetChannels([]Channel) Guild
	UnsetChannels() Guild

	Presences() []guild.PresenceUpdate
	PresencesIsSet() bool
	SetPresences([]guild.PresenceUpdate) Guild
	UnsetPresences() Guild

	MaxPresences() uint32
	MaxPresencesIsSet() bool
	MaxPresencesIsNull() bool
	SetMaxPresences(max uint32) Guild
	SetNullMaxPresences() Guild
	UnsetMaxPresences() Guild

	MaxMembers() uint32
	MaxMembersIsSet() bool
	SetMaxMembers(max uint32) Guild
	UnsetMaxMembers() Guild

	VanityURLCode() string
	VanityURLCodeIsNull() bool
	SetVanityURLCode(code string) Guild
	SetNullVanityURLCode() Guild

	Description() string
	DescriptionIsNull() string
	SetDescription(desc string) Guild
	SetNullDescription() Guild

	Banner() comm.ImageHash
	BannerIsNull() bool
	SetBanner(hash comm.ImageHash) Guild
	SetNullBanner() Guild

	PremiumTier() guild.PremiumTier
	SetPremiumTier(guild.PremiumTier) Guild

	PremiumSubscriptionCount() uint32
	PremiumSubscriptionCountIsSet() bool
	SetPremiumSubscriptionCount(count uint32) Guild
	UnsetPremiumSubscriptionCount() Guild

	PreferredLocale() string
	SetPreferredLocale(locale string) Guild

	PublicUpdatesChannelID() dlib.Snowflake
	PublicUpdatesChannelIDIsNull() bool
	SetPublicUpdatesChannelID(id dlib.Snowflake) Guild
	SetNullPublicUpdatesChannelID() Guild

	MaxVideoChannelUsers() uint32
	MaxVideoChannelUsersIsSet() bool
	SetMaxVideoChannelUsers(max uint32) Guild
	UnsetMaxVideoChannelUsers() Guild

	ApproximateMemberCount() uint32
	ApproximateMemberCountIsSet() bool
	SetApproximateMemberCount(count uint32) Guild
	UnsetApproximateMemberCount() Guild

	ApproximatePresenceCount() uint32
	ApproximatePresenceCountIsSet() bool
	SetApproximatePresenceCount(count uint32) Guild
	UnsetApproximatePresenceCount() Guild
}
