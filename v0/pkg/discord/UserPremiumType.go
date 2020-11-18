package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"io"
)

var (
	ErrBadPremiumType = errors.New("unrecognized premium type value")
)

type UserPremiumType uint8

const (
	PremiumTypeNone UserPremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

func (p *UserPremiumType) JSONSize() uint32 {
	if p == nil {
		return js.NullSize
	}

	return 1
}

func (p *UserPremiumType) IsNil() bool {
	return p == nil
}

func (p UserPremiumType) IsValid() bool {
	return nil == p.Validate()
}

func (p UserPremiumType) Validate() error {
	if p > PremiumTypeNitro {
		return ErrBadPremiumType
	}

	return nil
}

func (p UserPremiumType) MarshalJSON() ([]byte, error) {
	return bytify.Uint8ToByteSlice(uint8(p)), nil
}

func (p *UserPremiumType) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, (*uint8)(p))
}

func (p UserPremiumType) ToJSONBytes() []byte {
	return bytify.Uint8ToByteSlice(uint8(p))
}

func (p *UserPremiumType) AppendJSONBytes(writer io.Writer) (err error) {
	_, err = writer.Write(p.ToJSONBytes())
	return
}
