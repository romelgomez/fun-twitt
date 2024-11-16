package helper

import "time"

func DateTimeToStr(dt *time.Time) *string {
	if dt == nil {
		return nil
	}
	// UTC time zone format
	formatted := dt.UTC().Format("2006-01-02T15:04:05-07:00")
	return &formatted
}
