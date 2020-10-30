package message

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"
)

type ChannelMentionSlice []discord.ChannelMention

func (c ChannelMentionSlice) BufferSize() uint32 {
	out := uint32(2) + uint32(len(c) - 1)

	for i := range c {
		out += c[i].BufferSize()
	}

	return out
}

func (c ChannelMentionSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range c {
		enc.AddObject(c[i])
	}
}

func (c ChannelMentionSlice) IsNil() bool {
	return false
}

func (c *ChannelMentionSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewChannelMention()

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*c = append(*c, tmp)

	return nil
}

func NewChannelMention() discord.ChannelMention {
	return new(channelMention)
}

type channelMention struct {
	id      discord.Snowflake
	guildID discord.Snowflake
	kind    discord.ChannelType
	name    discord.ChannelName
}

const (
	chanMentionBaseSize = uint32(2 +
		len(serial.KeyID) + 3 +
		len(serial.KeyGuildID) + 4 +
		len(serial.KeyType) + 4 +
		len(serial.KeyName) + 4)
)
func (c *channelMention) BufferSize() uint32 {
	return chanMentionBaseSize +
		utils.OptionalSize(c.id) +
		utils.OptionalSize(c.guildID) +
		utils.OptionalSize(c.kind) +
		utils.OptionalSize(c.name)
}

func (c *channelMention) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, c.BufferSize()))
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
	panic("implement me")
}

func (c *channelMention) IsNil() bool {
	panic("implement me")
}

func (c *channelMention) UnmarshalJSONObject(decoder *gojay.Decoder, s string) error {
	panic("implement me")
}

func (c *channelMention) NKeys() int {
	panic("implement me")
}

func (c *channelMention) IsValid() bool {
	panic("implement me")
}

func (c *channelMention) Validate() error {
	panic("implement me")
}

func (c *channelMention) ID() discord.Snowflake {
	panic("implement me")
}

func (c *channelMention) SetID(snowflake discord.Snowflake) discord.ChannelMention {
	panic("implement me")
}

func (c *channelMention) GuildID() discord.Snowflake {
	panic("implement me")
}

func (c *channelMention) SetGuildID(snowflake discord.Snowflake) discord.ChannelMention {
	panic("implement me")
}

func (c *channelMention) Type() discord.ChannelType {
	panic("implement me")
}

func (c *channelMention) SetType(channelType discord.ChannelType) discord.ChannelMention {
	panic("implement me")
}

func (c *channelMention) Name() string {
	panic("implement me")
}

func (c *channelMention) SetName(s string) discord.ChannelMention {
	panic("implement me")
}
