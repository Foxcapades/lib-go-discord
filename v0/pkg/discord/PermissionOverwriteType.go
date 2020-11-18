package discord

import (
	"errors"
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/ugly"
	"io"
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

func (o *OverwriteType) JSONSize() uint32 {
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
	return o.ToJSONBytes(), nil
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

func (o *OverwriteType) ToJSONBytes() []byte {
	if o == nil {
		return js.NullBytesBuf
	}

	switch *o {
	case OverwriteTypeRole:
		return []byte{'0'}

	case OverwriteTypeMember:
		return []byte{'1'}

	default:
		return bytify.Uint8ToByteSlice(uint8(*o))
	}
}

func (o *OverwriteType) AppendJSONBytes(writer io.Writer) (err error) {
	_, err = writer.Write(o.ToJSONBytes())
	return
}
