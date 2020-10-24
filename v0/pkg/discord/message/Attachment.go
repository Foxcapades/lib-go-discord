package message

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type Attachment interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the attachment's id.
	ID() discord.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(discord.Snowflake) Attachment

	// Filename returns the current value of this record's `filename` field.
	//
	// The `filename` field contains the name of the file attached.
	Filename() string

	// SetFilename overwrites the current value of this record's `filename` field.
	SetFilename(string) Attachment

	// Size returns the current value of this record's `size` field.
	//
	// The `size` field contains the size of the file in bytes.
	Size() uint64

	// SetSize overwrites the current value of this record's `size` field.
	SetSize(uint64) Attachment

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the source url of the file.
	URL() string

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) Attachment

	// ProxyURL returns the current value of this record's `proxy_url` field.
	//
	// The `proxy_url` field contains a proxied url of the file.
	ProxyURL() string

	// SetProxyURL overwrites the current value of this record's `proxy_url`
	// field.
	SetProxyURL(string) Attachment

	// Height returns the current value of this record's `height` field.
	//
	// The `height` field contains the height of the file (if image).
	Height() uint16

	// SetHeight overwrites the current value of this record's `height` field.
	SetHeight(uint16) Attachment

	// Width returns the current value of this record's `width` field.
	//
	// The `width` field contains the width of the file (if image).
	Width() uint16

	// SetWidth overwrites the current value of this record's `width` field.
	SetWidth(uint16) Attachment
}
