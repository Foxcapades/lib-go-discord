package types

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewWidgetImpl() (out *WidgetImpl) {
	out = new(WidgetImpl)
	out.id.SetNull()

	return
}

type WidgetImpl struct {
	enabled bool
	id      NullableSnowflake
}

func (w *WidgetImpl) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{}, 2)
	out[serial.KeyEnabled] = w.enabled
	out[serial.KeyChannelID] = w.id
	return json.Marshal(out)
}

func (w *WidgetImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]json.RawMessage, 2)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if val, ok := tmp[serial.KeyEnabled]; ok {
		if err := json.Unmarshal(val, &w.enabled); err != nil {
			return err
		}
	}

	if val, ok := tmp[serial.KeyChannelID]; ok {
		if err := w.id.UnmarshalJSON(val); err != nil {
			return err
		}
	} else {
		w.id.SetNull()
	}

	return nil
}

func (w *WidgetImpl) Enabled() bool {
	return w.enabled
}

func (w *WidgetImpl) SetEnabled(b bool) Widget {
	w.enabled = b
	return w
}

func (w *WidgetImpl) ChannelID() Snowflake {
	return w.id.Get().(Snowflake)
}

func (w *WidgetImpl) ChannelIDIsNull() bool {
	return w.id.IsNull()
}

func (w *WidgetImpl) SetChannelID(id Snowflake) Widget {
	w.id.Set(id)
	return w
}

func (w *WidgetImpl) SetNullChannelID() Widget {
	w.id.SetNull()
	return w
}

