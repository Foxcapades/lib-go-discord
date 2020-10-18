package message

// EmbedVideo
// TODO: Document me
type EmbedVideo interface {
	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the source url of the video.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use URLIsSet to check if the field is present before use.
	URL() string

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) EmbedVideo

	// UnsetURL removes this record's `url` field.
	UnsetURL() EmbedVideo

	// Height returns the current value of this record's `height` field.
	//
	// The `height` field contains the height of the video.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use HeightIsSet to check if the field is present before use.
	Height() uint16

	// HeightIsSet returns whether this record's `height` field is currently
	// present.
	HeightIsSet() bool

	// SetHeight overwrites the current value of this record's `height` field.
	SetHeight(uint16) EmbedVideo

	// UnsetHeight removes this record's `height` field.
	UnsetHeight() EmbedVideo

	// Width returns the current value of this record's `width` field.
	//
	// The `width` field contains width of the video.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use WidthIsSet to check if the field is present before use.
	Width() uint16

	// WidthIsSet returns whether this record's `width` field is currently
	// present.
	WidthIsSet() bool

	// SetWidth overwrites the current value of this record's `width` field.
	SetWidth(uint16) EmbedVideo

	// UnsetWidth removes this record's `width` field.
	UnsetWidth() EmbedVideo
}
