package discord

import (
	"errors"
	"fmt"
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/francoispqt/gojay"
	"io"
)

var (
	ErrBadChannelType = errors.New("unrecognized channel type value")
)

type ChannelType uint8

const (
	// A text channel within a server
	ChannelTypeGuildText ChannelType = iota

	// A direct message between users
	ChannelTypeDM

	// A voice channel within a server
	ChannelTypeGuildVoice

	// A direct message between multiple users
	ChannelTypeGroupDM

	// An organizational category that contains up to 50 channels
	ChannelTypeGuildCategory

	// A channel that users can follow and crosspost into their own server
	ChannelTypeGuildNews

	// A channel in which game developers can sell their game on Discord
	ChannelTypeGuildStore
)

func (t ChannelType) JSONSize() uint32 {
	switch true {
	case t < 10:
		return 1
	case t < 100:
		return 2
	default:
		return 3
	}
}

func (t ChannelType) IsValid() bool {
	return nil == t.Validate()
}

func (t ChannelType) Validate() error {
	if t > ChannelTypeGuildStore {
		return ErrBadChannelType
	}

	return nil
}

func DecodeChannelType(dec *gojay.Decoder) (ChannelType, error) {
	var tmp uint8

	if err := dec.DecodeUint8(&tmp); err != nil {
		return 0, err
	}

	return ChannelType(tmp), nil
}

func (t ChannelType) MarshalJSON() ([]byte, error) {
	return t.ToJSONBytes(), nil
}

func (t *ChannelType) UnmarshalJSON(bytes []byte) error {
	if len(bytes) != 1 {
		return fmt.Errorf("incorrect number of bytes for a ChannelType value: expected 1, got %d", len(bytes))
	}

	if bytes[0] < '0' || bytes[0] > '9' {
		return fmt.Errorf(
			"incorrect byte value: expected a value in [%d..%d], got %d",
			ChannelTypeGuildText,
			ChannelTypeGuildStore,
			int16(bytes[0]) - '0',
		)
	}

	*t = ChannelType(bytes[0] - '0')

	return nil
}

func (t *ChannelType) IsNil() bool {
	return t == nil
}

func (t ChannelType) ToJSONBytes() []byte {
	return bytify.Uint8ToByteSlice(uint8(t))
}

func (t ChannelType) AppendJSONBytes(writer io.Writer) error {
	bytify.Uint8ToBuf(uint8(t), writer)
	return nil
}
