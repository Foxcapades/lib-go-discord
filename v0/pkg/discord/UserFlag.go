package discord

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/foxcapades/go-bytify/v0/bytify"
)

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

func (f *UserFlag) JSONSize() uint32 {
	if f == nil {
		return 4
	}

	return uint32(bytify.Uint32StringSize(uint32(*f)))
}

func (f *UserFlag) IsNil() bool {
	return f == nil
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

func (f UserFlag) MarshalJSON() ([]byte, error) {
	return bytify.Uint32ToByteSlice(uint32(f)), nil
}

func (f *UserFlag) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, (*uint32)(f))
}

func (f UserFlag) ToJSONBytes() []byte {
	return bytify.Uint32ToByteSlice(uint32(f))
}

func (f *UserFlag) AppendJSONBytes(writer io.Writer) (err error) {
	_, err = writer.Write(f.ToJSONBytes())
	return
}
