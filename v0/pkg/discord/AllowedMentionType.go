package discord

import "errors"

var (
	ErrBadAllowedMentionType = errors.New("unrecognized allowed mention type value")
)

type AllowedMentionType string

const (
	// Controls role mentions.
	AllowedMentionTypeRoles AllowedMentionType = "roles"

	// Controls user mentions.
	AllowedMentionTypeUsers AllowedMentionType = "users"

	// Controls @everyone and @here mentions.
	AllowedMentionTypeEveryone AllowedMentionType = "everyone"
)

func (a AllowedMentionType) IsValid() bool {
	return nil == a.Validate()
}

func (a AllowedMentionType) Validate() error {
	switch a {
	case AllowedMentionTypeRoles, AllowedMentionTypeUsers, AllowedMentionTypeEveryone:
		return nil
	}

	return ErrBadAllowedMentionType
}

