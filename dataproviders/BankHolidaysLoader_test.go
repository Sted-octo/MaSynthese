package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"Octoptimist/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Date_In_SourceMap_ShouldBe_Holiday(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}

	state := bankHolidays.IsHoliday(time.Date(2022, time.May, 26, 0, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris")))

	assert.True(t, state, "26 may 2022 should be a holiday")
}

func Test_Date_Not_In_SourceMap_Should_NotBe_Holiday(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}

	state := bankHolidays.IsHoliday(time.Date(2022, time.May, 25, 0, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris")))

	assert.False(t, state, "25 may 2022 should be a holiday")
}
