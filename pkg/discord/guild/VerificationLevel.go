package guild

import "errors"

var (
	ErrBadVerificationLevel = errors.New("unrecognized verification level value")
)

type VerificationLevel uint8

const (
	// unrestricted
	VerificationLevelNone VerificationLevel = iota

	// must have verified email on account
	VerificationLevelLow

	// must be registered on Discord for longer than 5 minutes
	VerificationLevelMedium

	// must be a member of the server for longer than 10 minutes
	VerificationLevelHigh

	// must have a verified phone number
	VerificationLevelVeryHigh
)

func (v VerificationLevel) IsValid() bool {
	return nil == v.Validate()
}

func (v VerificationLevel) Validate() error {
	if v > VerificationLevelVeryHigh {
		return ErrBadVerificationLevel
	}

	return nil
}
