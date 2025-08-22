package usecases

import "Octoptimist/domain"

func FiscalPeriodPreviousGetter(periodFiscal *domain.Period) *domain.Period {
	dayBeforeCurrentFiscalPeriod := periodFiscal.Start.AddDate(0, 0, -1)
	return FiscalPeriodGetter(dayBeforeCurrentFiscalPeriod, periodFiscal.BankHolidayManager)
}
