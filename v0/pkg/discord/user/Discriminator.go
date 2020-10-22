package user

import "errors"

var (
	ErrDiscriminatorLength = errors.New("discriminator values must be " +
		"exactly 4 digits in length")
)

type Discriminator uint16

func (d Discriminator) IsValid() bool {
	return nil == d.Validate()
}

func (d Discriminator) Validate() error {
	if d < 1000 || d > 9999 {
		return ErrDiscriminatorLength
	}

	return nil
}

