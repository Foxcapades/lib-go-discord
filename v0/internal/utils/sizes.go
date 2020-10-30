package utils

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"time"
)

type DataObject interface {
	IsNil() bool
	lib.Sized
}


func OptionalStringSize(s *string) uint32 {
	if s == nil {
		return 0
	}

	return StringSize(*s)
}

func StringSize(s string) uint32 {
	return uint32(len(s) + 2)
}

func OptionalSize(o DataObject) uint32 {


	return o.BufferSize()
}

func BoolSize(b bool) uint32 {
	if b {
		return 4
	} else {
		return 5
	}
}

func OptionalBoolSize(b *bool) uint32 {
	if b == nil {
		return 0
	}

	return BoolSize(*b)
}

const timeSize = uint8(len(time.RFC3339Nano)) + 2
func OptionalTimeSize(t *time.Time) uint32 {
	if t == nil {
		return 0
	}

	return uint32(timeSize)
}