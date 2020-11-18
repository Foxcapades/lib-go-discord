package discord

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
	"github.com/francoispqt/gojay"
)

// PermissionOverwrite
// TODO: document me
type PermissionOverwrite interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	dmeta.Sized
	dmeta.Validatable

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field
	ID() Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) PermissionOverwrite

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field
	Type() OverwriteType

	// SetType overwrites the current value of this record's `type` field.
	SetType(OverwriteType) PermissionOverwrite

	// Allow returns the current value of this record's `allow` field.
	//
	// The `allow` field
	Allow() Permission

	// SetAllow overwrites the current value of this record's `allow` field.
	SetAllow(Permission) PermissionOverwrite

	// Deny returns the current value of this record's `deny` field.
	//
	// The `deny` field
	Deny() Permission

	// SetDeny overwrites the current value of this record's `deny` field.
	SetDeny(Permission) PermissionOverwrite
}
