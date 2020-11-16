package guild

import (
	"bytes"
	"strconv"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
)

type FeatureSlice []discord.GuildFeature

func (f FeatureSlice) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, f.JSONSize()))
	enc := gojay.BorrowEncoder(buf)
	defer enc.Release()

	if err := enc.EncodeArray(f); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (f *FeatureSlice) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeArray(f)
}

func (f FeatureSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for i := range f {
		enc.AddString(string(f[i]))
	}
}

func (f FeatureSlice) IsNil() bool {
	return false
}

func (f *FeatureSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp string

	if err := dec.DecodeString(&tmp); err != nil {
		return err
	}

	*f = append(*f, discord.GuildFeature(tmp))

	return nil
}

func (f FeatureSlice) JSONSize() (size int) {
	size = 2 + uint32(len(f)-1)

	for i := range f {
		size += f[i].JSONSize()
	}

	return
}

func (f FeatureSlice) IsValid() bool {
	return nil == f.Validate()
}

func (f FeatureSlice) Validate() error {
	out := lib.NewValidationErrorSet()
	ln := len(f)

	for i := 0; i < ln; i++ {
		if err := f[i].Validate(); err != nil {
			out.AppendRawKeyedError(serial.Key("["+strconv.Itoa(i)+"]"), err)
		}
	}

	if out.Len() > 0 {
		return out
	}

	return nil
}
