package guild

import (
	"bytes"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/francoispqt/gojay"
	"strconv"
)

type EmojiSlice []discord.Emoji

func (e EmojiSlice) IsValid() bool {
	return nil == e.Validate()
}

func (e EmojiSlice) Validate() error {
	out := lib.NewValidationErrorSet()
	ln  := len(e)

	for i := 0; i < ln; i++ {
		if err := e[i].Validate(); err != nil {
			key := "[" + strconv.Itoa(i) + "]"
			out.AppendRawKeyedError(serial.Key(key), err)
		}
	}

	if out.Len() > 0 {
		return out
	}

	return nil
}

func (e EmojiSlice) JSONSize() (size int) {
	size = 2 + uint32(len(e) - 1)

	for _, v := range e {
		size += v.JSONSize()
	}

	return
}

func (e EmojiSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, v := range e {
		enc.AddObject(v)
	}
}

func (e EmojiSlice) IsNil() bool {
	return false
}

func (e *EmojiSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	tmp := NewEmoji()

	if err := dec.DecodeObject(tmp); err != nil {
		return nil
	}

	*e = append(*e, tmp)

	return nil
}

func (e EmojiSlice) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, e.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeArray(e); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (e *EmojiSlice) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeArray(e)
}
