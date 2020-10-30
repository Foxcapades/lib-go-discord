package discord

import "errors"

var (
	ErrInvalidFlag = errors.New("unrecognized user flag value(s)")
)

type UserFlag uint32

const FlagNone UserFlag = 0

const (
	UserFlagDiscordEmployee UserFlag = 1 << iota
	UserFlagPartneredServerOwner
	UserFlagHypeSquadEvents
	UserFlagBugHunterLevel1
)

const (
	UserFlagHouseBravery UserFlag = 1 << (6 + iota)
	UserFlagHouseBrilliance
	UserFlagHouseBalance
	UserFlagEarlySupporter
	UserFlagTeamUser
)

const (
	UserFlagSystem                    UserFlag = 1 << 12
	UserFlagBugHunterLevel2           UserFlag = 1 << 14
	UserFlagVerifiedBot               UserFlag = 1 << 16
	UserFlagEarlyVerifiedBotDeveloper UserFlag = 1 << 17
)

func (f UserFlag) BufferSize() uint32 {
	switch true {
	case f < 10:
		return 1
	case f < 100:
		return 2
	case f < 1_000:
		return 3
	case f < 10_000:
		return 4
	case f < 100_000:
		return 5
	default:
		return 6
	}
}

// IsValid returns whether the current UserFlag consists exclusively of valid
// values.
//
// This method will work on single flags as well as combination values (Flags
// that combined via bitwise OR)
func (f UserFlag) IsValid() bool {
	return nil == f.Validate()
}

// Validate checks that the current UserFlag consists exclusively of valid values.
//
// This method will work on single flags as well as combination values (Flags
// that combined via bitwise OR)
func (f UserFlag) Validate() error {
	valid := [...]UserFlag{
		UserFlagDiscordEmployee,
		UserFlagPartneredServerOwner,
		UserFlagHypeSquadEvents,
		UserFlagBugHunterLevel1,
		UserFlagHouseBravery,
		UserFlagHouseBrilliance,
		UserFlagHouseBalance,
		UserFlagEarlySupporter,
		UserFlagTeamUser,
		UserFlagSystem,
		UserFlagBugHunterLevel2,
		UserFlagVerifiedBot,
		UserFlagEarlyVerifiedBotDeveloper,
	}

	for _, v := range valid {
		f &= ^v
	}

	if f != 0 {
		return ErrInvalidFlag
	}

	return nil
}
