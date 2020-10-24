package types

import (
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
)

type OutBuilder map[serial.Key]interface{}

func (o OutBuilder) AppendNullable(k serial.Key, v NullableField) OutBuilder {
	o[k] = v
	return o
}

func (o OutBuilder) AppendOptional(k serial.Key, v OptionalField) OutBuilder {
	if v.IsSet() {
		o[k] = v
	}
	return o
}

func (o OutBuilder) AppendTriState(k serial.Key, v TriStateField) OutBuilder {
	if v.IsSet() {
		o[k] = v
	}

	return o
}

func (o OutBuilder) AppendRaw(k serial.Key, v interface{}) OutBuilder {
	o[k] = v
	return o
}
