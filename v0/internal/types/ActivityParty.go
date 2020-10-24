package types

import (
	"encoding/json"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type tmpParty struct {
	ID   *string    `json:"id"`
	Size *[2]uint16 `json:"size"`
}

type ActivityPartyImpl struct {
	id  OptionalString
	cur OptionalUint16
	max OptionalUint16
}

func (p *ActivityPartyImpl) MarshalJSON() ([]byte, error) {
	out := new(tmpParty)

	if p.id.IsSet() {
		tmp := p.id.Get()
		out.ID = &tmp
	}

	if p.cur.IsSet() {
		tmp := [2]uint16{p.cur.Get(), p.max.Get()}
		out.Size = &tmp
	}

	return json.Marshal(out)
}

func (p *ActivityPartyImpl) UnmarshalJSON(bytes []byte) error {
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

func (p *ActivityPartyImpl) ID() string {
	return p.id.Get()
}

func (p *ActivityPartyImpl) IDIsSet() bool {
	return p.id.IsSet()
}

func (p *ActivityPartyImpl) SetID(s string) ActivityParty {
	p.id.Set(s)
	return p
}

func (p *ActivityPartyImpl) UnsetID() ActivityParty {
	p.id.Unset()
	return p
}

func (p *ActivityPartyImpl) CurrentSize() uint16 {
	return p.cur.Get()
}

func (p *ActivityPartyImpl) CurrentSizeIsSet() bool {
	return p.cur.IsSet()
}

func (p *ActivityPartyImpl) SetCurrentSize(u uint16) ActivityParty {
	p.cur.Set(u)
	return p
}

func (p *ActivityPartyImpl) UnsetCurrentSize() ActivityParty {
	p.cur.Unset()
	p.max.Unset()
	return p
}

func (p *ActivityPartyImpl) MaxSize() uint16 {
	return p.max.Get()
}

func (p *ActivityPartyImpl) MaxSizeIsSet() bool {
	return p.max.IsSet()
}

func (p *ActivityPartyImpl) SetMaxSize(u uint16) ActivityParty {
	p.max.Set(u)
	return p
}

func (p *ActivityPartyImpl) UnsetMaxSize() ActivityParty {
	p.cur.Unset()
	p.max.Unset()
	return p
}
