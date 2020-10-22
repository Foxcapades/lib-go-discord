package discord

import (
	"errors"
	"strings"
)

var (
	ErrBadGuildNameLength = errors.New("guild names must be 2-100 characters" +
		" in length excluding leading/trailing whitespace")
)

type GuildName string

func (n GuildName) IsValid() bool {
	return nil == n.Validate()
}

func (n GuildName) Validate() error {
	v := strings.TrimSpace(string(n))
	l := len(v)

	if l < 2 || l > 100 {
		return ErrBadGuildNameLength
	}

	return nil
}

