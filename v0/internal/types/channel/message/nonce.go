package message

import (
	"encoding/json"
	"github.com/francoispqt/gojay"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func DecodeNonce(dec *gojay.Decoder) (Nonce, error) {
	tmp := make(gojay.EmbeddedJSON, 0, 32)
	if err := dec.EmbeddedJSON(&tmp); err != nil {
		return nil, err
	}

	out := NewNonce()
	if err := out.UnmarshalJSON(tmp); err != nil {
		return nil, err
	}

	return out, nil
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
