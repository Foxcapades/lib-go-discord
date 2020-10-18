package comm

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidColorValue = errors.New("rgb color value cannot be larger than 0x00FFFFFF")
)

const (
	redMask   = 0x00FF0000
	greenMask = 0x0000FF00
	blueMask  = 0x000000FF
	max       = 0x00FFFFFF
)

type Color interface {
	json.Marshaler
	json.Unmarshaler

	RawValue() uint32
	SetRawValue(val uint32) Color

	Red() uint8
	Green() uint8
	Blue() uint8
}

func NewColor() Color {
	return new(colorImpl)
}

type colorImpl struct {
	value uint32
}

func (c *colorImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *colorImpl) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &c.value)
}

func (c *colorImpl) RawValue() uint32 {
	return c.value
}

func (c *colorImpl) SetRawValue(val uint32) Color {
	if val > max {
		panic(ErrInvalidColorValue)
	}

	c.value = val

	return c
}

func (c *colorImpl) Red() uint8 {
	return uint8((c.value & redMask) >> 16)
}

func (c *colorImpl) Green() uint8 {
	return uint8((c.value & greenMask) >> 8)
}

func (c *colorImpl) Blue() uint8 {
	return uint8(c.value & blueMask)
}

