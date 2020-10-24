package types

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type BanImpl struct {
	Validate bool

	reason NullableString
	user   User
}

func (b *BanImpl) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{}, 2)
	out[serial.KeyReason] = b.reason
	out[serial.KeyUser] = b.user

	return json.Marshal(out)
}

func (b *BanImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 2)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if v, ok := tmp[serial.KeyReason]; ok {
		if err := b.reason.UnmarshalJSON(v); err != nil {
			return err
		}
	}

	if v, ok := tmp[serial.KeyUser]; ok {
		b.user = &UserImpl{validate: b.Validate}

		if err := b.user.UnmarshalJSON(v); err != nil {
			return err
		}
	}

	return nil
}

func (b *BanImpl) Reason() string {
	return b.reason.Get()
}

func (b *BanImpl) ReasonIsNull() bool {
	return b.reason.IsNull()
}

func (b *BanImpl) SetReason(s string) Ban {
	b.reason.Set(s)

	return b
}

func (b *BanImpl) SetNullReason() Ban {
	b.reason.SetNull()

	return b
}

func (b *BanImpl) User() User {
	return b.user
}

func (b *BanImpl) SetUser(u User) Ban {
	if u == nil {
		panic(ErrNilUser)
	}

	b.user = u

	return b
}
