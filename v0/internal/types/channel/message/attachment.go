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

func NewAttachment() discord.MessageAttachment {
	return new(attachment)
}

type attachment struct {
	id       discord.Snowflake
	fileName string
	size     uint64
	url      string
	proxyUrl string
	height   *uint16
	width    *uint16
}

func (a *attachment) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (a *attachment) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (a *attachment) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (a *attachment) IsNil() bool {
	panic("implement me")
}

func (a *attachment) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (a *attachment) NKeys() int {
	panic("implement me")
}

func (a *attachment) IsValid() bool {
	panic("implement me")
}

func (a *attachment) Validate() error {
	panic("implement me")
}

func (a *attachment) ID() discord.Snowflake {
	panic("implement me")
}

func (a *attachment) SetID(id discord.Snowflake) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) Filename() string {
	panic("implement me")
}

func (a *attachment) SetFilename(s string) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) Size() uint64 {
	panic("implement me")
}

func (a *attachment) SetSize(u uint64) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) URL() string {
	panic("implement me")
}

func (a *attachment) SetURL(s string) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) ProxyURL() string {
	panic("implement me")
}

func (a *attachment) SetProxyURL(s string) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) Height() uint16 {
	panic("implement me")
}

func (a *attachment) SetHeight(u uint16) discord.MessageAttachment {
	panic("implement me")
}

func (a *attachment) Width() uint16 {
	panic("implement me")
}

func (a *attachment) SetWidth(u uint16) discord.MessageAttachment {
	panic("implement me")
}
