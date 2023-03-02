package main

import (
	"Octoptimist/domain"
	"testing"
)

var synthesisLines ByAssending

func Test_Two_Empty_SynthsisLine_Less_Souldby_True(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{})

	isLess := synthesisLines.Less(0, 1)
	expected := true

	if isLess != expected {
		t.Errorf("Less return value should be true")
	}
	synthesisLines = nil
}

func Test_One_Without_Reference_One_With_Reference_Less_Souldby_True(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Reference: "aaa"})

	isLess := synthesisLines.Less(0, 1)
	expected := true

	if isLess != expected {
		t.Errorf("Less return value should be true")
	}
	synthesisLines = nil
}

func Test_One_Reference_bbb_One_Reference_aaa_Less_Souldby_False(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Reference: "bbb"})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Reference: "aaa"})

	isLess := synthesisLines.Less(0, 1)

	expected := false

	if isLess != expected {
		t.Errorf("Less return value should be false")
	}
	synthesisLines = nil
}

func Test_One_Reference_aaa_One_Reference_bbb_Less_Souldby_true(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Reference: "aaa"})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Reference: "bbb"})

	isLess := synthesisLines.Less(0, 1)
	expected := true

	if isLess != expected {
		t.Errorf("Less return value should be true")
	}
	synthesisLines = nil
}

func Test_One_Billable_One_OtherKing_Less_Souldby_true(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_BILLABLE})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_PERMANENT})

	isLess := synthesisLines.Less(0, 1)

	expected := true

	if isLess != expected {
		t.Errorf("Less return value should be true")
	}
	synthesisLines = nil
}

func Test_One_OtherKing_One_Billable_Less_Souldby_False(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_PERMANENT})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_BILLABLE})

	isLess := synthesisLines.Less(0, 1)

	expected := false

	if isLess != expected {
		t.Errorf("Less return value should be true")
	}
	synthesisLines = nil
}

func Test_Same_Kind_Different_Title_Less_Souldby_False(t *testing.T) {
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_BILLABLE, Title: "bbb"})
	synthesisLines = append(synthesisLines, domain.SynthesisLine{Kind: domain.KIND_BILLABLE, Title: "aaa"})

	isLess := synthesisLines.Less(0, 1)

	expected := false

	if isLess != expected {
		t.Errorf("Less return value should be %v", expected)
	}
	synthesisLines = nil
}
