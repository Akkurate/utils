package datetime

import "time"

// TimeToMicroseconds TimeToMicroseconds
func TimeToMicroseconds(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Microsecond)
}

// MicrosecondsToTime to Time
func MicrosecondsToTime(tm int64) time.Time {
	return time.Unix(0, tm*1000)
}

// ParseTime ParseTime
func ParseTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05Z", t)
	return tm
}

// NanosecondsToTime NanosecondsToTime
func NanosecondsToTime(tm int64) time.Time {
	return time.Unix(0, tm*int64(time.Nanosecond))
}

// FirstOfMonth returns a TZ timestamp of first day of given month
func FirstOfMonth(t time.Time) time.Time {
	return ParseTime(t.Format("2006-01") + "-01T00:00:00Z")
}
