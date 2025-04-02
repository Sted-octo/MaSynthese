package domain

import (
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_No_TimeInput_Clone_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)

	timeInputCloned := timeInputs.Clone()

	if timeInputCloned != nil {
		t.Errorf("timeInputCloned should be nil")
	}
}

func Test_One_TimeInput_Permanent_Clone_count_Shoulbe_One(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	timeInputCloned := timeInputs.Clone()

	expected := 1

	if len(*timeInputCloned) != expected {
		t.Errorf("timeInputCloned count should be %d but was %d", expected, len(synthesisLines))
	}
}

func Test_One_TimeInput_Name_Modification_After_Clone_Should_Have_Different_Value(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	timeInputCloned := timeInputs.Clone()

	(*timeInputs)[0].Activity.Title = "Modified"

	assert.NotEqual(t, (*timeInputs)[0].Activity.Title, (*timeInputCloned)[0].Activity.Title, "The title of the cloned TimeInput should not be modified")
}

func Test_One_TimeInput_TimeInDays_Modification_After_Clone_Shoulkd_Have_Different_Value(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	timeInputCloned := timeInputs.Clone()

	(*timeInputs)[0].TimeInDays = 0.25

	assert.NotEqual(t, (*timeInputs)[0].TimeInDays, (*timeInputCloned)[0].TimeInDays, "The timeInDays of the cloned TimeInput should not be modified")
}

func Test_One_TimeInput_Day_Modification_After_Clone_Should_Have_Different_Value(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillableAt(123, "intercontrat", 1, tools.DateSimple(2022, time.June, 10)))

	timeInputCloned := timeInputs.Clone()

	(*timeInputs)[0].Day = tools.DateSimple(2022, time.June, 11)

	assert.NotEqual(t, (*timeInputs)[0].Day, (*timeInputCloned)[0].Day, "The Day of the cloned TimeInput should not be modified")
}

func Test_One_TimeInput_ProjectReference_Modification_After_Clone_Should_Have_Different_Value(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementBillable(123, "Intercontrat", 0.5, "OctoMobile", "123456"))

	timeInputCloned := timeInputs.Clone()

	(*timeInputs)[0].Activity.Project.Reference = "Octo"

	assert.NotEqual(t, (*timeInputs)[0].Activity.Project.Reference, (*timeInputCloned)[0].Activity.Project.Reference, "The Project Reference of the cloned TimeInput should not be modified")
}

func Test_One_TimeInput_Customer_Modification_After_Clone_Should_Have_Different_Value(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementBillable(123, "Intercontrat", 0.5, "OctoMobile", "123456"))

	timeInputCloned := timeInputs.Clone()

	(*timeInputs)[0].Activity.Project.Customer.Name = "Octo"

	assert.NotEqual(t, (*timeInputs)[0].Activity.Project.Customer.Name, (*timeInputCloned)[0].Activity.Project.Customer.Name, "The Customer Name of the cloned TimeInput should not be modified")
}
