package discord

import (
	"encoding/json"
)

type Widget interface {
	json.Marshaler
	json.Unmarshaler

	// Enabled returns the value of the `enabled` field currently set on this
	// guild widget.
	//
	// The `enabled` flag indicates whether the current widget is enabled.
	Enabled() bool

	// SetEnabled overwrites the current value of this guild widget's `enabled`
	// field with the given param value.
	SetEnabled(bool) Widget

	// ChannelID returns the value of the `channel_id` field currently set on this
	// guild widget.
	//
	// The `channel_id` field contains the widget channel ID.
	//
	// If the guild widget's `channel_id` field is null when this method is
	// called, this method will panic.  Check the field's null state with the
	// ChannelIDIsNull method.
	ChannelID() Snowflake

	// ChannelIDIsNull returns whether the current value of this guild widget
	// object's `channel_id` field is `null`.
	ChannelIDIsNull() bool

	// SetChannelID overwrites the current value of this guild widget's
	// `channel_id` field with the given param value.
	SetChannelID(id Snowflake) Widget

	// SetNullChannelID sets the current value of this guild widget's `channel_id`
	// field to `null`.
	SetNullChannelID() Widget
}

func NewWidget() Widget {
	out := new(widgetImpl)
	out.id.SetNull()
	return out
}

type widgetImpl struct {
	enabled bool
	id      NullableSnowflake
}

func (w *widgetImpl) MarshalJSON() ([]byte, error) {
	out := make(map[FieldKey]interface{}, 2)
	out[FieldKeyEnabled] = w.enabled
	out[FieldKeyChannelID] = w.id
	return json.Marshal(out)
}

func (w *widgetImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]json.RawMessage, 2)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	if val, ok := tmp[FieldKeyEnabled]; ok {
		if err := json.Unmarshal(val, &w.enabled); err != nil {
			return err
		}
	}

	if val, ok := tmp[FieldKeyChannelID]; ok {
		if err := w.id.UnmarshalJSON(val); err != nil {
			return err
		}
	} else {
		w.id.SetNull()
	}

	return nil
}

func (w *widgetImpl) Enabled() bool {
	return w.enabled
}

func (w *widgetImpl) SetEnabled(b bool) Widget {
	w.enabled = b
	return w
}

func (w *widgetImpl) ChannelID() Snowflake {
	return w.id.Get()
}

func (w *widgetImpl) ChannelIDIsNull() bool {
	return w.id.IsNull()
}

func (w *widgetImpl) SetChannelID(id Snowflake) Widget {
	w.id.Set(id)
	return w
}

func (w *widgetImpl) SetNullChannelID() Widget {
	w.id.SetNull()
	return w
}
