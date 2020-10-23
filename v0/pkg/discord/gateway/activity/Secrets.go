package activity

import (
	"encoding/json"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type Secrets interface {
	json.Marshaler
	json.Unmarshaler

	// Join returns the current value of this record's `join` field.
	//
	// The `join` field contains the secret for joining a party.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use JoinIsSet to check if the field is present before use.
	Join() string

	// JoinIsSet returns whether this record's `join` field is currently present.
	JoinIsSet() bool

	// SetJoin overwrites the current value of this record's `join` field.
	SetJoin(string) Secrets

	// UnsetJoin removes this record's `join` field.
	UnsetJoin() Secrets

	// Spectate returns the current value of this record's `spectate` field.
	//
	// The `spectate` field contains the secret for spectating a game.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use SpectateIsSet to check if the field is present before use.
	Spectate() string

	// SpectateIsSet returns whether this record's `spectate` field is currently present.
	SpectateIsSet() bool

	// SetSpectate overwrites the current value of this record's `spectate` field.
	SetSpectate(string) Secrets

	// UnsetSpectate removes this record's `spectate` field.
	UnsetSpectate() Secrets

	// Match returns the current value of this record's `match` field.
	//
	// The `match` field contains the secret for a specific instanced match.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use MatchIsSet to check if the field is present before use.
	Match() string

	// MatchIsSet returns whether this record's `match` field is currently present.
	MatchIsSet() bool

	// SetMatch overwrites the current value of this record's `match` field.
	SetMatch(string) Secrets

	// UnsetMatch removes this record's `match` field.
	UnsetMatch() Secrets
}

func NewSecrets() Secrets {
	return new(secretsImpl)
}

type secretsImpl struct {
	join     dlib.OptionalString
	spectate dlib.OptionalString
	match    dlib.OptionalString
}

func (s *secretsImpl) MarshalJSON() ([]byte, error) {
	out := make(map[FieldKey]string)

	if s.join.IsSet() {
		out[FieldKeyJoin] = s.join.Get()
	}

	if s.spectate.IsSet() {
		out[FieldKeySpectate] = s.spectate.Get()
	}

	if s.match.IsSet() {
		out[FieldKeyMatch] = s.match.Get()
	}

	return json.Marshal(out)
}

func (s *secretsImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[FieldKey]string, 3)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for k, v := range tmp {
		switch k {
		case FieldKeyJoin:
			s.join.Set(v)
		case FieldKeySpectate:
			s.spectate.Set(v)
		case FieldKeyMatch:
			s.spectate.Set(v)
		}
	}

	return nil
}

func (s *secretsImpl) Join() string {
	return s.join.Get()
}

func (s *secretsImpl) JoinIsSet() bool {
	return s.join.IsSet()
}

func (s *secretsImpl) SetJoin(s2 string) Secrets {
	s.join.Set(s2)
	return s
}

func (s *secretsImpl) UnsetJoin() Secrets {
	s.join.Unset()
	return s
}

func (s *secretsImpl) Spectate() string {
	return s.spectate.Get()
}

func (s *secretsImpl) SpectateIsSet() bool {
	return s.spectate.IsSet()
}

func (s *secretsImpl) SetSpectate(s2 string) Secrets {
	s.spectate.Set(s2)
	return s
}

func (s *secretsImpl) UnsetSpectate() Secrets {
	s.spectate.Unset()
	return s
}

func (s *secretsImpl) Match() string {
	return s.match.Get()
}

func (s *secretsImpl) MatchIsSet() bool {
	return s.match.IsSet()
}

func (s *secretsImpl) SetMatch(s2 string) Secrets {
	s.match.Set(s2)
	return s
}

func (s *secretsImpl) UnsetMatch() Secrets {
	s.match.Unset()
	return s
}

