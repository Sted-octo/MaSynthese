package main

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
)

var bankHolidays *domain.BankHolidays

func CreateBankHolydays() {
	bankHolidays = &domain.BankHolidays{Loader: usecases.BankHolidayLoader}
}

func GetBankHolidaysInstance() *domain.BankHolidays {
	if bankHolidays == nil {
		CreateBankHolydays()
	}
	bankHolidays.Init()
	return bankHolidays
}
