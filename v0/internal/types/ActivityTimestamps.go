package types

import (
	"encoding/json"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

type tmpTimes struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type ActivityTimestampsImpl struct {
	start OptionalTime
	end   OptionalTime
}

func (t *ActivityTimestampsImpl) MarshalJSON() ([]byte, error) {
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

func (t *ActivityTimestampsImpl) UnmarshalJSON(bytes []byte) error {
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

func (t *ActivityTimestampsImpl) Start() time.Time {
	return t.start.Get()
}

func (t *ActivityTimestampsImpl) StartIsSet() bool {
	return t.start.IsSet()
}

func (t *ActivityTimestampsImpl) SetStart(t2 time.Time) ActivityTimestamps {
	t.start.Set(t2)
	return t
}

func (t *ActivityTimestampsImpl) UnsetStart() ActivityTimestamps {
	t.start.Unset()
	return t
}

func (t *ActivityTimestampsImpl) End() time.Time {
	return t.end.Get()
}

func (t *ActivityTimestampsImpl) EndIsSet() bool {
	return t.end.IsSet()
}

func (t *ActivityTimestampsImpl) SetEnd(t2 time.Time) ActivityTimestamps {
	t.end.Set(t2)
	return t
}

func (t *ActivityTimestampsImpl) UnsetEnd() ActivityTimestamps {
	t.end.Unset()
	return t
}

