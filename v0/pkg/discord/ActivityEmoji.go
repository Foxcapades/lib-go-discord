package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
	"github.com/francoispqt/gojay"
)

type ActivityEmoji interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	dmeta.Validatable

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the emoji.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) ActivityEmoji

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the id of the emoji.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IDIsSet to check if the field is present before use.
	ID() Snowflake

	// IDIsSet returns whether this record's `id` field is currently present.
	IDIsSet() bool

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) ActivityEmoji

	// UnsetID removes this record's `id` field.
	UnsetID() ActivityEmoji

	// Animated returns the current value of this record's `animated` field.
	//
	// The `animated` field indicates whether this emoji is animated.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AnimatedIsSet to check if the field is present before use.
	Animated() bool

	// AnimatedIsSet returns whether this record's `animated` field is currently present.
	AnimatedIsSet() bool

	// SetAnimated overwrites the current value of this record's `animated` field.
	SetAnimated(bool) ActivityEmoji

	// UnsetAnimated removes this record's `animated` field.
	UnsetAnimated() ActivityEmoji
}
