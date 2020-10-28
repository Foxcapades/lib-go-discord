package activity

import (
	"encoding/json"
	"time"

	"github.com/francoispqt/gojay"

	"github.com/foxcapades/lib-go-discord/v0/internal/utils"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

func NewTimestamps() ActivityTimestamps {
	return new(timestamps)
}

type timestamps struct {
	start *time.Time
	end   *time.Time
}

func (t *timestamps) MarshalJSONObject(enc *gojay.Encoder) {
	panic("implement me")
}

func (t *timestamps) IsNil() bool {
	panic("implement me")
}

func (t *timestamps) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	panic("implement me")
}

func (t *timestamps) NKeys() int {
	panic("implement me")
}

func (t *timestamps) IsValid() bool {
	panic("implement me")
}

func (t *timestamps) Validate() error {
	panic("implement me")
}

func (t *timestamps) MarshalJSON() ([]byte, error) {
	out := tmpTimes{}

	if t.start != nil {
		millis := uint64(t.start.UnixNano() / 1_000_000)
		out.Start = &millis
	}

	if t.end != nil {
		millis := uint64(t.end.UnixNano() / 1_000_000)
		out.End = &millis
	}

	return json.Marshal(out)
}

func (t *timestamps) UnmarshalJSON(bytes []byte) error {
	in := tmpTimes{}

	if err := json.Unmarshal(bytes, &in); err != nil {
		return err
	}

	if in.Start != nil {
		tmp := utils.MillisToUTCTime(*in.Start)
		t.start = &tmp
	}

	if in.End != nil {
		tmp := utils.MillisToUTCTime(*in.End)
		t.end = &tmp
	}

	return nil
}

func (t *timestamps) Start() time.Time {
	return *t.start
}

func (t *timestamps) StartIsSet() bool {
	return t.start != nil
}

func (t *timestamps) SetStart(t2 time.Time) ActivityTimestamps {
	t.start = &t2
	return t
}

func (t *timestamps) UnsetStart() ActivityTimestamps {
	t.start = nil
	return t
}

func (t *timestamps) End() time.Time {
	return *t.end
}

func (t *timestamps) EndIsSet() bool {
	return t.end != nil
}

func (t *timestamps) SetEnd(t2 time.Time) ActivityTimestamps {
	t.end = &t2
	return t
}

func (t *timestamps) UnsetEnd() ActivityTimestamps {
	t.end = nil
	return t
}

type tmpTimes struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}
