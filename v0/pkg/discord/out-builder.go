package discord

type outBuilder map[FieldKey]interface{}

func (o outBuilder) AppendNullable(k FieldKey, v NullableField) outBuilder {
	o[k] = v
	return o
}

func (o outBuilder) AppendOptional(k FieldKey, v OptionalField) outBuilder {
	if v.IsSet() {
		o[k] = v
	}
	return o
}

func (o outBuilder) AppendTriState(k FieldKey, v TriStateField) outBuilder {
	if v.IsSet() {
		o[k] = v
	}

	return o
}

func (o outBuilder) AppendRaw(k FieldKey, v interface{}) outBuilder {
	o[k] = v
	return o
}
