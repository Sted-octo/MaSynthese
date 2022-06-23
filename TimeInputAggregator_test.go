package main

import (
	"testing"
)

var timeInputs *TimeInput

func Test_NoT_TimeInput_SynthesisLine_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines != nil {
		t.Errorf("SynthesisLines should be nil")
	}
}

func Test_One_TimeInput_Permanent_SynthesisLine_count_Shoulbe_One(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := 1

	if len(synthesisLines) != expected {
		t.Errorf("SynthesisLines count should be %d but was %d", expected, len(synthesisLines))
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_ActivityId_Shoulbe_123(t *testing.T) {
	timeInputs = new(TimeInput)

	var expected int64 = 123
	timeInputs.Add(timeInputElementNotBillable(expected, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines[0].ActivityID != expected {
		t.Errorf("First SynthesisLine activityId shouldBe %d but was %d", expected, synthesisLines[0].ActivityID)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Title_Shoulbe_Intercontrat(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := "Intercontrat"
	timeInputs.Add(timeInputElementNotBillable(123, expected, 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines[0].Title != expected {
		t.Errorf("First SynthesisLine Title shouldBe %s but was %s", expected, synthesisLines[0].Title)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_TimeSum_Shoulbe_dot5(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := 0.5
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", expected))

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines[0].TimeSum != expected {
		t.Errorf("First SynthesisLine TimeSum shouldBe %f but was %f", expected, synthesisLines[0].TimeSum)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Kind_Shoulbe_Permanant(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := "permanent"

	if synthesisLines[0].Kind != expected {
		t.Errorf("First SynthesisLine kind shouldBe %s but was %s", expected, synthesisLines[0].Kind)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_CustomerName_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := ""

	if synthesisLines[0].CustomerName != expected {
		t.Errorf("First SynthesisLine Customer Name shouldBe %s but was %s", expected, synthesisLines[0].CustomerName)
	}
}

func Test_One_TimeInput_Permanent_First_SynthesisLine_Reference_Shoulbe_Empty(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := ""

	if synthesisLines[0].Reference != expected {
		t.Errorf("First SynthesisLine Reference Name shouldBe %s but was %s", expected, synthesisLines[0].Reference)
	}
}

func Test_Two_TimeInput_Permanent_SynthesisLine_count_Shoulbe_One(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	if len(synthesisLines) != 1 {
		t.Errorf("SynthesisLines count should be 1 but was %d", len(synthesisLines))
	}
}

func Test_Two_TimeInput_Permanent_First_SynthesisLine_TimeSum_Shoulbe_1(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := 1.0

	if synthesisLines[0].TimeSum != expected {
		t.Errorf("First SynthesisLine TimeSum shouldBe %f but was %f", expected, synthesisLines[0].TimeSum)
	}
}

func Test_Two_Differents_TimeInput_Permanent_SynthesisLine_count_Shoulbe_Two(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementNotBillable(123, "Intercontrat", 0.5))
	timeInputs.Add(timeInputElementNotBillable(456, "Shadowing", 0.5))

	synthesisLines := timeInputs.timeInputAggregator()

	expected := 2

	if len(synthesisLines) != expected {
		t.Errorf("SynthesisLines count should be %d but was %d", expected, len(synthesisLines))
	}
}

func Test_One_TimeInput_Billable_First_SynthesisLine_Reference_Shoulbe_123456(t *testing.T) {
	timeInputs = new(TimeInput)

	expected := "123456"
	timeInputs.Add(timeInputElementBillable(123, "Intercontrat", 0.5, "OctoMobile", expected))

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines[0].Reference != expected {
		t.Errorf("First SynthesisLine Reference Name shouldBe %s but was %s", expected, synthesisLines[0].Reference)
	}
}

func Test_One_TimeInput_Billable_First_SynthesisLine_CustomerName_Shoulbe_OctoMobile(t *testing.T) {
	timeInputs = new(TimeInput)
	expected := "OctoMobile"
	timeInputs.Add(timeInputElementBillable(123, "Intercontrat", 0.5, expected, "123456"))

	synthesisLines := timeInputs.timeInputAggregator()

	if synthesisLines[0].CustomerName != expected {
		t.Errorf("First SynthesisLine Customer Name shouldBe %s but was %s", expected, synthesisLines[0].CustomerName)
	}
}