package timestamp

import (
	"strconv"
	"time"

	jsoniter "github.com/aperturerobotics/json-iterator-lite"
	json "github.com/aperturerobotics/protobuf-go-lite/json"
)

// MarshalProtoJSON marshals the Timestamp message to JSON.
func (t *Timestamp) MarshalProtoJSON(s *json.MarshalState) {
	if t == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if t.TimeUnixMs != 0 || s.HasField("timeUnixMs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("timeUnixMs")
		s.WriteUint64(t.TimeUnixMs)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the Timestamp to JSON.
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(t.ToRFC3339())), nil
}

// UnmarshalProtoJSON unmarshals the Timestamp message from JSON.
//
// Supports string (unix milliseconds large value or RFC3339 timestamp), number (unix milliseconds)
func (t *Timestamp) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}

	nextTok := s.WhatIsNext()
	if nextTok == jsoniter.StringValue {
		str := s.ReadString()
		// try to parse as RFC3339
		tt, err := time.Parse(time.RFC3339, str)
		if err == nil {
			t.TimeUnixMs = ToUnixMs(tt)
			return
		}

		timeMs, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			t.TimeUnixMs = timeMs
			return
		}

		// otherwise set error
		s.SetError(err)
		return
	}

	if nextTok == jsoniter.NumberValue {
		t.TimeUnixMs = s.ReadUint64()
		return
	}

	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "time_unix_ms", "timeUnixMs":
			s.AddField("time_unix_ms")

			// note: this also supports string encoding!
			t.TimeUnixMs = s.ReadUint64()
		}
	})
}

// UnmarshalJSON unmarshals the Timestamp from JSON.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, t)
}

// _ is a type assertion
var (
	_ json.Marshaler   = ((*Timestamp)(nil))
	_ json.Unmarshaler = ((*Timestamp)(nil))
)
