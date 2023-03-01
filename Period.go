package main

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"errors"
	"time"
)

type Period struct {
	Start              time.Time
	End                time.Time
	BankHolidayManager *domain.BankHolidays
}

func NewPeriod(startDate, endDate time.Time, bankHolydays *domain.BankHolidays) *Period {
	return &Period{
		Start:              startDate,
		End:                endDate,
		BankHolidayManager: bankHolydays,
	}
}

func (p *Period) IsValid() bool {
	return !p.End.Before(p.Start)
}

func (p *Period) TotalWorkDaysGetter() (int, error) {

	if !p.IsValid() {
		return 0, errors.New("end date should be after start date")
	}

	totalWorkDays := 0

	for currentDate := p.Start; currentDate.Before(p.End) || DatesEquals(currentDate, p.End); currentDate = currentDate.AddDate(0, 0, 1) {
		if p.BankHolidayManager != nil && p.BankHolidayManager.IsHoliday(currentDate) {
			continue
		}
		if usecases.IsWorkDay(currentDate) {
			totalWorkDays++
		}
	}

	return totalWorkDays, nil
}

func (p *Period) IncludeDate(dateToTest time.Time) bool {
	return p.Start == dateToTest || (p.Start.Before(dateToTest) && dateToTest.Before(p.End))
}
