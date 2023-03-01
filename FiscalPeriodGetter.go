package main

import (
	"Octoptimist/domain"
	"time"
)

func FiscalPeriodGetter(day time.Time, bankHolydays *domain.BankHolidays) *Period {
	if day.IsZero() {
		return nil
	}
	startFiscalYear := time.Date(day.Year()-1, time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	endFiscalYear := time.Date(day.Year(), time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	if day.After(endFiscalYear) {
		startFiscalYear = time.Date(day.Year(), time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
		endFiscalYear = time.Date(day.Year()+1, time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	}
	periodFiscal := NewPeriod(startFiscalYear, endFiscalYear, bankHolydays)
	return periodFiscal
}
