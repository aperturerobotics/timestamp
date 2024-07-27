package timestamp

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

func TestNewTimestamp(t *testing.T) {
	ts := time.Now()
	times := ToTimestamp(ts)
	ta := times.ToTime()
	resolution := time.Millisecond * 1
	diff := ts.Sub(ta)
	if diff < 0 {
		diff *= -1
	}
	if diff > resolution {
		t.Logf("%s != %s", ts.String(), ta.String())
		t.Fail()
	}
	t.Logf("%d", times.GetTimeUnixMs())
}

func TestTimestampJSON(t *testing.T) {
	exampleVal := "2009-11-10T23:00:00Z"
	ts, _ := time.Parse(time.RFC3339, exampleVal)
	times := ToTimestamp(ts)
	dat, err := json.Marshal(&times)
	if err != nil {
		t.Fatal(err.Error())
	}
	datStr := string(dat)
	if datStr != strconv.Quote(exampleVal) {
		t.FailNow()
	}

	// test parsing the unix ms wrapped in a string
	ms := times.GetTimeUnixMs()
	msStr := strconv.FormatUint(ms, 10)
	datStr = strconv.Quote(msStr)
	parseTimestamp := &Timestamp{}
	if err := json.Unmarshal([]byte(datStr), parseTimestamp); err != nil {
		t.Fatal(err.Error())
	}
	if !parseTimestamp.EqualVT(times) {
		t.FailNow()
	}

	// test parsing the unix ms as a number
	parseTimestamp = &Timestamp{}
	if err := json.Unmarshal([]byte(msStr), parseTimestamp); err != nil {
		t.Fatal(err.Error())
	}
	if !parseTimestamp.EqualVT(times) {
		t.FailNow()
	}
}
