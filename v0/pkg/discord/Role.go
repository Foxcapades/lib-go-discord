package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Types                                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

var (
	ErrNilColor = errors.New("called SetColor with a nil value")
)

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Types                                                        ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// Role
// TODO: document me
type Role interface {
	json.Marshaler
	json.Unmarshaler

	// role id
	ID() Snowflake
	SetID(id Snowflake) Role

	// role name
	Name() string
	SetName(name string) Role

	// integer representation of hexadecimal color code
	Color() comm.Color
	SetColor(color comm.Color) Role
	SetRawColor(val uint32) Role

	// if this role is pinned in the user listing
	Hoist() bool
	SetHoist(bool) Role

	// position of this role
	Position() uint16
	SetPosition(pos uint16) Role

	// permission bit set
	Permissions() comm.Permission
	SetPermissions(perm comm.Permission) Role
	AddPermission(perm comm.Permission) Role
	RemovePermission(perm comm.Permission) Role
	PermissionsContains(perm comm.Permission) bool

	// whether this role is managed by an integration
	Managed() bool
	SetManaged(bool) Role

	// whether this role is mentionable
	Mentionable() bool
	SetMentionable(bool) Role
}

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Public Functions                                                    ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

// NewRole
func NewRole(validate bool) Role {
	return &roleImpl{
		validate: validate,
		color:    comm.NewColor(),
	}
}

// ╔════════════════════════════════════════════════════════════════════════╗ //
// ║                                                                        ║ //
// ║    Private Types                                                       ║ //
// ║                                                                        ║ //
// ╚════════════════════════════════════════════════════════════════════════╝ //

type roleImpl struct {
	validate bool

	id          Snowflake
	name        string
	color       comm.Color
	hoist       bool
	position    uint16
	permissions comm.Permission
	managed     bool
	mentionable bool
}

func (r *roleImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[FieldKey]interface{}{
		FieldKeyID:          r.id,
		FieldKeyName:        r.name,
		FieldKeyColor:       r.color,
		FieldKeyHoist:       r.hoist,
		FieldKeyPosition:    r.position,
		FieldKeyPermissions: r.permissions,
		FieldKeyManaged:     r.managed,
		FieldKeyMentionable: r.mentionable,
	})
}

func (r *roleImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]json.RawMessage, 8)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	jsons := map[FieldKey]json.Unmarshaler{
		FieldKeyID:          &r.id,
		FieldKeyColor:       r.color,
		FieldKeyPermissions: &r.permissions,
	}

	for k, v := range jsons {
		if raw, ok := tmp[k]; ok {
			if err := v.UnmarshalJSON(raw); err != nil {
				return err
			}
		}
	}

	raws := map[FieldKey]interface{}{
		FieldKeyName:        &r.name,
		FieldKeyHoist:       &r.hoist,
		FieldKeyPosition:    &r.position,
		FieldKeyManaged:     &r.managed,
		FieldKeyMentionable: &r.mentionable,
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

func (r *roleImpl) ID() Snowflake {
	return r.id
}

func (r *roleImpl) SetID(id Snowflake) Role {
	r.id = id

	return r
}

func (r *roleImpl) Name() string {
	return r.name
}

func (r *roleImpl) SetName(name string) Role {
	r.name = name

	return r
}

func (r *roleImpl) Color() comm.Color {
	return r.color
}

func (r *roleImpl) SetColor(color comm.Color) Role {
	if color == nil {
		panic(ErrNilColor)
	}

	r.color = color

	return r
}

func (r *roleImpl) SetRawColor(val uint32) Role {
	r.color.SetRawValue(val)

	return r
}

func (r *roleImpl) Hoist() bool {
	return r.hoist
}

func (r *roleImpl) SetHoist(b bool) Role {
	r.hoist = b

	return r
}

func (r *roleImpl) Position() uint16 {
	return r.position
}

func (r *roleImpl) SetPosition(pos uint16) Role {
	r.position = pos

	return r
}

func (r *roleImpl) Permissions() comm.Permission {
	return r.permissions
}

func (r *roleImpl) SetPermissions(perm comm.Permission) Role {
	if r.validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions = perm

	return r
}

func (r *roleImpl) AddPermission(perm comm.Permission) Role {
	if r.validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions |= perm

	return r
}

func (r *roleImpl) RemovePermission(perm comm.Permission) Role {
	if r.validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	r.permissions &= ^perm

	return r
}

func (r *roleImpl) PermissionsContains(perm comm.Permission) bool {
	if r.validate {
		if err := perm.Validate(); err != nil {
			panic(err)
		}
	}

	return r.permissions|perm == perm
}

func (r *roleImpl) Managed() bool {
	return r.managed
}

func (r *roleImpl) SetManaged(b bool) Role {
	r.managed = b

	return r
}

func (r *roleImpl) Mentionable() bool {
	return r.mentionable
}

func (r *roleImpl) SetMentionable(b bool) Role {
	r.mentionable = b

	return r
}
