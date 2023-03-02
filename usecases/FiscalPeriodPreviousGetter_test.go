package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Previous_20220710_Should_Return_20200901_To_20210831(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}

	day := tools.DateSimple(2022, time.July, 10)

	fiscalPeriod := FiscalPeriodGetter(day, &bankHolidays)
	fiscalPeriod.Previous()

	assert.Equal(t, tools.DateSimple(2020, time.September, 1), fiscalPeriod.Start, "Start fiscal period should be 2020-09-01 for 2022-07-10")
	assert.Equal(t, tools.DateSimple(2021, time.August, 31), fiscalPeriod.End, "End fiscal period should be 2021-08-31 for 2022-07-10")
}
