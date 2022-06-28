package main

import (
	"errors"
	"time"
)

func TotalWorkDaysGetter(startPeriode time.Time, endPeriode time.Time) (int, error) {

	if endPeriode.Before(startPeriode) {
		return 0, errors.New("end date should be after start date")
	}

	if datesEquals(startPeriode, endPeriode) && isWorkDay(startPeriode) {
		return 1, nil
	}

	totalWorkDays := 0

	for currentDate := startPeriode; currentDate.Before(endPeriode); currentDate = currentDate.AddDate(0, 0, 1) {
		if isWorkDay(currentDate) {
			totalWorkDays++
		}
	}

	return totalWorkDays, nil
}

func isWorkDay(date time.Time) bool {
	day := date.Weekday()
	return day == time.Monday || day == time.Tuesday || day == time.Wednesday || day == time.Thursday || day == time.Friday
}

func datesEquals(startPeriode time.Time, endPeriode time.Time) bool {
	sy, sm, sd := startPeriode.Date()
	ey, em, ed := endPeriode.Date()
	if sy == ey && sm == em && sd == ed {
		return true
	}
	return false
}
