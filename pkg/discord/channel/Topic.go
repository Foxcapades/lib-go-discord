package channel

import "errors"

var (
	ErrBadTopicLength = errors.New("channel topics cannot be longer than 1024 characters")
)

type Topic string

func (t Topic) IsValid() bool {
	return nil == t.Validate()
}

func (t Topic) Validate() error {
	if len(t) > 1024 {
		return ErrBadTopicLength
	}

	return nil
}


