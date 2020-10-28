package user

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/francoispqt/gojay"
)

type Slice []discord.User

func (u Slice) MarshalJSON() ([]byte, error) {
	return json.Marshal([]discord.User(u))
}

func (u *Slice) UnmarshalJSON(bytes []byte) error {
	tmp := make([]*user, 0, 20)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for _, v := range tmp {
		*u = append(*u, v)
	}

	return nil
}

func (u Slice) MarshalJSONArray(enc *gojay.Encoder) {
	for _, v := range u {
		enc.AddObject(v)
	}
}

func (u Slice) IsNil() bool {
	return false
}

func (u *Slice) UnmarshalJSONArray(dec *gojay.Decoder) error {
	var tmp *user

	if err := dec.DecodeObject(tmp); err != nil {
		return err
	}

	*u = append(*u, tmp)

	return nil
}
