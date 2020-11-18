package discord

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dmeta"
	"github.com/francoispqt/gojay"
)

var (
	ErrNilUser = errors.New("called SetUser with a nil value")
)

type Ban interface {
	json.Marshaler
	json.Unmarshaler

	gojay.MarshalerJSONObject
	gojay.UnmarshalerJSONObject

	dmeta.Validatable

	// the reason for the ban
	Reason() string
	ReasonIsNull() bool
	SetReason(string) Ban
	SetNullReason() Ban

	// the banned user
	User() User
	SetUser(User) Ban
}
