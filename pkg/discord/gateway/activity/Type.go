package activity

import "errors"

var (
	ErrBadActivityType = errors.New("unrecognized activity type value")
)

type Type uint8

const (
	// Format: Playing {name}
	// Example: "Playing Rocket League"
	TypeGame Type = iota

	// Format: Streaming {details}
	// Example: "Streaming Rocket League"
	//
	// Note: The streaming type currently only supports Twitch and YouTube. Only
	// https://twitch.tv/ and https://youtube.com/ urls will work.
	TypeStreaming

	// Format: Listening to {name}
	// Example: "Listening to Spotify"
	TypeListening

	// Format: {emoji} {name}
	// Example: ":smiley: I am cool"
	TypeCustom Type = 1 + iota

	// Format: Competing in {name}
	// Example: "Competing in Arena World Champions"
	TypeCompeting
)

func (a Type) IsValid() bool {
	return nil == a.Validate()
}

func (a Type) Validate() error {
	if a == 3 || a > 5 {
		return ErrBadActivityType
	}

	return nil
}

