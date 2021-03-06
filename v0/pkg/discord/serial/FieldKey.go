package serial

import (
	"io"

	"github.com/foxcapades/lib-go-discord/v0/internal/js"
)

type Key string

const (
	KeyAFKChannelID                Key = "afk_channel_id"
	KeyAFKTimeout                  Key = "afk_timeout"
	KeyAccessToken                 Key = "access_token"
	KeyActivity                    Key = "activity"
	KeyAllow                       Key = "allow"
	KeyAnimated                    Key = "animated"
	KeyApplication                 Key = "application"
	KeyApplicationID               Key = "application_id"
	KeyApproximateMemberCount      Key = "approximate_member_count"
	KeyApproximatePresenceCount    Key = "approximate_presence_count"
	KeyAssets                      Key = "assets"
	KeyAttachments                 Key = "attachments"
	KeyAuthor                      Key = "author"
	KeyAvatar                      Key = "avatar"
	KeyBanner                      Key = "banner"
	KeyBitrate                     Key = "bitrate"
	KeyBot                         Key = "bot"
	KeyChannelID                   Key = "channel_id"
	KeyChannels                    Key = "channels"
	KeyColor                       Key = "color"
	KeyContent                     Key = "content"
	KeyCreatedAt                   Key = "created_at"
	KeyCustom                      Key = "custom"
	KeyDeaf                        Key = "deaf"
	KeyDefaultMessageNotifications Key = "default_message_notifications"
	KeyDeny                        Key = "deny"
	KeyDeprecated                  Key = "deprecated"
	KeyDescription                 Key = "description"
	KeyDetails                     Key = "details"
	KeyDiscoverySplash             Key = "discovery_splash"
	KeyDiscriminator               Key = "discriminator"
	KeyEditedTimestamp             Key = "edited_timestamp"
	KeyEmail                       Key = "email"
	KeyEmbed                       Key = "embed"
	KeyEmbeds                      Key = "embeds"
	KeyEmoji                       Key = "emoji"
	KeyEmojis                      Key = "emojis"
	KeyEnabled                     Key = "enabled"
	KeyEnd                         Key = "end"
	KeyExplicitContentFilter       Key = "explicit_content_filter"
	KeyFeatures                    Key = "features"
	KeyFilename                    Key = "filename"
	KeyFlags                       Key = "flags"
	KeyFriendSync                  Key = "friend_sync"
	KeyGuildID                     Key = "guild_id"
	KeyHeight                      Key = "height"
	KeyHoist                       Key = "hoist"
	KeyID                          Key = "id"
	KeyIcon                        Key = "icon"
	KeyIconHash                    Key = "icon_hash"
	KeyInstance                    Key = "instance"
	KeyIntegrations                Key = "revoked"
	KeyJoin                        Key = "join"
	KeyJoinedAt                    Key = "joined_at"
	KeyLarge                       Key = "large"
	KeyLargeImage                  Key = "large_image"
	KeyLargeText                   Key = "large_text"
	KeyLastMessageID               Key = "last_message_id"
	KeyLastPinTimestamp            Key = "last_pin_timestamp"
	KeyLocale                      Key = "locale"
	KeyMFAEnabled                  Key = "mfa_enabled"
	KeyManaged                     Key = "managed"
	KeyMatch                       Key = "match"
	KeyMaxMembers                  Key = "max_members"
	KeyMaxPresences                Key = "max_presences"
	KeyMaxVideoChannelUsers        Key = "max_video_channel_users"
	KeyMember                      Key = "member"
	KeyMemberCount                 Key = "member_count"
	KeyMembers                     Key = "members"
	KeyMentionable                 Key = "mentionable"
	KeyMentionChannels             Key = "mention_channels"
	KeyMentionEveryone             Key = "mention_everyone"
	KeyMentionRoles                Key = "mention_roles"
	KeyMentions                    Key = "mentions"
	KeyMessageReference            Key = "message_reference"
	KeyMfaLevel                    Key = "mfa_level"
	KeyMute                        Key = "mute"
	KeyNSFW                        Key = "nsfw"
	KeyName                        Key = "name"
	KeyNick                        Key = "nick"
	KeyNonce                       Key = "nonce"
	KeyOptimal                     Key = "optimal"
	KeyOwner                       Key = "owner"
	KeyOwnerID                     Key = "owner_id"
	KeyParentID                    Key = "parent_id"
	KeyParty                       Key = "party"
	KeyPermissionOverwrites        Key = "permission_overwrites"
	KeyPermissions                 Key = "permissions"
	KeyPinned                      Key = "pinned"
	KeyPosition                    Key = "position"
	KeyPreferredLocale             Key = "preferred_locale"
	KeyPremiumSubscriptionCount    Key = "premium_subscription_count"
	KeyPremiumTier                 Key = "premium_tier"
	KeyPremiumType                 Key = "premium_type"
	KeyPresences                   Key = "presences"
	KeyProxyURL                    Key = "proxy_url"
	KeyPublicFlags                 Key = "public_flags"
	KeyPublicUpdatesChannelID      Key = "public_updates_channel_id"
	KeyRateLimitPerUser            Key = "rate_limit_per_user"
	KeyReactions                   Key = "reactions"
	KeyReason                      Key = "reason"
	KeyRecipients                  Key = "recipients"
	KeyRegion                      Key = "region"
	KeyRevoked                     Key = "revoked"
	KeyRoles                       Key = "roles"
	KeyRulesChannelID              Key = "rules_channel_id"
	KeySecrets                     Key = "secrets"
	KeySelfDeaf                    Key = "self_deaf"
	KeySelfMute                    Key = "self_mute"
	KeySelfStream                  Key = "self_stream"
	KeySelfVideo                   Key = "self_video"
	KeySessionID                   Key = "session_id"
	KeyShowActivity                Key = "show_activity"
	KeySize                        Key = "size"
	KeySmallImage                  Key = "small_image"
	KeySmallText                   Key = "small_text"
	KeySpectate                    Key = "spectate"
	KeySplash                      Key = "splash"
	KeyStart                       Key = "start"
	KeyState                       Key = "state"
	KeySuppress                    Key = "suppress"
	KeySystem                      Key = "system"
	KeySystemChannelFlags          Key = "system_channel_flags"
	KeySystemChannelID             Key = "system_channel_id"
	KeyTimestamp                   Key = "timestamp"
	KeyTimestamps                  Key = "timestamps"
	KeyToken                       Key = "token"
	KeyTopic                       Key = "topic"
	KeyTTS                         Key = "tts"
	KeyType                        Key = "type"
	KeyUnavailable                 Key = "unavailable"
	KeyURL                         Key = "url"
	KeyUser                        Key = "user"
	KeyUserID                      Key = "user_id"
	KeyUserLimit                   Key = "user_limit"
	KeyUsername                    Key = "username"
	KeyVIP                         Key = "vip"
	KeyVanityURLCode               Key = "vanity_url_code"
	KeyVerificationLevel           Key = "verification_level"
	KeyVerified                    Key = "verified"
	KeyVisibility                  Key = "visibility"
	KeyVoiceStates                 Key = "voice_states"
	KeyWebhookID                   Key = "webhook_id"
	KeyWidgetChannelID             Key = "widget_channel_id"
	KeyWidgetEnabled               Key = "widget_enabled"
	KeyWidth                       Key = "width"
)

func (k *Key) JSONSize() uint32 {
	if k == nil {
		return 4
	}

	return uint32(len(*k)) + js.QuoteSize
}

func (k Key) ToJSONBytes() []byte {
	sz := k.JSONSize()
	out := make([]byte, sz)
	out[0] = '"'
	out[sz-1] = '"'

	copy(out[1:], k)

	return out
}

func (k Key) AppendJSONBytes(writer io.Writer) (e error) {
	_, e = writer.Write(js.SingleQuoteBuf)

	if e == nil {
		_, e = writer.Write([]byte(k))
	}

	if e == nil {
		_, e = writer.Write(js.SingleQuoteBuf)
	}

	return
}
