package activity

import (
	"encoding/json"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/serial"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewSecrets() ActivitySecrets {
	return new(secrets)
}

type secrets struct {
	join     *string
	spectate *string
	match    *string
}

func (s *secrets) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (s *secrets) IsNil() bool {
	panic("implement me")
}

func (s *secrets) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	panic("implement me")
}

func (s *secrets) NKeys() int {
	panic("implement me")
}

func (s *secrets) MarshalJSON() ([]byte, error) {
	out := make(map[serial.Key]string)

	if s.join != nil {
		out[serial.KeyJoin] = *s.join
	}

	if s.spectate != nil {
		out[serial.KeySpectate] = *s.spectate
	}

	if s.match != nil {
		out[serial.KeyMatch] = *s.match
	}

	return json.Marshal(out)
}

func (s *secrets) UnmarshalJSON(bytes []byte) error {
	tmp := make(map[serial.Key]string, 3)

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for k, v := range tmp {
		switch k {
		case serial.KeyJoin:
			s.join = &v
		case serial.KeySpectate:
			s.spectate = &v
		case serial.KeyMatch:
			s.spectate = &v
		}
	}

	return nil
}

func (s *secrets) Join() string {
	return *s.join
}

func (s *secrets) JoinIsSet() bool {
	return s.join != nil
}

func (s *secrets) SetJoin(s2 string) ActivitySecrets {
	s.join = &s2
	return s
}

func (s *secrets) UnsetJoin() ActivitySecrets {
	s.join = nil
	return s
}

func (s *secrets) Spectate() string {
	return *s.spectate
}

func (s *secrets) SpectateIsSet() bool {
	return s.spectate != nil
}

func (s *secrets) SetSpectate(s2 string) ActivitySecrets {
	s.spectate = &s2
	return s
}

func (s *secrets) UnsetSpectate() ActivitySecrets {
	s.spectate = nil
	return s
}

func (s *secrets) Match() string {
	return *s.match
}

func (s *secrets) MatchIsSet() bool {
	return s.match != nil
}

func (s *secrets) SetMatch(s2 string) ActivitySecrets {
	s.match = &s2
	return s
}

func (s *secrets) UnsetMatch() ActivitySecrets {
	s.match = nil
	return s
}
