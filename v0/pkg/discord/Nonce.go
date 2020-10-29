package discord

import (
	"encoding/json"
	"errors"
)

var (
	ErrNonceNotInt = errors.New("called IntegerValue on a string nonce")
	ErrNonceNotStr = errors.New("called StringValue on an integer nonce")
)

type Nonce interface {
	json.Marshaler
	json.Unmarshaler

	// IsString returns whether the nonce value is a string.
	IsString() bool

	// IsInteger returns whether the nonce value is an integer.
	IsInteger() bool

	// StringValue returns the value of this nonce if it is a string.
	//
	// If the nonce value is not a string, this method will panic.
	StringValue() string

	// IntegerValue returns the value of this nonce if it is an integer.
	//
	// If the nonce value is not an integer, this method will panic.
	IntegerValue() int64

	// SetString sets the nonce value to the given string and the nonce type to
	// string.
	SetString(val string) Nonce

	// SetInteger sets the nonce value to the given integer and the nonce type to
	// integer.
	SetInteger(val int64) Nonce
}
