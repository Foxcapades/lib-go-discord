package permission

import (
	"github.com/foxcapades/lib-go-discord/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
)

// Overwrite
// TODO: document me
type Overwrite interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field
	ID() dlib.Snowflake

	// SetID overwrites the current value of this record's `id` field.
	SetID(dlib.Snowflake) Overwrite

	// Type returns the current value of this record's `type` field.
	//
	// The `type` field
	Type() OverwriteType

	// SetType overwrites the current value of this record's `type` field.
	SetType(OverwriteType) Overwrite

	// Allow returns the current value of this record's `allow` field.
	//
	// The `allow` field
	Allow() comm.Permission

	// SetAllow overwrites the current value of this record's `allow` field.
	SetAllow(comm.Permission) Overwrite

	// Deny returns the current value of this record's `deny` field.
	//
	// The `deny` field
	Deny() comm.Permission

	// SetDeny overwrites the current value of this record's `deny` field.
	SetDeny(comm.Permission) Overwrite
}
