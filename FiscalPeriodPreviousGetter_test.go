package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Previous_20220710_Should_Return_20200901_To_20210831(t *testing.T) {
	bankHolidays := BankHolidays{Loader: mockBankHolidayLoader}

	day := time.Date(2022, time.July, 10, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)
	fiscalPeriod.Previous()

	assert.Equal(t, time.Date(2020, time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")), fiscalPeriod.Start, "Start fiscal period should be 2020-09-01 for 2022-07-10")
	assert.Equal(t, time.Date(2021, time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")), fiscalPeriod.End, "End fiscal period should be 2021-08-31 for 2022-07-10")
}
