package dio

import (
	"errors"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
	"strconv"
)

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

func (r RecordLimit) String() string {
	return strconv.FormatUint(uint64(r), 10)
}

type OptionalRecordLimit struct {
	value *RecordLimit
}

func (o *OptionalRecordLimit) Unset() {
	o.value = nil
}

func (o *OptionalRecordLimit) IsSet() bool {
	return o.value != nil
}

func (o *OptionalRecordLimit) IsUnset() bool {
	return o.value == nil
}

func (o *OptionalRecordLimit) Set(limit RecordLimit) {
	o.value = &limit
}

func (o *OptionalRecordLimit) Get() RecordLimit {
	if o.value == nil {
		panic(dlib.ErrUnsetField)
	}

	return *o.value
}
