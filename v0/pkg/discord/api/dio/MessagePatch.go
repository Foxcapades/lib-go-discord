package dio

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"
	"github.com/foxcapades/lib-go-discord/v0/pkg/e"
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
	Content() discord.MessageContent

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
	SetContent(discord.MessageContent) MessagePatch

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
	Embed() discord.MessageEmbed

	// EmbedIsNull returns whether this request's `embed` field is currently null.
	EmbedIsNull() bool

	// EmbedIsSet returns whether this request's `embed` field is currently
	// present.
	EmbedIsSet() bool

	// EmbedIsReadable returns whether this request's `embed` field is currently
	// set and non-null.
	EmbedIsReadable() bool

	// SetEmbed overwrites the current value of this request's `embed` field.
	SetEmbed(discord.MessageEmbed) MessagePatch

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
	Flags() discord.MessageFlag

	// FlagsIsNull returns whether this request's `flags` field is currently null.
	FlagsIsNull() bool

	// FlagsIsSet returns whether this request's `flags` field is currently present.
	FlagsIsSet() bool

	// FlagsIsReadable returns whether this request's `flags` field is currently set
	// and non-null.
	FlagsIsReadable() bool

	// SetFlags overwrites the current value of this request's `flags` field.
	//
	// Only the message.MsgFlagSuppressEmbeds may be changed on a messages flags.
	//
	// Example usages:
	//     patch.SetFlags(msg.GetFlags().Remove(message.MsgFlagSuppressEmbeds))
	//
	//     patch.SetFlags(msg.GetFlags().Add(message.MsgFlagSuppressEmbeds))
	SetFlags(discord.MessageFlag) MessagePatch

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

	cont  *discord.MessageContent
	flag  *discord.MessageFlag
	embed discord.MessageEmbed

	// TODO: merge these
	nullCont bool
	nullFlag bool
	nullEmbed bool
}

func (m *messagePatch) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]interface{})

	if m.cont != nil {
		if m.ContentIsNull() {
			out[serial.KeyContent] = nil
		} else {
			out[serial.KeyContent] = m.cont
		}
	}

	if m.flag != nil {
		if m.FlagsIsNull() {
			out[serial.KeyFlags] = nil
		} else {
			out[serial.KeyFlags] = m.flag
		}
	}

	if m.embed != nil {
		if m.EmbedIsNull() {
			out[serial.KeyEmbed] = nil
		} else {
			out[serial.KeyEmbed] = m.embed
		}
	}

	return json.Marshal(out)
}

func (m *messagePatch) Content() discord.MessageContent {
	if m.cont == nil {
		if m.nullCont {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *m.cont
}

func (m *messagePatch) ContentIsNull() bool {
	return m.cont == nil && m.nullCont
}

func (m *messagePatch) ContentIsSet() bool {
	return m.cont != nil
}

func (m *messagePatch) ContentIsReadable() bool {
	return m.cont != nil
}

func (m *messagePatch) SetContent(content discord.MessageContent) MessagePatch {
	if m.validate {
		if err := content.Validate(); err != nil {
			panic(err)
		}
	}

	m.cont = &content
	m.nullCont = false

	return m
}

func (m *messagePatch) SetNullContent() MessagePatch {
	m.cont = nil
	m.nullCont = true

	return m
}

func (m *messagePatch) UnsetContent() MessagePatch {
	m.cont = nil
	m.nullCont = false

	return m
}

func (m *messagePatch) Embed() discord.MessageEmbed {
	if m.embed == nil {
		if m.nullEmbed {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return m.embed
}

func (m *messagePatch) EmbedIsNull() bool {
	return m.embed == nil && m.nullEmbed
}

func (m *messagePatch) EmbedIsSet() bool {
	return m.embed != nil
}

func (m *messagePatch) EmbedIsReadable() bool {
	return m.embed != nil
}

func (m *messagePatch) SetEmbed(embed discord.MessageEmbed) MessagePatch {
	if embed == nil {
		panic(e.ErrSetNil)
	}

	m.embed = embed

	return m
}

func (m *messagePatch) SetNullEmbed() MessagePatch {
	m.embed = nil
	m.nullEmbed = true

	return m
}

func (m *messagePatch) UnsetEmbed() MessagePatch {
	m.embed = nil
	m.nullEmbed = false

	return m
}

func (m *messagePatch) Flags() discord.MessageFlag {
	if m.flag == nil {
		if m.nullFlag {
			panic(e.ErrGetNull)
		}

		panic(e.ErrGetUnset)
	}

	return *m.flag
}

func (m *messagePatch) FlagsIsNull() bool {
	return m.flag == nil && m.nullFlag
}

func (m *messagePatch) FlagsIsSet() bool {
	return m.flag != nil
}

func (m *messagePatch) FlagsIsReadable() bool {
	return m.flag != nil
}

func (m *messagePatch) SetFlags(flag discord.MessageFlag) MessagePatch {
	m.flag = &flag
	m.nullFlag = false
	return m
}

func (m *messagePatch) SetNullFlags() MessagePatch {
	m.flag = nil
	m.nullFlag = true
	return m
}

func (m *messagePatch) UnsetFlags() MessagePatch {
	m.flag = nil
	m.nullFlag = false
	return m
}
