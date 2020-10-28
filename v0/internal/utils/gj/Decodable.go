package gj

import "github.com/francoispqt/gojay"

type Decodable interface {
	DecodeJSONValue(dec *gojay.Decoder) error
}
