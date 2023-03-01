package main

import (
	"Octoptimist/tools"
	"testing"
	"time"
)

func Test_Empty_TimeInput_Should_return_an_Empty_Map(t *testing.T) {
	timeInputs = new(TimeInput)

	dictionnary := timeInputs.toMapConverter()

	if dictionnary == nil || len(*dictionnary) != 0 {
		t.Errorf("Empty timeInput List should return empty map")
	}
}

func Test_OneDay_TimeInput_Should_return_an_Map_withOne_Key_of_One_TimeInput(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 1.0, "OCTO", "123456"))

	dictionnary := timeInputs.toMapConverter()

	if len(*dictionnary) != 1 {
		t.Errorf("One day timeInput List should return map with 1 key, it return %d", len(*dictionnary))
	}

	if len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]) != 1 {
		t.Errorf("One day timeInput List should return map with 1 element, it return %d", len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]))
	}
}

func Test_OneDay_TwoTimeInputs_Should_return_an_Map_with_One_Key_of_two_TimeInput(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 0.5, "OCTO", "123456"))
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 0.5, "OCTO", "123456"))

	dictionnary := timeInputs.toMapConverter()

	if len(*dictionnary) != 1 {
		t.Errorf("One day timeInput List should return map with 1 key, it return %d", len(*dictionnary))
	}

	if len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]) != 2 {
		t.Errorf("One day of 2 timeInputs should return map with 2 elements, it return %d", len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]))
	}
}

func Test_TwoDays_OneTimeInput_Should_return_an_Map_with_Two_Keys_of_one_TimeInput(t *testing.T) {
	timeInputs = new(TimeInput)
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 1.0, "OCTO", "123456"))
	timeInputs.Add(timeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 13), 123, "Audit", 1.0, "OCTO", "123456"))

	dictionnary := timeInputs.toMapConverter()

	if len(*dictionnary) != 2 {
		t.Errorf("Two days timeInput List should return map with 2 keys, it return %d", len(*dictionnary))
	}

	if len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]) != 1 {
		t.Errorf("Two days  of one timeInput List should return map with one element, it return %d", len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 12))]))
	}

	if len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 13))]) != 1 {
		t.Errorf("Two days  of one timeInput List should return map with one element, it return %d", len((*dictionnary)[tools.DateToString(tools.DateSimple(2023, time.January, 13))]))
	}
}
