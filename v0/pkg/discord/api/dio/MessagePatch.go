package dio

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/channel/message"
)

type MessagePatch interface {
	json.Marshaler

	// Content returns the current value of this request's `content` field.
	//
	// The `content` field contains the new message contents (up to 2000
	// characters).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use ContentIsReadable to check if the field is
	// present and non-null before use.
	Content() message.Content

	// ContentIsNull returns whether this request's `content` field is currently
	// null.
	ContentIsNull() bool

	// ContentIsSet returns whether this request's `content` field is currently
	// present.
	ContentIsSet() bool

	// ContentIsReadable returns whether this request's `content` field is
	// currently set and non-null.
	ContentIsReadable() bool

	// SetContent overwrites the current value of this request's `content` field.
	SetContent(message.Content) MessagePatch

	// SetNullContent overwrites the current value of this request's `content`
	// field with `null`.
	SetNullContent() MessagePatch

	// UnsetContent removes this request's `content` field.
	UnsetContent() MessagePatch

	// Embed returns the current value of this request's `embed` field.
	//
	// The `embed` field contains the embedded rich content.
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use EmbedIsReadable to check if the field is
	// present and non-null before use.
	Embed() message.Embed

	// EmbedIsNull returns whether this request's `embed` field is currently null.
	EmbedIsNull() bool

	// EmbedIsSet returns whether this request's `embed` field is currently
	// present.
	EmbedIsSet() bool

	// EmbedIsReadable returns whether this request's `embed` field is currently
	// set and non-null.
	EmbedIsReadable() bool

	// SetEmbed overwrites the current value of this request's `embed` field.
	SetEmbed(message.Embed) MessagePatch

	// SetNullEmbed overwrites the current value of this request's `embed` field
	// with `null`.
	SetNullEmbed() MessagePatch

	// UnsetEmbed removes this request's `embed` field.
	UnsetEmbed() MessagePatch

	// Flags returns the current value of this request's `flags` field.
	//
	// The `flags` field contains updated flags for a message (only
	// SUPPRESS_EMBEDS can currently be set/unset).
	//
	// If this method is called on a field that is unset or contains a null value,
	// this method will panic.  Use FlagsIsReadable to check if the field is
	// present and non-null before use.
	Flags() message.Flag

	// FlagsIsNull returns whether this request's `flags` field is currently null.
	FlagsIsNull() bool

	// FlagsIsSet returns whether this request's `flags` field is currently present.
	FlagsIsSet() bool

	// FlagsIsReadable returns whether this request's `flags` field is currently set
	// and non-null.
	FlagsIsReadable() bool

	// SetFlags overwrites the current value of this request's `flags` field.
	//
	// Only the message.FlagSuppressEmbeds may be changed on a messages flags.
	//
	// Example usages:
	//     patch.SetFlags(msg.GetFlags().Remove(message.FlagSuppressEmbeds))
	//
	//     patch.SetFlags(msg.GetFlags().Add(message.FlagSuppressEmbeds))
	SetFlags(message.Flag) MessagePatch

	// SetNullFlags overwrites the current value of this request's `flags` field
	// with `null`.
	SetNullFlags() MessagePatch

	// UnsetFlags removes this request's `flags` field.
	UnsetFlags() MessagePatch
}

func NewMessagePatch(validate bool) MessagePatch {
	return &messagePatch{ validate: validate }
}

type messagePatch struct {
	validate bool

	cont  message.TriStateContentField
	flag  message.TriStateFlagField
	embed message.TriStateEmbedField
}

func (m *messagePatch) MarshalJSON() ([]byte, error) {
	out := make(map[message.FieldKey]interface{})

	if m.cont.IsSet() {
		if m.cont.IsNull() {
			out[message.FieldKeyContent] = nil
		} else {
			out[message.FieldKeyContent] = m.cont.Get()
		}
	}

	if m.flag.IsSet() {
		if m.flag.IsNull() {
			out[message.FieldKeyFlags] = nil
		} else {
			out[message.FieldKeyFlags] = m.flag.Get()
		}
	}

	if m.embed.IsSet() {
		if m.embed.IsNull() {
			out[message.FieldKeyEmbed] = nil
		} else {
			out[message.FieldKeyEmbed] = m.embed.Get()
		}
	}

	return json.Marshal(out)
}

func (m *messagePatch) Content() message.Content {
	return m.cont.Get()
}

func (m *messagePatch) ContentIsNull() bool {
	return m.cont.IsNull()
}

func (m *messagePatch) ContentIsSet() bool {
	return m.cont.IsSet()
}

func (m *messagePatch) ContentIsReadable() bool {
	return m.cont.IsReadable()
}

func (m *messagePatch) SetContent(content message.Content) MessagePatch {
	if m.validate {
		if err := content.Validate(); err != nil {
			panic(err)
		}
	}

	m.cont.Set(content)
	return m
}

func (m *messagePatch) SetNullContent() MessagePatch {
	m.cont.SetNull()
	return m
}

func (m *messagePatch) UnsetContent() MessagePatch {
	m.cont.Unset()
	return m
}

func (m *messagePatch) Embed() message.Embed {
	return m.embed.Get()
}

func (m *messagePatch) EmbedIsNull() bool {
	return m.embed.IsNull()
}

func (m *messagePatch) EmbedIsSet() bool {
	return m.embed.IsSet()
}

func (m *messagePatch) EmbedIsReadable() bool {
	return m.embed.IsReadable()
}

func (m *messagePatch) SetEmbed(embed message.Embed) MessagePatch {
	m.embed.Set(embed)
	return m
}

func (m *messagePatch) SetNullEmbed() MessagePatch {
	m.embed.SetNull()
	return m
}

func (m *messagePatch) UnsetEmbed() MessagePatch {
	m.embed.Unset()
	return m
}

func (m *messagePatch) Flags() message.Flag {
	return m.flag.Get()
}

func (m *messagePatch) FlagsIsNull() bool {
	return m.flag.IsNull()
}

func (m *messagePatch) FlagsIsSet() bool {
	return m.flag.IsSet()
}

func (m *messagePatch) FlagsIsReadable() bool {
	return m.flag.IsReadable()
}

func (m *messagePatch) SetFlags(flag message.Flag) MessagePatch {
	if m.validate {
		if err := flag.Validate(); err != nil {
			panic(err)
		}
	}

	m.flag.Set(flag)
	return m
}

func (m *messagePatch) SetNullFlags() MessagePatch {
	m.flag.SetNull()
	return m
}

func (m *messagePatch) UnsetFlags() MessagePatch {
	m.flag.Unset()
	return m
}
