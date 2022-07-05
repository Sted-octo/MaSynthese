package main

import (
	"strconv"
	"time"
)

func (timeInput *TimeInput) ActivityRateCalculator(pivot time.Time, totalWorkDays int) (*ActivityRate, error) {
	activityRate := ActivityRate{}

	billableTimeTotal := 0.0

	workDaywWithoutDayBreak := totalWorkDays

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]
		if currentTimeInput.Activity.Kind == "billable" {
			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				billableTimeTotal += decimal
			}
		}
		if currentTimeInput.Activity.ID == 2140318361 {
			day, _ := time.Parse("2006-01-02", currentTimeInput.Day)
			if day.After(pivot) {
				if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
					billableTimeTotal += decimal
				}
			}
		}
		if currentTimeInput.Activity.Title == "absence" {
			workDaywWithoutDayBreak -= 1
		}
	}
	if workDaywWithoutDayBreak > 0 {
		activityRate.Value = billableTimeTotal / float64(workDaywWithoutDayBreak)
	}

	return &activityRate, nil
}
