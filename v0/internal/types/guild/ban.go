package guild

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types/user"
	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewBan() Ban {
	return new(ban)
}

type ban struct {
	reason *string
	user   User
}

func (b *ban) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (b *ban) IsNil() bool {
	panic("implement me")
}

func (b *ban) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	panic("implement me")
}

func (b *ban) NKeys() int {
	panic("implement me")
}

func (b *ban) IsValid() bool {
	panic("implement me")
}

func (b *ban) Validate() error {
	panic("implement me")
}

func (b *ban) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{}, 2)
	out[serial.KeyReason] = b.reason
	out[serial.KeyUser] = b.user

	return json.Marshal(out)
}

func (b *ban) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 2)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if v, ok := tmp[serial.KeyReason]; ok {
		if err := json.Unmarshal(v, b.reason); err != nil {
			return err
		}
	}

	if v, ok := tmp[serial.KeyUser]; ok {
		b.user = user.NewUser()

		if err := b.user.UnmarshalJSON(v); err != nil {
			return err
		}
	}

	return nil
}

func (b *ban) Reason() string {
	return *b.reason
}

func (b *ban) ReasonIsNull() bool {
	return b.reason == nil
}

func (b *ban) SetReason(s string) Ban {
	b.reason = &s
	return b
}

func (b *ban) SetNullReason() Ban {
	b.reason = nil
	return b
}

func (b *ban) User() User {
	return b.user
}

func (b *ban) SetUser(u User) Ban {
	if u == nil {
		panic(ErrNilUser)
	}

	b.user = u

	return b
}
