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

func (p UserPremiumType) IsValid() bool {
	return nil == p.Validate()
}

func (p UserPremiumType) Validate() error {
	if p > PremiumTypeNitro {
		return ErrBadPremiumType
	}

	return nil
}