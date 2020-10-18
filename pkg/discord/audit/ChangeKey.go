package audit

// TODO: this is a temporary measure until a more type-safe way to handle
//       changes is added.
type ChangeKey string

const (
	// String	name changed.
	//
	// Applies to: Guild
	ChangeKeyName = "name"
	// String	icon changed.
	//
	// Applies to: Guild
	ChangeKeyIconHash = "icon_hash"
	// String	invite splash page artwork changed.
	//
	// Applies to: Guild
	ChangeKeySplashHash = "splash_hash"
	// Snowflake	owner changed.
	//
	// Applies to: Guild
	ChangeKeyOwnerId = "owner_id"
	// String	region changed.
	//
	// Applies to: Guild
	ChangeKeyRegion = "region"
	// Snowflake	afk channel changed.
	//
	// Applies to: Guild
	ChangeKeyAfkChannelId = "afk_channel_id"
	// Integer	afk timeout duration changed.
	//
	// Applies to: Guild
	ChangeKeyAfkTimeout = "afk_timeout"
	// Integer	two-factor auth requirement changed.
	//
	// Applies to: Guild
	ChangeKeyMfaLevel = "mfa_level"
	// Integer	required verification level changed.
	//
	// Applies to: Guild
	ChangeKeyVerificationLevel = "verification_level"
	// Integer	change in whose messages are scanned and deleted for explicit content in the server.
	//
	// Applies to: Guild
	ChangeKeyExplicitContentFilter = "explicit_content_filter"
	// Integer	default message notification level changed.
	//
	// Applies to: Guild
	ChangeKeyDefaultMessageNotifications = "default_message_notifications"
	// String	guild invite vanity url changed.
	//
	// Applies to: Guild
	ChangeKeyVanityUrlCode = "vanity_url_code"
	// Array of partial role objects	new role added.
	//
	// Applies to: Guild
	ChangeKeyRoleAdd = "add"
	// Array of partial role objects	role removed.
	//
	// Applies to: Guild
	ChangeKeyRoleRemove = "remove"
	// Integer	change in number of days after which inactive and role-unassigned members are kicked.
	//
	// Applies to: Guild
	ChangeKeyPruneDeleteDays = "prune_delete_days"
	// Boolean	server widget enabled/disable.
	//
	// Applies to: Guild
	ChangeKeyWidgetEnabled = "widget_enabled"
	// Snowflake	channel id of the server widget changed.
	//
	// Applies to: Guild
	ChangeKeyWidgetChannelId = "widget_channel_id"
	// Snowflake	id of the system channel changed.
	//
	// Applies to: Guild
	ChangeKeySystemChannelId = "system_channel_id"
	// Integer	text or voice channel position changed.
	//
	// Applies to: Channel
	ChangeKeyPosition = "position"
	// String	text channel topic changed.
	//
	// Applies to: Channel
	ChangeKeyTopic = "topic"
	// Integer	voice channel bitrate changed.
	//
	// Applies to: Channel
	ChangeKeyBitrate = "bitrate"
	// Array of channel overwrite objects	permissions on a channel changed.
	//
	// Applies to: Channel
	ChangeKeyPermissionOverwrites = "permission_overwrites"
	// Boolean	channel nsfw restriction changed.
	//
	// Applies to: Channel
	ChangeKeyNsfw = "nsfw"
	// Snowflake	application id of the added or removed webhook or bot.
	//
	// Applies to: Channel
	ChangeKeyApplicationId = "application_id"
	// Integer	amount of seconds a user has to wait before sending another message changed.
	//
	// Applies to: Channel
	ChangeKeyRateLimitPerUser = "rate_limit_per_user"
	// String	permissions for a role changed.
	//
	// Applies to: Role
	ChangeKeyPermissions = "permissions"
	// Integer	role color changed.
	//
	// Applies to: Role
	ChangeKeyColor = "color"
	// Boolean	role is now displayed/no longer displayed separate from online users.
	//
	// Applies to: Role
	ChangeKeyHoist = "hoist"
	// Boolean	role is now mentionable/unmentionable.
	//
	// Applies to: Role
	ChangeKeyMentionable = "mentionable"
	// String	a permission on a text or voice channel was allowed for a role.
	//
	// Applies to: Role
	ChangeKeyAllow = "allow"
	// String	a permission on a text or voice channel was denied for a role.
	//
	// Applies to: Role
	ChangeKeyDeny = "deny"
	// String	invite code changed.
	//
	// Applies to: Invite
	ChangeKeyCode = "code"
	// Snowflake	channel for invite code changed.
	//
	// Applies to: Invite
	ChangeKeyChannelId = "channel_id"
	// Snowflake	person who created invite code changed.
	//
	// Applies to: Invite
	ChangeKeyInviterId = "inviter_id"
	// Integer	change to max number of times invite code can be used.
	//
	// Applies to: Invite
	ChangeKeyMaxUses = "max_uses"
	// Integer	number of times invite code used changed.
	//
	// Applies to: Invite
	ChangeKeyUses = "uses"
	// Integer	how long invite code lasts changed.
	//
	// Applies to: Invite
	ChangeKeyMaxAge = "max_age"
	// Boolean	invite code is temporary/never expires.
	//
	// Applies to: Invite
	ChangeKeyTemporary = "temporary"
	// Boolean	user server deafened/undeafened.
	//
	// Applies to: User
	ChangeKeyDeaf = "deaf"
	// Boolean	user server muted/unmuted.
	//
	// Applies to: User
	ChangeKeyMute = "mute"
	// String	user nickname changed.
	//
	// Applies to: User
	ChangeKeyNick = "nick"
	// String	user avatar changed.
	//
	// Applies to: User
	ChangeKeyAvatarHash = "avatar_hash"
	// Snowflake	the id of the changed entity - sometimes used in conjunction with other keys.
	//
	// Applies to: Any
	ChangeKeyId = "id"
	// Integer (channel type) or string	type of entity created.
	//
	// Applies to: Any
	ChangeKeyType = "type"
	// Boolean	integration emoticons enabled/disabled.
	//
	// Applies to: Integration
	ChangeKeyEnableEmoticons = "enable_emoticons"
	// Integer	integration expiring subscriber behavior changed.
	//
	// Applies to: Integration
	ChangeKeyExpireBehavior = "expire_behavior"
	// Integer	integration expire grace period changed.
	//
	// Applies to: Integration
	ChangeKeyExpireGracePeriod = "expire_grace_period"
)
