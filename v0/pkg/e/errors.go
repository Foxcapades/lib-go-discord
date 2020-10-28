package e

import "errors"

var (
	// ErrSetNil represents the error case where a nil value was passed to one of
	// a type's Set{Field}() methods.
	//
	// Due to the fact that Discord's API has 3-state fields, a value of nil can
	// ambiguous in it's meaning, and cannot be used on 3-state fields.  For
	// uniformity and due to the fact that it may not be clear which fields have
	// 3 states, passing nil values to field setters is forbidden for all types.
	ErrSetNil = errors.New("attempted to set a nil value through normal a field setter")

	// ErrGetUnset represents the error case where a field is absent from or unset
	// on a type and that field's Get{Field}() method was called.
	//
	// Use {Field}IsSet() methods to determine if a field is set and readable.
	ErrGetUnset = errors.New("attempted to get the value of an absent field from a field getter")

	// ErrGetNull represents the error case where a field is set to null and that
	// field's Get{Field}() method was called.
	//
	// Use {Field}IsNull() methods to determine if a field is non-null and
	// readable.
	ErrGetNull = errors.New("attempted to get the value of a null field from a field getter")
)
