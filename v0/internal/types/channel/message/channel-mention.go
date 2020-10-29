package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type ChannelMentionSlice []discord.ChannelMention

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

func (c channelMention) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (c channelMention) UnmarshalJSON(bytes []byte) error {
	panic("implement me")
}

func (c channelMention) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (c channelMention) IsNil() bool {
	panic("implement me")
}

func (c channelMention) UnmarshalJSONObject(decoder *gojay.Decoder, s string) error {
	panic("implement me")
}

func (c channelMention) NKeys() int {
	panic("implement me")
}

func (c channelMention) IsValid() bool {
	panic("implement me")
}

func (c channelMention) Validate() error {
	panic("implement me")
}

func (c channelMention) ID() discord.Snowflake {
	panic("implement me")
}

func (c channelMention) SetID(snowflake discord.Snowflake) discord.ChannelMention {
	panic("implement me")
}

func (c channelMention) GuildID() discord.Snowflake {
	panic("implement me")
}

func (c channelMention) SetGuildID(snowflake discord.Snowflake) discord.ChannelMention {
	panic("implement me")
}

func (c channelMention) Type() discord.ChannelType {
	panic("implement me")
}

func (c channelMention) SetType(channelType discord.ChannelType) discord.ChannelMention {
	panic("implement me")
}

func (c channelMention) Name() string {
	panic("implement me")
}

func (c channelMention) SetName(s string) discord.ChannelMention {
	panic("implement me")
}
