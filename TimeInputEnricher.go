package main

import (
	"Octoptimist/tools"
	"Octoptimist/usecases"
	"strconv"
	"time"
)

func (timeInput *TimeInput) timeInputEnricher(period *Period, pivot time.Time) *TimeInput {
	if period == nil {
		return nil
	}

	if period.End.Before(pivot) {
		return timeInput
	}

	startDate := pivot.AddDate(0, 0, 1)
	if startDate.Before(period.Start) {
		startDate = period.Start
	}

	dictionnary := timeInput.toMapConverter()

	for currentDate := startDate; currentDate.Before(period.End) || tools.DatesEquals(currentDate, period.End); currentDate = currentDate.AddDate(0, 0, 1) {
		if period.BankHolidayManager != nil && period.BankHolidayManager.IsHoliday(currentDate) {
			continue
		}
		if !usecases.IsWorkDay(currentDate) {
			continue
		}
		currentDayActivityAccumulation := 0.0
		currentDateString := tools.DateToString(currentDate)

		for _, timeInputElement := range (*dictionnary)[currentDateString] {
			if decimal, err := strconv.ParseFloat(timeInputElement.TimeInDays, 64); err == nil {
				currentDayActivityAccumulation += decimal
			}
		}
		if currentDayActivityAccumulation != 1.0 {
			timeInDaysToAdd := 1.0 - currentDayActivityAccumulation
			timeInput.Add(timeInputElementPermanentAt(ACTIVITY_ID_INTERCONTRAT, "Intercontrat", timeInDaysToAdd, currentDate))
		}
	}

	return timeInput
}
