package gj

import (
	"github.com/francoispqt/gojay"
	"time"
)

func DecodeReqTime(dec *gojay.Decoder) (time.Time, error) {
	var tmp string

	if err := dec.String(&tmp); err != nil {
		return time.Time{}, err
	}

	if out, err := time.Parse(time.RFC3339Nano, tmp); err != nil {
		return time.Time{}, err
	} else {
		return out, nil
	}
}

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

func EncodeTime(t *time.Time) *string {
	if t == nil {
		return nil
	}

	tmp := t.Format(time.RFC3339Nano)
	return &tmp
}

func DecodeOptUint8(dec *gojay.Decoder) (*uint8, error) {
	var tmp *uint8

	if err := dec.Uint8Null(&tmp); err != nil {
		return nil, err
	}

	return tmp, nil
}