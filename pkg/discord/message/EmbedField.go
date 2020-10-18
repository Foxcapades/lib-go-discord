package message

type EmbedField interface {
	// Name returns the current value of this record's `name` field.
	//
	// The `name` field contains the name of the field.
	Name() string

	// SetName overwrites the current value of this record's `name` field.
	SetName(string) EmbedField

	// Value returns the current value of this record's `value` field.
	//
	// The `value` field contains the value of the field.
	Value() string

	// SetValue overwrites the current value of this record's `value` field.
	SetValue(string) EmbedField

	// Inline returns the current value of this record's `inline` field.
	//
	// The `inline` field indicates whether or not this field should display
	// inline.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use InlineIsSet to check if the field is present before use.
	Inline() bool

	// InlineIsSet returns whether this record's `inline` field is currently
	// present.
	InlineIsSet() bool

	// SetInline overwrites the current value of this record's `inline` field.
	SetInline(bool) EmbedField

	// UnsetInline removes this record's `inline` field.
	UnsetInline() EmbedField
}
