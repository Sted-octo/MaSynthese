package main

import (
	"strconv"
	"time"
)

var GENERAL_PURPOSE_PROJECT_ID int = 2146913740

func (timeInput *TimeInput) ActivityRateInternalCalculator(pivot time.Time, totalWorkDays int) (*ActivityRate, error) {
	activityRate := ActivityRate{}

	billableTimeTotal := 0.0

	workDaywWithoutDayBreak := totalWorkDays

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]
		if currentTimeInput.Activity.Kind == KIND_BILLABLE ||
			currentTimeInput.Activity.Kind == KIND_INTERNAL ||
			(currentTimeInput.Activity.Project != nil && currentTimeInput.Activity.Project.ID == int64(GENERAL_PURPOSE_PROJECT_ID)) {
			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				billableTimeTotal += decimal
			}
		}

		if currentTimeInput.Activity.IsDayBreak() {
			workDaywWithoutDayBreak -= 1
		}
	}
	if workDaywWithoutDayBreak > 0 {
		activityRate.Value = billableTimeTotal / float64(workDaywWithoutDayBreak)
	}

	return &activityRate, nil
}
