package dio

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/permission"
)

type ChannelPermissionPatch interface {
	// Allow returns the current value of this request's `allow` field.
	//
	// The `allow` field contains the bitwise value of all allowed permissions.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AllowIsSet to check if the field is present before use.
	Allow() comm.Permission

	// AllowIsSet returns whether this request's `allow` field is currently present.
	AllowIsSet() bool

	// SetAllow overwrites the current value of this request's `allow` field.
	SetAllow(comm.Permission) ChannelPermissionPatch

	// UnsetAllow removes this request's `allow` field.
	UnsetAllow() ChannelPermissionPatch

	// Deny returns the current value of this request's `deny` field.
	//
	// The `deny` field contains the bitwise value of all disallowed permissions.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use DenyIsSet to check if the field is present before use.
	Deny() comm.Permission

	// DenyIsSet returns whether this request's `deny` field is currently present.
	DenyIsSet() bool

	// SetDeny overwrites the current value of this request's `deny` field.
	SetDeny(comm.Permission) ChannelPermissionPatch

	// UnsetDeny removes this request's `deny` field.
	UnsetDeny() ChannelPermissionPatch

	// Type returns the current value of this request's `type` field.
	//
	// The `type` field contains the permission type being changed, either role or
	// user.
	Type() permission.Type

	// SetType overwrites the current value of this request's `type` field.
	SetType(permission.Type) ChannelPermissionPatch
}
