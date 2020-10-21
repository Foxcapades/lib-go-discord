package dio

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm/image"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/guild"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type GuildPatch interface {
	// Name returns the current value of this request's `name` field.
	//
	// The `name` field contains the guild name.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NameIsSet to check if the field is present before use.
	Name() string

	// NameIsSet returns whether this request's `name` field is currently present.
	NameIsSet() bool

	// SetName overwrites the current value of this request's `name` field.
	SetName(string) GuildPatch

	// UnsetName removes this request's `name` field.
	UnsetName() GuildPatch

	// Region returns the current value of this request's `region` field.
	//
	// The `region` field contains the guild voice region id.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use RegionIsReadable to check if the field is
	// present and non-null before use.
	Region() string

	// RegionIsNull returns whether this request's `region` field is currently
	// null.
	RegionIsNull() bool

	// RegionIsSet returns whether this request's `region` field is currently
	// present.
	RegionIsSet() bool

	// RegionIsReadable returns whether this request's `region` field is currently
	// set and non-null.
	RegionIsReadable() bool

	// SetRegion overwrites the current value of this request's `region` field.
	SetRegion(string) GuildPatch

	// SetNullRegion overwrites the current value of this request's `region` field
	// with `null`.
	SetNullRegion() GuildPatch

	// UnsetRegion removes this request's `region` field.
	UnsetRegion() GuildPatch

	// VerificationLevel returns the current value of this request's
	// `verification_level` field.
	//
	// The `verification_level` field contains the verification level.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use VerificationLevelIsReadable to check if the
	// field is present and non-null before use.
	VerificationLevel() guild.VerificationLevel

	// VerificationLevelIsNull returns whether this request's `verification_level`
	// field is currently null.
	VerificationLevelIsNull() bool

	// VerificationLevelIsSet returns whether this request's `verification_level`
	// field is currently present.
	VerificationLevelIsSet() bool

	// VerificationLevelIsReadable returns whether this request's
	// `verification_level` field is currently set and non-null.
	VerificationLevelIsReadable() bool

	// SetVerificationLevel overwrites the current value of this request's
	// `verification_level` field.
	SetVerificationLevel(guild.VerificationLevel) GuildPatch

	// SetNullVerificationLevel overwrites the current value of this request's
	// `verification_level` field with `null`.
	SetNullVerificationLevel() GuildPatch

	// UnsetVerificationLevel removes this request's `verification_level` field.
	UnsetVerificationLevel() GuildPatch

	// DefaultMessageNotifications returns the current value of this request's
	// `default_message_notifications` field.
	//
	// The `default_message_notifications` field contains the default message
	// notification level.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use DefaultMessageNotificationsIsReadable to check
	// if the field is present and non-null before use.
	DefaultMessageNotifications() guild.MessageNotificationLevel

	// DefaultMessageNotificationsIsNull returns whether this request's
	// `default_message_notifications` field is currently null.
	DefaultMessageNotificationsIsNull() bool

	// DefaultMessageNotificationsIsSet returns whether this request's
	// `default_message_notifications` field is currently present.
	DefaultMessageNotificationsIsSet() bool

	// DefaultMessageNotificationsIsReadable returns whether this request's
	// `default_message_notifications` field is currently set and non-null.
	DefaultMessageNotificationsIsReadable() bool

	// SetDefaultMessageNotifications overwrites the current value of this
	// request's `default_message_notifications` field.
	SetDefaultMessageNotifications(guild.MessageNotificationLevel) GuildPatch

	// SetNullDefaultMessageNotifications overwrites the current value of this
	// request's `default_message_notifications` field with `null`.
	SetNullDefaultMessageNotifications() GuildPatch

	// UnsetDefaultMessageNotifications removes this request's
	// `default_message_notifications` field.
	UnsetDefaultMessageNotifications() GuildPatch

	// ExplicitContentFilter returns the current value of this request's
	// `explicit_content_filter` field.
	//
	// The `explicit_content_filter` field contains the explicit content filter
	// level.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use ExplicitContentFilterIsReadable to check if
	// the field is present and non-null before use.
	ExplicitContentFilter() guild.ExplicitContentFilterLevel

	// ExplicitContentFilterIsNull returns whether this request's
	// `explicit_content_filter` field is currently null.
	ExplicitContentFilterIsNull() bool

	// ExplicitContentFilterIsSet returns whether this request's
	// `explicit_content_filter` field is currently present.
	ExplicitContentFilterIsSet() bool

	// ExplicitContentFilterIsReadable returns whether this request's
	// `explicit_content_filter` field is currently set and non-null.
	ExplicitContentFilterIsReadable() bool

	// SetExplicitContentFilter overwrites the current value of this request's
	// `explicit_content_filter` field.
	SetExplicitContentFilter(guild.ExplicitContentFilterLevel) GuildPatch

	// SetNullExplicitContentFilter overwrites the current value of this request's
	// `explicit_content_filter` field with `null`.
	SetNullExplicitContentFilter() GuildPatch

	// UnsetExplicitContentFilter removes this request's `explicit_content_filter`
	// field.
	UnsetExplicitContentFilter() GuildPatch

	// AFKChannelID returns the current value of this request's `afk_channel_id`
	// field.
	//
	// The `afk_channel_id` field contains the id for the afk channel.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use AFKChannelIDIsReadable to check if the field is
	// present and non-null before use.
	AFKChannelID() dlib.Snowflake

	// AFKChannelIDIsNull returns whether this request's `afk_channel_id` field is
	// currently null.
	AFKChannelIDIsNull() bool

	// AFKChannelIDIsSet returns whether this request's `afk_channel_id` field is
	// currently present.
	AFKChannelIDIsSet() bool

	// AFKChannelIDIsReadable returns whether this request's `afk_channel_id`
	// field is currently set and non-null.
	AFKChannelIDIsReadable() bool

	// SetAFKChannelID overwrites the current value of this request's
	// `afk_channel_id` field.
	SetAFKChannelID(dlib.Snowflake) GuildPatch

	// SetNullAFKChannelID overwrites the current value of this request's
	// `afk_channel_id` field with `null`.
	SetNullAFKChannelID() GuildPatch

	// UnsetAFKChannelID removes this request's `afk_channel_id` field.
	UnsetAFKChannelID() GuildPatch

	// AFKTimeout returns the current value of this request's `afk_timeout` field.
	//
	// The `afk_timeout` field contains the afk timeout in seconds.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AFKTimeoutIsSet to check if the field is present before use.
	AFKTimeout() uint64

	// AFKTimeoutIsSet returns whether this request's `afk_timeout` field is
	// currently present.
	AFKTimeoutIsSet() bool

	// SetAFKTimeout overwrites the current value of this request's `afk_timeout`
	// field.
	SetAFKTimeout(uint64) GuildPatch

	// UnsetAFKTimeout removes this request's `afk_timeout` field.
	UnsetAFKTimeout() GuildPatch

	// Icon returns the current value of this request's `icon` field.
	//
	// The `icon` field contains a base64 1024x1024 png/jpeg/gif image for the
	// guild icon (can be animated gif when the server has ANIMATED_ICON feature).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use IconIsReadable to check if the field is
	// present and non-null before use.
	Icon() image.Data

	// IconIsNull returns whether this request's `icon` field is currently null.
	IconIsNull() bool

	// IconIsSet returns whether this request's `icon` field is currently present.
	IconIsSet() bool

	// IconIsReadable returns whether this request's `icon` field is currently set
	// and non-null.
	IconIsReadable() bool

	// SetIcon overwrites the current value of this request's `icon` field.
	SetIcon(image.Data) GuildPatch

	// SetNullIcon overwrites the current value of this request's `icon` field
	// with `null`.
	SetNullIcon() GuildPatch

	// UnsetIcon removes this request's `icon` field.
	UnsetIcon() GuildPatch

	// OwnerID returns the current value of this request's `owner_id` field.
	//
	// The `owner_id` field contains the user id to transfer guild ownership to
	// (current user must be the current guild owner).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use OwnerIDIsSet to check if the field is present before use.
	OwnerID() dlib.Snowflake

	// OwnerIDIsSet returns whether this request's `owner_id` field is currently
	// present.
	OwnerIDIsSet() bool

	// SetOwnerID overwrites the current value of this request's `owner_id` field.
	SetOwnerID(dlib.Snowflake) GuildPatch

	// UnsetOwnerID removes this request's `owner_id` field.
	UnsetOwnerID() GuildPatch

	// Splash returns the current value of this request's `splash` field.
	//
	// The `splash` field contains a base64 16:9 png/jpeg image for the guild
	// splash (when the server has INVITE_SPLASH feature).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use SplashIsReadable to check if the field is
	// present and non-null before use.
	Splash() image.Data

	// SplashIsNull returns whether this request's `splash` field is currently
	// null.
	SplashIsNull() bool

	// SplashIsSet returns whether this request's `splash` field is currently
	// present.
	SplashIsSet() bool

	// SplashIsReadable returns whether this request's `splash` field is currently
	// set and non-null.
	SplashIsReadable() bool

	// SetSplash overwrites the current value of this request's `splash` field.
	SetSplash(image.Data) GuildPatch

	// SetNullSplash overwrites the current value of this request's `splash` field
	// with `null`.
	SetNullSplash() GuildPatch

	// UnsetSplash removes this request's `splash` field.
	UnsetSplash() GuildPatch

	// Banner returns the current value of this request's `banner` field.
	//
	// The `banner` field contains the base64 16:9 png/jpeg image for the guild
	// banner (when the server has BANNER feature).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use BannerIsReadable to check if the field is
	// present and non-null before use.
	Banner() image.Data

	// BannerIsNull returns whether this request's `banner` field is currently
	// null.
	BannerIsNull() bool

	// BannerIsSet returns whether this request's `banner` field is currently
	// present.
	BannerIsSet() bool

	// BannerIsReadable returns whether this request's `banner` field is currently
	// set and non-null.
	BannerIsReadable() bool

	// SetBanner overwrites the current value of this request's `banner` field.
	SetBanner(image.Data) GuildPatch

	// SetNullBanner overwrites the current value of this request's `banner` field
	// with `null`.
	SetNullBanner() GuildPatch

	// UnsetBanner removes this request's `banner` field.
	UnsetBanner() GuildPatch

	// SystemChannelID returns the current value of this request's `system_channel_id` field.
	//
	// The `system_channel_id` field contains the
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use SystemChannelIDIsReadable to check if the field is
	// present and non-null before use.
	SystemChannelID() dlib.Snowflake

	// SystemChannelIDIsNull returns whether this request's `system_channel_id` field is currently null.
	SystemChannelIDIsNull() bool

	// SystemChannelIDIsSet returns whether this request's `system_channel_id` field is currently present.
	SystemChannelIDIsSet() bool

	// SystemChannelIDIsReadable returns whether this request's `system_channel_id` field is currently set
	// and non-null.
	SystemChannelIDIsReadable() bool

	// SetSystemChannelID overwrites the current value of this request's `system_channel_id` field.
	SetSystemChannelID(dlib.Snowflake) GuildPatch

	// SetNullSystemChannelID overwrites the current value of this request's `system_channel_id` field
	// with `null`.
	SetNullSystemChannelID() GuildPatch

	// UnsetSystemChannelID removes this request's `system_channel_id` field.
	UnsetSystemChannelID() GuildPatch

	// RulesChannelID returns the current value of this request's `rules_channel_id` field.
	//
	// The `rules_channel_id` field contains the
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use RulesChannelIDIsReadable to check if the field is
	// present and non-null before use.
	RulesChannelID() dlib.Snowflake

	// RulesChannelIDIsNull returns whether this request's `rules_channel_id` field is currently null.
	RulesChannelIDIsNull() bool

	// RulesChannelIDIsSet returns whether this request's `rules_channel_id` field is currently present.
	RulesChannelIDIsSet() bool

	// RulesChannelIDIsReadable returns whether this request's `rules_channel_id` field is currently set
	// and non-null.
	RulesChannelIDIsReadable() bool

	// SetRulesChannelID overwrites the current value of this request's `rules_channel_id` field.
	SetRulesChannelID(dlib.Snowflake) GuildPatch

	// SetNullRulesChannelID overwrites the current value of this request's `rules_channel_id` field
	// with `null`.
	SetNullRulesChannelID() GuildPatch

	// UnsetRulesChannelID removes this request's `rules_channel_id` field.
	UnsetRulesChannelID() GuildPatch

	// PublicUpdatesChannelID returns the current value of this request's `public_updates_channel_id` field.
	//
	// The `public_updates_channel_id` field contains the
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use PublicUpdatesChannelIDIsReadable to check if the field is
	// present and non-null before use.
	PublicUpdatesChannelID() dlib.Snowflake

	// PublicUpdatesChannelIDIsNull returns whether this request's `public_updates_channel_id` field is currently null.
	PublicUpdatesChannelIDIsNull() bool

	// PublicUpdatesChannelIDIsSet returns whether this request's `public_updates_channel_id` field is currently present.
	PublicUpdatesChannelIDIsSet() bool

	// PublicUpdatesChannelIDIsReadable returns whether this request's `public_updates_channel_id` field is currently set
	// and non-null.
	PublicUpdatesChannelIDIsReadable() bool

	// SetPublicUpdatesChannelID overwrites the current value of this request's `public_updates_channel_id` field.
	SetPublicUpdatesChannelID(dlib.Snowflake) GuildPatch

	// SetNullPublicUpdatesChannelID overwrites the current value of this request's `public_updates_channel_id` field
	// with `null`.
	SetNullPublicUpdatesChannelID() GuildPatch

	// UnsetPublicUpdatesChannelID removes this request's `public_updates_channel_id` field.
	UnsetPublicUpdatesChannelID() GuildPatch

	// PreferredLocale returns the current value of this request's `preferred_locale` field.
	//
	// The `preferred_locale` field contains the
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use PreferredLocaleIsReadable to check if the field is
	// present and non-null before use.
	PreferredLocale() string

	// PreferredLocaleIsNull returns whether this request's `preferred_locale` field is currently null.
	PreferredLocaleIsNull() bool

	// PreferredLocaleIsSet returns whether this request's `preferred_locale` field is currently present.
	PreferredLocaleIsSet() bool

	// PreferredLocaleIsReadable returns whether this request's `preferred_locale` field is currently set
	// and non-null.
	PreferredLocaleIsReadable() bool

	// SetPreferredLocale overwrites the current value of this request's `preferred_locale` field.
	SetPreferredLocale(string) GuildPatch

	// SetNullPreferredLocale overwrites the current value of this request's `preferred_locale` field
	// with `null`.
	SetNullPreferredLocale() GuildPatch

	// UnsetPreferredLocale removes this request's `preferred_locale` field.
	UnsetPreferredLocale() GuildPatch
}
