package guild

import (
	"bytes"
	"encoding/json"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/types/com"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/internal/utils/gj"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

const (
	widgetBaseSize = uint32(2 +
		len(serial.KeyEnabled) + 3 +
		len(serial.KeyID) + 4)
)

func NewWidget() Widget {
	return new(widget)
}

type widget struct {
	enabled bool
	id      Snowflake
}

func (w *widget) MarshalJSONObject(enc *gojay.Encoder) {
	gj.NewEncWrapper(enc).
		AddBool(serial.KeyEnabled, w.enabled).
		AddOptionalSnowflake(serial.KeyChannelID, w.id)
}

func (w *widget) IsNil() bool {
	return w == nil
}

func (w *widget) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch serial.Key(key) {
	case serial.KeyEnabled:
		return dec.DecodeBool(&w.enabled)
	case serial.KeyChannelID:
		var err error
		w.id, err = com.DecodeSnowflake(dec)
		return err
	}

	return nil
}

func (w *widget) NKeys() int {
	return 0
}

func (w *widget) JSONSize() int {
	return widgetBaseSize +
		utils.BoolSize(w.enabled) +
		utils.NullableSize(w.id)
}

func (w *widget) IsValid() bool {
	panic("implement me")
}

func (w *widget) Validate() error {
	panic("implement me")
}

func (w *widget) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{}, 2)
	out[serial.KeyEnabled] = w.enabled
	out[serial.KeyChannelID] = w.id
	return json.Marshal(out)
}

func (w *widget) UnmarshalJSON(in []byte) error {
	dec := gojay.BorrowDecoder(bytes.NewBuffer(in))
	defer dec.Release()

	return dec.DecodeObject(w)
}

func (w *widget) Enabled() bool {
	return w.enabled
}

func (w *widget) SetEnabled(b bool) Widget {
	w.enabled = b
	return w
}

func (w *widget) ChannelID() Snowflake {
	if w.id == nil {
		panic(e.ErrGetNull)
	}

	return w.id
}

func (w *widget) ChannelIDIsNull() bool {
	return w.id == nil
}

func (w *widget) SetChannelID(id Snowflake) Widget {
	if id == nil {
		panic(e.ErrSetNil)
	}

	w.id = id

	return w
}

func (w *widget) SetNullChannelID() Widget {
	w.id = nil
	return w
}
