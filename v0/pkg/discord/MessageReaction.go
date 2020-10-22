package discord

// MessageReaction
// TODO: document me
type MessageReaction interface {
	// Count returns the current value of this record's `count` field.
	//
	// The `count` field contains the times this emoji has been used to react.
	Count() uint16

	// SetCount overwrites the current value of this record's `count` field.
	SetCount(uint16) MessageReaction

	// Me returns the current value of this record's `me` field.
	//
	// The `me` field indicates whether the current user reacted using this emoji.
	Me() bool

	// SetMe overwrites the current value of this record's `me` field.
	SetMe(bool) MessageReaction

	// Emoji returns the current value of this record's `emoji` field.
	//
	// The `emoji` field contains emoji information.
	Emoji() Emoji

	// SetEmoji overwrites the current value of this record's `emoji` field.
	SetEmoji(Emoji) MessageReaction
}
