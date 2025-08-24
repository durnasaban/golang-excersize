package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	layout := "1/2/2006 15:04:05"

	parsedDate, _ := time.Parse(layout, date)

	return parsedDate

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	layout := "January 2, 2006 15:04:05"

	parsedDate, _ := time.Parse(layout, date)

	return !parsedDate.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {

	layout := "Monday, January 2, 2006 15:04:05"

	parsedDate, _ := time.Parse(layout, date)

	return parsedDate.Hour() < 18 && parsedDate.Hour() >= 12 
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {

	parsedDate := Schedule(date)

	return fmt.Sprintf(
		"You have an appointment on %v, %v %v, %v, at %v:%v.",
		parsedDate.Weekday(),
		parsedDate.Month(),
		parsedDate.Day(),
		parsedDate.Year(),
		parsedDate.Hour(),
		parsedDate.Minute(),
	)

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {

	layout := "1/2/2006 15:04:05"

	anniversary := fmt.Sprintf("9/15/%v 00:00:00", time.Now().Year())

	anniversaryDate, _ := time.Parse(layout, anniversary)

	return anniversaryDate

}