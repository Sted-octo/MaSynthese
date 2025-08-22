package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"time"
)

func Period250WorkDaysGetter() domain.Period {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2024, time.January, 1)
	end := tools.DateSimple(2024, time.December, 13)
	period := domain.NewPeriod(start, end, &bankHolidays)

	return *period
}
