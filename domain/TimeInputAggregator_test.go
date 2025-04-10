package domain

import (
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NoT_TimeInput_SynthesisLine_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines != nil {
		t.Errorf("SynthesisLines should be nil")
	}
}

func Test_One_TimeInput_Permanent_SynthesisLine_count_Shoulbe_One(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := 1

	if len(synthesisLines) != expected {
		t.Errorf("SynthesisLines count should be %d but was %d", expected, len(synthesisLines))
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_ActivityId_Shoulbe_123(t *testing.T) {
	timeInputs = new(TimeInput)

	var expected int64 = 123
	timeInputs.Add(TimeInputElementNotBillable(expected, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines[0].ActivityID != expected {
		t.Errorf("First SynthesisLine activityId shouldBe %d but was %d", expected, synthesisLines[0].ActivityID)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Title_Shoulbe_Intercontrat(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := "Intercontrat"
	timeInputs.Add(TimeInputElementNotBillable(123, expected, 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines[0].Title != expected {
		t.Errorf("First SynthesisLine Title shouldBe %s but was %s", expected, synthesisLines[0].Title)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_TimeSum_Shoulbe_dot5(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := 0.5
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", expected))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines[0].TimeSum != expected {
		t.Errorf("First SynthesisLine TimeSum shouldBe %f but was %f", expected, synthesisLines[0].TimeSum)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Kind_Shoulbe_Permanant(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementPermanent(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := KIND_PERMANENT

	if synthesisLines[0].Kind != expected {
		t.Errorf("First SynthesisLine kind shouldBe %s but was %s", expected, synthesisLines[0].Kind)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_CustomerName_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := ""

	if synthesisLines[0].CustomerName != expected {
		t.Errorf("First SynthesisLine Customer Name shouldBe %s but was %s", expected, synthesisLines[0].CustomerName)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Reference_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := ""

	if synthesisLines[0].Reference != expected {
		t.Errorf("First SynthesisLine Reference Name shouldBe %s but was %s", expected, synthesisLines[0].Reference)
	}
}

func Test_Two_TimeInput_Permanent_SynthesisLine_count_Shoulbe_One(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if len(synthesisLines) != 1 {
		t.Errorf("SynthesisLines count should be 1 but was %d", len(synthesisLines))
	}
}

func Test_Two_TimeInput_Permanent_First_SynthesisLine_TimeSum_Shoulbe_1(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := 1.0

	if synthesisLines[0].TimeSum != expected {
		t.Errorf("First SynthesisLine TimeSum shouldBe %f but was %f", expected, synthesisLines[0].TimeSum)
	}
}

func Test_Two_Differents_TimeInput_Permanent_SynthesisLine_count_Shoulbe_Two(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(TimeInputElementNotBillable(456, "Shadowing", 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := 2

	if len(synthesisLines) != expected {
		t.Errorf("SynthesisLines count should be %d but was %d", expected, len(synthesisLines))
	}
}

func Test_One_TimeInput_Billable_First_SynthesisLine_Reference_Shoulbe_123456(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := "123456"
	timeInputs.Add(TimeInputElementBillable(123, "Intercontrat", 0.5, "OctoMobile", expected))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines[0].Reference != expected {
		t.Errorf("First SynthesisLine Reference Name shouldBe %s but was %s", expected, synthesisLines[0].Reference)
	}
}

func Test_One_TimeInput_Billable_First_SynthesisLine_CustomerName_Shoulbe_OctoMobile(t *testing.T) {
	timeInputs = new(TimeInput)
	expected := "OctoMobile"
	timeInputs.Add(TimeInputElementBillable(123, "Intercontrat", 0.5, expected, "123456"))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	if synthesisLines[0].CustomerName != expected {
		t.Errorf("First SynthesisLine Customer Name shouldBe %s but was %s", expected, synthesisLines[0].CustomerName)
	}
}

func Test_One_Permanent_RTT_Absence_Shoulbe_Kind_Absence(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementAbsence(ACTIVITY_ID_RTT, 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := KIND_ABSENCE
	if synthesisLines[0].Kind != expected {
		t.Errorf("First SynthesisLine permanent title absence should be kind %s but was %s", expected, synthesisLines[0].Kind)
	}
}

func Test_One_Permanent_CongesPaye_Absence_Shoulbe_Kind_Absence(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementAbsence(ACTIVITY_ID_CONGES_PAYE, 0.5))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	expected := KIND_ABSENCE
	if synthesisLines[0].Kind != expected {
		t.Errorf("First SynthesisLine permanent title absence should be kind %s but was %s", expected, synthesisLines[0].Kind)
	}
}

func Test_One_Permanent_Before_Pivot_TimeDone_Shoulbe_1(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillableAt(123, "intercontrat", 1, tools.DateSimple(2022, time.June, 10)))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.Equal(t, 1.0, synthesisLines[0].TimeSumDone, "First SynthesisLine time sum done should be 1")
	assert.Equal(t, 0.0, synthesisLines[0].TimeSumTodo, "First SynthesisLine time sum todo should be 0")
}

func Test_One_Permanent_After_Pivot_TimeDone_Shoulbe_1(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillableAt(123, "intercontrat", 1, tools.DateSimple(2022, time.July, 10)))

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.Equal(t, 0.0, synthesisLines[0].TimeSumDone, "First SynthesisLine time sum done should be 0")
	assert.Equal(t, 1.0, synthesisLines[0].TimeSumTodo, "First SynthesisLine time sum todo should be 1")
}

func Test_One_GlobalPurpose_TimeInput_should_survive_aggregation(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputElement := TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 10))
	timeInputElement.Activity.GlobalPurpose = true
	timeInputs.Add(timeInputElement)

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.True(t, synthesisLines[0].IsGlobalPurpose, "Global purpose boolean from Time Input shoul be kept in synthesis line")
}

func Test_Two_GlobalPurpose_TimeInputs_should_survive_aggregation(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 9)))
	timeInputElement := TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 10))
	timeInputElement.Activity.GlobalPurpose = true
	timeInputs.Add(timeInputElement)

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.True(t, synthesisLines[0].IsGlobalPurpose, "Global purpose boolean from Time Input shoul be kept in synthesis line")
}

func Test_One_Discount_TimeInput_should_survive_aggregation(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputElement := TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 10))
	timeInputElement.Activity.Discount = true
	timeInputs.Add(timeInputElement)

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.True(t, synthesisLines[0].IsDiscount, "Discount boolean from Time Input shoul be kept in synthesis line")
}

func Test_Two_Discount_TimeInputs_should_survive_aggregation(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 9)))
	timeInputElement := TimeInputElementNotBillableAt(123, "Mecenat", 1, tools.DateSimple(2022, time.July, 10))
	timeInputElement.Activity.Discount = true
	timeInputs.Add(timeInputElement)

	synthesisLines := timeInputs.TimeInputAggregator(PIVOT_DATE)

	assert.True(t, synthesisLines[0].IsDiscount, "Discount boolean from Time Input shoul be kept in synthesis line")
}
