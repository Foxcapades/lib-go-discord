package message

import (
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type ChannelMentionSlice []discord.ChannelMention

func (c ChannelMentionSlice) JSONSize() uint32 {
	// len(c) - 1 -> number of commas between elements.
	out := js.BracketSize + uint32(len(c)-1)

	for i := range c {
		out += c[i].JSONSize()
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
