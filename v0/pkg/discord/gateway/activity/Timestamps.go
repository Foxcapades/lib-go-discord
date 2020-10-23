package activity

import (
	"encoding/json"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"
	"github.com/foxcapades/lib-go-discord/v0/pkg/dlib"
)

type Timestamps interface {
	json.Marshaler
	json.Unmarshaler

	// Start returns the current value of this record's `start` field.
	//
	// The `start` field contains the unix time (in milliseconds) of when the
	// activity started.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use StartIsSet to check if the field is present before use.
	Start() time.Time

	// StartIsSet returns whether this record's `start` field is currently present.
	StartIsSet() bool

	// SetStart overwrites the current value of this record's `start` field.
	SetStart(time.Time) Timestamps

	// UnsetStart removes this record's `start` field.
	UnsetStart() Timestamps

	// End returns the current value of this record's `end` field.
	//
	// The `end` field contains the unix time (in milliseconds) of when the
	// activity ends.
	//
	// If this method is called on a field that is unset, this method will panic.
	// Use EndIsSet to check if the field is present before use.
	End() time.Time

	// EndIsSet returns whether this record's `end` field is currently present.
	EndIsSet() bool

	// SetEnd overwrites the current value of this record's `end` field.
	SetEnd(time.Time) Timestamps

	// UnsetEnd removes this record's `end` field.
	UnsetEnd() Timestamps
}

func NewTimestamps() Timestamps {
	return new(timestampsImpl)
}

type tmpTimes struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type timestampsImpl struct {
	start dlib.OptionalTime
	end   dlib.OptionalTime
}

func (t *timestampsImpl) MarshalJSON() ([]byte, error) {
	out := tmpTimes{}

	if t.start.IsSet() {
		millis := uint64(t.start.Get().UnixNano() / 1_000_000)
		out.Start = &millis
	}

	if t.end.IsSet() {
		millis := uint64(t.end.Get().UnixNano() / 1_000_000)
		out.End = &millis
	}

	return json.Marshal(out)
}

func (t *timestampsImpl) UnmarshalJSON(bytes []byte) error {
	in := tmpTimes{}

	if err := json.Unmarshal(bytes, &in); err != nil {
		return err
	}

	if in.Start != nil {
		t.start.Set(utils.MillisToUTCTime(*in.Start))
	}
	if in.End != nil {
		t.end.Set(utils.MillisToUTCTime(*in.End))
	}

	return nil
}

func (t *timestampsImpl) Start() time.Time {
	return t.start.Get()
}

func (t *timestampsImpl) StartIsSet() bool {
	return t.start.IsSet()
}

func (t *timestampsImpl) SetStart(t2 time.Time) Timestamps {
	t.start.Set(t2)
	return t
}

func (t *timestampsImpl) UnsetStart() Timestamps {
	t.start.Unset()
	return t
}

func (t *timestampsImpl) End() time.Time {
	return t.end.Get()
}

func (t *timestampsImpl) EndIsSet() bool {
	return t.end.IsSet()
}

func (t *timestampsImpl) SetEnd(t2 time.Time) Timestamps {
	t.end.Set(t2)
	return t
}

func (t *timestampsImpl) UnsetEnd() Timestamps {
	t.end.Unset()
	return t
}

