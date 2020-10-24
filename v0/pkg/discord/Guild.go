package discord

import (
	"time"
)

type Guild interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the guild id.
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) Guild

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the guild name (2-100 characters, excluding
	// trailing and leading whitespace).
	Name() GuildName

	// SetName overwrites the current value of this record's `name` field.
	SetName(GuildName) Guild

	// Icon returns the current value of this record's `icon` field.
	//
	// The `icon` field contains the icon hash.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use IconIsNull to check if the field is null before use.
	Icon() ImageHash

	// IconIsNull returns whether this record's `icon` field is currently null.
	IconIsNull() bool

	// SetIcon overwrites the current value of this record's `icon` field.
	SetIcon(ImageHash) Guild

	// SetNullIcon overwrites the current value of this record's `icon` field
	// with `null`.
	SetNullIcon() Guild

	// IconHash returns the current value of this record's `icon_hash` field.
	//
	// The `icon_hash` field contains the icon hash, returned when in the template
	// object.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use IconHashIsNull to check if the field is null before use.
	IconHash() ImageHash

	// IconHashIsNull returns whether this record's `icon_hash` field is currently
	// null.
	IconHashIsNull() bool

	// SetIconHash overwrites the current value of this record's `icon_hash`
	// field.
	SetIconHash(ImageHash) Guild

	// SetNullIconHash overwrites the current value of this record's `icon_hash`
	// field with `null`.
	SetNullIconHash() Guild

	// Splash returns the current value of this record's `splash` field.
	//
	// The `splash` field contains the splash hash.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use SplashIsNull to check if the field is null before use.
	Splash() ImageHash

	// SplashIsNull returns whether this record's `splash` field is currently
	// null.
	SplashIsNull() bool

	// SetSplash overwrites the current value of this record's `splash` field.
	SetSplash(ImageHash) Guild

	// SetNullSplash overwrites the current value of this record's `splash` field
	// with `null`.
	SetNullSplash() Guild

	// DiscoverySplash returns the current value of this record's
	// `discovery_splash` field.
	//
	// The `discovery_splash` field contains the discovery splash hash; only
	// present for guilds with the "DISCOVERABLE" feature.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use DiscoverySplashIsNull to check if the field is null before use.
	DiscoverySplash() ImageHash

	// DiscoverySplashIsNull returns whether this record's `discovery_splash`
	// field is currently null.
	DiscoverySplashIsNull() bool

	// SetDiscoverySplash overwrites the current value of this record's
	// `discovery_splash` field.
	SetDiscoverySplash(ImageHash) Guild

	// SetNullDiscoverySplash overwrites the current value of this record's
	// `discovery_splash` field with `null`.
	SetNullDiscoverySplash() Guild

	// Owner returns the current value of this record's `owner` field.
	//
	// The `owner` field indicates whether the user is the owner of the guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use OwnerIsSet to check if the field is present before use.
	Owner() bool

	// OwnerIsSet returns whether this record's `owner` field is currently
	// present.
	OwnerIsSet() bool

	// SetOwner overwrites the current value of this record's `owner` field.
	SetOwner(bool) Guild

	// UnsetOwner removes this record's `owner` field.
	UnsetOwner() Guild

	// OwnerID returns the current value of this record's `owner_id` field.
	//
	// The `owner_id` field contains the id of the owner.
	OwnerID() Snowflake

	// SetOwnerID overwrites the current value of this record's `owner_id` field.
	SetOwnerID(Snowflake) Guild

	// Permissions returns the current value of this record's `permissions` field.
	//
	// The `permissions` field contains the total permissions for the user in the
	// guild (excludes overrides).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PermissionsIsSet to check if the field is present before use.
	Permissions() Permission

	// PermissionsIsSet returns whether this record's `permissions` field is
	// currently present.
	PermissionsIsSet() bool

	// SetPermissions overwrites the current value of this record's `permissions`
	// field.
	SetPermissions(Permission) Guild

	// UnsetPermissions removes this record's `permissions` field.
	UnsetPermissions() Guild

	// Region returns the current value of this record's `region` field.
	//
	// The `region` field contains the voice region id for the guild.
	Region() string

	// SetRegion overwrites the current value of this record's `region` field.
	SetRegion(string) Guild

	// AFKChannelID returns the current value of this record's `afk_channel_id`
	// field.
	//
	// The `afk_channel_id` field contains the id of afk the channel.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use AFKChannelIDIsNull to check if the field is null before use.
	AFKChannelID() Snowflake

	// AFKChannelIDIsNull returns whether this record's `afk_channel_id` field is
	// currently null.
	AFKChannelIDIsNull() bool

	// SetAFKChannelID overwrites the current value of this record's
	// `afk_channel_id` field.
	SetAFKChannelID(Snowflake) Guild

	// SetNullAFKChannelID overwrites the current value of this record's
	// `afk_channel_id` field with `null`.
	SetNullAFKChannelID() Guild

	// AFKTimeout returns the current value of this record's `afk_timeout` field.
	//
	// The `afk_timeout` field contains the afk timeout in seconds.
	AFKTimeout() uint64

	// SetAFKTimeout overwrites the current value of this record's `afk_timeout`
	// field.
	SetAFKTimeout(uint64) Guild

	// WidgetEnabled returns the current value of this record's `widget_enabled`
	// field.
	//
	// The `widget_enabled` field indicates true if the server widget is enabled.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use WidgetEnabledIsSet to check if the field is present before use.
	WidgetEnabled() bool

	// WidgetEnabledIsSet returns whether this record's `widget_enabled` field is
	// currently present.
	WidgetEnabledIsSet() bool

	// SetWidgetEnabled overwrites the current value of this record's
	// `widget_enabled` field.
	SetWidgetEnabled(bool) Guild

	// UnsetWidgetEnabled removes this record's `widget_enabled` field.
	UnsetWidgetEnabled() Guild

	// WidgetChannelID returns the current value of this record's
	// `widget_channel_id` field.
	//
	// The `widget_channel_id` field contains the the channel id that the widget
	// will generate an invite to, or null if set to no invite.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use WidgetChannelIDIsReadable to check if the
	// field is present and non-null before use.
	WidgetChannelID() Snowflake

	// WidgetChannelIDIsNull returns whether this record's `widget_channel_id`
	// field is currently null.
	WidgetChannelIDIsNull() bool

	// WidgetChannelIDIsSet returns whether this record's `widget_channel_id`
	// field is currently present.
	WidgetChannelIDIsSet() bool

	// WidgetChannelIDIsReadable returns whether this record's `widget_channel_id`
	// field is currently set and non-null.
	WidgetChannelIDIsReadable() bool

	// SetWidgetChannelID overwrites the current value of this record's
	// `widget_channel_id` field.
	SetWidgetChannelID(Snowflake) Guild

	// SetNullWidgetChannelID overwrites the current value of this record's
	// `widget_channel_id` field with `null`.
	SetNullWidgetChannelID() Guild

	// UnsetWidgetChannelID removes this record's `widget_channel_id` field.
	UnsetWidgetChannelID() Guild

	// VerificationLevel returns the current value of this record's
	// `verification_level` field.
	//
	// The `verification_level` field contains the verification level required for
	// the guild.
	VerificationLevel() VerificationLevel

	// SetVerificationLevel overwrites the current value of this record's
	// `verification_level` field.
	SetVerificationLevel(VerificationLevel) Guild

	// DefaultMessageNotifications returns the current value of this record's
	// `default_message_notifications` field.
	//
	// The `default_message_notifications` field contains the default message
	// notifications level.
	DefaultMessageNotifications() MessageNotificationLevel

	// SetDefaultMessageNotifications overwrites the current value of this
	// record's `default_message_notifications` field.
	SetDefaultMessageNotifications(MessageNotificationLevel) Guild

	// ExplicitContentFilter returns the current value of this record's
	// `explicit_content_filter` field.
	//
	// The `explicit_content_filter` field contains the explicit content filter
	// level.
	ExplicitContentFilter() ExplicitContentFilterLevel

	// SetExplicitContentFilter overwrites the current value of this record's
	// `explicit_content_filter` field.
	SetExplicitContentFilter(ExplicitContentFilterLevel) Guild

	// Roles returns the current value of this record's `roles` field.
	//
	// The `roles` field contains the roles in the guild.
	Roles() []Role

	// SetRoles overwrites the current value of this record's `roles` field.
	SetRoles([]Role) Guild

	// Emojis returns the current value of this record's `emojis` field.
	//
	// The `emojis` field contains the custom guild emojis.
	Emojis() []Emoji

	// SetEmojis overwrites the current value of this record's `emojis` field.
	SetEmojis([]Emoji) Guild

	// Features returns the current value of this record's `features` field.
	//
	// The `features` field contains the enabled guild features.
	Features() []GuildFeature

	// SetFeatures overwrites the current value of this record's `features` field.
	SetFeatures([]GuildFeature) Guild

	// MFALevel returns the current value of this record's `mfa_level` field.
	//
	// The `mfa_level` field contains the required MFA level for the guild.
	MFALevel() MFALevel

	// SetMFALevel overwrites the current value of this record's `mfa_level`
	// field.
	SetMFALevel(MFALevel) Guild

	// ApplicationID returns the current value of this record's `application_id`
	// field.
	//
	// The `application_id` field contains the application id of the guild creator
	// if it is bot-created.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use ApplicationIDIsNull to check if the field is null before use.
	ApplicationID() Snowflake

	// ApplicationIDIsNull returns whether this record's `application_id` field is
	// currently null.
	ApplicationIDIsNull() bool

	// SetApplicationID overwrites the current value of this record's
	// `application_id` field.
	SetApplicationID(Snowflake) Guild

	// SetNullApplicationID overwrites the current value of this record's
	// `application_id` field with `null`.
	SetNullApplicationID() Guild

	// SystemChannelID returns the current value of this record's
	// `system_channel_id` field.
	//
	// The `system_channel_id` field contains the the id of the channel where
	// guild notices such as welcome messages and boost events are posted.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use SystemChannelIDIsNull to check if the field is null before use.
	SystemChannelID() Snowflake

	// SystemChannelIDIsNull returns whether this record's `system_channel_id`
	// field is currently null.
	SystemChannelIDIsNull() bool

	// SetSystemChannelID overwrites the current value of this record's
	// `system_channel_id` field.
	SetSystemChannelID(Snowflake) Guild

	// SetNullSystemChannelID overwrites the current value of this record's
	// `system_channel_id` field with `null`.
	SetNullSystemChannelID() Guild

	// SystemChannelFlags returns the current value of this record's
	// `system_channel_flags` field.
	//
	// The `system_channel_flags` field contains the system channel flags.
	SystemChannelFlags() SystemChannelFlag

	// SetSystemChannelFlags overwrites the current value of this record's
	// `system_channel_flags` field.
	SetSystemChannelFlags(SystemChannelFlag) Guild

	// AddSystemChannelFlag adds the given flag to this record's current
	// `system_channel_flags` value.
	AddSystemChannelFlag(SystemChannelFlag) Guild

	// RemoveSystemChannelFlag removes the given flag to this record's current
	// `system_channel_flags` value.
	RemoveSystemChannelFlag(SystemChannelFlag) Guild

	// SystemChannelFlagsContains returns whether the given flag is already in
	// this record's `system_channel_flags` value.
	SystemChannelFlagsContains(SystemChannelFlag) bool

	// RulesChannelID returns the current value of this record's
	// `rules_channel_id` field.
	//
	// The `rules_channel_id` field contains the the id of the channel where
	// Community guilds can display rules and/or guidelines.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use RulesChannelIDIsNull to check if the field is null before use.
	RulesChannelID() Snowflake

	// RulesChannelIDIsNull returns whether this record's `rules_channel_id` field
	// is currently null.
	RulesChannelIDIsNull() bool

	// SetRulesChannelID overwrites the current value of this record's
	// `rules_channel_id` field.
	SetRulesChannelID(Snowflake) Guild

	// SetNullRulesChannelID overwrites the current value of this record's
	// `rules_channel_id` field with `null`.
	SetNullRulesChannelID() Guild

	// JoinedAt returns the current value of this record's `joined_at` field.
	//
	// The `joined_at` field contains when this guild was joined at.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use JoinedAtIsSet to check if the field is present before use.
	JoinedAt() time.Time

	// JoinedAtIsSet returns whether this record's `joined_at` field is currently
	// present.
	JoinedAtIsSet() bool

	// SetJoinedAt overwrites the current value of this record's `joined_at`
	// field.
	SetJoinedAt(time.Time) Guild

	// UnsetJoinedAt removes this record's `joined_at` field.
	UnsetJoinedAt() Guild

	// Large returns the current value of this record's `large` field.
	//
	// The `large` field indicates if this is considered a large guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use LargeIsSet to check if the field is present before use.
	Large() bool

	// LargeIsSet returns whether this record's `large` field is currently
	// present.
	LargeIsSet() bool

	// SetLarge overwrites the current value of this record's `large` field.
	SetLarge(bool) Guild

	// UnsetLarge removes this record's `large` field.
	UnsetLarge() Guild

	// Unavailable returns the current value of this record's `unavailable` field.
	//
	// The `unavailable` field indicates true if this guild is unavailable due to
	// an outage.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use UnavailableIsSet to check if the field is present before use.
	Unavailable() bool

	// UnavailableIsSet returns whether this record's `unavailable` field is
	// currently present.
	UnavailableIsSet() bool

	// SetUnavailable overwrites the current value of this record's `unavailable`
	// field.
	SetUnavailable(bool) Guild

	// UnsetUnavailable removes this record's `unavailable` field.
	UnsetUnavailable() Guild

	// MemberCount returns the current value of this record's `member_count`
	// field.
	//
	// The `member_count` field contains the total number of members in this
	// guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MemberCountIsSet to check if the field is present before use.
	MemberCount() uint32

	// MemberCountIsSet returns whether this record's `member_count` field is
	// currently present.
	MemberCountIsSet() bool

	// SetMemberCount overwrites the current value of this record's `member_count`
	// field.
	SetMemberCount(uint32) Guild

	// UnsetMemberCount removes this record's `member_count` field.
	UnsetMemberCount() Guild

	// VoiceStates returns the current value of this record's `voice_states`
	// field.
	//
	// The `voice_states` field contains the states of members currently in voice
	// channels; lacks the guild_id key.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use VoiceStatesIsSet to check if the field is present before use.
	VoiceStates() []VoiceState

	// VoiceStatesIsSet returns whether this record's `voice_states` field is
	// currently present.
	VoiceStatesIsSet() bool

	// SetVoiceStates overwrites the current value of this record's `voice_states`
	// field.
	SetVoiceStates([]VoiceState) Guild

	// UnsetVoiceStates removes this record's `voice_states` field.
	UnsetVoiceStates() Guild

	// Members returns the current value of this record's `members` field.
	//
	// The `members` field contains the users in the guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MembersIsSet to check if the field is present before use.
	Members() GuildMember

	// MembersIsSet returns whether this record's `members` field is currently
	// present.
	MembersIsSet() bool

	// SetMembers overwrites the current value of this record's `members` field.
	SetMembers(GuildMember) Guild

	// UnsetMembers removes this record's `members` field.
	UnsetMembers() Guild

	// Channels returns the current value of this record's `channels` field.
	//
	// The `channels` field contains the channels in the guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ChannelsIsSet to check if the field is present before use.
	Channels() []Channel

	// ChannelsIsSet returns whether this record's `channels` field is currently
	// present.
	ChannelsIsSet() bool

	// SetChannels overwrites the current value of this record's `channels` field.
	SetChannels([]Channel) Guild

	// UnsetChannels removes this record's `channels` field.
	UnsetChannels() Guild

	// Presences returns the current value of this record's `presences` field.
	//
	// The `presences` field contains the presences of the members in the guild,
	// will only include non-offline members if the size is greater than large
	// threshold.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PresencesIsSet to check if the field is present before use.
	Presences() []PresenceUpdate

	// PresencesIsSet returns whether this record's `presences` field is currently
	// present.
	PresencesIsSet() bool

	// SetPresences overwrites the current value of this record's `presences`
	// field.
	SetPresences([]PresenceUpdate) Guild

	// UnsetPresences removes this record's `presences` field.
	UnsetPresences() Guild

	// MaxPresences returns the current value of this record's `max_presences`
	// field.
	//
	// The `max_presences` field contains the maximum number of presences for the
	// guild (the default value, currently 25000, is in effect when null is
	// returned).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use MaxPresencesIsReadable to check if the field
	// is present and non-null before use.
	MaxPresences() uint32

	// MaxPresencesIsNull returns whether this record's `max_presences` field is
	// currently null.
	MaxPresencesIsNull() bool

	// MaxPresencesIsSet returns whether this record's `max_presences` field is
	// currently present.
	MaxPresencesIsSet() bool

	// MaxPresencesIsReadable returns whether this record's `max_presences` field
	// is currently set
	// and non-null.
	MaxPresencesIsReadable() bool

	// SetMaxPresences overwrites the current value of this record's
	// `max_presences` field.
	SetMaxPresences(uint32) Guild

	// SetNullMaxPresences overwrites the current value of this record's
	// `max_presences` field with `null`.
	SetNullMaxPresences() Guild

	// UnsetMaxPresences removes this record's `max_presences` field.
	UnsetMaxPresences() Guild

	// MaxMembers returns the current value of this record's `max_members` field.
	//
	// The `max_members` field contains the the maximum number of members for the
	// guild.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MaxMembersIsSet to check if the field is present before use.
	MaxMembers() uint32

	// MaxMembersIsSet returns whether this record's `max_members` field is
	// currently present.
	MaxMembersIsSet() bool

	// SetMaxMembers overwrites the current value of this record's `max_members`
	// field.
	SetMaxMembers(uint32) Guild

	// UnsetMaxMembers removes this record's `max_members` field.
	UnsetMaxMembers() Guild

	// VanityURLCode returns the current value of this record's `vanity_url_code`
	// field.
	//
	// The `vanity_url_code` field contains the the vanity url code for the guild.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use VanityURLCodeIsNull to check if the field is null before use.
	VanityURLCode() string

	// VanityURLCodeIsNull returns whether this record's `vanity_url_code` field
	// is currently null.
	VanityURLCodeIsNull() bool

	// SetVanityURLCode overwrites the current value of this record's
	// `vanity_url_code` field.
	SetVanityURLCode(string) Guild

	// SetNullVanityURLCode overwrites the current value of this record's
	// `vanity_url_code` field with `null`.
	SetNullVanityURLCode() Guild

	// Description returns the current value of this record's `description` field.
	//
	// The `description` field contains the the description for the guild, if the
	// guild is discoverable.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use DescriptionIsNull to check if the field is null before use.
	Description() string

	// DescriptionIsNull returns whether this record's `description` field is
	// currently null.
	DescriptionIsNull() bool

	// SetDescription overwrites the current value of this record's `description`
	// field.
	SetDescription(string) Guild

	// SetNullDescription overwrites the current value of this record's
	// `description` field with `null`.
	SetNullDescription() Guild

	// Banner returns the current value of this record's `banner` field.
	//
	// The `banner` field contains the banner hash.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use BannerIsNull to check if the field is null before use.
	Banner() ImageHash

	// BannerIsNull returns whether this record's `banner` field is currently
	// null.
	BannerIsNull() bool

	// SetBanner overwrites the current value of this record's `banner` field.
	SetBanner(ImageHash) Guild

	// SetNullBanner overwrites the current value of this record's `banner` field
	// with `null`.
	SetNullBanner() Guild

	// PremiumTier returns the current value of this record's `premium_tier`
	// field.
	//
	// The `premium_tier` field contains the premium tier (Server Boost level).
	PremiumTier() PremiumTier

	// SetPremiumTier overwrites the current value of this record's `premium_tier`
	// field.
	SetPremiumTier(PremiumTier) Guild

	// PremiumSubscriptionCount returns the current value of this record's
	// `premium_subscription_count` field.
	//
	// The `premium_subscription_count` field contains the the number of boosts
	// this guild currently has.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PremiumSubscriptionCountIsSet to check if the field is present before
	// use.
	PremiumSubscriptionCount() uint16

	// PremiumSubscriptionCountIsSet returns whether this record's
	// `premium_subscription_count` field is currently present.
	PremiumSubscriptionCountIsSet() bool

	// SetPremiumSubscriptionCount overwrites the current value of this record's
	// `premium_subscription_count` field.
	SetPremiumSubscriptionCount(uint16) Guild

	// UnsetPremiumSubscriptionCount removes this record's
	// `premium_subscription_count` field.
	UnsetPremiumSubscriptionCount() Guild

	// PreferredLocale returns the current value of this record's
	// `preferred_locale` field.
	//
	// The `preferred_locale` field contains the the preferred locale of a
	// Community guild; used in server discovery and notices from Discord;
	// defaults to "en-US".
	PreferredLocale() string

	// SetPreferredLocale overwrites the current value of this record's
	// `preferred_locale` field.
	SetPreferredLocale(string) Guild

	// PublicUpdatesChannelID returns the current value of this record's
	// `public_updates_channel_id` field.
	//
	// The `public_updates_channel_id` field contains the id of the channel
	// where admins and moderators of Community guilds receive notices from
	// Discord.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use PublicUpdatesChannelIDIsNull to check if the field is null
	// before use.
	PublicUpdatesChannelID() Snowflake

	// PublicUpdatesChannelIDIsNull returns whether this record's
	// `public_updates_channel_id` field is currently null.
	PublicUpdatesChannelIDIsNull() bool

	// SetPublicUpdatesChannelID overwrites the current value of this record's
	// `public_updates_channel_id` field.
	SetPublicUpdatesChannelID(Snowflake) Guild

	// SetNullPublicUpdatesChannelID overwrites the current value of this record's
	// `public_updates_channel_id` field with `null`.
	SetNullPublicUpdatesChannelID() Guild

	// MaxVideoChannelUsers returns the current value of this record's
	// `max_video_channel_users` field.
	//
	// The `max_video_channel_users` field contains the maximum amount of
	// users in a video channel.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MaxVideoChannelUsersIsSet to check if the field is present before use.
	MaxVideoChannelUsers() uint16

	// MaxVideoChannelUsersIsSet returns whether this record's
	// `max_video_channel_users` field is currently present.
	MaxVideoChannelUsersIsSet() bool

	// SetMaxVideoChannelUsers overwrites the current value of this record's
	// `max_video_channel_users` field.
	SetMaxVideoChannelUsers(uint16) Guild

	// UnsetMaxVideoChannelUsers removes this record's `max_video_channel_users`
	// field.
	UnsetMaxVideoChannelUsers() Guild

	// ApproximateMemberCount returns the current value of this record's
	// `approximate_member_count` field.
	//
	// The `approximate_member_count` field contains the approximate number of
	// members in this guild, returned from the GET /guild/<id> endpoint when
	// with_counts is true.
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
	SetApproximateMemberCount(uint32) Guild

	// UnsetApproximateMemberCount removes this record's
	// `approximate_member_count` field.
	UnsetApproximateMemberCount() Guild

	// ApproximatePresenceCount returns the current value of this record's
	// `approximate_presence_count` field.
	//
	// The `approximate_presence_count` field contains the approximate number of
	// non-offline members in this guild, returned from the GET /guild/<id>
	// endpoint when with_counts is true.
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
	SetApproximatePresenceCount(uint32) Guild

	// UnsetApproximatePresenceCount removes this record's
	// `approximate_presence_count` field.
	UnsetApproximatePresenceCount() Guild
}
