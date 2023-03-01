package main

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
)

var bankHolidays *domain.BankHolidays

func CreateBankHolydays() {
	bankHolidays = &domain.BankHolidays{Loader: dataproviders.BankHolidaysLoader}
}

func GetBankHolidaysInstance() *domain.BankHolidays {
	if bankHolidays == nil {
		CreateBankHolydays()
	}
	bankHolidays.Init()
	return bankHolidays
}
