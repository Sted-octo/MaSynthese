package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_EndDate_Before_StartDate_Should_Return_Error(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 17)
	period := domain.NewPeriod(start, end, &bankHolidays)

	_, err := period.TotalWorkDaysGetter()

	assert.Errorf(t, err, "end date should be after start date")
}

func Test_StartDate_Monday_Equal_EndDate_Should_Return_1(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.June, 27)
	end := tools.DateSimple(2022, time.June, 27)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 1, totalDays, "Total Days should be 1 when start date equal end date on monday")
}

func Test_StartDate_Tuesday_Equal_EndDate_Should_Return_1(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.June, 28)
	end := tools.DateSimple(2022, time.June, 28)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 1, totalDays, "Total Days should be 1 when start date equal end date on tuesday")
}

func Test_StartDate_Wednesday_Equal_EndDate_Should_Return_1(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.June, 29)
	end := tools.DateSimple(2022, time.June, 29)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 1, totalDays, "Total Days should be 1 when start date equal end date on wednesday")
}

func Test_StartDate_Thursday_Equal_EndDate_Should_Return_1(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.June, 30)
	end := tools.DateSimple(2022, time.June, 30)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 1, totalDays, "Total Days should be 1 when start date equal end date on thursday")
}

func Test_StartDate_Friday_Equal_EndDate_Should_Return_1(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.July, 1)
	end := tools.DateSimple(2022, time.July, 1)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 1, totalDays, "Total Days should be 1 when start date equal end date on friday")
}

func Test_StartDate_Saturday_Equal_EndDate_Should_Return_0(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.July, 2)
	end := tools.DateSimple(2022, time.July, 2)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 0, totalDays, "Total Days should be 0 when start date equal end date on saturday")
}

func Test_StartDate_Sunday_Equal_EndDate_Should_Return_0(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.July, 3)
	end := tools.DateSimple(2022, time.July, 3)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 0, totalDays, "Total Days should be 0 when start date equal end date on sunday")
}

func Test_Two_Days_Except_Weekend_Should_Return_2(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.July, 4)
	end := tools.DateSimple(2022, time.July, 5)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 2, totalDays, "Total Days should be 2 for a two days period except weekend")
}

func Test_Two_Days_Full_Weekend_Should_Return_0(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.July, 2)
	end := time.Date(2022, time.July, 3, 23, 59, 59, 0, tools.TimeZoneGetter("Europe/Paris"))
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 0, totalDays, "Total Days should be 0 for a two days period full weekend")
}

func Test_One_Break_Day_Should_Return_0(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.May, 26)
	end := time.Date(2022, time.May, 26, 23, 59, 59, 0, tools.TimeZoneGetter("Europe/Paris"))
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 0, totalDays, "Total Days should be 0 for one holiday")
}
func Test_One_Week_With_OneHoliday_outside_weekend_Should_Return_4(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.May, 23)
	end := time.Date(2022, time.May, 29, 23, 59, 59, 0, tools.TimeZoneGetter("Europe/Paris"))
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 4, totalDays, "Total Days should be 0 for one holiday")
}

func Test_One_Week_outside_weekend_Should_Return_5(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2022, time.June, 20)
	end := time.Date(2022, time.June, 24, 23, 59, 59, 0, tools.TimeZoneGetter("Europe/Paris"))
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 5, totalDays, "Total Days should be 5 for one week without holiday outside weekend")
}

func Test_250_Workdays_Without_Holiday_Should_Return_TotalWorkDays_250(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	start := tools.DateSimple(2024, time.January, 1)
	end := tools.DateSimple(2024, time.December, 13)
	period := domain.NewPeriod(start, end, &bankHolidays)

	totalDays, _ := period.TotalWorkDaysGetter()

	assert.Equal(t, 250, totalDays, "Total Days should be 250 for 250 work days without holiday")
}
