package user

import "errors"

var (
	ErrBadPremiumType = errors.New("unrecognized premium type value")
)

type PremiumType uint8

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

func (p PremiumType) IsValid() bool {
	return nil == p.Validate()
}

func (p PremiumType) Validate() error {
	if p > PremiumTypeNitro {
		return ErrBadPremiumType
	}

	return nil
}
