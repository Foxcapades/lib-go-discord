package discord

import "errors"

var (
	ErrBadActivityType = errors.New("unrecognized activity type value")
)

type ActivityType uint8

const (
	// Format: Playing {name}
	// Example: "Playing Rocket League"
	ActivityTypeGame ActivityType = iota

	// Format: Streaming {details}
	// Example: "Streaming Rocket League"
	//
	// Note: The streaming type currently only supports Twitch and YouTube. Only
	// https://twitch.tv/ and https://youtube.com/ urls will work.
	ActivityTypeStreaming

	// Format: Listening to {name}
	// Example: "Listening to Spotify"
	ActivityTypeListening

	// Format: {emoji} {name}
	// Example: ":smiley: I am cool"
	ActivityTypeCustom ActivityType = 1 + iota

	// Format: Competing in {name}
	// Example: "Competing in Arena World Champions"
	ActivityTypeCompeting
)

func (a ActivityType) IsValid() bool {
	return nil == a.Validate()
}

func (a ActivityType) Validate() error {
	if a == 3 || a > 5 {
		return ErrBadActivityType
	}

	return nil
}

