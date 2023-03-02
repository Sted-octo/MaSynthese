package main

import (
	"Octoptimist/domain"
	"strconv"
	"time"
)

var ACTIVITY_ID_INTERCONTRAT int64 = 2140318361

func (timeInput *TimeInput) ActivityRateCalculator(pivot time.Time, totalWorkDays int) (*ActivityRate, error) {
	activityRate := ActivityRate{}

	billableTimeTotal := 0.0

	workDaywWithoutDayBreak := float64(totalWorkDays)

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]
		if currentTimeInput.Activity.Kind == domain.KIND_BILLABLE {
			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				billableTimeTotal += decimal
			}
		}
		if currentTimeInput.Activity.ID == ACTIVITY_ID_INTERCONTRAT {
			day, _ := time.Parse("2006-01-02", currentTimeInput.Day)
			if day.After(pivot) {
				if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
					billableTimeTotal += decimal
				}
			}
		}
		if currentTimeInput.Activity.IsDayBreak() {
			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				workDaywWithoutDayBreak -= decimal
			}
		}
	}
	if workDaywWithoutDayBreak > 0 {
		activityRate.Value = billableTimeTotal / workDaywWithoutDayBreak
	}

	return &activityRate, nil
}
