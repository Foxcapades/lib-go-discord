package discord

import (
	"errors"
	"github.com/francoispqt/gojay"
)

var (
	ErrBadBitrate = errors.New("bitrates must be between 8000 to 96000 (128000 for VIP servers) inclusive")
)

type Bitrate uint32

func (b Bitrate) IsValid() bool {
	return nil == b.Validate()
}

func (b Bitrate) Validate() error {
	if b < 8000 || b > 128000 {
		return ErrBadBitrate
	}

	return nil
}

func DecodeBitrate(dec *gojay.Decoder) (*Bitrate, error) {
	var tmp *uint32

	if err := dec.Uint32Null(&tmp); err != nil {
		return nil, err
	}

	return (*Bitrate)(tmp), nil
}
