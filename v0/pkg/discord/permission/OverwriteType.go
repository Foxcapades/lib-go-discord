package permission

import "errors"

var (
	ErrBadOverwriteType = errors.New("unrecognized overwrite type value")
)

type OverwriteType uint8

const (
	OverwriteTypeRole OverwriteType = iota
	OverwriteTypeMember
)

func (o OverwriteType) IsValid() bool {
	return nil == o.Validate()
}

func (o OverwriteType) Validate() error {
	if o > 1 {
		return ErrBadOverwriteType
	}

	return nil
}

