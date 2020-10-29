package message

import (
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
)

type EmbedSlice []discord.MessageEmbed

func (e EmbedSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range e {
		enc.AddObject(e[i])
	}
}

func (e EmbedSlice) IsNil() bool {
	return false
}

func (e *EmbedSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewEmbed()

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*e = append(*e, tmp)

	return nil
}

func NewEmbed() discord.MessageEmbed {
	return new(embed)
}

type embed struct{}

func (e *embed) MarshalJSON() ([]byte, error) {
	panic("implement me")
}

func (e *embed) UnmarshalJSON(in []byte) error {
	panic("implement me")
}

func (e *embed) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (e *embed) IsNil() bool {
	panic("implement me")
}

func (e *embed) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (e *embed) NKeys() int {
	panic("implement me")
}

func (e *embed) IsValid() bool {
	panic("implement me")
}

func (e *embed) Validate() error {
	panic("implement me")
}

func (e *embed) Title() string {
	panic("implement me")
}

func (e *embed) TitleIsSet() bool {
	panic("implement me")
}

func (e *embed) SetTitle(s string) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetTitle() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Type() discord.EmbedType {
	panic("implement me")
}

func (e *embed) TypeIsSet() bool {
	panic("implement me")
}

func (e *embed) SetType(kind discord.EmbedType) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetType() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Description() string {
	panic("implement me")
}

func (e *embed) DescriptionIsSet() bool {
	panic("implement me")
}

func (e *embed) SetDescription(s string) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetDescription() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) URL() string {
	panic("implement me")
}

func (e *embed) URLIsSet() bool {
	panic("implement me")
}

func (e *embed) SetURL(s string) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetURL() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Timestamp() time.Time {
	panic("implement me")
}

func (e *embed) TimestampIsSet() bool {
	panic("implement me")
}

func (e *embed) SetTimestamp(t time.Time) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetTimestamp() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Color() lib.Color {
	panic("implement me")
}

func (e *embed) ColorIsSet() bool {
	panic("implement me")
}

func (e *embed) SetColor(color lib.Color) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetColor() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Footer() discord.EmbedFooter {
	panic("implement me")
}

func (e *embed) FooterIsSet() bool {
	panic("implement me")
}

func (e *embed) SetFooter(footer discord.EmbedFooter) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetFooter() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Image() discord.EmbedImage {
	panic("implement me")
}

func (e *embed) ImageIsSet() bool {
	panic("implement me")
}

func (e *embed) SetImage(image discord.EmbedImage) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetImage() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Thumbnail() discord.EmbedThumbnail {
	panic("implement me")
}

func (e *embed) ThumbnailIsSet() bool {
	panic("implement me")
}

func (e *embed) SetThumbnail(thumbnail discord.EmbedThumbnail) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetThumbnail() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Video() discord.EmbedVideo {
	panic("implement me")
}

func (e *embed) VideoIsSet() bool {
	panic("implement me")
}

func (e *embed) SetVideo(video discord.EmbedVideo) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetVideo() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Provider() discord.EmbedProvider {
	panic("implement me")
}

func (e *embed) ProviderIsSet() bool {
	panic("implement me")
}

func (e *embed) SetProvider(provider discord.EmbedProvider) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetProvider() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Author() discord.EmbedAuthor {
	panic("implement me")
}

func (e *embed) AuthorIsSet() bool {
	panic("implement me")
}

func (e *embed) SetAuthor(author discord.EmbedAuthor) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetAuthor() discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) Fields() []discord.EmbedField {
	panic("implement me")
}

func (e *embed) FieldsIsSet() bool {
	panic("implement me")
}

func (e *embed) SetFields(fields []discord.EmbedField) discord.MessageEmbed {
	panic("implement me")
}

func (e *embed) UnsetFields() discord.MessageEmbed {
	panic("implement me")
}
