package usecases

import "Octoptimist/domain"

func FiscalPeriodNextGetter(periodFiscal *domain.Period) *domain.Period {

	dayAfterCurrentFiscalPeriod := periodFiscal.End.AddDate(0, 0, 1)
	return FiscalPeriodGetter(dayAfterCurrentFiscalPeriod, periodFiscal.BankHolidayManager)
}
