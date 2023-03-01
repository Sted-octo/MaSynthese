package infrastructure

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
)

var bankHolidays *domain.BankHolidays

func createBankHolydays() {
	bankHolidays = &domain.BankHolidays{Loader: dataproviders.BankHolidaysLoader}
}

func BankHolidaysSingletonGetter() *domain.BankHolidays {
	if bankHolidays == nil {
		createBankHolydays()
	}
	bankHolidays.Init()
	return bankHolidays
}
