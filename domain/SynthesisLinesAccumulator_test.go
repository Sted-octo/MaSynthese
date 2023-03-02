package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Accumulate_Empty_SynthsisLines_Should_Return_Empty_List(t *testing.T) {
	synthesisLines := SynthesisLines{}

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 0, len(synthesisLines), "Length list shouldbe 0")
}

func Test_Accumulate_OneLine_SynthsisLines_len_List_ShouldBe_3(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission", TimeSum: 1})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 3, len(synthesisLines), "List len shouldbe 3")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_len_List_ShouldBe_4(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1 ", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 2 ", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 4, len(synthesisLines), "List len shouldbe 3")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_SubTotalLines_ShouldBe_1(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1 ", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 2 ", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if line.IsLineSubTotal {
			subTotalCount++
		}
	}

	assert.Equal(t, 1, subTotalCount, "Only one subTotal line expected")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_TotalLines_ShouldBe_1(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1 ", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 2 ", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()
	totalCount := 0
	for _, line := range synthesisLines {
		if line.IsLineTotal {
			totalCount++
		}
	}

	assert.Equal(t, 1, totalCount, "Only one Total line expected")
}

func Test_Accumulate_TwoLines_SameKind_SynthsisLines_Count_NormalLines_ShouldBe_2(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1 ", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 2 ", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if !line.IsLineTotal && !line.IsLineSubTotal {
			subTotalCount++
		}
	}

	assert.Equal(t, 2, subTotalCount, "Only one subTotal line expected")
}

func Test_Accumulate_TwoLines_DifferentKind_SynthsisLines_len_List_ShouldBe_5(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_PERMANENT, Title: "Intercontrat", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()

	assert.Equal(t, 5, len(synthesisLines), "List len shouldbe 5")
}

func Test_Accumulate_TwoLines_DifferentKind_SynthsisLines_Count_SubTotalLines_ShouldBe_2(t *testing.T) {
	synthesisLines := SynthesisLines{}
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_BILLABLE, Title: "Mission 1", TimeSum: 1, ActivityID: 1})
	synthesisLines = append(synthesisLines, SynthesisLine{Kind: KIND_PERMANENT, Title: "Intercontrat", TimeSum: 1, ActivityID: 2})

	synthesisLines = synthesisLines.Accumulate()
	subTotalCount := 0
	for _, line := range synthesisLines {
		if line.IsLineSubTotal {
			subTotalCount++
		}
	}

	assert.Equal(t, 2, subTotalCount, "No subTotal line expected with 2 lines of different kinds")
}
