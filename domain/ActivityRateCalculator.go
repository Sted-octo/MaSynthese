package domain

import (
	"time"
)

func (timeInput *TimeInput) ActivityRateCalculator(pivot time.Time, totalWorkDays int) (*ActivityRate, error) {
	activityRate := ActivityRate{}

	billableTimeTotal := 0.0

	workDaywWithoutDayBreak := float64(totalWorkDays)

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]
		if currentTimeInput.Activity.Kind == KIND_BILLABLE {
			billableTimeTotal += currentTimeInput.TimeInDays
		}
		if currentTimeInput.Activity.ID == ACTIVITY_ID_INTERCONTRAT {
			if currentTimeInput.Day.After(pivot) {
				billableTimeTotal += currentTimeInput.TimeInDays
			}
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
