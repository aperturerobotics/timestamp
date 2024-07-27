package timestamp

import (
	"encoding/json"
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
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{"RFC3339", `"2009-11-10T23:00:00Z"`, time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"UnixMsString", `"1257894000000"`, time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
		{"UnixMsNumber", `1257894000000`, time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing JSON unmarshaling for input: %s", tt.input)

			var ts Timestamp
			err := json.Unmarshal([]byte(tt.input), &ts)
			if err != nil {
				t.Fatalf("Failed to unmarshal JSON: %v", err)
			}

			gotTime := ts.ToTime()
			if !gotTime.Equal(tt.expected) {
				t.Errorf("Expected time %v, but got %v", tt.expected, gotTime)
			}

			t.Logf("Successfully unmarshaled to time: %v", gotTime)

			// Test marshaling
			marshaledJSON, err := json.Marshal(&ts)
			if err != nil {
				t.Fatalf("Failed to marshal Timestamp: %v", err)
			}

			t.Logf("Marshaled JSON: %s", string(marshaledJSON))

			// Verify that unmarshaling the marshaled JSON gives the same result
			var ts2 Timestamp
			err = json.Unmarshal(marshaledJSON, &ts2)
			if err != nil {
				t.Fatalf("Failed to unmarshal marshaled JSON: %v", err)
			}

			if !ts.ToTime().Equal(ts2.ToTime()) {
				t.Errorf("Marshaling and unmarshaling resulted in different times. Original: %v, After: %v", ts.ToTime(), ts2.ToTime())
			}

			t.Logf("Successfully round-tripped through JSON")
		})
	}

	// Test error case
	t.Run("InvalidInput", func(t *testing.T) {
		invalidInput := `"not a timestamp"`
		var ts Timestamp
		err := json.Unmarshal([]byte(invalidInput), &ts)
		if err == nil {
			t.Error("Expected an error for invalid input, but got none")
		} else {
			t.Logf("Correctly received error for invalid input: %v", err)
		}
	})
}
