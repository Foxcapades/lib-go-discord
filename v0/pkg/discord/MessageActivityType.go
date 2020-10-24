package discord

// ActivityType
// TODO: Document me
type MessageActivityType uint8

const (
	MsgActivityTypeJoin        MessageActivityType = 1
	MsgActivityTypeSpectate    MessageActivityType = 2
	MsgActivityTypeListen      MessageActivityType = 3
	MsgActivityTypeJoinRequest MessageActivityType = 5
)

func (a MessageActivityType) IsValid() bool {
	return nil == a.Validate()
}

func (a MessageActivityType) Validate() error {
	if a < 1 || a > 5 || a == 4 {
		return ErrBadActivityType
	}

	return nil
}

