package discord

import (
	"encoding/json"
)

// VoiceRegion
// TODO: Documentation
type VoiceRegion interface {
	json.Marshaler
	json.Unmarshaler

	// ID returns the value of the `id` field currently set on this voice region.
	//
	// The `id` field contains the unique ID for the region.
	ID() string

	// SetID overwrites the current value of this voice region record's `id`
	// field with the given param value.
	SetID(id string) VoiceRegion

	// Name returns the value of the `name` field currently set on this voice
	// region.
	//
	// The `name` field contains the name of the region
	Name() string

	// SetName overwrites the current value of this voice region record's `name`
	// field with the given param value.
	SetName(name string) VoiceRegion

	// VIP returns the value of the `vip` field currently set on this voice
	// region.
	//
	// The `vip` flag indicates if this is a vip-only server.
	VIP() bool

	// SetVIP overwrites the current value of this voice region record's `vip`
	// field with the given param value.
	SetVIP(vip bool) VoiceRegion

	// Optimal returns the value of the `optimal` field currently set on this
	// voice region.
	//
	// The `optimal` flag indicates if this single server is closest to the
	// current user's client.
	Optimal() bool

	// SetOptimal overwrites the current value of this voice region record's
	// `optimal` field with the given param value.
	SetOptimal(opt bool) VoiceRegion

	// Deprecated returns the value of the `deprecated` field currently set on
	// this voice region.
	//
	// The `deprecated` flag indicates whether this is a deprecated voice region
	// (avoid switching to these).
	Deprecated() bool

	// SetDeprecated overwrites the current value of this voice region record's
	// `deprecated` field with the given param value.
	SetDeprecated(dep bool) VoiceRegion

	// Custom returns the value of the `custom` field currently set on this voice
	// region.
	//
	// The `custom` flag indicates whether this is a custom voice region
	// (used for events/etc).
	Custom() bool

	// SetCustom overwrites the current value of this voice region record's
	// `custom` field with the given param value.
	SetCustom(cus bool) VoiceRegion
}
