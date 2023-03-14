package domain

import (
	"time"
)

var GENERAL_PURPOSE_PROJECT_ID int = 2146913740

func (timeInput *TimeInput) ActivityRateInternalCalculator(pivot time.Time, totalWorkDays int) (*ActivityRate, error) {
	activityRate := ActivityRate{}

	billableTimeTotal := 0.0

	workDaywWithoutDayBreak := float64(totalWorkDays)

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]
		if currentTimeInput.Activity.Kind == KIND_BILLABLE ||
			currentTimeInput.Activity.GlobalPurpose {
			billableTimeTotal += currentTimeInput.TimeInDays
		}

		if currentTimeInput.Activity.IsDayBreak() {
			workDaywWithoutDayBreak -= currentTimeInput.TimeInDays
		}
	}
	if workDaywWithoutDayBreak > 0 {
		activityRate.Value = billableTimeTotal / workDaywWithoutDayBreak
	}

	return &activityRate, nil
}
