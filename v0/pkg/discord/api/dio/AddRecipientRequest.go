package dio

import (
	"encoding/json"
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

var (
	ErrNoChanID      = errors.New("request missing required field ChannelID")
	ErrNoUserID      = errors.New("request missing required field UserID")
	ErrNoAccessToken = errors.New("request missing required field AccessToken")
	ErrNoUserNick    = errors.New("request missing required field Nick")
)

type AddRecipientRequest interface {
	json.Marshaler
	lib.Validatable

	// AccessToken returns the access token of a user that has granted your app
	// the gdm.join scope.
	AccessToken() string

	// SetAccessToken sets the access token of a user that has granted your app
	// the gdm.join scope.
	SetAccessToken(string) AddRecipientRequest

	Nick() string
	SetNick(string) AddRecipientRequest

	// ChannelID returns the group DM channel ID.
	ChannelID() Snowflake

	// SetChannelID sets the group DM channel ID.
	SetChannelID(Snowflake) AddRecipientRequest

	// UserID returns the target user ID.
	UserID() Snowflake

	// SetUserID returns the target user ID.
	SetUserID(Snowflake) AddRecipientRequest
}

func NewAddRecipientRequest(validate bool) AddRecipientRequest {
	return &addRecipientRequest{validate: validate}
}

type addRecipientRequest struct {
	validate bool

	channelID   Snowflake
	userID      Snowflake
	accessToken string
	nick        string
}

func (a *addRecipientRequest) MarshalJSON() ([]byte, error) {
	if a.validate {
		if err := a.Validate(); err != nil {
			return nil, err
		}
	}

	return json.Marshal(map[serial.Key]interface{}{
		serial.KeyNick:        &a.nick,
		serial.KeyAccessToken: &a.accessToken,
	})
}

func (a *addRecipientRequest) IsValid() bool {
	return nil == a.Validate()
}

func (a *addRecipientRequest) Validate() error {
	if a.channelID.RawValue() == 0 {
		return ErrNoChanID
	}

	if a.userID.RawValue() == 0 {
		return ErrNoUserID
	}

	if len(a.accessToken) == 0 {
		return ErrNoAccessToken
	}

	// TODO: is this actually required?  The API docs are unclear.
	if len(a.nick) == 0 {
		return ErrNoUserNick
	}

	return nil
}

func (a *addRecipientRequest) AccessToken() string {
	return a.accessToken
}

func (a *addRecipientRequest) SetAccessToken(s string) AddRecipientRequest {
	a.accessToken = s
	return a
}

func (a *addRecipientRequest) Nick() string {
	return a.nick
}

func (a *addRecipientRequest) SetNick(s string) AddRecipientRequest {
	a.nick = s
	return a
}

func (a *addRecipientRequest) ChannelID() Snowflake {
	return a.channelID
}

func (a *addRecipientRequest) SetChannelID(id Snowflake) AddRecipientRequest {
	a.channelID = id
	return a
}

func (a *addRecipientRequest) UserID() Snowflake {
	return a.userID
}

func (a *addRecipientRequest) SetUserID(id Snowflake) AddRecipientRequest {
	a.userID = id
	return a
}
