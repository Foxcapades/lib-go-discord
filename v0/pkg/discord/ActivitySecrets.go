package discord

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
)

type ActivitySecrets interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	// Join returns the current value of this record's `join` field.
	//
	// The `join` field contains the secret for joining a party.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use JoinIsSet to check if the field is present before use.
	Join() string

	// JoinIsSet returns whether this record's `join` field is currently present.
	JoinIsSet() bool

	// SetJoin overwrites the current value of this record's `join` field.
	SetJoin(string) ActivitySecrets

	// UnsetJoin removes this record's `join` field.
	UnsetJoin() ActivitySecrets

	// Spectate returns the current value of this record's `spectate` field.
	//
	// The `spectate` field contains the secret for spectating a game.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SpectateIsSet to check if the field is present before use.
	Spectate() string

	// SpectateIsSet returns whether this record's `spectate` field is currently present.
	SpectateIsSet() bool

	// SetSpectate overwrites the current value of this record's `spectate` field.
	SetSpectate(string) ActivitySecrets

	// UnsetSpectate removes this record's `spectate` field.
	UnsetSpectate() ActivitySecrets

	// Match returns the current value of this record's `match` field.
	//
	// The `match` field contains the secret for a specific instanced match.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MatchIsSet to check if the field is present before use.
	Match() string

	// MatchIsSet returns whether this record's `match` field is currently present.
	MatchIsSet() bool

	// SetMatch overwrites the current value of this record's `match` field.
	SetMatch(string) ActivitySecrets

	// UnsetMatch removes this record's `match` field.
	UnsetMatch() ActivitySecrets
}
