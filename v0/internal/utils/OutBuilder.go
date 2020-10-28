package utils

import "github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

type OutBuilder map[serial.Key]interface{}

func (o OutBuilder) RawValue(k serial.Key, v interface{}) OutBuilder {
	o[k] = v
	return o
}

func (o OutBuilder) Nullable(k serial.Key, v interface{}) OutBuilder {
	o[k] = v
	return o
}

func (o OutBuilder) Optional(k serial.Key, v interface{}) OutBuilder {
	if v != nil {
		o[k] = v
	}

	return o
}

func (o OutBuilder) TriState(k serial.Key, null bool, v interface{}) OutBuilder {
	if null {
		return o.Nullable(k, v)
	} else {
		return o.Optional(k, v)
	}
}

func (o OutBuilder) OptSlice(k serial.Key, empty bool, v interface{}) OutBuilder {
	// If the given slice is nil
	if v == nil {
		// And has been marked as absent
		if empty {
			// do not add to the output
			return o
		}

		// And has _not_ been marked as absent
		// add a constructed empty slice to the output
		return o.RawValue(k, []uint8{})
	}

	// If the given slice is _not_ nil
	// append the value to the output
	return o.RawValue(k, v)
}

func (o OutBuilder) NullSlice(k serial.Key, null bool, v interface{}) OutBuilder {
	// If the given slice is nil
	if v == nil {
		// And has been marked as null
		if null {
			// Add a nil to the output
			return o.RawValue(k, nil)
		}

		// and has _not_ been marked as nil
		// add an empty constructed slice to the output
		return o.RawValue(k, []uint8{})
	}

	// If the given slice is _not_ nil
	// add the value to the output
	return o.RawValue(k, v)
}
