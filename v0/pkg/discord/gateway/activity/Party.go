package activity

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/pkg/dlib"
)

type Party interface {
	json.Marshaler
	json.Unmarshaler

	// ID returns the current value of this record's `id` field.
	//
	// The `id` field contains the
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use IDIsSet to check if the field is present before use.
	ID() string

	// IDIsSet returns whether this record's `id` field is currently present.
	IDIsSet() bool

	// SetID overwrites the current value of this record's `id` field.
	SetID(string) Party

	// UnsetID removes this record's `id` field.
	UnsetID() Party

	// CurrentSize returns the current first value of this record's `size` field.
	//
	// The `size` field contains the party's current and maximum size.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use CurrentSizeIsSet to check if the field is present before use.
	CurrentSize() uint16

	// CurrentSizeIsSet returns whether this record's `size` field is currently
	// present.
	CurrentSizeIsSet() bool

	// SetCurrentSize overwrites the current first value of this record's `size`
	// field.
	SetCurrentSize(uint16) Party

	// UnsetCurrentSize removes this record's `size` field.
	//
	// NOTE: This will also unset the MaxSize value.
	UnsetCurrentSize() Party

	// MaxSize returns the current second value of this record's `size` field.
	//
	// The `size` field contains the party's current and maximum size.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MaxSizeIsSet to check if the field is present before use.
	MaxSize() uint16

	// MaxSizeIsSet returns whether this record's `size` field is currently
	// present.
	MaxSizeIsSet() bool

	// SetMaxSize overwrites the current second value of this record's `size`
	// field.
	SetMaxSize(uint16) Party

	// UnsetMaxSize removes this record's `size` field.
	//
	// NOTE: This will also unset the CurrentSize value.
	UnsetMaxSize() Party
}

func NewParty() Party {
	return new(partyImpl)
}

type tmpParty struct {
	ID   *string    `json:"id"`
	Size *[2]uint16 `json:"size"`
}

type partyImpl struct {
	id dlib.StrOptionalField
	cur dlib.U16OptionalField
	max dlib.U16OptionalField
}

func (p *partyImpl) MarshalJSON() ([]byte, error) {
	out := new(tmpParty)

	if p.id.IsSet() {
		tmp := p.id.Value()
		out.ID = &tmp
	}

	if p.cur.IsSet() {
		tmp := [2]uint16{p.cur.Value(), p.max.Value()}
		out.Size = &tmp
	}

	return json.Marshal(out)
}

func (p *partyImpl) UnmarshalJSON(bytes []byte) error {
	tmp := new(tmpParty)

	if err := json.Unmarshal(bytes, tmp); err != nil {
		return err
	}

	if tmp.ID != nil {
		p.id.Set(*tmp.ID)
	}

	if tmp.Size != nil {
		v := *tmp.Size
		p.cur.Set(v[0])
		p.max.Set(v[1])
	}

	return nil
}

func (p *partyImpl) ID() string {
	return p.id.Value()
}

func (p *partyImpl) IDIsSet() bool {
	return p.id.IsSet()
}

func (p *partyImpl) SetID(s string) Party {
	p.id.Set(s)
	return p
}

func (p *partyImpl) UnsetID() Party {
	p.id.Unset()
	return p
}

func (p *partyImpl) CurrentSize() uint16 {
	return p.cur.Value()
}

func (p *partyImpl) CurrentSizeIsSet() bool {
	return p.cur.IsSet()
}

func (p *partyImpl) SetCurrentSize(u uint16) Party {
	p.cur.Set(u)
	return p
}

func (p *partyImpl) UnsetCurrentSize() Party {
	p.cur.Unset()
	p.max.Unset()
	return p
}

func (p *partyImpl) MaxSize() uint16 {
	return p.max.Value()
}

func (p *partyImpl) MaxSizeIsSet() bool {
	return p.max.IsSet()
}

func (p *partyImpl) SetMaxSize(u uint16) Party {
	p.max.Set(u)
	return p
}

func (p *partyImpl) UnsetMaxSize() Party {
	p.cur.Unset()
	p.max.Unset()
	return p
}

