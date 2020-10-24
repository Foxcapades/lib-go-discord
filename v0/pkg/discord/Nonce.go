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

// NewStringNonce creates a new Nonce value containing the given string.
func NewStringNonce(val string) Nonce {
	return &nonceImpl{val, nonceTypeStr}
}

// NewIntNonce creates a new Nonce value containing the given integer.
func NewIntNonce(val int64) Nonce {
	return &nonceImpl{val, nonceTypeInt}
}

// NewNonce creates a new empty Nonce value.
//
// Note: Nonce values created by this method are empty and will return false on
// calls to both IsString() and IsInteger().
func NewNonce() Nonce {
	return new(nonceImpl)
}

const (
	nonceTypeInt uint8 = 1 + iota
	nonceTypeStr
)

type nonceImpl struct {
	value interface{}
	nType uint8
}

func (n *nonceImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}

func (n *nonceImpl) UnmarshalJSON(bytes []byte) error {
	if bytes[0] == '"' {
		var tmp string

		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}

		n.value = tmp
		n.nType = nonceTypeStr
	} else {
		var tmp int64

		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}

		n.value = tmp
		n.nType = nonceTypeInt
	}

	return nil
}

func (n *nonceImpl) IsString() bool {
	return n.nType == nonceTypeStr
}

func (n *nonceImpl) IsInteger() bool {
	return n.nType == nonceTypeInt
}

func (n *nonceImpl) StringValue() string {
	if n.nType != nonceTypeStr {
		panic(ErrNonceNotStr)
	}

	return n.value.(string)
}

func (n *nonceImpl) IntegerValue() int64 {
	if n.nType != nonceTypeInt {
		panic(ErrNonceNotInt)
	}

	return n.value.(int64)
}

func (n *nonceImpl) SetString(val string) Nonce {
	n.value = val
	n.nType = nonceTypeStr

	return n
}

func (n *nonceImpl) SetInteger(val int64) Nonce {
	n.value = val
	n.nType = nonceTypeInt

	return n
}



