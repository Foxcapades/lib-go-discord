package discord

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrBadUsernameLength   = errors.New("usernames must be 2-32 characters in length")
	ErrBadUsernameContents = errors.New(`usernames cannot contain "@", ":", or "` + "```" + `"`)
	ErrBadUsernameValue    = errors.New(`usernames cannot be "discordtag", "everyone", "here"`)
)

var excessSpace *regexp.Regexp

func init() {
	var err error

	if excessSpace, err = regexp.Compile(" {2,}"); err != nil {
		panic(err)
	}
}

type Username string

func (u Username) JSONSize() int {
	return uint32(len(u) + 2)
}

func (u Username) IsValid() bool {
	return nil == u.Validate()
}

func (u Username) Validate() error {
	ln := len(scrubName(string(u)))

	if ln < 2 || ln > 32 {
		return ErrBadUsernameLength
	}

	switch u {
	case disallowedName1, disallowedName2, disallowedName3:
		return ErrBadUsernameValue
	}

	if strings.ContainsAny(string(u), disallowedNameChars) {
		return ErrBadUsernameContents
	}

	if strings.Contains(string(u), disallowedNameSubst) {
		return ErrBadUsernameContents
	}

	return nil
}

const (
	disallowedName1 = "discordtag"
	disallowedName2 = "everyone"
	disallowedName3 = "here"
)

const (
	disallowedNameChars = "@#:"
	disallowedNameSubst = "```"
)

func scrubName(in string) string {
	return excessSpace.ReplaceAllString(strings.TrimSpace(in), " ")
}