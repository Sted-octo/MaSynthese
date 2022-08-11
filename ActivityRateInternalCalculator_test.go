package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_No_TimeInput_ActivityRateInternal_Should_not_be_Nil(t *testing.T) {
	timeInputs = new(TimeInput)

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.NotNil(t, activityRate, "ActivityRateIntenalCalculator should return a not nil objet")
}

func Test_No_TimeInput_ActivityRateInternal_Value_Shoulbe_0(t *testing.T) {
	timeInputs = new(TimeInput)

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 0.0, activityRate.Value, "Activity Rate value should be 0 when no input time")
}

func Test_One_Billable_Day_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}

func Test_Two_Billable_Days_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 2.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 2/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Billable_And_One_NotBillable_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 1))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Billable_And_One_Absence_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementNotBillable(ACTIVITY_ID_RTT, "absence", 1))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22-1), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22-1))
}

func Test_All_Absences_ActivityRateInternal_value_shouldbe_0(t *testing.T) {
	timeInputs = new(TimeInput)
	for i := 0; i < TOTAL_WORKDAYS_FY22; i++ {
		timeInputs.Add(timeInputElementNotBillable(ACTIVITY_ID_RTT, "absence", 1))
	}

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 0.0, activityRate.Value, "Activity Rate value should be 0 when all absences")
}

func Test_One_Internal_Day_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementInternal(123, "Ma mission Interne", 1, "OctoMobile", "123456"))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Billable_One_Internal_Days_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillable(123, "Ma mission", 1, "OctoMobile", "123456"))
	timeInputs.Add(timeInputElementInternal(456, "Ma mission interne", 1, "OctoMobile", "123456789"))

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 2.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 2/%d", TOTAL_WORKDAYS_FY22))
}

func Test_One_Not_Billable_Days_Project_GeneralPurpose_ActivityRateInternal_value_shouldbe_Correct(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputElement := timeInputElementNotBillable(123, "Ma mission", 1)
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.ID = int64(GENERAL_PURPOSE_PROJECT_ID)
	timeInputs.Add(timeInputElement)

	activityRate, _ := timeInputs.ActivityRateInternalCalculator(PIVOT_DATE, TOTAL_WORKDAYS_FY22)

	assert.Equal(t, 1.0/float64(TOTAL_WORKDAYS_FY22), activityRate.Value, fmt.Sprintf("Activity Rate value should be 1/%d", TOTAL_WORKDAYS_FY22))
}
