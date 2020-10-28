package types

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
	"strconv"
	"time"

	. "github.com/foxcapades/lib-go-discord/v0/pkg/discord"
)

const (
	timeMask   uint64 = 0xFFFF_FFFF_FFC0_0000
	workIdMask uint64 = 0x0000_0000_003E_0000
	procIdMask uint64 = 0x0000_0000_0001_F000
	countMask  uint64 = 0x0000_0000_0000_0FFF

	timeShift   uint8 = 22
	workIdShift uint8 = 17
	procIdShift uint8 = 12

	ns2ms = 1_000_000
)

func NewSnowflake() Snowflake {
	return new(snowflake)
}

// DecodeSnowflake attempts to decode a JSON representation of a Snowflake via
// gojay.
//
// This method may return nil if the input json value was `null`.
func DecodeSnowflake(dec *gojay.Decoder) (Snowflake, error) {
	var tmp *string

	if err := dec.StringNull(&tmp); err != nil {
		return nil, err
	}

	if tmp == nil {
		return nil, nil
	}

	sn := NewSnowflake()

	if err := sn.UnmarshalString(*tmp); err != nil {
		return nil, err
	}

	return sn, nil
}

// UnmarshalJSONSnowflake attempts to unmarshal a JSON representation of a
// Snowflake via the standard library json package.
func UnmarshalJSONSnowflake(buf []byte) (out Snowflake, err error) {
	var tmp *string

	if err = json.Unmarshal(buf, &tmp); err != nil {
		return
	}

	if tmp == nil {
		return
	}

	out = NewSnowflake()

	if err = out.UnmarshalString(*tmp); err != nil {
		return nil, err
	}

	return
}

type snowflake struct {
	raw uint64
}

func (s *snowflake) RawValue() uint64 {
	return s.raw
}

func (s *snowflake) SetRawValue(val uint64) Snowflake {
	s.raw = val

	return s
}

func (s *snowflake) Timestamp() time.Time {
	tmp := (s.raw >> timeShift) + EpochMillis
	sec := tmp / 1000
	nano := (tmp % 1000) * ns2ms

	return time.Unix(int64(sec), int64(nano))
}

func (s *snowflake) SetTimestamp(t time.Time) Snowflake {
	tmp := uint64(t.UnixNano())/ns2ms - EpochMillis

	s.raw &= ^timeMask
	s.raw |= (tmp << timeShift) & timeMask

	return s
}

func (s *snowflake) InternalWorkerID() uint8 {
	return uint8((s.raw & workIdMask) >> workIdShift)
}

func (s *snowflake) SetInternalWorkerID(id uint8) Snowflake {
	s.raw &= ^workIdMask
	s.raw |= (uint64(id) << workIdShift) & workIdMask

	return s
}

func (s *snowflake) InternalProcessID() uint8 {
	return uint8((s.raw & procIdMask) >> procIdShift)
}

func (s *snowflake) SetInternalProcessID(id uint8) Snowflake {
	s.raw &= ^procIdMask
	s.raw |= (uint64(id) << procIdShift) & procIdMask

	return s
}

func (s *snowflake) CounterValue() uint16 {
	return uint16(s.raw & countMask)
}

func (s *snowflake) SetCounterValue(id uint16) Snowflake {
	s.raw &= ^countMask
	s.raw |= uint64(id) & countMask

	return s
}

func (s *snowflake) String() string {
	return strconv.FormatUint(s.raw, 10)
}

func (s *snowflake) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *snowflake) UnmarshalJSON(bytes []byte) (err error) {
	var tmp string

	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	return s.UnmarshalString(tmp)
}

func (s *snowflake) UnmarshalString(val string) (err error) {
	s.raw, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}

	return
}

func (s *snowflake) IsValid() bool {
	return nil == s.Validate()
}

func (s *snowflake) Validate() error {
	if s.raw == 0 {
		return ErrEmptySnowflake
	}

	if s.raw&timeMask == 0 {
		return ErrNoSnowflakeTime
	}

	if s.raw&workIdMask == 0 {
		return ErrNoSnowflakeWorkerID
	}

	if s.raw&procIdMask == 0 {
		return ErrNoSnowflakeProcID
	}

	if s.raw&countMask == 0 {
		return ErrNoSnowflakeCounter
	}

	return nil
}
