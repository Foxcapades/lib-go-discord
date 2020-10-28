package discord

import "github.com/francoispqt/gojay"

type ImageHash string

func DecodeImageHash(dec *gojay.Decoder) (*ImageHash, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	return (*ImageHash)(tmp), nil
}
