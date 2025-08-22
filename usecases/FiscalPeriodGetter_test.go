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
	assert.Equal(t, "22", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '22' for 2022-07-10")
}

func Test_PeriodFiscalGetter_20210901_Should_Return_20210901_To_20220831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2021, time.September, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2021, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2021-09-01 for 2021-09-01")
	assert.Equal(t, tools.DateSimple(2022, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2022-08-31 for 2021-09-01")
	assert.Equal(t, "22", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '22' for 2021-09-01")
}

func Test_PeriodFiscalGetter_20220831_Should_Return_20210901_To_20220831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.August, 31)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2021, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2021-09-01 for 2022-08-31")
	assert.Equal(t, tools.DateSimple(2022, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2022-08-31 for 2022-08-31")
	assert.Equal(t, "22", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '22' for 2022-08-31")
}

func Test_PeriodFiscalGetter_20230101_Should_Return_20220901_To_20230831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2023, time.January, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2022, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2022-09-01 for 2023-01-01")
	assert.Equal(t, tools.DateSimple(2023, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2023-08-31 for 2023-01-01")
	assert.Equal(t, "23", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '23' for 2023-01-01")
}

func Test_PeriodFiscalGetter_20221231_Should_Return_20220901_To_20230831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.December, 31)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2022, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2022-09-01 for 2022-12-31")
	assert.Equal(t, tools.DateSimple(2023, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2023-08-31 for 2022-12-31")
	assert.Equal(t, "23", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '23' for 2022-12-31")
}

func Test_PeriodFiscalGetter_20240901_Should_Return_20240901_To_20260228(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2024, time.September, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2024, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2024-09-01 for 2024-09-01")
	assert.Equal(t, tools.DateSimple(2026, time.February, 28), fiscalPeriod.End, "End fiscal period should be 2026-02-28 for 2024-09-01")
	assert.Equal(t, "25", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '25' for 2024-09-01")
}

func Test_PeriodFiscalGetter_20260301_Should_Return_20260301_To_20270228(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2026, time.March, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2026, time.March, 1), fiscalPeriod.Start, "Start fiscal period should be 2026-03-01 for 2026-03-01")
	assert.Equal(t, tools.DateSimple(2027, time.February, 28), fiscalPeriod.End, "End fiscal period should be 2027-02-28 for 2026-03-01")
	assert.Equal(t, "26", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '26' for 2026-03-01")
}

func Test_PeriodFiscalGetter_20270101_Should_Return_20260301_To_20270228(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2027, time.January, 1)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2026, time.March, 1), fiscalPeriod.Start, "Start fiscal period should be 2026-03-01 for 2027-01-01")
	assert.Equal(t, tools.DateSimple(2027, time.February, 28), fiscalPeriod.End, "End fiscal period should be 2027-02-28 for 2027-01-01")
	assert.Equal(t, "26", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '26' for 2027-01-01")
}

func Test_PeriodFiscalGetter_20270404_Should_Return_20260301_To_20270228(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2027, time.April, 4)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)

	assert.Equal(t, tools.DateSimple(2027, time.March, 1), fiscalPeriod.Start, "Start fiscal period should be 2027-03-01 for 2027-04-04")
	assert.Equal(t, tools.DateSimple(2028, time.February, 29), fiscalPeriod.End, "End fiscal period should be 2028-02-29 for 2027-04-04")
	assert.Equal(t, "27", fiscalPeriod.FiscalYearFormatYY, "Fiscal year format should be '27' for 2027-04-04")
}
