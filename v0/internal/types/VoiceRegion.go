package types

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

const (
	vrVIPMask uint8 = 1 << iota
	vrOptimalMask
	vrDeprecatedMask
	vrCustomMask
)

type VoiceRegionImpl struct {
	id    string
	name  string
	flags uint8
}

func (v *VoiceRegionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[serial.Key]interface{}{
		serial.KeyID:         v.id,
		serial.KeyName:       v.name,
		serial.KeyVIP:        v.VIP(),
		serial.KeyOptimal:    v.Optimal(),
		serial.KeyDeprecated: v.Deprecated(),
		serial.KeyCustom:     v.Custom(),
	})
}

func (v *VoiceRegionImpl) UnmarshalJSON(bytes []byte) (err error) {
	tmp := make(map[serial.Key]json.RawMessage, 6)

	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	for k, raw := range tmp {
		switch k {
		// Handle string fields
		case serial.KeyID:
			err = json.Unmarshal(raw, &v.id)
		case serial.KeyName:
			err = json.Unmarshal(raw, &v.name)

		// Handle boolean flags
		default:
			var tb bool

			if err = json.Unmarshal(raw, &tb); err == nil {
				switch k {
				case serial.KeyVIP:
					v.SetVIP(tb)
				case serial.KeyOptimal:
					v.SetOptimal(tb)
				case serial.KeyDeprecated:
					v.SetDeprecated(tb)
				case serial.KeyCustom:
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

func (v *VoiceRegionImpl) ID() string {
	return v.id
}

func (v *VoiceRegionImpl) SetID(id string) VoiceRegion {
	v.id = id

	return v
}

func (v *VoiceRegionImpl) Name() string {
	return v.name
}

func (v *VoiceRegionImpl) SetName(name string) VoiceRegion {
	v.name = name

	return v
}

func (v *VoiceRegionImpl) VIP() bool {
	return v.flags&vrVIPMask == vrVIPMask
}

func (v *VoiceRegionImpl) SetVIP(vip bool) VoiceRegion {
	if vip {
		v.flags |= vrVIPMask
	} else {
		v.flags &= ^vrVIPMask
	}

	return v
}

func (v *VoiceRegionImpl) Optimal() bool {
	return v.flags&vrOptimalMask == vrOptimalMask
}

func (v *VoiceRegionImpl) SetOptimal(opt bool) VoiceRegion {
	if opt {
		v.flags |= vrOptimalMask
	} else {
		v.flags &= ^vrOptimalMask
	}

	return v
}

func (v *VoiceRegionImpl) Deprecated() bool {
	return v.flags&vrDeprecatedMask == vrDeprecatedMask
}

func (v *VoiceRegionImpl) SetDeprecated(dep bool) VoiceRegion {
	if dep {
		v.flags |= vrDeprecatedMask
	} else {
		v.flags &= ^vrDeprecatedMask
	}

	return v
}

func (v *VoiceRegionImpl) Custom() bool {
	return v.flags|vrCustomMask == vrCustomMask
}

func (v *VoiceRegionImpl) SetCustom(cus bool) VoiceRegion {
	if cus {
		v.flags |= vrCustomMask
	} else {
		v.flags &= ^vrCustomMask
	}

	return v
}
