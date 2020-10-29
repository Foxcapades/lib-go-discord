package guild

import (
	"encoding/json"
	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type RoleSlice []Role

func (r *RoleSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewRole()

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*r = append(*r, tmp)

	return nil
}

func NewRole() Role {
	return new(role)
}

type role struct {
	id          Snowflake
	name        string
	color       lib.Color
	hoist       bool
	position    uint16
	permissions Permission
	managed     bool
	mentionable bool
}

func (r *role) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (r *role) IsNil() bool {
	panic("implement me")
}

func (r *role) UnmarshalJSONObject(decoder *gojay.Decoder, s string) error {
	panic("implement me")
}

func (r *role) NKeys() int {
	panic("implement me")
}

func (r *role) IsValid() bool {
	panic("implement me")
}

func (r *role) Validate() error {
	panic("implement me")
}

func (r *role) MarshalJSON() ([]byte, error) {
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

func (r *role) UnmarshalJSON(bytes []byte) error {
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

func (r *role) ID() Snowflake {
	return r.id
}

func (r *role) SetID(id Snowflake) Role {
	r.id = id

	return r
}

func (r *role) Name() string {
	return r.name
}

func (r *role) SetName(name string) Role {
	r.name = name

	return r
}

func (r *role) Color() lib.Color {
	return r.color
}

func (r *role) SetColor(color lib.Color) Role {
	if color == nil {
		panic(ErrNilColor)
	}

	r.color = color

	return r
}

func (r *role) SetRawColor(val uint32) Role {
	r.color.SetRawValue(val)

	return r
}

func (r *role) Hoist() bool {
	return r.hoist
}

func (r *role) SetHoist(b bool) Role {
	r.hoist = b

	return r
}

func (r *role) Position() uint16 {
	return r.position
}

func (r *role) SetPosition(pos uint16) Role {
	r.position = pos

	return r
}

func (r *role) Permissions() Permission {
	return r.permissions
}

func (r *role) SetPermissions(perm Permission) Role {
	r.permissions = perm

	return r
}

func (r *role) AddPermission(perm Permission) Role {
	r.permissions |= perm

	return r
}

func (r *role) RemovePermission(perm Permission) Role {
	r.permissions &= ^perm

	return r
}

func (r *role) PermissionsContains(perm Permission) bool {
	return r.permissions|perm == perm
}

func (r *role) Managed() bool {
	return r.managed
}

func (r *role) SetManaged(b bool) Role {
	r.managed = b

	return r
}

func (r *role) Mentionable() bool {
	return r.mentionable
}

func (r *role) SetMentionable(b bool) Role {
	r.mentionable = b

	return r
}
