package message

import "errors"

var (
	ErrBadActivityType = errors.New("unrecognized activity type value")
)

// ActivityType
// TODO: Document me
type ActivityType uint8

const (
	ActivityTypeJoin        ActivityType = 1
	ActivityTypeSpectate    ActivityType = 2
	ActivityTypeListen      ActivityType = 3
	ActivityTypeJoinRequest ActivityType = 5
)

func (a ActivityType) IsValid() bool {
	return nil == a.Validate()
}

func (a ActivityType) Validate() error {
	if a < 1 || a > 5 || a == 4 {
		return ErrBadActivityType
	}

	return nil
}

