package timestamp

import (
	"testing"
	"time"
)

func TestNewTimestamp(t *testing.T) {
	ts := time.Now()
	times := ToTimestamp(ts)
	ta := times.ToTime()
	var resolution = time.Millisecond * 1
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
