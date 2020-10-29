package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"
)

type MessageAttachment interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the attachment's id.
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) MessageAttachment

	// Filename returns the current value of this record's `filename` field.
	//
	// The `filename` field contains the name of the file attached.
	Filename() string

	// SetFilename overwrites the current value of this record's `filename` field.
	SetFilename(string) MessageAttachment

	// Size returns the current value of this record's `size` field.
	//
	// The `size` field contains the size of the file in bytes.
	Size() uint64

	// SetSize overwrites the current value of this record's `size` field.
	SetSize(uint64) MessageAttachment

	// URL returns the current value of this record's `url` field.
	//
	// The `url` field contains the source url of the file.
	URL() string

	// SetURL overwrites the current value of this record's `url` field.
	SetURL(string) MessageAttachment

	// ProxyURL returns the current value of this record's `proxy_url` field.
	//
	// The `proxy_url` field contains a proxied url of the file.
	ProxyURL() string

	// SetProxyURL overwrites the current value of this record's `proxy_url`
	// field.
	SetProxyURL(string) MessageAttachment

	// Height returns the current value of this record's `height` field.
	//
	// The `height` field contains the height of the file (if image).
	Height() uint16

	// SetHeight overwrites the current value of this record's `height` field.
	SetHeight(uint16) MessageAttachment

	// Width returns the current value of this record's `width` field.
	//
	// The `width` field contains the width of the file (if image).
	Width() uint16

	// SetWidth overwrites the current value of this record's `width` field.
	SetWidth(uint16) MessageAttachment
}
