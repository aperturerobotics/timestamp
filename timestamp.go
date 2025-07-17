package timestamp

import "time"

// ToTimestamp generates a millisecond timestamp from the time object.
func ToTimestamp(t time.Time) *Timestamp {
	return &Timestamp{
		TimeUnixMs: ToUnixMs(t),
	}
}

// ToUnixMs converts a time to unix ms.
func ToUnixMs(t time.Time) uint64 {
	return uint64(t.UnixNano() / 1000000) //nolint:gosec
}

// ToTime generates a time object from a millisecond timestamp.
func ToTime(t uint64) time.Time {
	return time.Unix(0, int64(t)*1000000).UTC() //nolint:gosec
}

// ToTime converts the Timestamp to a time.Time
func (t *Timestamp) ToTime() time.Time {
	if t.GetTimeUnixMs() == 0 {
		return time.Time{}
	}
	return ToTime(t.GetTimeUnixMs())
}

// Now returns a timestamp for now.
func Now() *Timestamp {
	return ToTimestamp(time.Now().UTC())
}

// Clone copies the timestamp.
func (t *Timestamp) Clone() *Timestamp {
	if t == nil {
		return nil
	}

	return &Timestamp{TimeUnixMs: t.TimeUnixMs}
}

// Equals checks if the timestamp equals another timestamp.
func (t *Timestamp) Equals(ot *Timestamp) bool {
	return t.GetTimeUnixMs() == ot.GetTimeUnixMs()
}

// Validate checks the timestamp.
func (t *Timestamp) Validate(allowEmpty bool) error {
	if !allowEmpty && t.GetEmpty() {
		return ErrEmptyTimestamp
	}
	return nil
}

// GetEmpty checks if the timestamp is empty.
func (t *Timestamp) GetEmpty() bool {
	return t.GetTimeUnixMs() == 0
}

// ToRFC3339 formats the timestamp as RFC3339
func (t *Timestamp) ToRFC3339() string {
	return t.Format(time.RFC3339)
}

// Format formats using time.Format.
func (t *Timestamp) Format(formatStr string) string {
	return t.ToTime().Format(formatStr)
}
