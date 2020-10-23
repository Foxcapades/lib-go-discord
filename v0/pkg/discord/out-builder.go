package discord

import "github.com/foxcapades/lib-go-discord/v0/pkg/dlib"

type outBuilder map[FieldKey]interface{}

func (o outBuilder) AppendNullable(k FieldKey, v dlib.NullableField) outBuilder {
	o[k] = v
	return o
}

func (o outBuilder) AppendOptional(k FieldKey, v dlib.OptionalField) outBuilder {
	if v.IsSet() {
		o[k] = v
	}
	return o
}

func (o outBuilder) AppendTriState(k FieldKey, v dlib.TriStateField) outBuilder {
	if v.IsSet() {
		o[k] = v
	}

	return o
}

func (o outBuilder) AppendRaw(k FieldKey, v interface{}) outBuilder {
	o[k] = v
	return o
}
