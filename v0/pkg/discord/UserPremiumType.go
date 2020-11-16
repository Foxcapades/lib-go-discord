package discord

import "errors"

var (
	ErrBadPremiumType = errors.New("unrecognized premium type value")
)

type UserPremiumType uint8

const (
	PremiumTypeNone UserPremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

func (p *UserPremiumType) JSONSize() int {
	if p == nil {
		return 4
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
