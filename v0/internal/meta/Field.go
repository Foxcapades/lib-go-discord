package meta

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
)

type Field interface {
	json.Marshaler
	json.Unmarshaler

	Sized
	Nillable
	Bytable
}

type ObjectField interface {
	Field

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject
}

type ArrayField interface {
	Field

	gojay.MarshalerJSONArray
	gojay.UnmarshalerJSONArray
}