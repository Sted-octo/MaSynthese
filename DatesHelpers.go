package main

import (
	"Octoptimist/tools"
	"time"
)

func DatesEquals(startPeriode time.Time, endPeriode time.Time) bool {
	sy, sm, sd := startPeriode.Date()
	ey, em, ed := endPeriode.Date()
	if sy == ey && sm == em && sd == ed {
		return true
	}
	return false
}

func DateSimple(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))
}

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}
