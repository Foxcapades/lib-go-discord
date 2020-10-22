package discord

type EmbedProvider interface {
	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the provider.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NameIsSet to check if the field is present before use.
	Name() string

	// NameIsSet returns whether this record's `name` field is currently present.
	NameIsSet() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) EmbedProvider

	// UnsetName removes this record's `name` field.
	UnsetName() EmbedProvider

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the url of the provider.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use URLIsSet to check if the field is present before use.
	URL() string

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) EmbedProvider

	// UnsetURL removes this record's `url` field.
	UnsetURL() EmbedProvider
}
