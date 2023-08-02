package timestamp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/valyala/fastjson"
)

// ToTimestamp generates a millisecond timestamp from the time object.
func ToTimestamp(t time.Time) *Timestamp {
	return &Timestamp{
		TimeUnixMs: ToUnixMs(t),
	}
}

// ToUnixMs converts a time to unix ms.
func ToUnixMs(t time.Time) uint64 {
	return uint64(t.UnixNano() / 1000000)
}

// ToTime generates a time object from a millisecond timestamp.
func ToTime(t uint64) time.Time {
	return time.Unix(0, int64(t)*1000000).UTC()
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

// UnmarshalJSON unmarshals json.
// Supports string (unix milliseconds large value or RFC3339 timestamp), number (unix milliseconds)
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	t.Reset()
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	if v.Type() == fastjson.TypeNumber {
		ms, err := v.Uint64()
		if err != nil {
			return err
		}
		t.TimeUnixMs = ms
		return nil
	}
	if v.Type() == fastjson.TypeObject {
		if v.Exists("timeUnixMs") {
			t.TimeUnixMs = v.GetUint64("timeUnixMs")
		}
		return nil
	}
	if v.Type() == fastjson.TypeString {
		str := string(v.GetStringBytes())
		// try to parse as RFC3339
		tt, err := time.Parse(time.RFC3339, str)
		if err == nil {
			t.TimeUnixMs = ToUnixMs(tt)
			return nil
		}

		timeMs, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			t.TimeUnixMs = timeMs
			return nil
		}

		return errors.New("cannot parse timestamp string")
	}

	return nil
}

// ToRFC3339 formats the timestamp as RFC3339
func (t *Timestamp) ToRFC3339() string {
	return t.Format(time.RFC3339)
}

// Format formats using time.Format.
func (t *Timestamp) Format(formatStr string) string {
	return t.ToTime().Format(formatStr)
}

// MarshalJSON marshals to a JSON RFC3339 timestamp.
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(t.ToRFC3339())), nil
}

// _ is a type assertion
var (
	_ json.Unmarshaler = ((*Timestamp)(nil))
	_ json.Marshaler   = ((*Timestamp)(nil))
)
