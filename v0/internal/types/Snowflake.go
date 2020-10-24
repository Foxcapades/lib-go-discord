package types

import (
	"encoding/json"
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

	ns2ms uint64 = 1_000_000
)

func NewSnowflakeImpl(validate bool) *SnowflakeImpl {
	return &SnowflakeImpl{validate: validate}
}

type SnowflakeImpl struct {
	validate bool
	raw      uint64
}

func (s *SnowflakeImpl) RawValue() uint64 {
	return s.raw
}

func (s *SnowflakeImpl) SetRawValue(val uint64) Snowflake {
	s.raw = val

	if s.validate {
		if err := s.Validate(); err != nil {
			panic(err)
		}
	}

	return s
}

func (s *SnowflakeImpl) Timestamp() time.Time {
	tmp := (s.raw >> timeShift) + EpochMillis
	sec := tmp / 1000
	nano := (tmp % 1000) * ns2ms

	return time.Unix(int64(sec), int64(nano))
}

func (s *SnowflakeImpl) SetTimestamp(t time.Time) Snowflake {
	tmp := uint64(t.UnixNano())/ns2ms - EpochMillis

	if s.validate && tmp&MaxTimestamp > 0 {
		panic(ErrSnowflakeBitLoss)
	}

	s.raw &= ^timeMask
	s.raw |= (tmp << timeShift) & timeMask

	return s
}

func (s *SnowflakeImpl) InternalWorkerID() uint8 {
	return uint8((s.raw & workIdMask) >> workIdShift)
}

func (s *SnowflakeImpl) SetInternalWorkerID(id uint8) Snowflake {
	if s.validate && id&MaxWorkerID > 0 {
		panic(ErrSnowflakeBitLoss)
	}

	s.raw &= ^workIdMask
	s.raw |= (uint64(id) << workIdShift) & workIdMask

	return s
}

func (s *SnowflakeImpl) InternalProcessID() uint8 {
	return uint8((s.raw & procIdMask) >> procIdShift)
}

func (s *SnowflakeImpl) SetInternalProcessID(id uint8) Snowflake {
	if s.validate && id&MaxProcessID > 0 {
		panic(ErrSnowflakeBitLoss)
	}

	s.raw &= ^procIdMask
	s.raw |= (uint64(id) << procIdShift) & procIdMask

	return s
}

func (s *SnowflakeImpl) CounterValue() uint16 {
	return uint16(s.raw & countMask)
}

func (s *SnowflakeImpl) SetCounterValue(id uint16) Snowflake {
	if s.validate && id&MaxCounterValue > 0 {
		panic(ErrSnowflakeBitLoss)
	}

	s.raw &= ^countMask
	s.raw |= uint64(id) & countMask

	return s
}

func (s *SnowflakeImpl) String() string {
	return strconv.FormatUint(s.raw, 10)
}

func (s *SnowflakeImpl) MarshalJSON() ([]byte, error) {
	if s.validate {
		if err := s.Validate(); err != nil {
			return nil, err
		}
	}

	return []byte(s.String()), nil
}

func (s *SnowflakeImpl) UnmarshalJSON(bytes []byte) (err error) {
	var tmp string

	if err = json.Unmarshal(bytes, &tmp); err != nil {
		return
	}

	return s.UnmarshalString(tmp)
}

func (s *SnowflakeImpl) UnmarshalString(val string) (err error) {
	s.raw, err = strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}

	if s.validate {
		err = s.Validate()
	}

	return
}

func (s *SnowflakeImpl) IsValid() bool {
	return nil == s.Validate()
}

func (s *SnowflakeImpl) Validate() error {
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
