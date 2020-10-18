package message

import "github.com/foxcapades/lib-go-discord/pkg/discord/comm"

// Reaction
// TODO: document me
type Reaction interface {
	// Count returns the current value of this record's `count` field.
	//
	// The `count` field contains the times this emoji has been used to react.
	Count() uint16

	// SetCount overwrites the current value of this record's `count` field.
	SetCount(uint16) Reaction

	// Me returns the current value of this record's `me` field.
	//
	// The `me` field indicates whether the current user reacted using this emoji.
	Me() bool

	// SetMe overwrites the current value of this record's `me` field.
	SetMe(bool) Reaction

	// Emoji returns the current value of this record's `emoji` field.
	//
	// The `emoji` field contains emoji information.
	Emoji() comm.Emoji

	// SetEmoji overwrites the current value of this record's `emoji` field.
	SetEmoji(comm.Emoji) Reaction
}
