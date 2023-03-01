package main

import (
	"Octoptimist/domain"
	"time"
)

func FiscalPeriodGetter(day time.Time, bankHolydays *domain.BankHolidays) *Period {
	if day.IsZero() {
		return nil
	}
	startFiscalYear := DateSimple(day.Year()-1, time.September, 1)
	endFiscalYear := DateSimple(day.Year(), time.August, 31)
	if day.After(endFiscalYear) {
		startFiscalYear = DateSimple(day.Year(), time.September, 1)
		endFiscalYear = DateSimple(day.Year()+1, time.August, 31)
	}
	periodFiscal := NewPeriod(startFiscalYear, endFiscalYear, bankHolydays)
	return periodFiscal
}
