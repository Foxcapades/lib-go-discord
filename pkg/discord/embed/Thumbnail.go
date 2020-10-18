package embed

// EmbedThumbnail
// TODO: Document me
type Thumbnail interface {
	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the source url of thumbnail (only supports http(s)
	// and attachments).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use URLIsSet to check if the field is present before use.
	URL() string

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) Thumbnail

	// UnsetURL removes this record's `url` field.
	UnsetURL() Thumbnail

	// ProxyUrl returns the current value of this record's `proxy_url` field.
	//
	// The `proxy_url` field contains a proxied url of the thumbnail.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ProxyUrlIsSet to check if the field is present before use.
	ProxyUrl() string

	// ProxyUrlIsSet returns whether this record's `proxy_url` field is currently
	// present.
	ProxyUrlIsSet() bool

	// SetProxyUrl overwrites the current value of this record's `proxy_url`
	// field.
	SetProxyUrl(string) Thumbnail

	// UnsetProxyUrl removes this record's `proxy_url` field.
	UnsetProxyUrl() Thumbnail

	// Height returns the current value of this record's `height` field.
	//
	// The `height` field contains the height of the thumbnail.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use HeightIsSet to check if the field is present before use.
	Height() uint16

	// HeightIsSet returns whether this record's `height` field is currently
	// present.
	HeightIsSet() bool

	// SetHeight overwrites the current value of this record's `height` field.
	SetHeight(uint16) Thumbnail

	// UnsetHeight removes this record's `height` field.
	UnsetHeight() Thumbnail

	// Width returns the current value of this record's `width` field.
	//
	// The `width` field contains the width of the thumbnail.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use WidthIsSet to check if the field is present before use.
	Width() uint16

	// WidthIsSet returns whether this record's `width` field is currently
	// present.
	WidthIsSet() bool

	// SetWidth overwrites the current value of this record's `width` field.
	SetWidth(uint16) Thumbnail

	// UnsetWidth removes this record's `width` field.
	UnsetWidth() Thumbnail
}
