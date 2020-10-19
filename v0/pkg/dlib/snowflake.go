package dlib

import (
	"encoding/json"
	"strconv"
)

type Snowflake struct {
	raw uint64
}

func (s *Snowflake) SetRawValue(val uint64) {
	s.raw = val
}

func (s Snowflake) RawValue() uint64 {
	return s.raw
}

func (s Snowflake) String() string {
	return strconv.FormatUint(s.raw, 10)
}

func (s Snowflake) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Snowflake) UnmarshalJSON(bytes []byte) (err error) {
	var tmp string

	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	return s.UnmarshalString(tmp)
}

func (s *Snowflake) UnmarshalString(val string) (err error) {
	s.raw, err = strconv.ParseUint(val, 10, 64)
	return
}