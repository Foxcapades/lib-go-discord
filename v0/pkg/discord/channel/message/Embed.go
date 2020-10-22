package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/embed"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"time"
)

// Embed
// TODO: document me
type Embed interface {
	// Title returns the current value of this record's `title` field.
	//
	// The `title` field contains the title of this embed.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TitleIsSet to check if the field is present before use.
	Title() string

	// TitleIsSet returns whether this record's `title` field is currently
	// present.
	TitleIsSet() bool

	// SetTitle overwrites the current value of this record's `title` field.
	SetTitle(string) Embed

	// UnsetTitle removes this record's `title` field.
	UnsetTitle() Embed

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of this embed (always "rich" for webhook
	// embeds).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TypeIsSet to check if the field is present before use.
	Type() embed.Type

	// TypeIsSet returns whether this record's `type` field is currently present.
	TypeIsSet() bool

	// SetType overwrites the current value of this record's `type` field.
	SetType(embed.Type) Embed

	// UnsetType removes this record's `type` field.
	UnsetType() Embed

	// Description returns the current value of this record's `description` field.
	//
	// The `description` field contains the description of this embed.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use DescriptionIsSet to check if the field is present before use.
	Description() string

	// DescriptionIsSet returns whether this record's `description` field is
	// currently present.
	DescriptionIsSet() bool

	// SetDescription overwrites the current value of this record's `description`
	// field.
	SetDescription(string) Embed

	// UnsetDescription removes this record's `description` field.
	UnsetDescription() Embed

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the url of this embed.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use URLIsSet to check if the field is present before use.
	URL() string

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) Embed

	// UnsetURL removes this record's `url` field.
	UnsetURL() Embed

	// Timestamp returns the current value of this record's `timestamp` field.
	//
	// The `timestamp` field contains the timestamp of this embed's content.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TimestampIsSet to check if the field is present before use.
	Timestamp() time.Time

	// TimestampIsSet returns whether this record's `timestamp` field is currently
	// present.
	TimestampIsSet() bool

	// SetTimestamp overwrites the current value of this record's `timestamp`
	// field.
	SetTimestamp(time.Time) Embed

	// UnsetTimestamp removes this record's `timestamp` field.
	UnsetTimestamp() Embed

	// Color returns the current value of this record's `color` field.
	//
	// The `color` field contains the color code of this embed.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ColorIsSet to check if the field is present before use.
	Color() comm.Color

	// ColorIsSet returns whether this record's `color` field is currently
	// present.
	ColorIsSet() bool

	// SetColor overwrites the current value of this record's `color` field.
	SetColor(comm.Color) Embed

	// UnsetColor removes this record's `color` field.
	UnsetColor() Embed

	// Footer returns the current value of this record's `footer` field.
	//
	// The `footer` field contains this embed's footer information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use FooterIsSet to check if the field is present before use.
	Footer() embed.Footer

	// FooterIsSet returns whether this record's `footer` field is currently
	// present.
	FooterIsSet() bool

	// SetFooter overwrites the current value of this record's `footer` field.
	SetFooter(embed.Footer) Embed

	// UnsetFooter removes this record's `footer` field.
	UnsetFooter() Embed

	// Image returns the current value of this record's `image` field.
	//
	// The `image` field contains this embed's image information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ImageIsSet to check if the field is present before use.
	Image() embed.Image

	// ImageIsSet returns whether this record's `image` field is currently
	// present.
	ImageIsSet() bool

	// SetImage overwrites the current value of this record's `image` field.
	SetImage(embed.Image) Embed

	// UnsetImage removes this record's `image` field.
	UnsetImage() Embed

	// Thumbnail returns the current value of this record's `thumbnail` field.
	//
	// The `thumbnail` field contains this embed's thumbnail information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ThumbnailIsSet to check if the field is present before use.
	Thumbnail() embed.Thumbnail

	// ThumbnailIsSet returns whether this record's `thumbnail` field is currently
	// present.
	ThumbnailIsSet() bool

	// SetThumbnail overwrites the current value of this record's `thumbnail`
	// field.
	SetThumbnail(embed.Thumbnail) Embed

	// UnsetThumbnail removes this record's `thumbnail` field.
	UnsetThumbnail() Embed

	// Video returns the current value of this record's `video` field.
	//
	// The `video` field contains this embed's video information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use VideoIsSet to check if the field is present before use.
	Video() embed.Video

	// VideoIsSet returns whether this record's `video` field is currently
	// present.
	VideoIsSet() bool

	// SetVideo overwrites the current value of this record's `video` field.
	SetVideo(embed.Video) Embed

	// UnsetVideo removes this record's `video` field.
	UnsetVideo() Embed

	// Provider returns the current value of this record's `provider` field.
	//
	// The `provider` field contains this embed's provider information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ProviderIsSet to check if the field is present before use.
	Provider() embed.Provider

	// ProviderIsSet returns whether this record's `provider` field is currently
	// present.
	ProviderIsSet() bool

	// SetProvider overwrites the current value of this record's `provider` field.
	SetProvider(embed.Provider) Embed

	// UnsetProvider removes this record's `provider` field.
	UnsetProvider() Embed

	// Author returns the current value of this record's `author` field.
	//
	// The `author` field contains this embed's author information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AuthorIsSet to check if the field is present before use.
	Author() embed.Author

	// AuthorIsSet returns whether this record's `author` field is currently
	// present.
	AuthorIsSet() bool

	// SetAuthor overwrites the current value of this record's `author` field.
	SetAuthor(embed.Author) Embed

	// UnsetAuthor removes this record's `author` field.
	UnsetAuthor() Embed

	// Fields returns the current value of this record's `fields` field.
	//
	// The `fields` field contains this embed's fields information.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use FieldsIsSet to check if the field is present before use.
	Fields() []embed.Field

	// FieldsIsSet returns whether this record's `fields` field is currently
	// present.
	FieldsIsSet() bool

	// SetFields overwrites the current value of this record's `fields` field.
	SetFields([]embed.Field) Embed

	// UnsetFields removes this record's `fields` field.
	UnsetFields() Embed
}


type TriStateEmbedField struct {
	value Embed
	null  bool
}

func (t TriStateEmbedField) IsSet() bool {
	return t.value != nil
}

func (t TriStateEmbedField) IsUnset() bool {
	return t.value == nil && !t.null
}

func (t TriStateEmbedField) IsNull() bool {
	return t.value == nil && t.null
}

func (t TriStateEmbedField) IsNotNull() bool {
	return !t.null
}

func (t TriStateEmbedField) IsReadable() bool {
	return t.value != nil
}

func (t TriStateEmbedField) SetNull() {
	t.value = nil
	t.null = true
}

func (t TriStateEmbedField) Unset() {
	t.value = nil
	t.null = false
}

func (t TriStateEmbedField) Get() Embed {
	if t.value == nil {
		if t.null {
			panic(dlib.ErrNullField)
		} else {
			panic(dlib.ErrUnsetField)
		}
	}

	return t.value
}

func (t TriStateEmbedField) Set(e Embed) {
	if e == nil {
		panic(dlib.ErrSetNilTriState)
	}

	t.value = e
	t.null = false
}
