/* Helper functions for handling time.Time objects.
 */
package datetime

import "time"

// Converts time.Time object to microseconds.
func TimeToMicroseconds(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Microsecond)
}

// Converts microseconds to time.Time object.
func MicrosecondsToTime(tm int64) time.Time {
	return time.Unix(0, tm*1000)
}

// Converts string to time.Time object. Time format:
//  "2006-01-02T15:04:05Z"
func ParseTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05Z", t)
	return tm
}

// Converts nanoseconds to time.Time object.
func NanosecondsToTime(tm int64) time.Time {
	return time.Unix(0, tm*int64(time.Nanosecond))
}

// Returns a time.Time timestamp of first day of given month. Time format:
//  "2006-01-02T15:04:05Z"
func FirstOfMonth(t time.Time) time.Time {
	return ParseTime(t.Format("2006-01") + "-01T00:00:00Z")
}
