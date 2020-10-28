package activity

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/internal/types"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
)

func NewEmoji() discord.ActivityEmoji {
	return new(emoji)
}

type emoji struct {
	name     string
	id       discord.Snowflake
	animated *bool
}

func (a *emoji) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddString(serial.KeyName, a.name).
		AddOptionalSnowflake(serial.KeyID, a.id).
		AddOptionalBool(serial.KeyAnimated, a.animated)
}

func (a *emoji) IsNil() bool {
	return a == nil
}

func (a *emoji) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch serial.Key(k) {

	case serial.KeyName:
		return dec.DecodeString(&a.name)

	case serial.KeyID:
		var tmp string

		if err := dec.DecodeString(&tmp); err != nil {
			return err
		}

		a.id = types.NewSnowflake()

		if err := a.id.UnmarshalString(tmp); err != nil {
			return nil
		}

	case serial.KeyAnimated:
		return dec.DecodeBool(a.animated)
	}

	return nil
}

func (a *emoji) NKeys() int {
	return 0
}

func (a *emoji) MarshalJSON() ([]byte, error) {
	return json.Marshal(make(utils.OutBuilder, 3).
		RawValue(serial.KeyName, a.name).
		Optional(serial.KeyID, a.id).
		Optional(serial.KeyAnimated, a.animated))
}

func (a *emoji) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 3)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if val, ok := tmp[serial.KeyName]; ok {
		if err := json.Unmarshal(val, &a.name); err != nil {
			return err
		}
	}

	if val, ok := tmp[serial.KeyID]; ok {
		a.id = types.NewSnowflake()
		if err := a.id.UnmarshalJSON(val); err != nil {
			return err
		}
	}

	if val, ok := tmp[serial.KeyAnimated]; ok {
		if err := json.Unmarshal(val, a.animated); err != nil {
			return err
		}
	}

	return nil
}

func (a *emoji) IsValid() bool {
	return nil == a.Validate()
}

func (a *emoji) Validate() error {

	if a.id != nil {
		return lib.NewValidationErrorSet().
			AppendRawKeyedError(serial.KeyID, a.id.Validate())
	}

	return nil
}

func (a *emoji) Name() string {
	return a.name
}

func (a *emoji) SetName(s string) discord.ActivityEmoji {
	a.name = s
	return a
}

func (a *emoji) ID() discord.Snowflake {
	if a.id == nil {
		panic(e.ErrGetUnset)
	}

	return a.id
}

func (a *emoji) IDIsSet() bool {
	return a.id != nil
}

func (a *emoji) SetID(id discord.Snowflake) discord.ActivityEmoji {
	if id == nil {
		panic(e.ErrSetNil)
	}

	a.id = id

	return a
}

func (a *emoji) UnsetID() discord.ActivityEmoji {
	a.id = nil
	return a
}

func (a *emoji) Animated() bool {
	if a.animated == nil {
		panic(e.ErrGetUnset)
	}

	return *a.animated
}

func (a *emoji) AnimatedIsSet() bool {
	return a.animated != nil
}

func (a *emoji) SetAnimated(b bool) discord.ActivityEmoji {
	a.animated = &b
	return a
}

func (a *emoji) UnsetAnimated() discord.ActivityEmoji {
	a.animated = nil
	return a
}
