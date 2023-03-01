package main

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"Octoptimist/usecases"
	"testing"
	"time"
)

func Test_NilPeriod_Should_return_nil_and_state_false(t *testing.T) {
	timeInputs = new(TimeInput)
	pivotDate := time.Date(2023, time.January, 31, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(nil, pivotDate)

	if timeInput != nil {
		t.Errorf("timeInputEnricher with nil input pariod should return nil")
	}
}

func Test_Period_before_pivotDate_Should_not_enrich_timeInput_list(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 1.0, "OCTO", "123456"))
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 8)
	end := tools.DateSimple(2023, time.January, 10)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 31, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 2 {
		t.Errorf("timeInputEnricher should not enrich timeInputs when period is before pivot date")
	}
}

func Test_Period_After_pivotDate_Should_enrich_when_cumulation_par_day_lower_than_1(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 10, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 2 {
		t.Errorf("timeInputEnricher should add a timeInput when period is after pivot date and cumumalation per day is lower than 1")
	}
}

func Test_NewTimeInput_ActivityId_Should_Be_Intercontrat_Activity_Id(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 10, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if (*timeInput)[1].Activity.ID != ACTIVITY_ID_INTERCONTRAT {
		t.Errorf("New time input should have Activity ID = %d", ACTIVITY_ID_INTERCONTRAT)
	}
}

func Test_NewTimeInput_ActivityTitle_Should_Be_Intercontrat(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 10, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if (*timeInput)[1].Activity.Title != "Intercontrat" {
		t.Errorf("New time input should have Activity Title = Intercontrat")
	}
}

func Test_NewTimeInput_TimeInDays_Should_Be_ZeroPointFive_ForExistingTimeInputTimeInDAyx_ZeroPointFive(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 10, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if (*timeInput)[1].TimeInDays != "0.5" {
		t.Errorf("New time input should have TimeInDays = 0.5 when existing time input have timeInDays = 0.5, but has %s", (*timeInput)[1].TimeInDays)
	}
}

func Test_Period_After_pivotDate_Should_Create_newOne_not_Exist_for_a_specific_day(t *testing.T) {
	timeInputs = new(TimeInput)
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 10, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 1 {
		t.Errorf("timeInputEnricher should add a timeInput when period is after pivot date and no one exist for a specific day")
	}
}

func Test_Period_include_pivotDate_Should_only_enrich_timeInput_After_DatePivot(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 11), 123, "Audit", 0.5, "OCTO", "123456"))
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 11)
	end := tools.DateSimple(2023, time.January, 13)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 12, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 3 {
		t.Errorf("timeInputEnricher should only enrich timeInput when day after date pivot, %d", len(*timeInput))
	}
}

func Test_Period_include_pivotDate_Should_only_enrich_timeInput_in_WorkingDay(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 0.5, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.January, 13)
	end := tools.DateSimple(2023, time.January, 15)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 13, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 1 {
		t.Errorf("timeInputEnricher should only enrich timeInput when day is working day, %d", len(*timeInput))
	}
}

func Test_Period_include_pivotDate_Should_not_enrich_timeInput_in_HoliDay(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.May, 16), 123, "Audit", 1.0, "OCTO", "123456"))
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.May, 17), 123, "Audit", 1.0, "OCTO", "123456"))
	bankHolidays := domain.BankHolidays{Loader: usecases.MockBankHolidaysLoader}
	start := tools.DateSimple(2023, time.May, 16)
	end := tools.DateSimple(2023, time.May, 19)
	period := NewPeriod(start, end, &bankHolidays)
	pivotDate := time.Date(2023, time.January, 16, 10, 0, 0, 0, tools.TimeZoneGetter("Europe/Paris"))

	timeInput := timeInputs.timeInputEnricher(period, pivotDate)

	if len(*timeInput) != 3 {
		t.Errorf("timeInputEnricher should not enrich timeInput when day is holiday, %d", len(*timeInput))
	}
}
