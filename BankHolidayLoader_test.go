package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Date_In_SourceMap_ShouldBe_Holiday(t *testing.T) {
	bankHolidays := BankHolidays{Loader: mockBankHolidayLoader}

	state := bankHolidays.IsHoliday(time.Date(2022, time.May, 26, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")))

	assert.True(t, state, "26 may 2022 should be a holiday")
}

func Test_Date_Not_In_SourceMap_Should_NotBe_Holiday(t *testing.T) {
	bankHolidays := BankHolidays{Loader: mockBankHolidayLoader}

	state := bankHolidays.IsHoliday(time.Date(2022, time.May, 25, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")))

	assert.False(t, state, "25 may 2022 should be a holiday")
}

func mockBankHolidayLoader() (map[int][]BankHoliday, error) {
	bankHolidayMap := make(map[int][]BankHoliday)
	bankHolidayMap[2021] = append(bankHolidayMap[2021], BankHoliday{DayDate: time.Date(2021, time.May, 13, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")),
		Year: 2021, Zone: "Métropole", Name: "Ascension"})
	bankHolidayMap[2022] = append(bankHolidayMap[2022], BankHoliday{DayDate: time.Date(2022, time.May, 26, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")),
		Year: 2022, Zone: "Métropole", Name: "Ascension"})
	bankHolidayMap[2023] = append(bankHolidayMap[2023], BankHoliday{DayDate: time.Date(2023, time.May, 18, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")),
		Year: 2023, Zone: "Métropole", Name: "Ascension"})
	return bankHolidayMap, nil
}
