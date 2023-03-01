package main

import (
	"Octoptimist/domain"
	"time"
)

func FiscalPeriodGetter(day time.Time, bankHolydays *domain.BankHolidays) *Period {
	if day.IsZero() {
		return nil
	}
	startFiscalYear := dateSimple(day.Year()-1, time.September, 1)
	endFiscalYear := dateSimple(day.Year(), time.August, 31)
	if day.After(endFiscalYear) {
		startFiscalYear = dateSimple(day.Year(), time.September, 1)
		endFiscalYear = dateSimple(day.Year()+1, time.August, 31)
	}
	periodFiscal := NewPeriod(startFiscalYear, endFiscalYear, bankHolydays)
	return periodFiscal
}
