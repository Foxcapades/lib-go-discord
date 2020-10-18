package audit

import "errors"

var (
	ErrBadReasonLength = errors.New("audit log entry reason values cannot be greater than 512 characters")
)

type LogReason string

func (l LogReason) IsValid() bool {
	return nil == l.Validate()
}

func (l LogReason) Validate() error {
	if len(l) > 512 {
		return ErrBadReasonLength
	}

	return nil
}


