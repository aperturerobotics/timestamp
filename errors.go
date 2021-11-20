package timestamp

import "errors"

var (
	// ErrEmptyTimestamp is returned if the timestamp was empty.
	ErrEmptyTimestamp = errors.New("timestamp is empty")
)
