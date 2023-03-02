package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParameterDate_Before_StartDate_Should_Return_false(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 20)
	period := domain.NewPeriod(start, end, &bankHolidays)

	dateToTest := tools.DateSimple(2022, time.March, 20)

	isParameterDateInPeriod := period.IncludeDate(dateToTest)

	assert.False(t, isParameterDateInPeriod, "IncludeDate should return false when date is before period start date")
}

func Test_ParameterDate_Equal_StartDate_Should_Return_true(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 20)
	period := domain.NewPeriod(start, end, &bankHolidays)

	dateToTest := start

	isParameterDateInPeriod := period.IncludeDate(dateToTest)

	assert.True(t, isParameterDateInPeriod, "IncludeDate should return true when date is equal to period start date")
}

func Test_ParameterDate_Between_StartDate_And_EndDate_Should_Return_true(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 20)
	period := domain.NewPeriod(start, end, &bankHolidays)

	dateToTest := tools.DateSimple(2022, time.March, 19)

	isParameterDateInPeriod := period.IncludeDate(dateToTest)

	assert.True(t, isParameterDateInPeriod, "IncludeDate should return true when date is between period start and end dates")
}

func Test_ParameterDate_Equal_EndDate_Should_Return_false(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 20)
	period := domain.NewPeriod(start, end, &bankHolidays)

	dateToTest := end
	isParameterDateInPeriod := period.IncludeDate(dateToTest)

	assert.False(t, isParameterDateInPeriod, "IncludeDate should return false when date is equal to period end dates")
}
