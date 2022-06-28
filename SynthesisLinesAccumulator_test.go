package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Accumulate_Empty_SynthsisLines_Should_Return_Empty_List(t *testing.T) {
	synthesisLines := SynthesisLines{}

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 0, len(synthesisLines), "Length list shouldbe 0")
}

func Test_Accumulate_OneLine_SynthsisLines_len_List_ShouldBe_1(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission", TimeSum: 1, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 1, len(synthesisLines), "List len shouldbe 1")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_len_List_ShouldBe_4(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1 ", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 2 ", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 4, len(synthesisLines), "List len shouldbe 3")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_SubTotalLines_ShouldBe_1(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1 ", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 2 ", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if line.TypeLine == LineSubTotal {
			subTotalCount++
		}
	}

	assert.Equal(t, 1, subTotalCount, "Only one subTotal line expected")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_TotalLines_ShouldBe_1(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1 ", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 2 ", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()
	totalCount := 0
	for _, line := range synthesisLines {
		if line.TypeLine == LineTotal {
			totalCount++
		}
	}

	assert.Equal(t, 1, totalCount, "Only one Total line expected")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_NormalLines_ShouldBe_2(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1 ", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 2 ", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if line.TypeLine == LineNormal {
			subTotalCount++
		}
	}

	assert.Equal(t, 2, subTotalCount, "Only one subTotal line expected")
}

func Test_Accumulate_TwoLines_DifferentKind_SynthsisLines_len_List_ShouldBe_3(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "permanent", Title: "Intercontrat", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 3, len(synthesisLines), "List len shouldbe 3")
}

func Test_Accumulate_TwoLines_DifferentKind_SynthsisLines_Count_SubTotalLines_ShouldBe_0(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "billable", Title: "Mission 1", TimeSum: 1, ActivityID: 1, TypeLine: LineNormal})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: "permanent", Title: "Intercontrat", TimeSum: 1, ActivityID: 2, TypeLine: LineNormal})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if line.TypeLine == LineSubTotal {
			subTotalCount++
		}
	}

	assert.Equal(t, 0, subTotalCount, "No subTotal line expected with 2 lines of different kinds")
}
