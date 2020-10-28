package activity

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

// NewParty returns a new instance of the ActivityParty type.
func NewParty() ActivityParty {
	return new(party)
}

// party implements discord.ActivityParty.
type party struct {
	// This could be any string
	id   *string
	size *sizeArray
}

func (p *party) MarshalJSONObject(enc *gojay.Encoder) {
	if p.IDIsSet() {
		enc.AddStringKey(string(serial.KeyID), *p.id)
	}

	if p.SizesAreSet() {
		enc.AddArrayKey(string(serial.KeySize), p.size)
	}
}

func (p *party) IsNil() bool {
	return p == nil
}

func (p *party) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch serial.Key(k) {
	case serial.KeyID:
		return dec.DecodeString(p.id)
	case serial.KeySize:

	}

	panic("implement me")
}

func (p *party) NKeys() int {
	panic("implement me")
}

// IsValid implements lib.Validatable.IsValid.
func (p *party) IsValid() bool {
	return nil == p.Validate()
}

// Validate implements lib.Validatable.Validate.
func (p *party) Validate() error {
	if p.size[0] > p.size[1] {
		return lib.NewValidationErrorSet().
			AppendKeyedError(serial.KeySize, "size[0] must be less than or equal to size[1]")
	}

	return nil
}

// MarshalJSON implements json.Marshaler.MarshalJSON.
func (p *party) MarshalJSON() ([]byte, error) {
	out := new(tmpParty)

	if p.IDIsSet() {
		out.ID = p.id
	}

	if p.SizesAreSet() {
		out.Size = p.size
	}

	return json.Marshal(out)
}

// UnmarshalJSON implements json.Unmarshaler.UnmarshalJSON.
func (p *party) UnmarshalJSON(bytes []byte) error {
	tmp := new(tmpParty)

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	p.id = tmp.ID
	p.size = tmp.Size

	return nil
}

// ID implements discord.ActivityParty.ID.
func (p *party) ID() string {
	if p.id == nil {
		panic(types.ErrUnsetField)
	}

	return *p.id
}

// IDIsSet implements discord.ActivityParty.IDIsSet.
func (p *party) IDIsSet() bool {
	return p.id != nil
}

// SetID implements discord.ActivityParty.SetID.
func (p *party) SetID(s string) ActivityParty {
	p.id = &s
	return p
}

// UnsetID implements discord.ActivityParty.UnsetID.
func (p *party) UnsetID() ActivityParty {
	p.id = nil
	return p
}

// CurrentSize implements discord.ActivityParty.CurrentSize.
func (p *party) CurrentSize() uint16 {
	if p.size == nil {
		panic(types.ErrUnsetField)
	}

	return p.size[0]
}

// MaxSize implements discord.ActivityParty.MaxSize.
func (p *party) MaxSize() uint16 {
	if p.size == nil {
		panic(types.ErrUnsetField)
	}

	return p.size[1]
}

// SizesAreSet implements discord.ActivityParty.SizesAreSet.
func (p *party) SizesAreSet() bool {
	panic("implement me")
}

// SetSizes implements discord.ActivityParty.SetSizes.
func (p *party) SetSizes(current, max uint16) ActivityParty {
	p.size = new(sizeArray)
	p.size[0] = current
	p.size[1] = max

	return p
}

// UnsetSizes implements discord.ActivityParty.UnsetSizes.
func (p *party) UnsetSizes() ActivityParty {
	p.size = nil

	return p
}

type tmpParty struct {
	ID   *string    `json:"id,omitempty"`
	Size *sizeArray `json:"size,omitempty"`
}

type sizeArray [2]uint16

func (s *sizeArray) MarshalJSONArray(enc *gojay.Encoder) {
	enc.AddUint16(s[0])
	enc.AddUint16(s[1])
}

func (s *sizeArray) IsNil() bool {
	return s == nil
}

func (s *sizeArray) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp uint16

	if err := dec.DecodeUint16(&tmp); err != nil {
		return err
	}

	s[dec.Index()] = tmp

	return nil
}
