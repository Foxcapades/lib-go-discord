package discord

type Emoji interface {
	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the emoji id.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use IDIsNull to check if the field is null before use.
	ID() Snowflake

	// IDIsNull returns whether this record's `id` field is currently null.
	IDIsNull() bool

	// SetID overwrites the current value of this record's `id` field.
	SetID(Snowflake) Emoji

	// SetNullID overwrites the current value of this record's `id` field
	// with `null`.
	SetNullID() Emoji

	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the emoji name.
	//
	// If this method is called on a field with a null value, this method will
	// panic.  Use NameIsNull to check if the field is null before use.
	Name() string

	// NameIsNull returns whether this record's `name` field is currently null.
	NameIsNull() bool

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) Emoji

	// SetNullName overwrites the current value of this record's `name` field
	// with `null`.
	SetNullName() Emoji

	// Roles returns the current value of this record's `roles` field.
	//
	// The `roles` field contains the roles this emoji is whitelisted to.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RolesIsSet to check if the field is present before use.
	Roles() []Role

	// RolesIsSet returns whether this record's `roles` field is currently
	// present.
	RolesIsSet() bool

	// SetRoles overwrites the current value of this record's `roles` field.
	SetRoles([]Role) Emoji

	// UnsetRoles removes this record's `roles` field.
	UnsetRoles() Emoji

	// User returns the current value of this record's `user` field.
	//
	// The `user` field contains the user that created this emoji.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use UserIsSet to check if the field is present before use.
	User() User

	// UserIsSet returns whether this record's `user` field is currently present.
	UserIsSet() bool

	// SetUser overwrites the current value of this record's `user` field.
	SetUser(User) Emoji

	// UnsetUser removes this record's `user` field.
	UnsetUser() Emoji

	// RequireColons returns the current value of this record's `require_colons`
	// field.
	//
	// The `require_colons` field indicates whether this emoji must be wrapped in
	// colons.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use RequireColonsIsSet to check if the field is present before use.
	RequireColons() bool

	// RequireColonsIsSet returns whether this record's `require_colons` field is
	// currently present.
	RequireColonsIsSet() bool

	// SetRequireColons overwrites the current value of this record's
	// `require_colons` field.
	SetRequireColons(bool) Emoji

	// UnsetRequireColons removes this record's `require_colons` field.
	UnsetRequireColons() Emoji

	// Managed returns the current value of this record's `managed` field.
	//
	// The `managed` field indicates whether this emoji is managed.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ManagedIsSet to check if the field is present before use.
	Managed() bool

	// ManagedIsSet returns whether this record's `managed` field is currently present.
	ManagedIsSet() bool

	// SetManaged overwrites the current value of this record's `managed` field.
	SetManaged(bool) Emoji

	// UnsetManaged removes this record's `managed` field.
	UnsetManaged() Emoji

	// Animated returns the current value of this record's `animated` field.
	//
	// The `animated` field indicates whether this emoji is animated.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AnimatedIsSet to check if the field is present before use.
	Animated() bool

	// AnimatedIsSet returns whether this record's `animated` field is currently
	// present.
	AnimatedIsSet() bool

	// SetAnimated overwrites the current value of this record's `animated` field.
	SetAnimated(bool) Emoji

	// UnsetAnimated removes this record's `animated` field.
	UnsetAnimated() Emoji

	// Available returns the current value of this record's `available` field.
	//
	// The `available` field indicates whether this emoji can be used, may be
	// false due to loss of Server Boosts.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AvailableIsSet to check if the field is present before use.
	Available() bool

	// AvailableIsSet returns whether this record's `available` field is currently
	// present.
	AvailableIsSet() bool

	// SetAvailable overwrites the current value of this record's `available`
	// field.
	SetAvailable(bool) Emoji

	// UnsetAvailable removes this record's `available` field.
	UnsetAvailable() Emoji
}
