package discord

import (
	"encoding/json"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
)

type ActivityTimestamps interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	dmeta.Validatable

	// Start returns the current value of this record's `start` field.
	//
	// The `start` field contains the unix time (in milliseconds) of when the
	// activity started.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use StartIsSet to check if the field is present before use.
	Start() time.Time

	// StartIsSet returns whether this record's `start` field is currently present.
	StartIsSet() bool

	// SetStart overwrites the current value of this record's `start` field.
	SetStart(time.Time) ActivityTimestamps

	// UnsetStart removes this record's `start` field.
	UnsetStart() ActivityTimestamps

	// End returns the current value of this record's `end` field.
	//
	// The `end` field contains the unix time (in milliseconds) of when the
	// activity ends.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use EndIsSet to check if the field is present before use.
	End() time.Time

	// EndIsSet returns whether this record's `end` field is currently present.
	EndIsSet() bool

	// SetEnd overwrites the current value of this record's `end` field.
	SetEnd(time.Time) ActivityTimestamps

	// UnsetEnd removes this record's `end` field.
	UnsetEnd() ActivityTimestamps
}
