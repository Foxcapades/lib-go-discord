package integration

import "errors"

var (
	ErrBadExpireBehavior = errors.New("unrecognized expire_behavior value")
)

type ExpireBehavior uint8

const (
	ExpireBehaviorRemoveRole ExpireBehavior = iota
	ExpireBehaviorKick
)

func (e ExpireBehavior) IsValid() bool {
	return nil == e.Validate()
}

func (e ExpireBehavior) Validate() error {
	if e > ExpireBehaviorKick {
		return ErrBadExpireBehavior
	}

	return nil
}

