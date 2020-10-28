package utils

import "time"

const (
	Millis2Seconds = 1_000
	Millis2Nano    = 1_000_000
)

func MillisToUTCTime(millis uint64) time.Time {
	a := millis / Millis2Seconds
	b := (millis % Millis2Seconds) * Millis2Nano

	return time.Unix(int64(a), int64(b)).UTC()
}

func UTCTimeToMillis(t *time.Time) *uint64 {
	if t == nil {
		return nil
	}

	tmp := uint64(t.UTC().UnixNano()) / Millis2Nano

	return &tmp
}
