package main

import "time"

func FiscalPeriodGetter() *Period {
	startFiscalYear := time.Date(time.Now().Year()-1, time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	endFiscalYear := time.Date(time.Now().Year(), time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	if time.Now().After(endFiscalYear) {
		startFiscalYear = time.Date(time.Now().Year(), time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
		endFiscalYear = time.Date(time.Now().Year()+1, time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))
	}
	periodFiscal := NewPeriod(startFiscalYear, endFiscalYear, GetBankHolidayInstance())
	return periodFiscal
}
