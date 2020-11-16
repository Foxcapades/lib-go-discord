package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type AttachmentSlice []discord.MessageAttachment

func (a AttachmentSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range a {
		enc.AddObject(a[i])
	}
}

func (a AttachmentSlice) IsNil() bool {
	return false
}

func (a *AttachmentSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewAttachment()

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*a = append(*a, tmp)

	return nil
}
