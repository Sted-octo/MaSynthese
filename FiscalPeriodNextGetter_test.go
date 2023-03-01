package main

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Previous_20220710_Should_Return_20220901_To_20230831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidayLoader}

	day := time.Date(2022, time.July, 10, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)
	fiscalPeriod.Next()

	assert.Equal(t, time.Date(2022, time.September, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")), fiscalPeriod.Start, "Start fiscal period should be 2022-09-01 for 2022-07-10")
	assert.Equal(t, time.Date(2023, time.August, 31, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris")), fiscalPeriod.End, "End fiscal period should be 2023-08-31 for 2022-07-10")
}
