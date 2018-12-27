package timestamp

import (
	"errors"
	"time"
)

// ToTimestamp generates a millisecond timestamp from the time object.
func ToTimestamp(t time.Time) Timestamp {
	return Timestamp{
		TimeUnixMs: uint64(t.UnixNano() / 1000000),
	}
}

// ToTime generates a time object from a millisecond timestamp.
func ToTime(t uint64) time.Time {
	return time.Unix(0, int64(t)*1000000)
}

// ToTime converts the Timestamp to a time.Time
func (t *Timestamp) ToTime() time.Time {
	return ToTime(t.GetTimeUnixMs())
}

// Now returns a timestamp for now.
func Now() Timestamp {
	return ToTimestamp(time.Now())
}

// Validate checks the timestamp.
func (t *Timestamp) Validate() error {
	if t == nil || t.TimeUnixMs == 0 {
		return errors.New("timestamp is empty")
	}
	return nil
}
