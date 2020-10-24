package discord

import "errors"

var (
	ErrBadPremiumTier = errors.New("unrecognized premium_tier value")
)

type PremiumTier uint8

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

func (p PremiumTier) IsValid() bool {
	return nil == p.Validate()
}

func (p PremiumTier) Validate() error {
	if p > PremiumTier3 {
		return ErrBadPremiumTier
	}

	return nil
}

