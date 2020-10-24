package types

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"time"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewChannelImpl(val bool) Channel {
	return &ChannelImpl{validate: val}
}

type ChannelImpl struct {
	validate bool

	id          Snowflake
	kind        ChannelType
	guildID     OptionalSnowflake
	position    OptionalUint16
	permOver    OptionalAny
	name        OptionalString
	topic       TriStateString
	nsfw        OptionalBool
	lastMsgID   TriStateSnowflake
	bitrate     OptionalUint32
	userLim     OptionalUint8
	rLimPerUser OptionalUint16
	recipients  OptionalAny
	icon        TriStateString
	ownerID     OptionalSnowflake
	appID       OptionalSnowflake
	parentID    TriStateSnowflake
	lastPin     TriStateTime
}

func (c *ChannelImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(OutBuilder, 17).
		AppendRaw(serial.KeyID, &c.id).
		AppendRaw(serial.KeyType, &c.kind).
		AppendOptional(serial.KeyGuildID, &c.guildID).
		AppendOptional(serial.KeyPosition, &c.position).
		AppendOptional(serial.KeyPermissionOverwrites, &c.permOver).
		AppendOptional(serial.KeyName, &c.name).
		AppendTriState(serial.KeyTopic, &c.topic).
		AppendOptional(serial.KeyNSFW, &c.nsfw).
		AppendTriState(serial.KeyLastMessageID, &c.lastMsgID).
		AppendOptional(serial.KeyBitrate, &c.bitrate).
		AppendOptional(serial.KeyUserLimit, &c.userLim).
		AppendOptional(serial.KeyRateLimitPerUser, &c.rLimPerUser).
		AppendOptional(serial.KeyRecipients, &c.recipients).
		AppendTriState(serial.KeyIcon, &c.icon).
		AppendOptional(serial.KeyOwnerID, &c.ownerID).
		AppendOptional(serial.KeyApplicationID, &c.appID).
		AppendTriState(serial.KeyParentID, &c.parentID).
		AppendTriState(serial.KeyLastPinTimestamp, &c.lastPin))
}

func (c *ChannelImpl) UnmarshalJSON(bytes []byte) error {
	panic("implement me")
}

func (c *ChannelImpl) IsValid() bool {
	panic("implement me")
}

func (c *ChannelImpl) Validate() error {
	panic("implement me")
}

func (c *ChannelImpl) ID() Snowflake {
	return c.id
}

func (c *ChannelImpl) SetID(id Snowflake) Channel {
	c.id = id
	return c
}

func (c *ChannelImpl) Type() ChannelType {
	return c.kind
}

func (c *ChannelImpl) SetType(kind ChannelType) Channel {
	c.kind = kind
	return c
}

func (c *ChannelImpl) GuildID() Snowflake {
	return c.guildID.Get().(Snowflake)
}

func (c *ChannelImpl) GuildIDIsSet() bool {
	return c.guildID.IsSet()
}

func (c *ChannelImpl) SetGuildID(id Snowflake) Channel {
	c.guildID.Set(id)
	return c
}

func (c *ChannelImpl) UnsetGuildID() Channel {
	c.guildID.Unset()
	return c
}

func (c *ChannelImpl) Position() uint16 {
	return c.position.Get()
}

func (c *ChannelImpl) PositionIsSet() bool {
	return c.position.IsSet()
}

func (c *ChannelImpl) SetPosition(u uint16) Channel {
	c.position.Set(u)
	return c
}

func (c *ChannelImpl) UnsetPosition() Channel {
	c.position.Unset()
	return c
}

func (c *ChannelImpl) PermissionOverwrites() []PermissionOverwrite {
	return c.permOver.Get().([]PermissionOverwrite)
}

func (c *ChannelImpl) PermissionOverwritesIsSet() bool {
	return c.permOver.IsSet()
}

func (c *ChannelImpl) SetPermissionOverwrites(o []PermissionOverwrite) Channel {
	c.permOver.Set(o)
	return c
}

func (c *ChannelImpl) UnsetPermissionOverwrites() Channel {
	c.permOver.Unset()
	return c
}

func (c *ChannelImpl) Name() ChannelName {
	return ChannelName(c.name.Get())
}

func (c *ChannelImpl) NameIsSet() bool {
	return c.name.IsSet()
}

func (c *ChannelImpl) SetName(name ChannelName) Channel {
	c.name.Set(string(name))
	return c
}

func (c *ChannelImpl) UnsetName() Channel {
	c.name.Unset()
	return c
}

func (c *ChannelImpl) Topic() ChannelTopic {
	return ChannelTopic(c.topic.Get())
}

func (c *ChannelImpl) TopicIsNull() bool {
	return c.topic.IsNull()
}

func (c *ChannelImpl) TopicIsSet() bool {
	return c.topic.IsSet()
}

func (c *ChannelImpl) TopicIsReadable() bool {
	return c.topic.IsReadable()
}

func (c *ChannelImpl) SetTopic(topic ChannelTopic) Channel {
	c.topic.Set(string(topic))
	return c
}

func (c *ChannelImpl) SetNullTopic() Channel {
	c.topic.SetNull()
	return c
}

func (c *ChannelImpl) UnsetTopic() Channel {
	c.topic.Unset()
	return c
}

func (c *ChannelImpl) NSFW() bool {
	return c.nsfw.Get()
}

func (c *ChannelImpl) NSFWIsSet() bool {
	return c.nsfw.IsSet()
}

func (c *ChannelImpl) SetNSFW(b bool) Channel {
	c.nsfw.Set(b)
	return c
}

func (c *ChannelImpl) UnsetNSFW() Channel {
	c.nsfw.Unset()
	return c
}

func (c *ChannelImpl) LastMessageID() Snowflake {
	return c.lastMsgID.Get().(Snowflake)
}

func (c *ChannelImpl) LastMessageIDIsNull() bool {
	return c.lastMsgID.IsNull()
}

func (c *ChannelImpl) LastMessageIDIsSet() bool {
	return c.lastMsgID.IsSet()
}

func (c *ChannelImpl) LastMessageIDIsReadable() bool {
	return c.lastMsgID.IsReadable()
}

func (c *ChannelImpl) SetLastMessageID(id Snowflake) Channel {
	c.lastMsgID.Set(id)
	return c
}

func (c *ChannelImpl) SetNullLastMessageID() Channel {
	c.lastMsgID.SetNull()
	return c
}

func (c *ChannelImpl) UnsetLastMessageID() Channel {
	c.lastMsgID.Unset()
	return c
}

func (c *ChannelImpl) Bitrate() Bitrate {
	return Bitrate(c.bitrate.Get())
}

func (c *ChannelImpl) BitrateIsSet() bool {
	return c.bitrate.IsSet()
}

func (c *ChannelImpl) SetBitrate(bitrate Bitrate) Channel {
	c.bitrate.Set(uint32(bitrate))
	return c
}

func (c *ChannelImpl) UnsetBitrate() Channel {
	c.bitrate.Unset()
	return c
}

func (c *ChannelImpl) UserLimit() UserLimit {
	return UserLimit(c.userLim.Get())
}

func (c *ChannelImpl) UserLimitIsSet() bool {
	return c.userLim.IsSet()
}

func (c *ChannelImpl) SetUserLimit(limit UserLimit) Channel {
	c.userLim.Set(uint8(limit))
	return c
}

func (c *ChannelImpl) UnsetUserLimit() Channel {
	c.userLim.Unset()
	return c
}

func (c *ChannelImpl) RateLimitPerUser() PerUserRateLimit {
	return PerUserRateLimit(c.rLimPerUser.Get())
}

func (c *ChannelImpl) RateLimitPerUserIsSet() bool {
	return c.rLimPerUser.IsSet()
}

func (c *ChannelImpl) SetRateLimitPerUser(limit PerUserRateLimit) Channel {
	c.rLimPerUser.Set(uint16(limit))
	return c
}

func (c *ChannelImpl) UnsetRateLimitPerUser() Channel {
	c.rLimPerUser.Unset()
	return c
}

func (c *ChannelImpl) Recipients() []User {
	return c.recipients.Get().([]User)
}

func (c *ChannelImpl) RecipientsIsSet() bool {
	return c.recipients.IsSet()
}

func (c *ChannelImpl) SetRecipients(users []User) Channel {
	c.recipients.Set(users)
	return c
}

func (c *ChannelImpl) UnsetRecipients() Channel {
	c.recipients.Unset()
	return c
}

func (c *ChannelImpl) Icon() ImageHash {
	return ImageHash(c.icon.Get())
}

func (c *ChannelImpl) IconIsNull() bool {
	return c.icon.IsNull()
}

func (c *ChannelImpl) IconIsSet() bool {
	return c.icon.IsSet()
}

func (c *ChannelImpl) IconIsReadable() bool {
	return c.icon.IsReadable()
}

func (c *ChannelImpl) SetIcon(h ImageHash) Channel {
	c.icon.Set(string(h))
	return c
}

func (c *ChannelImpl) SetNullIcon() Channel {
	c.icon.SetNull()
	return c
}

func (c *ChannelImpl) UnsetIcon() Channel {
	c.icon.Unset()
	return c
}

func (c *ChannelImpl) OwnerID() Snowflake {
	return c.ownerID.Get().(Snowflake)
}

func (c *ChannelImpl) OwnerIDIsSet() bool {
	return c.ownerID.IsSet()
}

func (c *ChannelImpl) SetOwnerID(id Snowflake) Channel {
	c.ownerID.Set(id)
	return c
}

func (c *ChannelImpl) UnsetOwnerID() Channel {
	c.ownerID.Unset()
	return c
}

func (c *ChannelImpl) ApplicationID() Snowflake {
	return c.appID.Get().(Snowflake)
}

func (c *ChannelImpl) ApplicationIDIsSet() bool {
	return c.appID.IsSet()
}

func (c *ChannelImpl) SetApplicationID(id Snowflake) Channel {
	c.appID.Set(id)
	return c
}

func (c *ChannelImpl) UnsetApplicationID() Channel {
	c.appID.Unset()
	return c
}

func (c *ChannelImpl) ParentID() Snowflake {
	return c.parentID.Get().(Snowflake)
}

func (c *ChannelImpl) ParentIDIsNull() bool {
	return c.parentID.IsNull()
}

func (c *ChannelImpl) ParentIDIsSet() bool {
	return c.parentID.IsSet()
}

func (c *ChannelImpl) ParentIDIsReadable() bool {
	return c.parentID.IsReadable()
}

func (c *ChannelImpl) SetParentID(id Snowflake) Channel {
	c.parentID.Set(id)
	return c
}

func (c *ChannelImpl) SetNullParentID() Channel {
	c.parentID.SetNull()
	return c
}

func (c *ChannelImpl) UnsetParentID() Channel {
	c.parentID.Unset()
	return c
}

func (c *ChannelImpl) LastPinTimestamp() time.Time {
	return c.lastPin.Get()
}

func (c *ChannelImpl) LastPinTimestampIsNull() bool {
	return c.lastPin.IsNull()
}

func (c *ChannelImpl) LastPinTimestampIsSet() bool {
	return c.lastPin.IsSet()
}

func (c *ChannelImpl) LastPinTimestampIsReadable() bool {
	return c.lastPin.IsReadable()
}

func (c *ChannelImpl) SetLastPinTimestamp(t time.Time) Channel {
	c.lastPin.Set(t)
	return c
}

func (c *ChannelImpl) SetNullLastPinTimestamp() Channel {
	c.lastPin.SetNull()
	return c
}

func (c *ChannelImpl) UnsetLastPinTimestamp() Channel {
	c.lastPin.Unset()
	return c
}
