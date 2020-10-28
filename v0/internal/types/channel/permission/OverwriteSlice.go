package permission

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type OverwriteSlice []discord.PermissionOverwrite

func (o *OverwriteSlice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp *overwrite

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*o = append(*o, tmp)

	return nil
}

func (o OverwriteSlice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, v := range o {
		enc.AddObject(v)
	}
}

func (o OverwriteSlice) IsNil() bool {
	return false
}

func (o *OverwriteSlice) UnmarshalJSON(bytes []byte) error {
	tmp := make([]*overwrite, 0, 10)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for _, v := range tmp {
		*o = append(*o, v)
	}

	return nil
}

func (o OverwriteSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal([]discord.PermissionOverwrite(o))
}
