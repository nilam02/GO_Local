package components

import (
	"time"
)

// GetDateTimeByZone returns the current date and time in the specified time zone.
func GetDateTimeByZone(timeZone string) (time.Time, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}




