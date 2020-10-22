package discord

import (
	"encoding/json"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Types                                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

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

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Functions                                                    ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// NewVoiceRegion
// TODO: Document me
func NewVoiceRegion() VoiceRegion {
	return new(voiceRegionImpl)
}

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Internals                                                           ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

const (
	vrVIPMask uint8 = 1 << iota
	vrOptimalMask
	vrDeprecatedMask
	vrCustomMask
)

type voiceRegionImpl struct {
	id    string
	name  string
	flags uint8
}

func (v *voiceRegionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[FieldKey]interface{}{
		FieldKeyID:         v.id,
		FieldKeyName:       v.name,
		FieldKeyVIP:        v.VIP(),
		FieldKeyOptimal:    v.Optimal(),
		FieldKeyDeprecated: v.Deprecated(),
		FieldKeyCustom:     v.Custom(),
	})
}

func (v *voiceRegionImpl) UnmarshalJSON(bytes []byte) (err error) {
	tmp := make(map[FieldKey]json.RawMessage, 6)

	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	for k, raw := range tmp {
		switch k {
		// Handle string fields
		case FieldKeyID:
			err = json.Unmarshal(raw, &v.id)
		case FieldKeyName:
			err = json.Unmarshal(raw, &v.name)

		// Handle boolean flags
		default:
			var tb bool

			if err = json.Unmarshal(raw, &tb); err == nil {
				switch k {
				case FieldKeyVIP:
					v.SetVIP(tb)
				case FieldKeyOptimal:
					v.SetOptimal(tb)
				case FieldKeyDeprecated:
					v.SetDeprecated(tb)
				case FieldKeyCustom:
					v.SetCustom(tb)
				}
			}
		}

		if err != nil {
			return
		}
	}

	return
}

func (v *voiceRegionImpl) ID() string {
	return v.id
}

func (v *voiceRegionImpl) SetID(id string) VoiceRegion {
	v.id = id

	return v
}

func (v *voiceRegionImpl) Name() string {
	return v.name
}

func (v *voiceRegionImpl) SetName(name string) VoiceRegion {
	v.name = name

	return v
}

func (v *voiceRegionImpl) VIP() bool {
	return v.flags&vrVIPMask == vrVIPMask
}

func (v *voiceRegionImpl) SetVIP(vip bool) VoiceRegion {
	if vip {
		v.flags |= vrVIPMask
	} else {
		v.flags &= ^vrVIPMask
	}

	return v
}

func (v *voiceRegionImpl) Optimal() bool {
	return v.flags&vrOptimalMask == vrOptimalMask
}

func (v *voiceRegionImpl) SetOptimal(opt bool) VoiceRegion {
	if opt {
		v.flags |= vrOptimalMask
	} else {
		v.flags &= ^vrOptimalMask
	}

	return v
}

func (v *voiceRegionImpl) Deprecated() bool {
	return v.flags&vrDeprecatedMask == vrDeprecatedMask
}

func (v *voiceRegionImpl) SetDeprecated(dep bool) VoiceRegion {
	if dep {
		v.flags |= vrDeprecatedMask
	} else {
		v.flags &= ^vrDeprecatedMask
	}

	return v
}

func (v *voiceRegionImpl) Custom() bool {
	return v.flags|vrCustomMask == vrCustomMask
}

func (v *voiceRegionImpl) SetCustom(cus bool) VoiceRegion {
	if cus {
		v.flags |= vrCustomMask
	} else {
		v.flags &= ^vrCustomMask
	}

	return v
}
