package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/foxcapades/lib-go-discord/v0/pkg/discord/lib"
)

var (
	ErrEmptySnowflake      = errors.New("empty snowflake value")
	ErrNoSnowflakeTime     = errors.New("snowflake missing timestamp value")
	ErrNoSnowflakeWorkerID = errors.New("snowflake missing internal worker id")
	ErrNoSnowflakeProcID   = errors.New("snowflake missing internal process id")
	ErrNoSnowflakeCounter  = errors.New("snowflake missing counter value")
	ErrSnowflakeBitLoss    = errors.New("value is greater than the max allowed value for this field")
)

const (
	// Discord epoch timestamp in milliseconds.
	EpochMillis     uint64 = 1420070400000

	// Max valid value for a raw snowflake timestamp.
	MaxTimestamp    uint64 = 0x0000_FFFF_FFFF_FFC0

	// Max valid value for a snowflake worker id.
	MaxWorkerID     uint8  = 0x3E

	// Max valid value for a snowflake process id.
	MaxProcessID    uint8  = 0x1F

	// Max valid value for a snowflake counter value.
	MaxCounterValue uint16 = 0x0FFF
)

type Snowflake interface {
	json.Marshaler
	json.Unmarshaler
	fmt.Stringer
	lib.Validatable

	// RawValue returns the raw uint64 value backing this Snowflake.
	RawValue() uint64

	// SetRawValue overwrites this Snowflakes value with the given uint64 value.
	SetRawValue(raw uint64) Snowflake

	// Timestamp returns the timestamp for when this Snowflake was issued with
	// millisecond precision.
	Timestamp() time.Time

	// SetTimestamp overwrites this Snowflake's creation timestamp with the given
	// time.
	//
	// Note: Snowflake timestamps use millisecond precision.  The input value will
	// be truncated to milliseconds if needed.
	SetTimestamp(time.Time) Snowflake

	// InternalWorkerID returns the Discord worker ID for the node that generated
	// this Snowflake.
	InternalWorkerID() uint8

	// SetInternalWorkerID overwrites this Snowflake's internal worker ID with the
	// given value.
	SetInternalWorkerID(uint8) Snowflake

	// InternalProcessID returns the Discord process ID that generated this
	// Snowflake.
	InternalProcessID() uint8

	// SetInternalProcessID overwrites this Snowflake's internal process ID with
	// the given value.
	SetInternalProcessID(uint8) Snowflake

	// CounterValue returns the counter value of this Snowflake.
	//
	// The counter value is an incrementing value assigned to each ID generated
	// by a specific process.
	CounterValue() uint16

	// SetCounterValue overwrites this Snowflake's counter value with the given
	// input value.
	SetCounterValue(uint16) Snowflake

	// UnmarshalString attempts to parse the given string as a stringified
	// Snowflake value.
	UnmarshalString(string) error
}
