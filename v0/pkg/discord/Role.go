package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
)

var (
	ErrNilColor = errors.New("called SetColor with a nil value")
)

// Role
// TODO: document me
type Role interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	dmeta.Validatable

	// role id
	ID() Snowflake
	SetID(id Snowflake) Role

	// role name
	Name() string
	SetName(name string) Role

	// integer representation of hexadecimal color code
	Color() lib.Color
	SetColor(color lib.Color) Role
	SetRawColor(val uint32) Role

	// if this role is pinned in the user listing
	Hoist() bool
	SetHoist(bool) Role

	// position of this role
	Position() uint16
	SetPosition(pos uint16) Role

	// permission bit set
	Permissions() Permission
	SetPermissions(perm Permission) Role
	AddPermission(perm Permission) Role
	RemovePermission(perm Permission) Role
	PermissionsContains(perm Permission) bool

	// whether this role is managed by an integration
	Managed() bool
	SetManaged(bool) Role

	// whether this role is mentionable
	Mentionable() bool
	SetMentionable(bool) Role
}
