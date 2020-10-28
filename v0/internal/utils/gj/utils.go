package gj

import (
	"github.com/francoispqt/gojay"
	"time"
)

func DecodeTime(dec *gojay.Decoder) (*time.Time, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	if tmp == nil {
		return nil, nil
	}

	if out, err := time.Parse(time.RFC3339Nano, *tmp); err != nil {
		return nil, err
	} else {
		return &out, nil
	}
}
