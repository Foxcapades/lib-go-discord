package types

import (
	"encoding/json"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type ActivitySecretsImpl struct {
	join     OptionalString
	spectate OptionalString
	match    OptionalString
}

func (s *ActivitySecretsImpl) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]string)

	if s.join.IsSet() {
		out[serial.KeyJoin] = s.join.Get()
	}

	if s.spectate.IsSet() {
		out[serial.KeySpectate] = s.spectate.Get()
	}

	if s.match.IsSet() {
		out[serial.KeyMatch] = s.match.Get()
	}

	return json.Marshal(out)
}

func (s *ActivitySecretsImpl) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]string, 3)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for k, v := range tmp {
		switch k {
		case serial.KeyJoin:
			s.join.Set(v)
		case serial.KeySpectate:
			s.spectate.Set(v)
		case serial.KeyMatch:
			s.spectate.Set(v)
		}
	}

	return nil
}

func (s *ActivitySecretsImpl) Join() string {
	return s.join.Get()
}

func (s *ActivitySecretsImpl) JoinIsSet() bool {
	return s.join.IsSet()
}

func (s *ActivitySecretsImpl) SetJoin(s2 string) ActivitySecrets {
	s.join.Set(s2)
	return s
}

func (s *ActivitySecretsImpl) UnsetJoin() ActivitySecrets {
	s.join.Unset()
	return s
}

func (s *ActivitySecretsImpl) Spectate() string {
	return s.spectate.Get()
}

func (s *ActivitySecretsImpl) SpectateIsSet() bool {
	return s.spectate.IsSet()
}

func (s *ActivitySecretsImpl) SetSpectate(s2 string) ActivitySecrets {
	s.spectate.Set(s2)
	return s
}

func (s *ActivitySecretsImpl) UnsetSpectate() ActivitySecrets {
	s.spectate.Unset()
	return s
}

func (s *ActivitySecretsImpl) Match() string {
	return s.match.Get()
}

func (s *ActivitySecretsImpl) MatchIsSet() bool {
	return s.match.IsSet()
}

func (s *ActivitySecretsImpl) SetMatch(s2 string) ActivitySecrets {
	s.match.Set(s2)
	return s
}

func (s *ActivitySecretsImpl) UnsetMatch() ActivitySecrets {
	s.match.Unset()
	return s
}

