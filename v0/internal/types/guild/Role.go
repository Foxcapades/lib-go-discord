package guild

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type RoleImpl struct {
	Validate bool

	id          Snowflake
	name        string
	color       lib.Color
	hoist       bool
	position    uint16
	permissions Permission
	managed     bool
	mentionable bool
}

func (r *RoleImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[serial.Key]interface{}{
		serial.KeyID:          r.id,
		serial.KeyName:        r.name,
		serial.KeyColor:       r.color,
		serial.KeyHoist:       r.hoist,
		serial.KeyPosition:    r.position,
		serial.KeyPermissions: r.permissions,
		serial.KeyManaged:     r.managed,
		serial.KeyMentionable: r.mentionable,
	})
}

func (r *RoleImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 8)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	jsons := map[serial.Key]json.Unmarshaler{
		serial.KeyID:          r.id,
		serial.KeyColor:       r.color,
		serial.KeyPermissions: &r.permissions,
	}

	for k, v := range jsons {
		if raw, ok := tmp[k]; ok {
			if err := v.UnmarshalJSON(raw); err != nil {
				return err
			}
		}
	}

	raws := map[serial.Key]interface{}{
		serial.KeyName:        &r.name,
		serial.KeyHoist:       &r.hoist,
		serial.KeyPosition:    &r.position,
		serial.KeyManaged:     &r.managed,
		serial.KeyMentionable: &r.mentionable,
	}

	for k, v := range raws {
		if raw, ok := tmp[k]; ok {
			if err := json.Unmarshal(raw, v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *RoleImpl) ID() Snowflake {
	return r.id
}

func (r *RoleImpl) SetID(id Snowflake) Role {
	r.id = id

	return r
}

func (r *RoleImpl) Name() string {
	return r.name
}

func (r *RoleImpl) SetName(name string) Role {
	r.name = name

	return r
}

func (r *RoleImpl) Color() lib.Color {
	return r.color
}

func (r *RoleImpl) SetColor(color lib.Color) Role {
	if color == nil {
		panic(ErrNilColor)
	}

	r.color = color

	return r
}

func (r *RoleImpl) SetRawColor(val uint32) Role {
	r.color.SetRawValue(val)

	return r
}

func (r *RoleImpl) Hoist() bool {
	return r.hoist
}

func (r *RoleImpl) SetHoist(b bool) Role {
	r.hoist = b

	return r
}

func (r *RoleImpl) Position() uint16 {
	return r.position
}

func (r *RoleImpl) SetPosition(pos uint16) Role {
	r.position = pos

	return r
}

func (r *RoleImpl) Permissions() Permission {
	return r.permissions
}

func (r *RoleImpl) SetPermissions(perm Permission) Role {
	if r.Validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions = perm

	return r
}

func (r *RoleImpl) AddPermission(perm Permission) Role {
	if r.Validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions |= perm

	return r
}

func (r *RoleImpl) RemovePermission(perm Permission) Role {
	if r.Validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions &= ^perm

	return r
}

func (r *RoleImpl) PermissionsContains(perm Permission) bool {
	if r.Validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	return r.permissions|perm == perm
}

func (r *RoleImpl) Managed() bool {
	return r.managed
}

func (r *RoleImpl) SetManaged(b bool) Role {
	r.managed = b

	return r
}

func (r *RoleImpl) Mentionable() bool {
	return r.mentionable
}

func (r *RoleImpl) SetMentionable(b bool) Role {
	r.mentionable = b

	return r
}
