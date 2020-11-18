package message

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
	"github.com/francoispqt/gojay"
)

func NewChannelMention() discord.ChannelMention {
	return &channelMention{
		id:      com.NewSnowflake(),
		guildID: com.NewSnowflake(),
	}
}

const (
	chanMentionBaseSize = js.BracketSize +
		uint32(len(serial.KeyID)) + js.FirstKeySize +
		uint32(len(serial.KeyGuildID)) + js.NextKeySize +
		uint32(len(serial.KeyType)) + js.NextKeySize +
		uint32(len(serial.KeyName)) + js.NextKeySize
)

type channelMention struct {
	id      discord.Snowflake
	guildID discord.Snowflake
	kind    discord.ChannelType
	name    discord.ChannelName
}

func (c *channelMention) JSONSize() uint32 {
	return chanMentionBaseSize +
		utils.OptionalSize(c.id) +
		utils.OptionalSize(c.guildID) +
		utils.OptionalSize(&c.kind) +
		utils.OptionalSize(&c.name)
}

func (c *channelMention) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, c.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeObject(c); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *channelMention) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(c)
}

func (c *channelMention) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddSnowflake(serial.KeyID, c.id).
		AddSnowflake(serial.KeyGuildID, c.guildID).
		AddUint8(serial.KeyType, uint8(c.kind)).
		AddString(serial.KeyName, string(c.name))
}

func (c *channelMention) IsNil() bool {
	return c == nil
}

func (c *channelMention) UnmarshalJSONObject(decoder *gojay.Decoder, s string) (err error) {
	switch serial.Key(s) {
	case serial.KeyID:
		c.id, err = com.DecodeSnowflake(decoder)
	case serial.KeyGuildID:
		c.guildID, err = com.DecodeSnowflake(decoder)
	case serial.KeyType:
		err = decoder.DecodeUint8((*uint8)(&c.kind))
	case serial.KeyName:
		err = decoder.DecodeString((*string)(&c.name))
	}

	return
}

func (c *channelMention) NKeys() int {
	return 0
}

func (c *channelMention) IsValid() bool {
	return c.Validate() == nil
}

func (c *channelMention) Validate() error {
	return lib.NewValidationErrorSet().
		AppendRawKeyedError(serial.KeyID, c.id.Validate()).
		AppendRawKeyedError(serial.KeyGuildID, c.guildID.Validate()).
		AppendRawKeyedError(serial.KeyType, c.kind.Validate()).
		AppendRawKeyedError(serial.KeyName, c.name.Validate())
}

func (c *channelMention) ID() discord.Snowflake {
	return c.id
}

func (c *channelMention) SetID(snowflake discord.Snowflake) discord.ChannelMention {
	if snowflake == nil {
		panic(e.ErrSetNil)
	}

	c.id = snowflake

	return c
}

func (c *channelMention) GuildID() discord.Snowflake {
	return c.guildID
}

func (c *channelMention) SetGuildID(snowflake discord.Snowflake) discord.ChannelMention {
	if snowflake == nil {
		panic(e.ErrSetNil)
	}

	c.guildID = snowflake

	return c
}

func (c *channelMention) Type() discord.ChannelType {
	return c.kind
}

func (c *channelMention) SetType(channelType discord.ChannelType) discord.ChannelMention {
	c.kind = channelType
	return c
}

func (c *channelMention) Name() discord.ChannelName {
	return c.name
}

func (c *channelMention) SetName(s discord.ChannelName) discord.ChannelMention {
	c.name = s
	return c
}
