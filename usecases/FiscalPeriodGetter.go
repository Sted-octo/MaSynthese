package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"time"
)

func FiscalPeriodGetter(day time.Time, bankHolydays *domain.BankHolidays) *domain.Period {
	if day.IsZero() {
		return nil
	}
	var periodFiscal *domain.Period
	fiscalYear25ChangeDate := tools.DateSimple(2024, time.September, 1)
	fiscalYear25EndDate := tools.DateSimple(2026, time.February, 28)

	if day.Before(fiscalYear25ChangeDate) {
		startFiscalYear := tools.DateSimple(day.Year()-1, time.September, 1)
		endFiscalYear := tools.DateSimple(day.Year(), time.August, 31)
		if day.After(endFiscalYear) {
			startFiscalYear = tools.DateSimple(day.Year(), time.September, 1)
			endFiscalYear = tools.DateSimple(day.Year()+1, time.August, 31)
		}
		periodFiscal = domain.NewPeriod(startFiscalYear, endFiscalYear, bankHolydays)
		periodFiscal.FiscalYearFormatYY = endFiscalYear.Format("06")
		return periodFiscal
	}

	if day.After(fiscalYear25EndDate) {
		startFiscalYear := tools.DateSimple(day.Year()-1, time.March, 1)
		// To deal with the last day of babruary in leap year
		endFiscalYear := tools.DateSimple(day.Year(), time.March, 1).AddDate(0, 0, -1)
		if day.After(endFiscalYear) {
			startFiscalYear = tools.DateSimple(day.Year(), time.March, 1)
			endFiscalYear = tools.DateSimple(day.Year()+1, time.March, 1).AddDate(0, 0, -1)
		}
		periodFiscal = domain.NewPeriod(startFiscalYear, endFiscalYear, bankHolydays)
		periodFiscal.FiscalYearFormatYY = startFiscalYear.Format("06")
		return periodFiscal
	}

	startFiscalYear25 := tools.DateSimple(2024, time.September, 1)
	endFiscalYear25 := tools.DateSimple(2026, time.February, 28)
	periodFiscal = domain.NewPeriod(startFiscalYear25, endFiscalYear25, bankHolydays)
	periodFiscal.FiscalYearFormatYY = "25"
	return periodFiscal
}
