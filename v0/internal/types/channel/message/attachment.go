package message

import (
	"bytes"
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/ugly"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"
	"io"
)

const (
	baseAttachSize = uint32(js.BracketSize +
		js.FirstFieldSize + len(KeyID) +
		js.NextFieldSize + len(KeyFilename) +
		js.NextFieldSize + len(KeySize) +
		js.NextFieldSize + len(KeyURL) +
		js.NextFieldSize + len(KeyProxyURL) +
		js.NextFieldSize + len(KeyHeight) +
		js.NextFieldSize + len(KeyWidth))
)

func NewAttachment() discord.MessageAttachment {
	return new(attachment)
}

type attachment struct {
	id       discord.Snowflake
	fileName ugly.PseudoString
	size     uint32
	url      ugly.PseudoString
	proxyUrl ugly.PseudoString
	height   *uint16
	width    *uint16
}

func (a *attachment) JSONSize() uint32 {
	return baseAttachSize +
		a.id.JSONSize() +
		a.fileName.JSONSize() +
		uint32(bytify.Uint32StringSize(a.size)) +
		a.url.JSONSize() +
		a.proxyUrl.JSONSize() +
		js.SizeUint16OrNull(a.height) +
		js.SizeUint16OrNull(a.width)
}

func (a *attachment) ToJSONBytes() []byte {
	buf := new(bytes.Buffer)
	buf.Grow(int(a.JSONSize()))
	buf.Reset()

	buf.WriteByte(js.OpenObject)

	_ = KeyID.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	_ = a.id.AppendJSONBytes(buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeyFilename.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	_ = a.fileName.AppendJSONBytes(buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeySize.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	bytify.Uint32ToBuf(a.size, buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeyURL.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	_ = a.url.AppendJSONBytes(buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeyProxyURL.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	_ = a.proxyUrl.AppendJSONBytes(buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeyHeight.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	bytify.Uint16ToBuf(*a.height, buf)

	buf.WriteByte(js.FieldDivider)

	_ = KeyWidth.AppendJSONBytes(buf)
	buf.WriteByte(js.PairSeparator)
	bytify.Uint16ToBuf(*a.width, buf)

	return nil
}

func (a *attachment) AppendJSONBytes(buf io.Writer) (e error) {
	_, e = buf.Write(a.ToJSONBytes())
	return
}

func (a *attachment) MarshalJSON() ([]byte, error) {
	return a.ToJSONBytes(), nil
}

func (a *attachment) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (a *attachment) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (a *attachment) IsNil() bool {
	return a == nil
}

func (a *attachment) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (a *attachment) NKeys() int {
	return 0
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

func (a *attachment) Size() uint32 {
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
