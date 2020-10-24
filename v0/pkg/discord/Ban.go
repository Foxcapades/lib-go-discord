package discord

import (
	"encoding/json"
	"errors"
)

var (
	ErrNilUser = errors.New("called SetUser with a nil value")
)

type Ban interface {
	json.Marshaler
	json.Unmarshaler

	// the reason for the ban
	Reason() string
	ReasonIsNull() bool
	SetReason(string) Ban
	SetNullReason() Ban

	// the banned user
	User() User
	SetUser(User) Ban
}
