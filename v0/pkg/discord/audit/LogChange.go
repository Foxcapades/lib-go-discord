package audit

type LogChange interface {
	// NewValue returns the current value of this record's `new_value` field.
	//
	// The `new_value` field contains the new value of the key.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NewValueIsSet to check if the field is present before use.
	NewValue() interface{}

	// NewValueIsSet returns whether this record's `new_value` field is currently
	// present.
	NewValueIsSet() bool

	// SetNewValue overwrites the current value of this record's `new_value`
	// field.
	SetNewValue(interface{}) LogChange

	// UnsetNewValue removes this record's `new_value` field.
	UnsetNewValue() LogChange

	// OldValue returns the current value of this record's `old_value` field.
	//
	// The `old_value` field contains the old value of the key.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use OldValueIsSet to check if the field is present before use.
	OldValue() interface{}

	// OldValueIsSet returns whether this record's `old_value` field is currently
	// present.
	OldValueIsSet() bool

	// SetOldValue overwrites the current value of this record's `old_value`
	// field.
	SetOldValue(interface{}) LogChange

	// UnsetOldValue removes this record's `old_value` field.
	UnsetOldValue() LogChange

	// Key returns the current value of this record's `key` field.
	//
	// The `key` field contains the
	Key() ChangeKey

	// SetKey overwrites the current value of this record's `key` field.
	SetKey(ChangeKey) LogChange
}
