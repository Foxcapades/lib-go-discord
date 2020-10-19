package utils

import "time"

func MillisToUTCTime(millis uint64) time.Time {
	a := millis / uint64(1000)
	b := (millis % uint64(1000)) * 1000000

	return time.Unix(int64(a), int64(b)).UTC()
}
