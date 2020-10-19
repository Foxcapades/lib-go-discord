package user

import "errors"

var (
	ErrInvalidFlag = errors.New("unrecognized user flag value(s)")
)

type Flag uint32

const FlagNone Flag = 0

const (
	FlagDiscordEmployee Flag = 1 << iota
	FlagPartneredServerOwner
	FlagHypeSquadEvents
	FlagBugHunterLevel1
)

const (
	FlagHouseBravery Flag = 1 << (6 + iota)
	FlagHouseBrilliance
	FlagHouseBalance
	FlagEarlySupporter
	FlagTeamUser
)

const (
	FlagSystem                    Flag = 1 << 12
	FlagBugHunterLevel2           Flag = 1 << 14
	FlagVerifiedBot               Flag = 1 << 16
	FlagEarlyVerifiedBotDeveloper Flag = 1 << 17
)

// IsValid returns whether the current Flag consists exclusively of valid
// values.
//
// This method will work on single flags as well as combination values (Flags
// that combined via bitwise OR)
func (f Flag) IsValid() bool {
	return nil == f.Validate()
}

// Validate checks that the current Flag consists exclusively of valid values.
//
// This method will work on single flags as well as combination values (Flags
// that combined via bitwise OR)
func (f Flag) Validate() error {
	valid := [...]Flag{
		FlagDiscordEmployee,
		FlagPartneredServerOwner,
		FlagHypeSquadEvents,
		FlagBugHunterLevel1,
		FlagHouseBravery,
		FlagHouseBrilliance,
		FlagHouseBalance,
		FlagEarlySupporter,
		FlagTeamUser,
		FlagSystem,
		FlagBugHunterLevel2,
		FlagVerifiedBot,
		FlagEarlyVerifiedBotDeveloper,
	}

	for _, v := range valid {
		f &= ^v
	}

	if f != 0 {
		return ErrInvalidFlag
	}

	return nil
}
