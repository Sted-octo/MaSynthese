package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var TOTAL_WORKDAYS_FY22 int = 20
var PIVOT_DATE time.Time = time.Date(2022, time.July, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))

func Test_No_TimeInput_ActivityRate_Shoul_not_be_Nil(t *testing.T) {
	timeInputs = new(TimeInput)

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.NotNil(t, activityRate, "ActivityRateCalculator should return a not nil objet")
}

func Test_No_TimeInput_ActivityRate_Value_Shoulbe_0(t *testing.T) {
	timeInputs = new(TimeInput)

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 0.0, activityRate.Value, "Activity Rate value should be 0 when no input time")
}

func Test_One_Billable_Day_ActivityRate_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}

func Test_Two_Billable_Days_ActivityRate_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 2.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 2/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Billable_And_One_NotBillable_ActivityRate_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 1))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Billable_And_One_Absence_ActivityRate_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementNotBillable(ACTIVITY_ID_RTT, "absence", 1))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22-1), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22-1))
}

func Test_All_Absences_ActivityRate_value_shouldbe_0(t *testing.T) {
	timeInputs = new(TimeInput)
	for i := 0; i < TOTAL_WORKDAYS_FY22; i++ {
		timeInputs.Add(timeInputElementNotBillable(ACTIVITY_ID_RTT, "absence", 1))
	}

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 0.0, activityRate.Value, "Activity Rate value should be 0 when all absences")
}

func Test_One_Intercontrat_Before_Pivot_ActivityRate_value_shouldbe_0(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillableAt(ACTIVITY_ID_INTERCONTRAT, "intercontrat", 1, time.Date(2022, time.June, 1, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 0.0, activityRate.Value, "Activity Rate value should be 0 when on intercontrat before pivot date")
}

func Test_One_Intercontrat_After_Pivot_ActivityRate_value_shouldbe_0(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillableAt(ACTIVITY_ID_INTERCONTRAT, "Intercontrat", 1, time.Date(2022, time.July, 10, 0, 0, 0, 0, TimeZoneGetter("Europe/Paris"))))

	activityRate, _ := timeInputs.ActivityRateCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}