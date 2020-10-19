package guild

import "github.com/foxcapades/lib-go-discord/pkg/dlib"

type Widget interface {
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
	ChannelID() dlib.Snowflake

	// ChannelIDIsNull returns whether the current value of this guild widget
	// object's `channel_id` field is `null`.
	ChannelIDIsNull() bool

	// SetChannelID overwrites the current value of this guild widget's
	// `channel_id` field with the given param value.
	SetChannelID(id dlib.Snowflake) Widget

	// SetNullChannelID sets the current value of this guild widget's `channel_id`
	// field to `null`.
	SetNullChannelID() Widget
}
