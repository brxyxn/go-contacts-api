package constants

import "time"

// DateTime formats
const (
	// LongDateTimeFormat 2006-01-02T15:04:05Z07:00
	LongDateTimeFormat = time.RFC3339
	// ShortDateTimeFormat 02 Jan 06 15:04 -0700
	ShortDateTimeFormat = time.RFC822Z
	// DateOnlyFormat 2006-01-02
	DateOnlyFormat = time.DateOnly
	// TimeOnlyFormat 15:04:05
	TimeOnlyFormat = time.TimeOnly
)
