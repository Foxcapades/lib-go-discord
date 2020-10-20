package dio

import "errors"

var (
	ErrBadRecordLimit = errors.New("invalid record limit value, must be in the range 1-100")
)

type RecordLimit uint8

func (r RecordLimit) IsValid() bool {
	return nil == r.Validate()
}

func (r RecordLimit) Validate() error {
	if r < 1 || r > 100 {
		return ErrBadRecordLimit
	}

	return nil
}



