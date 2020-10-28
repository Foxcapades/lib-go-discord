package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"
)

var (
	ErrCurrentIsGtMax = errors.New("activity party's current_size is greater than its max size")
)

type ActivityParty interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IDIsSet to check if the field is present before use.
	ID() string

	// IDIsSet returns whether this record's `id` field is currently present.
	IDIsSet() bool

	// SetID overwrites the current value of this record's `id` field.
	SetID(string) ActivityParty

	// UnsetID removes this record's `id` field.
	UnsetID() ActivityParty

	// CurrentSize returns the current first value of this record's `size` field.
	//
	// The `size` field contains the party's current and maximum size.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use CurrentSizeIsSet to check if the field is present before use.
	CurrentSize() uint16

	// MaxSize returns the current second value of this record's `size` field.
	//
	// The `size` field contains the party's current and maximum size.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MaxSizeIsSet to check if the field is present before use.
	MaxSize() uint16

	// SizesAreSet returns whether this record's `size` field is currently
	// present.
	SizesAreSet() bool

	// SetCurrentSize overwrites the current first value of this record's `size`
	// field.
	SetSizes(current, max uint16) ActivityParty

	// UnsetCurrentSize removes this record's `size` field.
	UnsetSizes() ActivityParty
}
