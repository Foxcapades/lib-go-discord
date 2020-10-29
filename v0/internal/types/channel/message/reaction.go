package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type ReactionSlice []discord.MessageReaction

func (r ReactionSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range r {
		enc.AddObject(r[i])
	}
}

func (r *ReactionSlice) IsNil() bool {
	return false
}

func (r *ReactionSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewReaction()

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*r = append(*r, tmp)

	return nil
}

func NewReaction() discord.MessageReaction {
	return new(reaction)
}

type reaction struct{}

func (r *reaction) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (r *reaction) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (r *reaction) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (r *reaction) IsNil() bool {
	panic("implement me")
}

func (r *reaction) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (r *reaction) NKeys() int {
	panic("implement me")
}

func (r *reaction) IsValid() bool {
	panic("implement me")
}

func (r *reaction) Validate() error {
	panic("implement me")
}

func (r *reaction) Count() uint16 {
	panic("implement me")
}

func (r *reaction) SetCount(u uint16) discord.MessageReaction {
	panic("implement me")
}

func (r *reaction) Me() bool {
	panic("implement me")
}

func (r *reaction) SetMe(b bool) discord.MessageReaction {
	panic("implement me")
}

func (r *reaction) Emoji() discord.Emoji {
	panic("implement me")
}

func (r *reaction) SetEmoji(emoji discord.Emoji) discord.MessageReaction {
	panic("implement me")
}
