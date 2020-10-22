package discord

type EmbedFooter interface {
	// Text returns the current value of this record's `text` field.
	//
	// The `text` field contains the footer text.
	Text() string

	// SetText overwrites the current value of this record's `text` field.
	SetText(string) EmbedFooter

	// IconURL returns the current value of this record's `icon_url` field.
	//
	// The `icon_url` field contains the url of the footer icon (only supports
	// http(s) and attachments).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IconURLIsSet to check if the field is present before use.
	IconURL() string

	// IconURLIsSet returns whether this record's `icon_url` field is currently
	// present.
	IconURLIsSet() bool

	// SetIconURL overwrites the current value of this record's `icon_url` field.
	SetIconURL(string) EmbedFooter

	// UnsetIconURL removes this record's `icon_url` field.
	UnsetIconURL() EmbedFooter

	// ProxyIconURL returns the current value of this record's `proxy_icon_url`
	// field.
	//
	// The `proxy_icon_url` field contains the a proxied url of the footer icon.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ProxyIconURLIsSet to check if the field is present before use.
	ProxyIconURL() string

	// ProxyIconURLIsSet returns whether this record's `proxy_icon_url` field is
	// currently present.
	ProxyIconURLIsSet() bool

	// SetProxyIconURL overwrites the current value of this record's
	// `proxy_icon_url` field.
	SetProxyIconURL(string) EmbedFooter

	// UnsetProxyIconURL removes this record's `proxy_icon_url` field.
	UnsetProxyIconURL() EmbedFooter
}
