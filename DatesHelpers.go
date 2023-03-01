package main

import (
	"time"
)

func datesEquals(startPeriode time.Time, endPeriode time.Time) bool {
	sy, sm, sd := startPeriode.Date()
	ey, em, ed := endPeriode.Date()
	if sy == ey && sm == em && sd == ed {
		return true
	}
	return false
}

func dateSimple(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
}

func dateToString(date time.Time) string {
	return date.Format("2006-01-02")
}
