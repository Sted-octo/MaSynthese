package main

import "Octoptimist/domain"

var bankHolidays *domain.BankHolidays

func CreateBankHolydays() {
	bankHolidays = &domain.BankHolidays{Loader: bankHolidayLoader}
}

func GetBankHolidaysInstance() *domain.BankHolidays {
	if bankHolidays == nil {
		CreateBankHolydays()
	}
	bankHolidays.Init()
	return bankHolidays
}
