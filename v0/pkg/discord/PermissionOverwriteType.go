package discord

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/ugly"
	"strconv"
)

var (
	ErrBadOverwriteType = errors.New("unrecognized overwrite type value")
)

type OverwriteType uint8

const (
	OverwriteTypeRole OverwriteType = iota
	OverwriteTypeMember
)

func (o *OverwriteType) JSONSize() int {
	if o == nil {
		return 4
	}

	return 1
}

func (o OverwriteType) IsValid() bool {
	return nil == o.Validate()
}

func (o OverwriteType) Validate() error {
	if o > 1 {
		return ErrBadOverwriteType
	}

	return nil
}

func (o *OverwriteType) MarshalJSON() ([]byte, error) {
	return o.ToBytes(), nil
}

func (o *OverwriteType) UnmarshalJSON(in []byte) error {
	tmp, err := strconv.ParseUint(ugly.UnsafeString(in), 10, 8)
	if err != nil {
		return err
	}

	*o = OverwriteType(tmp)

	return nil
}

func (o *OverwriteType) IsNil() bool {
	return o == nil
}

func (o *OverwriteType) ToBytes() []byte {
	if o == nil {
		return js.NullBytesBuf
	}

	switch *o {
	case OverwriteTypeRole:
		return []byte{'0'}

	case OverwriteTypeMember:
		return []byte{'1'}

	default:
		return []byte(strconv.FormatUint(uint64(*o), 10))
	}
}
