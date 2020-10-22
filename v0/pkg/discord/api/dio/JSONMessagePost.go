package dio

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/comm"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/message"
)

type JSONMessagePost interface {
	json.Marshaler
	json.Unmarshaler

	// Content returns the current value of this request's `content` field.
	//
	// The `content` field contains the message contents (up to 2000 characters).
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use ContentIsSet to check if the field is present before use.
	Content() discord.MessageContent

	// ContentIsSet returns whether this request's `content` field is currently
	// present.
	ContentIsSet() bool

	// SetContent overwrites the current value of this request's `content` field.
	SetContent(discord.MessageContent) JSONMessagePost

	// UnsetContent removes this request's `content` field.
	UnsetContent() JSONMessagePost

	// Nonce returns the current value of this request's `nonce` field.
	//
	// The `nonce` field contains a nonce that can be used for optimistic message
	// sending.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use NonceIsSet to check if the field is present before use.
	Nonce() comm.Nonce

	// NonceIsSet returns whether this request's `nonce` field is currently
	// present.
	NonceIsSet() bool

	// SetNonce overwrites the current value of this request's `nonce` field.
	SetNonce(comm.Nonce) JSONMessagePost

	// UnsetNonce removes this request's `nonce` field.
	UnsetNonce() JSONMessagePost

	// TTS returns the current value of this request's `tts` field.
	//
	// The `tts` field indicates whether this is a TTS message.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use TTSIsSet to check if the field is present before use.
	TTS() bool

	// TTSIsSet returns whether this request's `tts` field is currently present.
	TTSIsSet() bool

	// SetTTS overwrites the current value of this request's `tts` field.
	SetTTS(bool) JSONMessagePost

	// UnsetTTS removes this request's `tts` field.
	UnsetTTS() JSONMessagePost

	// Embed returns the current value of this request's `embed` field.
	//
	// The `embed` field contains the embedded rich content.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use EmbedIsSet to check if the field is present before use.
	Embed() message.Embed

	// EmbedIsSet returns whether this request's `embed` field is currently present.
	EmbedIsSet() bool

	// SetEmbed overwrites the current value of this request's `embed` field.
	//
	// For the embed object, you can set every field except `type` (it will be
	// rich regardless of if you try to set it), provider, video, and any height,
	// width, or proxy_url values for images.
	SetEmbed(message.Embed) JSONMessagePost

	// UnsetEmbed removes this request's `embed` field.
	UnsetEmbed() JSONMessagePost

	// AllowedMentions returns the current value of this request's `allowed_mentions` field.
	//
	// The `allowed_mentions` field contains the allowed mentions for a message.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use AllowedMentionsIsSet to check if the field is present before use.
	AllowedMentions() message.AllowedMentions

	// AllowedMentionsIsSet returns whether this request's `allowed_mentions` field is currently present.
	AllowedMentionsIsSet() bool

	// SetAllowedMentions overwrites the current value of this request's `allowed_mentions` field.
	SetAllowedMentions(message.AllowedMentions) JSONMessagePost

	// UnsetAllowedMentions removes this request's `allowed_mentions` field.
	UnsetAllowedMentions() JSONMessagePost
}

func NewJSONMessagePost(validate bool) JSONMessagePost {
	panic("implement me")
}