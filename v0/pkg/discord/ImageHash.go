package discord

import (
	"github.com/foxcapades/lib-go-discord/v0/internal/js"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/francoispqt/gojay"
	"io"
)

type ImageHash string

func (i ImageHash) MarshalJSON() (buf []byte, err error) {
	buf = make([]byte, len(i))
	copy(buf, i)

	return
}

func (i *ImageHash) UnmarshalJSON(bytes []byte) error {
	if tmp, err := utils.UnmarshalString(bytes); err != nil {
		return err
	} else {
		*i = ImageHash(tmp)
	}

	return nil
}

func (i *ImageHash) ToJSONBytes() (buf []byte) {
	if i.IsNil() {
		return js.NullBytesBuf
	}

	buf = make([]byte, len(*i))
	copy(buf, *i)

	return
}

func (i *ImageHash) AppendJSONBytes(writer io.Writer) (err error) {
	_, err = writer.Write(i.ToJSONBytes())
	return
}

func (i *ImageHash) IsNil() bool {
	return i == nil
}

func (i *ImageHash) JSONSize() uint32 {
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
