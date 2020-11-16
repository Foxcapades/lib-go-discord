package utils

import (
	"github.com/foxcapades/lib-go-discord/v0/internal/meta"
	"time"
)

// TriStateStringSize returns the size in bytes of the given value for the three
// possible states: absent, null, populated.
//
// If the given value is nil and null is false, returns 0 (absent).
//
// If the given value is nil and null is true, returns 4 (len("null")).
//
// If the given value is non-nil, len(*value) + 2 (for quotes) is returned.
func TriStateStringSize(s *string, null bool) uint32 {
	if s == nil {
		if null {
			return 4
		}

		return 0
	}

	return StringSize(*s)
}

// OptionalStringSize returns either 0 if the given value is nil, or len(*value)
// + 2 (for quotes).
func OptionalStringSize(s *string) uint32 {
	if s == nil {
		return 0
	}

	return StringSize(*s)
}

// StringSize returns len(string) + 2.
func StringSize(s string) uint32 {
	return uint32(len(s) + 2)
}

func NullableSize(o meta.Field) uint32 {
	if o.IsNil() {
		return 4
	}

	return o.JSONSize()
}

func OptionalSize(o meta.Field) uint32 {
	if o.IsNil() {
		return 0
	}

	return o.JSONSize()
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
