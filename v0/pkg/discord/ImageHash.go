package discord

import "github.com/francoispqt/gojay"

type ImageHash string

func (i *ImageHash) IsNil() bool {
	return i == nil
}

func (i *ImageHash) JSONSize() int {
	if i == nil {
		return 4
	}

	return uint32(len(*i)) + 2
}

func DecodeImageHash(dec *gojay.Decoder) (*ImageHash, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	return (*ImageHash)(tmp), nil
}
