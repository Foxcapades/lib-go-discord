package channel

import (
	"encoding/json"
	"fmt"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/channel/permission"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/user"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
)

func NewChannel() Channel {
	return new(channel)
}

const (
	// null topic
	maskNullTopic uint8 = 1 << iota

	// null last_message_id
	maskNullLastMsgId

	// null icon
	maskNullIcon

	// null parent_id
	maskNullParentId

	// null last_pin
	maskNullLastPin

	// absent permission_overwrites
	maskPermOver

	// absent recipients
	maskRecipients
)

type channel struct {
	metaFlags uint8

	id          Snowflake
	kind        ChannelType
	guildID     Snowflake
	position    *uint16
	permOver    permission.OverwriteSlice
	name        *ChannelName
	topic       *ChannelTopic
	nsfw        *bool
	lastMsgID   Snowflake
	bitrate     *Bitrate
	userLim     *UserLimit
	rLimPerUser *PerUserRateLimit
	recipients  user.Slice
	icon        *ImageHash
	ownerID     Snowflake
	appID       Snowflake
	parentID    Snowflake
	lastPin     *time.Time
}

func (c *channel) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(KeyID, c.id).
		AddUint8(KeyType, uint8(c.kind)).
		AddOptionalSnowflake(KeyGuildID, c.guildID).
		AddOptionalUint16(KeyPosition, c.position).
		AddOptionalSlice(
			KeyPermissionOverwrites,
			c.permOver,
			c.hasMetaFlag(maskPermOver),
		).
		AddOptionalString(KeyName, (*string)(c.name)).
		AddTriStateString(
			KeyTopic,
			(*string)(c.topic),
			c.hasMetaFlag(maskNullTopic),
		).
		AddOptionalBool(KeyNSFW, c.nsfw).
		AddTriStateSnowflake(
			KeyLastMessageID,
			c.lastMsgID,
			c.hasMetaFlag(maskNullLastMsgId),
		).
		AddOptionalUint32(KeyBitrate, (*uint32)(c.bitrate)).
		AddOptionalUint8(KeyUserLimit, (*uint8)(c.userLim)).
		AddOptionalUint16(KeyRateLimitPerUser, (*uint16)(c.rLimPerUser)).
		AddOptionalSlice(
			KeyRecipients,
			c.recipients,
			c.hasMetaFlag(maskRecipients),
		).
		AddTriStateString(
			KeyIcon,
			(*string)(c.icon),
			c.hasMetaFlag(maskNullIcon),
		).
		AddOptionalSnowflake(KeyOwnerID, c.ownerID).
		AddOptionalSnowflake(KeyApplicationID, c.appID).
		AddTriStateSnowflake(
			KeyParentID,
			c.parentID,
			c.hasMetaFlag(maskNullParentId),
		).
		AddTriStateUint64(
			KeyLastPinTimestamp,
			utils.UTCTimeToMillis(c.lastPin),
			c.hasMetaFlag(maskNullLastPin),
		)
}

func (c *channel) IsNil() bool {
	return c == nil
}

func (c *channel) UnmarshalJSONObject(dec *gojay.Decoder, key string) (err error) {
	switch Key(key) {

	case KeyID:
		c.id, err = com.DecodeSnowflake(dec)

	case KeyType:
		c.kind, err = DecodeChannelType(dec)

	case KeyGuildID:
		c.guildID, err = com.DecodeSnowflake(dec)

	case KeyPosition:
		return dec.Uint16Null(&c.position)

	case KeyPermissionOverwrites:
		c.permOver = make(permission.OverwriteSlice, 0, 10)
		return dec.DecodeArray(&c.permOver)

	case KeyName:
		c.name, err = DecodeChannelName(dec)

	case KeyTopic:
		c.topic, err = DecodeChannelTopic(dec)
		if c.topic == nil {
			c.setMetaFlag(maskNullTopic)
		}

	case KeyNSFW:
		return dec.BoolNull(&c.nsfw)

	case KeyLastMessageID:
		c.lastMsgID, err = com.DecodeSnowflake(dec)
		if c.lastMsgID == nil {
			c.setMetaFlag(maskNullLastMsgId)
		}

	case KeyBitrate:
		c.bitrate, err = DecodeBitrate(dec)

	case KeyUserLimit:
		c.userLim, err = DecodeUserLimit(dec)

	case KeyRateLimitPerUser:
		c.rLimPerUser, err = DecodeUserRateLimit(dec)

	case KeyRecipients:
		c.recipients = make(user.Slice, 0, 4)
		return dec.DecodeArray(&c.recipients)

	case KeyIcon:
		c.icon, err = DecodeImageHash(dec)
		if c.icon == nil {
			c.setMetaFlag(maskNullIcon)
		}

	case KeyOwnerID:
		c.ownerID, err = com.DecodeSnowflake(dec)

	case KeyApplicationID:
		c.appID, err = com.DecodeSnowflake(dec)

	case KeyParentID:
		c.parentID, err = com.DecodeSnowflake(dec)
		if c.parentID == nil {
			c.setMetaFlag(maskNullParentId)
		}

	case KeyLastPinTimestamp:
		c.lastPin, err = gj.DecodeTime(dec)
	}

	return
}

func (c *channel) NKeys() int {
	return 0
}

func (c *channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(utils.OutBuilder, 17).
		RawValue(KeyID, &c.id).
		RawValue(KeyType, &c.kind).
		Optional(KeyGuildID, c.guildID).
		Optional(KeyPosition, c.position).
		OptSlice(KeyPermissionOverwrites,
			c.hasMetaFlag(maskRecipients),
			c.permOver,
		).
		Optional(KeyName, c.name).
		TriState(KeyTopic, c.hasMetaFlag(maskNullTopic), c.topic).
		Optional(KeyNSFW, c.nsfw).
		TriState(
			KeyLastMessageID,
			c.hasMetaFlag(maskNullLastMsgId),
			c.lastMsgID,
		).
		Optional(KeyBitrate, c.bitrate).
		Optional(KeyUserLimit, c.userLim).
		Optional(KeyRateLimitPerUser, c.rLimPerUser).
		OptSlice(
			KeyRecipients,
			c.hasMetaFlag(maskRecipients),
			c.recipients,
		).
		TriState(KeyIcon, c.hasMetaFlag(maskNullIcon), c.icon).
		Optional(KeyOwnerID, c.ownerID).
		Optional(KeyApplicationID, c.appID).
		TriState(KeyParentID, c.hasMetaFlag(maskNullParentId), c.parentID).
		TriState(
			KeyLastPinTimestamp,
			c.hasMetaFlag(maskNullLastPin),
			c.lastPin,
		),
	)
}

func (c *channel) UnmarshalJSON(bytes []byte) (err error) {
	stage := make(map[Key]json.RawMessage, 8)

	if err := json.Unmarshal(bytes, &stage); err != nil {
		return err
	}

	// Simple types
	unm := map[Key]interface{}{
		KeyID:                   &c.id,
		KeyType:                 &c.kind,
		KeyPosition:             &c.position,
		KeyName:                 &c.name,
		KeyTopic:                &c.topic,
		KeyNSFW:                 &c.nsfw,
		KeyBitrate:              &c.bitrate,
		KeyUserLimit:            &c.userLim,
		KeyRateLimitPerUser:     &c.rLimPerUser,
		KeyIcon:                 &c.icon,
		KeyLastPinTimestamp:     &c.lastPin,
		KeyRecipients:           &c.recipients,
		KeyPermissionOverwrites: &c.permOver,
	}

	for k, v := range stage {
		if ptr, ok := unm[k]; ok {
			if err = json.Unmarshal(v, ptr); err != nil {
				return
			}

			continue
		}

		switch k {
		case KeyID:
			c.id = com.NewSnowflake()
			err = c.id.UnmarshalJSON(v)

		case KeyGuildID:
			c.guildID = com.NewSnowflake()
			err = c.guildID.UnmarshalJSON(v)

		case KeyLastMessageID:
			var tmp *string

			if err = json.Unmarshal(v, &tmp); err != nil {
				return
			}

			if tmp == nil {
				c.SetNullLastMessageID()
			} else {
				c.lastMsgID = com.NewSnowflake()
				err = c.lastMsgID.UnmarshalString(*tmp)
			}

		case KeyOwnerID:
			c.ownerID = com.NewSnowflake()
			err = c.ownerID.UnmarshalJSON(v)

		case KeyApplicationID:
			if c.appID, err = com.UnmarshalJSONSnowflake(v); err != nil {
				return
			}

			if c.appID == nil {
				c.SetNullParentID()
			}

		case KeyParentID:
			if c.parentID, err = com.UnmarshalJSONSnowflake(v); err != nil {
				return
			}

			if c.parentID == nil {
				c.SetNullParentID()
			}
		}
	}

	return
}

func (c *channel) Validate() error {
	out := lib.NewValidationErrorSet()

	if c.id == nil {
		out.AppendKeyedError(KeyID, "channel id is required")
	}

	singles := map[Key]lib.Validatable{
		KeyID:   c.id,
		KeyType: c.kind,
		KeyGuildID: c.guildID,
		KeyName: c.name,
		KeyTopic: c.topic,
		KeyLastMessageID: c.lastMsgID,
		KeyBitrate: c.bitrate,
		KeyUserLimit: c.userLim,
		KeyRateLimitPerUser: c.rLimPerUser,
		//KeyIcon: c.icon,
		KeyOwnerID: c.ownerID,
		KeyApplicationID: c.appID,
		KeyParentID: c.parentID,
	}

	for k, v := range singles {
		if v != nil {
			out.AppendRawKeyedError(k, v.Validate())
		}
	}

	for i := range c.permOver {
		if err := c.permOver[i].Validate(); err != nil {
			tmp := KeyPermissionOverwrites + Key(fmt.Sprintf("[%d]", i))
			out.AppendRawKeyedError(tmp, err)
		}
	}

	for i := range c.recipients {
		if err := c.recipients[i].Validate(); err != nil {
			tmp := KeyRecipients + Key(fmt.Sprintf("[%d]", i))
			out.AppendRawKeyedError(tmp, err)
		}
	}

	if out.Len() > 0 {
		return out
	}

	return nil
}

func (c *channel) IsValid() bool {
	return nil == c.Validate()
}

func (c *channel) ID() Snowflake {
	return c.id
}

func (c *channel) SetID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.id = id

	return c
}

func (c *channel) Type() ChannelType {
	return c.kind
}

func (c *channel) SetType(kind ChannelType) Channel {
	c.kind = kind
	return c
}

func (c *channel) GuildID() Snowflake {
	if c.guildID == nil {
		panic(e.ErrGetUnset)
	}

	return c.guildID
}

func (c *channel) GuildIDIsSet() bool {
	return c.guildID != nil
}

func (c *channel) SetGuildID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.guildID = id

	return c
}

func (c *channel) UnsetGuildID() Channel {
	c.guildID = nil
	return c
}

func (c *channel) Position() uint16 {
	if c.position == nil {
		panic(e.ErrGetUnset)
	}

	return *c.position
}

func (c *channel) PositionIsSet() bool {
	return c.position != nil
}

func (c *channel) SetPosition(u uint16) Channel {
	c.position = &u
	return c
}

func (c *channel) UnsetPosition() Channel {
	c.position = nil
	return c
}

func (c *channel) PermissionOverwrites() []PermissionOverwrite {
	if !c.PermissionOverwritesIsSet() {
		panic(e.ErrGetUnset)
	}

	return c.permOver
}

func (c *channel) PermissionOverwritesIsSet() bool {
	return c.hasMetaFlag(maskPermOver)
}

func (c *channel) SetPermissionOverwrites(o []PermissionOverwrite) Channel {
	c.permOver = o
	return c.unsetMetaFlag(maskPermOver)
}

func (c *channel) UnsetPermissionOverwrites() Channel {
	c.permOver = nil
	return c.setMetaFlag(maskPermOver)
}

func (c *channel) Name() ChannelName {
	if !c.NameIsSet() {
		panic(e.ErrGetUnset)
	}

	return *c.name
}

func (c *channel) NameIsSet() bool {
	return c.name != nil
}

func (c *channel) SetName(name ChannelName) Channel {
	c.name = &name
	return c
}

func (c *channel) UnsetName() Channel {
	c.name = nil
	return c
}

func (c *channel) Topic() ChannelTopic {
	if c.topic == nil {
		if c.hasMetaFlag(maskNullTopic) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return *c.topic
}

func (c *channel) TopicIsNull() bool {
	return c.topic == nil && c.hasMetaFlag(maskNullTopic)
}

func (c *channel) TopicIsSet() bool {
	return c.topic != nil
}

func (c *channel) TopicIsReadable() bool {
	return c.topic != nil
}

func (c *channel) SetTopic(topic ChannelTopic) Channel {
	c.topic = &topic
	return c.unsetMetaFlag(maskNullTopic)
}

func (c *channel) SetNullTopic() Channel {
	c.topic = nil
	return c.setMetaFlag(maskNullTopic)
}

func (c *channel) UnsetTopic() Channel {
	c.topic = nil
	return c.unsetMetaFlag(maskNullTopic)
}

func (c *channel) NSFW() bool {
	if c.nsfw == nil {
		panic(e.ErrGetUnset)
	}

	return *c.nsfw
}

func (c *channel) NSFWIsSet() bool {
	return c.nsfw != nil
}

func (c *channel) SetNSFW(b bool) Channel {
	c.nsfw = &b
	return c
}

func (c *channel) UnsetNSFW() Channel {
	c.nsfw = nil
	return c
}

func (c *channel) LastMessageID() Snowflake {
	if c.lastMsgID == nil {
		if c.hasMetaFlag(maskNullLastMsgId) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return c.lastMsgID
}

func (c *channel) LastMessageIDIsNull() bool {
	return c.lastMsgID == nil && c.hasMetaFlag(maskNullLastMsgId)
}

func (c *channel) LastMessageIDIsSet() bool {
	return c.lastMsgID != nil
}

func (c *channel) LastMessageIDIsReadable() bool {
	return c.lastMsgID != nil
}

func (c *channel) SetLastMessageID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.lastMsgID = id
	return c.unsetMetaFlag(maskNullLastMsgId)
}

func (c *channel) SetNullLastMessageID() Channel {
	c.lastMsgID = nil
	return c.setMetaFlag(maskNullLastMsgId)
}

func (c *channel) UnsetLastMessageID() Channel {
	c.lastMsgID = nil
	return c.unsetMetaFlag(maskNullLastMsgId)
}

func (c *channel) Bitrate() Bitrate {
	if c.bitrate == nil {
		panic(e.ErrGetUnset)
	}

	return *c.bitrate
}

func (c *channel) BitrateIsSet() bool {
	return c.bitrate != nil
}

func (c *channel) SetBitrate(bitrate Bitrate) Channel {
	c.bitrate = &bitrate
	return c
}

func (c *channel) UnsetBitrate() Channel {
	c.bitrate = nil
	return c
}

func (c *channel) UserLimit() UserLimit {
	if c.userLim == nil {
		panic(e.ErrGetUnset)
	}

	return *c.userLim
}

func (c *channel) UserLimitIsSet() bool {
	return c.userLim != nil
}

func (c *channel) SetUserLimit(limit UserLimit) Channel {
	c.userLim = &limit
	return c
}

func (c *channel) UnsetUserLimit() Channel {
	c.userLim = nil
	return c
}

func (c *channel) RateLimitPerUser() PerUserRateLimit {
	if c.rLimPerUser == nil {
		panic(e.ErrGetUnset)
	}

	return *c.rLimPerUser
}

func (c *channel) RateLimitPerUserIsSet() bool {
	return c.rLimPerUser != nil
}

func (c *channel) SetRateLimitPerUser(limit PerUserRateLimit) Channel {
	c.rLimPerUser = &limit
	return c
}

func (c *channel) UnsetRateLimitPerUser() Channel {
	c.rLimPerUser = nil
	return c
}

func (c *channel) Recipients() []User {
	if c.hasMetaFlag(maskRecipients) {
		panic(e.ErrGetUnset)
	}

	return c.recipients
}

func (c *channel) RecipientsIsSet() bool {
	return !c.hasMetaFlag(maskRecipients)
}

func (c *channel) SetRecipients(users []User) Channel {
	c.recipients = users
	return c.unsetMetaFlag(maskRecipients)
}

func (c *channel) UnsetRecipients() Channel {
	c.recipients = nil
	return c.setMetaFlag(maskRecipients)
}

func (c *channel) Icon() ImageHash {
	if c.id == nil {
		if c.hasMetaFlag(maskNullIcon) {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *c.icon
}

func (c *channel) IconIsNull() bool {
	return c.icon == nil && c.hasMetaFlag(maskNullIcon)
}

func (c *channel) IconIsSet() bool {
	return c.icon != nil
}

func (c *channel) IconIsReadable() bool {
	return c.icon != nil
}

func (c *channel) SetIcon(h ImageHash) Channel {
	c.icon = &h
	return c.unsetMetaFlag(maskNullIcon)
}

func (c *channel) SetNullIcon() Channel {
	c.icon = nil
	return c.setMetaFlag(maskNullIcon)
}

func (c *channel) UnsetIcon() Channel {
	c.icon = nil
	return c.unsetMetaFlag(maskNullIcon)
}

func (c *channel) OwnerID() Snowflake {
	if c.ownerID == nil {
		panic(e.ErrGetUnset)
	}

	return c.ownerID
}

func (c *channel) OwnerIDIsSet() bool {
	return c.ownerID != nil
}

func (c *channel) SetOwnerID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.ownerID = id

	return c
}

func (c *channel) UnsetOwnerID() Channel {
	c.ownerID = nil
	return c
}

func (c *channel) ApplicationID() Snowflake {
	if c.appID == nil {
		panic(e.ErrGetUnset)
	}

	return c.appID
}

func (c *channel) ApplicationIDIsSet() bool {
	return c.appID != nil
}

func (c *channel) SetApplicationID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.appID = id
	return c
}

func (c *channel) UnsetApplicationID() Channel {
	c.appID = nil
	return c
}

func (c *channel) ParentID() Snowflake {
	if c.parentID == nil {
		if c.hasMetaFlag(maskNullParentId) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return c.parentID
}

func (c *channel) ParentIDIsNull() bool {
	return c.parentID == nil && c.hasMetaFlag(maskNullParentId)
}

func (c *channel) ParentIDIsSet() bool {
	return c.parentID != nil
}

func (c *channel) ParentIDIsReadable() bool {
	return c.parentID != nil
}

func (c *channel) SetParentID(id Snowflake) Channel {
	if id == nil {
		panic(e.ErrSetNil)
	}

	c.id = id

	return c.unsetMetaFlag(maskNullParentId)
}

func (c *channel) SetNullParentID() Channel {
	c.parentID = nil
	return c.setMetaFlag(maskNullParentId)
}

func (c *channel) UnsetParentID() Channel {
	c.parentID = nil
	return c.unsetMetaFlag(maskNullLastPin)
}

func (c *channel) LastPinTimestamp() time.Time {
	if c.lastPin == nil {
		if c.hasMetaFlag(maskNullLastPin) {
			panic(e.ErrGetNull)
		} else {
			panic(e.ErrGetUnset)
		}
	}

	return *c.lastPin
}

func (c *channel) LastPinTimestampIsNull() bool {
	return c.lastPin == nil && c.hasMetaFlag(maskNullLastPin)
}

func (c *channel) LastPinTimestampIsSet() bool {
	return c.lastPin != nil
}

func (c *channel) LastPinTimestampIsReadable() bool {
	return c.lastPin != nil
}

func (c *channel) SetLastPinTimestamp(t time.Time) Channel {
	c.lastPin = &t
	return c.unsetMetaFlag(maskNullLastPin)
}

func (c *channel) SetNullLastPinTimestamp() Channel {
	c.lastPin = nil
	return c.setMetaFlag(maskNullLastPin)
}

func (c *channel) UnsetLastPinTimestamp() Channel {
	c.lastPin = nil
	return c.unsetMetaFlag(maskNullLastPin)
}

func (c *channel) hasMetaFlag(mask uint8) bool {
	return c.metaFlags&mask == mask
}

func (c *channel) setMetaFlag(mask uint8) Channel {
	c.metaFlags |= mask
	return c
}

func (c *channel) unsetMetaFlag(mask uint8) Channel {
	c.metaFlags &= ^mask
	return c
}
