package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	return parsedDate.Hour() >= 12.0 && parsedDate.Hour() < 18.0
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	layout = "Monday, January 2, 2006, at 15:04."
	msg := "You have an appointment on " + parsedDate.Format(layout)
	return msg
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
