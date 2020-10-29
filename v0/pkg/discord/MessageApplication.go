package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"
)

type MessageApplication interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the application.
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) MessageApplication

	// CoverImage returns the current value of this record's `cover_image` field.
	//
	// The `cover_image` field contains the id of the embed's image asset.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use CoverImageIsSet to check if the field is present before use.
	CoverImage() string

	// CoverImageIsSet returns whether this record's `cover_image` field is
	// currently present.
	CoverImageIsSet() bool

	// SetCoverImage overwrites the current value of this record's `cover_image`
	// field.
	SetCoverImage(string) MessageApplication

	// UnsetCoverImage removes this record's `cover_image` field.
	UnsetCoverImage() MessageApplication

	// Description returns the current value of this record's `description` field.
	//
	// The `description` field contains the application's description.
	Description() string

	// SetDescription overwrites the current value of this record's `description`
	// field.
	SetDescription(string) MessageApplication

	// Icon returns the current value of this record's `icon` field.
	//
	// The `icon` field contains the id of the application's icon.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use IconIsNull to check if the field is null before use.
	Icon() string

	// IconIsNull returns whether this record's `icon` field is currently null.
	IconIsNull() bool

	// SetIcon overwrites the current value of this record's `icon` field.
	SetIcon(string) MessageApplication

	// SetNullIcon overwrites the current value of this record's `icon` field
	// with `null`.
	SetNullIcon() MessageApplication

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the application.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) MessageApplication
}
