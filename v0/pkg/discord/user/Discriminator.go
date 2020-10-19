package user

import "errors"

var (
	ErrDiscriminatorLength = errors.New("discriminator values must be " +
		"exactly 4 characters in length")
	ErrDiscriminatorChar = errors.New("discriminator values must consist " +
		"of numeric characters only")
)

type Discriminator string

func (d Discriminator) IsValid() bool {
	return nil == d.Validate()
}

func (d Discriminator) Validate() error {
	if len(d) != 4 {
		return ErrDiscriminatorLength
	}

	for i := range d {
		if i < '0' || i > '9' {
			return ErrDiscriminatorChar
		}
	}

	return nil
}

