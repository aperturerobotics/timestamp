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
	ts = ts.Round(resolution)
	ta = ta.Round(resolution)
	if !ts.Equal(ta) {
		t.Logf("%s != %s", ts.String(), ta.String())
		t.Fail()
	}
	t.Logf("%d", times.GetTimeUnixMs())
}
