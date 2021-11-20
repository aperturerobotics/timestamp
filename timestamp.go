package timestamp

import (
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
	return time.Unix(0, int64(t)*1000000).UTC()
}

// ToTime converts the Timestamp to a time.Time
func (t *Timestamp) ToTime() time.Time {
	return ToTime(t.GetTimeUnixMs())
}

// Now returns a timestamp for now.
func Now() Timestamp {
	return ToTimestamp(time.Now().UTC())
}

// Validate checks the timestamp.
func (t *Timestamp) Validate(allowEmpty bool) error {
	if !allowEmpty && t.GetTimeUnixMs() == 0 {
		return ErrEmptyTimestamp
	}
	return nil
}
