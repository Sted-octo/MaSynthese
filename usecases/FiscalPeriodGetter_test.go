package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PeriodFiscalGetter_20220710_Should_Return_20210901_To_20220831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.July, 10)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2021, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2021-09-01 for 2022-07-10")
	assert.Equal(t, tools.DateSimple(2022, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2022-08-31 for 2022-07-10")
}

func Test_PeriodFiscalGetter_20210901_Should_Return_20210901_To_20220831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2021, time.September, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2021, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2021-09-01 for 2021-09-01")
	assert.Equal(t, tools.DateSimple(2022, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2022-08-31 for 2021-09-01")
}

func Test_PeriodFiscalGetter_20220831_Should_Return_20210901_To_20220831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.August, 31)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2021, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2021-09-01 for 2022-08-31")
	assert.Equal(t, tools.DateSimple(2022, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2022-08-31 for 2022-08-31")
}

func Test_PeriodFiscalGetter_20230101_Should_Return_20220901_To_20230831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2023, time.January, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2022, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2022-09-01 for 2023-01-01")
	assert.Equal(t, tools.DateSimple(2023, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2023-08-31 for 2023-01-01")
}

func Test_PeriodFiscalGetter_20221231_Should_Return_20220901_To_20230831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.December, 31)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2022, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2022-09-01 for 2022-12-31")
	assert.Equal(t, tools.DateSimple(2023, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2023-08-31 for 2022-12-31")
}
