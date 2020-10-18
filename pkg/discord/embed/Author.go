package embed

type Author interface {
	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the author.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NameIsSet to check if the field is present before use.
	Name() string

	// NameIsSet returns whether this record's `name` field is currently present.
	NameIsSet() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) Author

	// UnsetName removes this record's `name` field.
	UnsetName() Author

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the url of the author.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use URLIsSet to check if the field is present before use.
	URL() string

	// URLIsSet returns whether this record's `url` field is currently present.
	URLIsSet() bool

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) Author

	// UnsetURL removes this record's `url` field.
	UnsetURL() Author

	// IconURL returns the current value of this record's `icon_url` field.
	//
	// The `icon_url` field contains the url of the author icon (only supports
	// http(s) and attachments).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IconURLIsSet to check if the field is present before use.
	IconURL() string

	// IconURLIsSet returns whether this record's `icon_url` field is currently
	// present.
	IconURLIsSet() bool

	// SetIconURL overwrites the current value of this record's `icon_url` field.
	SetIconURL(string) Author

	// UnsetIconURL removes this record's `icon_url` field.
	UnsetIconURL() Author

	// ProxyIconURL returns the current value of this record's `proxy_icon_url`
	// field.
	//
	// The `proxy_icon_url` field contains a proxied url of the author icon.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ProxyIconURLIsSet to check if the field is present before use.
	ProxyIconURL() string

	// ProxyIconURLIsSet returns whether this record's `proxy_icon_url` field is
	// currently present.
	ProxyIconURLIsSet() bool

	// SetProxyIconURL overwrites the current value of this record's
	// `proxy_icon_url` field.
	SetProxyIconURL(string) Author

	// UnsetProxyIconURL removes this record's `proxy_icon_url` field.
	UnsetProxyIconURL() Author
}
