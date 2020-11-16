package guild

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
	"time"

	"github.com/francoispqt/gojay"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

const (
	maskGuildNullIconHash uint8 = 1 << iota
	maskGuildNullWidgetChanID
	maskGuildAbsentVoiceStates
	maskGuildAbsentMembers
	maskGuildAbsentChannels
	maskGuildAbsentPresences
	maskGuildNullMaxPresences
)

func NewGuild() Guild {
	return new(guild)
}

type guild struct {
	meta uint8

	id              Snowflake
	name            GuildName
	icon            *ImageHash
	iconHash        *ImageHash
	splash          *ImageHash
	discoverySplash *ImageHash
	owner           *bool
	ownerId         Snowflake
	permissions     *Permission
	region          string
	afkChannelID    Snowflake
	afkTimeout      uint64
	widgetEnabled   *bool
	widgetChannelID Snowflake
	verificationLvl VerificationLevel
	defMsgNotes     MessageNotificationLevel
	xxxFilter       ExplicitContentFilterLevel
	roles           RoleSlice
	emoji           EmojiSlice
	features        FeatureSlice
	mfa             MFALevel
	appID           Snowflake
	sysChanID       Snowflake
	sysChanFlags    SystemChannelFlag
	ruleChanID      Snowflake
	joinedAt        *time.Time
	large           *bool
	unavailable     *bool
	memberCount     *uint32
	voiceStates     []VoiceState     //FIXME: should be custom type
	members         []GuildMember    //FIXME: needs 2 b custom type
	channels        []Channel        //FIXME
	presences       []PresenceUpdate // FIXME
	maxPresences    *uint32
	maxMembers      *uint32
	vanityUrlCode   *string
	description     *string
	banner          *ImageHash
	premiumTier     PremiumTier
	premiumSubCount *uint16
	preferredLocale string
	pubUpdateChanID Snowflake
	maxVidChanUsers *uint16
	appxMemberCount *uint32
	appxPresCount   *uint32
}

func (g *guild) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, g.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeObject(g); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g *guild) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(g)
}

func (g *guild) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(KeyID, g.id).
		AddString(KeyName, string(g.name)).
		AddNullableString(KeyIcon, (*string)(g.icon)).
		AddTriStateString(KeyIconHash, (*string)(g.iconHash), g.IconHashIsNull()).
		AddNullableString(KeySplash, (*string)(g.splash)).
		AddNullableString(KeyDiscoverySplash, (*string)(g.discoverySplash)).
		AddOptionalBool(KeyOwner, g.owner).
		AddSnowflake(KeyOwnerID, g.ownerId).
		AddOptionalStringer(KeyPermissions, g.permissions)
}

func (g *guild) IsNil() bool {
	return g == nil
}

func (g *guild) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (g *guild) NKeys() int {
	return 0
}

func (g *guild) JSONSize() int {
	panic("implement me")
}

func (g *guild) IsValid() bool {
	panic("implement me")
}

func (g *guild) Validate() error {
	panic("implement me")
}

func (g *guild) ID() Snowflake {
	return g.id
}

func (g *guild) SetID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.id = id

	return g
}

func (g *guild) Name() GuildName {
	return g.name
}

func (g *guild) SetName(name GuildName) Guild {
	g.name = name
	return g
}

func (g *guild) Icon() ImageHash {
	if g.icon != nil {
		panic(e.ErrGetNull)
	}

	return *g.icon
}

func (g *guild) IconIsNull() bool {
	return g.icon == nil
}

func (g *guild) SetIcon(hash ImageHash) Guild {
	g.icon = &hash
	return g
}

func (g *guild) SetNullIcon() Guild {
	g.icon = nil
	return g
}

func (g *guild) IconHash() ImageHash {
	if g.iconHash == nil {
		if g.meta&maskGuildNullIconHash > 0 {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *g.iconHash
}

func (g *guild) IconHashIsNull() bool {
	return g.iconHash == nil && g.meta&maskGuildNullIconHash > 0
}

func (g *guild) IconHashIsSet() bool {
	return g.iconHash != nil
}

func (g *guild) IconHashIsReadable() bool {
	return g.iconHash != nil
}

func (g *guild) UnsetIconHash() Guild {
	g.iconHash = nil
	g.meta &= ^maskGuildNullIconHash

	return g
}

func (g *guild) SetIconHash(hash ImageHash) Guild {
	g.iconHash = &hash
	g.meta &= ^maskGuildNullIconHash

	return g
}

func (g *guild) SetNullIconHash() Guild {
	g.iconHash = nil
	g.meta |= maskGuildNullIconHash

	return g
}

func (g *guild) Splash() ImageHash {
	if g.splash == nil {
		panic(e.ErrGetNull)
	}

	return *g.splash
}

func (g *guild) SplashIsNull() bool {
	return g.splash == nil
}

func (g *guild) SetSplash(hash ImageHash) Guild {
	g.splash = &hash
	return g
}

func (g *guild) SetNullSplash() Guild {
	g.splash = nil
	return g
}

func (g *guild) DiscoverySplash() ImageHash {
	if g.discoverySplash == nil {
		panic(e.ErrGetNull)
	}

	return *g.discoverySplash
}

func (g *guild) DiscoverySplashIsNull() bool {
	return g.discoverySplash == nil
}

func (g *guild) SetDiscoverySplash(hash ImageHash) Guild {
	g.splash = &hash
	return g
}

func (g *guild) SetNullDiscoverySplash() Guild {
	g.discoverySplash = nil
	return g
}

func (g *guild) Owner() bool {
	if g.owner == nil {
		panic(e.ErrGetUnset)
	}

	return *g.owner
}

func (g *guild) OwnerIsSet() bool {
	return g.owner != nil
}

func (g *guild) SetOwner(b bool) Guild {
	g.owner = &b
	return g
}

func (g *guild) UnsetOwner() Guild {
	g.owner = nil
	return g
}

func (g *guild) OwnerID() Snowflake {
	return g.ownerId
}

func (g *guild) SetOwnerID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.ownerId = id

	return g
}

func (g *guild) Permissions() Permission {
	if g.permissions == nil {
		panic(e.ErrGetUnset)
	}

	return *g.permissions
}

func (g *guild) PermissionsIsSet() bool {
	return g.permissions != nil
}

func (g *guild) SetPermissions(permission Permission) Guild {
	g.permissions = &permission
	return g
}

func (g *guild) UnsetPermissions() Guild {
	g.permissions = nil
	return g
}

func (g *guild) Region() string {
	return g.region
}

func (g *guild) SetRegion(s string) Guild {
	g.region = s
	return g
}

func (g *guild) AFKChannelID() Snowflake {
	if g.afkChannelID == nil {
		panic(e.ErrGetNull)
	}

	return g.afkChannelID
}

func (g *guild) AFKChannelIDIsNull() bool {
	return g.afkChannelID == nil
}

func (g *guild) SetAFKChannelID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrGetNull)
	}

	g.afkChannelID = id

	return g
}

func (g *guild) SetNullAFKChannelID() Guild {
	g.afkChannelID = nil
	return g
}

func (g *guild) AFKTimeout() uint64 {
	return g.afkTimeout
}

func (g *guild) SetAFKTimeout(u uint64) Guild {
	g.afkTimeout = u
	return g
}

func (g *guild) WidgetEnabled() bool {
	if g.widgetEnabled == nil {
		panic(e.ErrGetUnset)
	}

	return *g.widgetEnabled
}

func (g *guild) WidgetEnabledIsSet() bool {
	return g.widgetEnabled != nil
}

func (g *guild) SetWidgetEnabled(b bool) Guild {
	g.widgetEnabled = &b
	return g
}

func (g *guild) UnsetWidgetEnabled() Guild {
	g.widgetEnabled = nil
	return g
}

func (g *guild) WidgetChannelID() Snowflake {
	if g.widgetChannelID == nil {
		if g.meta&maskGuildNullWidgetChanID > 0 {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return g.widgetChannelID
}

func (g *guild) WidgetChannelIDIsNull() bool {
	return g.widgetChannelID == nil && g.meta&maskGuildNullWidgetChanID > 0
}

func (g *guild) WidgetChannelIDIsSet() bool {
	return g.widgetChannelID != nil
}

func (g *guild) WidgetChannelIDIsReadable() bool {
	return g.widgetChannelID != nil
}

func (g *guild) SetWidgetChannelID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.widgetChannelID = id
	g.meta &= ^maskGuildNullWidgetChanID

	return g
}

func (g *guild) SetNullWidgetChannelID() Guild {
	g.widgetChannelID = nil
	g.meta |= maskGuildNullWidgetChanID

	return g
}

func (g *guild) UnsetWidgetChannelID() Guild {
	g.widgetChannelID = nil
	g.meta &= ^maskGuildNullWidgetChanID

	return g
}

func (g *guild) VerificationLevel() VerificationLevel {
	return g.verificationLvl
}

func (g *guild) SetVerificationLevel(level VerificationLevel) Guild {
	g.verificationLvl = level
	return g
}

func (g *guild) DefaultMessageNotifications() MessageNotificationLevel {
	return g.defMsgNotes
}

func (g *guild) SetDefaultMessageNotifications(level MessageNotificationLevel) Guild {
	g.defMsgNotes = level
	return g
}

func (g *guild) ExplicitContentFilter() ExplicitContentFilterLevel {
	return g.xxxFilter
}

func (g *guild) SetExplicitContentFilter(level ExplicitContentFilterLevel) Guild {
	g.xxxFilter = level
	return g
}

func (g *guild) Roles() []Role {
	return g.roles
}

func (g *guild) SetRoles(roles []Role) Guild {
	g.roles = roles
	return g
}

func (g *guild) Emojis() []Emoji {
	return g.emoji
}

func (g *guild) SetEmojis(emojis []Emoji) Guild {
	g.emoji = emojis
	return g
}

func (g *guild) Features() []GuildFeature {
	return g.features
}

func (g *guild) SetFeatures(features []GuildFeature) Guild {
	g.features = features
	return g
}

func (g *guild) MFALevel() MFALevel {
	return g.mfa
}

func (g *guild) SetMFALevel(level MFALevel) Guild {
	g.mfa = level
	return g
}

func (g *guild) ApplicationID() Snowflake {
	if g.appID == nil {
		panic(e.ErrGetNull)
	}

	return g.appID
}

func (g *guild) ApplicationIDIsNull() bool {
	return g.appID == nil
}

func (g *guild) SetApplicationID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.appID = id

	return g
}

func (g *guild) SetNullApplicationID() Guild {
	g.appID = nil
	return g
}

func (g *guild) SystemChannelID() Snowflake {
	if g.sysChanID == nil {
		panic(e.ErrGetNull)
	}

	return g.sysChanID
}

func (g *guild) SystemChannelIDIsNull() bool {
	return g.sysChanID == nil
}

func (g *guild) SetSystemChannelID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.sysChanID = id

	return g
}

func (g *guild) SetNullSystemChannelID() Guild {
	g.sysChanID = nil
	return g
}

func (g *guild) SystemChannelFlags() SystemChannelFlag {
	return g.sysChanFlags
}

func (g *guild) SetSystemChannelFlags(flag SystemChannelFlag) Guild {
	g.sysChanFlags = flag
	return g
}

func (g *guild) AddSystemChannelFlag(flag SystemChannelFlag) Guild {
	g.sysChanFlags |= flag
	return g
}

func (g *guild) RemoveSystemChannelFlag(flag SystemChannelFlag) Guild {
	g.sysChanFlags &= ^flag
	return g
}

func (g *guild) SystemChannelFlagsContains(flag SystemChannelFlag) bool {
	return g.sysChanFlags&flag == flag
}

func (g *guild) RulesChannelID() Snowflake {
	if g.ruleChanID == nil {
		panic(e.ErrGetNull)
	}

	return g.ruleChanID
}

func (g *guild) RulesChannelIDIsNull() bool {
	return g.ruleChanID == nil
}

func (g *guild) SetRulesChannelID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.ruleChanID = id

	return g
}

func (g *guild) SetNullRulesChannelID() Guild {
	g.ruleChanID = nil
	return g
}

func (g *guild) JoinedAt() time.Time {
	if g.joinedAt == nil {
		panic(e.ErrGetUnset)
	}

	return *g.joinedAt
}

func (g *guild) JoinedAtIsSet() bool {
	return g.joinedAt != nil
}

func (g *guild) SetJoinedAt(time time.Time) Guild {
	g.joinedAt = &time
	return g
}

func (g *guild) UnsetJoinedAt() Guild {
	g.joinedAt = nil
	return g
}

func (g *guild) Large() bool {
	if g.large == nil {
		panic(e.ErrGetUnset)
	}

	return *g.large
}

func (g *guild) LargeIsSet() bool {
	return g.large != nil
}

func (g *guild) SetLarge(b bool) Guild {
	g.large = &b
	return g
}

func (g *guild) UnsetLarge() Guild {
	g.large = nil
	return g
}

func (g *guild) Unavailable() bool {
	if g.unavailable == nil {
		panic(e.ErrGetUnset)
	}

	return *g.unavailable
}

func (g *guild) UnavailableIsSet() bool {
	return g.unavailable != nil
}

func (g *guild) SetUnavailable(b bool) Guild {
	g.unavailable = &b
	return g
}

func (g *guild) UnsetUnavailable() Guild {
	g.unavailable = nil
	return g
}

func (g *guild) MemberCount() uint32 {
	if g.memberCount == nil {
		panic(e.ErrGetUnset)
	}

	return *g.memberCount
}

func (g *guild) MemberCountIsSet() bool {
	return g.memberCount != nil
}

func (g *guild) SetMemberCount(u uint32) Guild {
	g.memberCount = &u
	return g
}

func (g *guild) UnsetMemberCount() Guild {
	g.memberCount = nil
	return g
}

func (g *guild) VoiceStates() []VoiceState {
	panic("implement me")
}

func (g *guild) VoiceStatesIsSet() bool {
	panic("implement me")
}

func (g *guild) SetVoiceStates(states []VoiceState) Guild {
	panic("implement me")
}

func (g *guild) UnsetVoiceStates() Guild {
	panic("implement me")
}

func (g *guild) Members() GuildMember {
	panic("implement me")
}

func (g *guild) MembersIsSet() bool {
	panic("implement me")
}

func (g *guild) SetMembers(guildMember GuildMember) Guild {
	panic("implement me")
}

func (g *guild) UnsetMembers() Guild {
	panic("implement me")
}

func (g *guild) Channels() []Channel {
	panic("implement me")
}

func (g *guild) ChannelsIsSet() bool {
	panic("implement me")
}

func (g *guild) SetChannels(channels []Channel) Guild {
	panic("implement me")
}

func (g *guild) UnsetChannels() Guild {
	panic("implement me")
}

func (g *guild) Presences() []PresenceUpdate {
	panic("implement me")
}

func (g *guild) PresencesIsSet() bool {
	panic("implement me")
}

func (g *guild) SetPresences(updates []PresenceUpdate) Guild {
	panic("implement me")
}

func (g *guild) UnsetPresences() Guild {
	panic("implement me")
}

func (g *guild) MaxPresences() uint32 {
	if g.maxPresences == nil {
		if g.meta&maskGuildNullMaxPresences > 0 {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *g.maxPresences
}

func (g *guild) MaxPresencesIsNull() bool {
	return g.maxPresences == nil && g.meta&maskGuildNullMaxPresences > 0
}

func (g *guild) MaxPresencesIsSet() bool {
	return g.maxPresences != nil
}

func (g *guild) MaxPresencesIsReadable() bool {
	return g.maxPresences != nil
}

func (g *guild) SetMaxPresences(u uint32) Guild {
	g.maxPresences = &u
	g.meta &= ^maskGuildNullMaxPresences

	return g
}

func (g *guild) SetNullMaxPresences() Guild {
	g.maxPresences = nil
	g.meta |= maskGuildNullMaxPresences

	return g
}

func (g *guild) UnsetMaxPresences() Guild {
	g.maxPresences = nil
	g.meta &= ^maskGuildNullMaxPresences

	return g
}

func (g *guild) MaxMembers() uint32 {
	if g.maxMembers == nil {
		panic(e.ErrGetUnset)
	}

	return *g.maxMembers
}

func (g *guild) MaxMembersIsSet() bool {
	return g.maxMembers != nil
}

func (g *guild) SetMaxMembers(u uint32) Guild {
	g.maxMembers = &u
	return g
}

func (g *guild) UnsetMaxMembers() Guild {
	g.maxVidChanUsers = nil
	return g
}

func (g *guild) VanityURLCode() string {
	if g.vanityUrlCode == nil {
		panic(e.ErrGetNull)
	}

	return *g.vanityUrlCode
}

func (g *guild) VanityURLCodeIsNull() bool {
	return g.vanityUrlCode == nil
}

func (g *guild) SetVanityURLCode(s string) Guild {
	g.vanityUrlCode = &s
	return g
}

func (g *guild) SetNullVanityURLCode() Guild {
	g.vanityUrlCode = nil
	return g
}

func (g *guild) Description() string {
	if g.description == nil {
		panic(e.ErrGetNull)
	}

	return *g.description
}

func (g *guild) DescriptionIsNull() bool {
	return g.description == nil
}

func (g *guild) SetDescription(s string) Guild {
	g.description = &s
	return g
}

func (g *guild) SetNullDescription() Guild {
	g.description = nil
	return g
}

func (g *guild) Banner() ImageHash {
	if g.banner == nil {
		panic(e.ErrGetNull)
	}

	return *g.banner
}

func (g *guild) BannerIsNull() bool {
	return g.banner == nil
}

func (g *guild) SetBanner(hash ImageHash) Guild {
	g.banner = &hash
	return g
}

func (g *guild) SetNullBanner() Guild {
	g.banner = nil
	return g
}

func (g *guild) PremiumTier() PremiumTier {
	return g.premiumTier
}

func (g *guild) SetPremiumTier(tier PremiumTier) Guild {
	g.premiumTier = tier
	return g
}

func (g *guild) PremiumSubscriptionCount() uint16 {
	if g.premiumSubCount == nil {
		panic(e.ErrGetUnset)
	}

	return *g.premiumSubCount
}

func (g *guild) PremiumSubscriptionCountIsSet() bool {
	return g.premiumSubCount != nil
}

func (g *guild) SetPremiumSubscriptionCount(u uint16) Guild {
	g.premiumSubCount = &u
	return g
}

func (g *guild) UnsetPremiumSubscriptionCount() Guild {
	g.premiumSubCount = nil
	return g
}

func (g *guild) PreferredLocale() string {
	return g.preferredLocale
}

func (g *guild) SetPreferredLocale(s string) Guild {
	g.preferredLocale = s
	return g
}

func (g *guild) PublicUpdatesChannelID() Snowflake {
	if g.pubUpdateChanID == nil {
		panic(e.ErrGetNull)
	}

	return g.pubUpdateChanID
}

func (g *guild) PublicUpdatesChannelIDIsNull() bool {
	return g.pubUpdateChanID == nil
}

func (g *guild) SetPublicUpdatesChannelID(id Snowflake) Guild {
	if id == nil {
		panic(e.ErrSetNil)
	}

	g.pubUpdateChanID = id

	return g
}

func (g *guild) SetNullPublicUpdatesChannelID() Guild {
	g.pubUpdateChanID = nil
	return g
}

func (g *guild) MaxVideoChannelUsers() uint16 {
	if g.maxVidChanUsers == nil {
		panic(e.ErrGetUnset)
	}

	return *g.maxVidChanUsers
}

func (g *guild) MaxVideoChannelUsersIsSet() bool {
	return g.maxVidChanUsers != nil
}

func (g *guild) SetMaxVideoChannelUsers(u uint16) Guild {
	g.maxVidChanUsers = &u
	return g
}

func (g *guild) UnsetMaxVideoChannelUsers() Guild {
	g.maxVidChanUsers = nil
	return g
}

func (g *guild) ApproximateMemberCount() uint32 {
	if g.appxMemberCount == nil {
		panic(e.ErrGetUnset)
	}

	return *g.appxMemberCount
}

func (g *guild) ApproximateMemberCountIsSet() bool {
	return g.appxMemberCount != nil
}

func (g *guild) SetApproximateMemberCount(u uint32) Guild {
	g.appxMemberCount = &u
	return g
}

func (g *guild) UnsetApproximateMemberCount() Guild {
	g.appxMemberCount = nil
	return g
}

func (g *guild) ApproximatePresenceCount() uint32 {
	if g.appxPresCount == nil {
		panic(e.ErrGetUnset)
	}

	return *g.appxPresCount
}

func (g *guild) ApproximatePresenceCountIsSet() bool {
	return g.appxPresCount != nil
}

func (g *guild) SetApproximatePresenceCount(u uint32) Guild {
	g.appxPresCount = &u
	return g
}

func (g *guild) UnsetApproximatePresenceCount() Guild {
	g.appxPresCount = nil
	return g
}
