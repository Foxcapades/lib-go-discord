package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"
)

// MessageActivity
// TODO: Document me
type MessageActivity interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	lib.Validatable

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field contains the type of message activity.
	Type() MessageActivityType

	// SetType overwrites the current value of this record's `type` field.
	SetType(MessageActivityType) MessageActivity

	// PartyID returns the current value of this record's `party_id` field.
	//
	// The `party_id` field contains the party_id from a Rich Presence event.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use PartyIDIsSet to check if the field is present before use.
	PartyID() string

	// PartyIDIsSet returns whether this record's `party_id` field is currently present.
	PartyIDIsSet() bool

	// SetPartyID overwrites the current value of this record's `party_id` field.
	SetPartyID(string) MessageActivity

	// UnsetPartyID removes this record's `party_id` field.
	UnsetPartyID() MessageActivity
}
