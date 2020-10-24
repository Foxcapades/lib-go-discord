package discord

import (
	"encoding/json"
	"errors"
)

var (
	ErrNilUser = errors.New("called SetUser with a nil value")
)

type Ban interface {
	json.Marshaler
	json.Unmarshaler

	// the reason for the ban
	Reason() string
	ReasonIsNull() bool
	SetReason(string) Ban
	SetNullReason() Ban

	// the banned user
	User() User
	SetUser(User) Ban
}

func NewBan(validate bool) Ban {
	return &banImpl{validate: validate}
}

type banImpl struct {
	validate bool

	reason NullableString
	user   User
}

func (b *banImpl) MarshalJSON() ([]byte, error) {
	out := make(map[FieldKey]interface{}, 2)
	out[FieldKeyReason] = b.reason
	out[FieldKeyUser] = b.user

	return json.Marshal(out)
}

func (b *banImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]json.RawMessage, 2)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if v, ok := tmp[FieldKeyReason]; ok {
		if err := b.reason.UnmarshalJSON(v); err != nil {
			return err
		}
	}

	if v, ok := tmp[FieldKeyUser]; ok {
		b.user = NewUser(b.validate)

		if err := b.user.UnmarshalJSON(v); err != nil {
			return err
		}
	}

	return nil
}

func (b *banImpl) Reason() string {
	return b.reason.Get()
}

func (b *banImpl) ReasonIsNull() bool {
	return b.reason.IsNull()
}

func (b *banImpl) SetReason(s string) Ban {
	b.reason.Set(s)

	return b
}

func (b *banImpl) SetNullReason() Ban {
	b.reason.SetNull()

	return b
}

func (b *banImpl) User() User {
	return b.user
}

func (b *banImpl) SetUser(u User) Ban {
	if u == nil {
		panic(ErrNilUser)
	}

	b.user = u

	return b
}
