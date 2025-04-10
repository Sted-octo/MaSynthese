package usecases

import (
	"Octoptimist/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WorkDaysPartSplitter_Should_Return_Empty_List_When_Input_List_Is_Empty(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 0, len(taceItems))
}

func Test_WorkDaysPartSplitter_One_People_Working_250_Days_Should_Return_A_One_Item_List(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 1, len(taceItems))
	assert.Equal(t, 250, taceItems[0].TotalWorkDays)
}

func Test_WorkDaysPartSplitter_Two_Peoples_Working_250_Days_Should_Return_A_One_Item_List(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 1, len(taceItems))
	assert.Equal(t, 250, taceItems[0].TotalWorkDays)
}

func Test_WorkDaysPartSplitter_Two_Peoples_Working_125_And_250_Days_Should_Return_A_Two_Item_List_Sorted_Desc(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 2, len(taceItems))
	assert.Equal(t, 250, taceItems[0].TotalWorkDays)
	assert.Equal(t, 125, taceItems[1].TotalWorkDays)
}

func Test_WorkDaysPartSplitter_Two_Peoples_Working_250_And_125_Days_Should_Return_A_Two_Item_List_Sorted_Desc(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 2, len(taceItems))
	assert.Equal(t, 250, taceItems[0].TotalWorkDays)
	assert.Equal(t, 125, taceItems[1].TotalWorkDays)
}

func Test_WorkDaysPartSplitter_One_People_Working_250_Days_Should_Return_A_One_Item_List_With_DeltaWorkDays_250(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 1, len(taceItems))
	assert.Equal(t, 250, taceItems[0].DeltaWorkDays)
}

func Test_WorkDaysPartSplitter_Three_Peoples_Different_Working_Days_Should_Return_A_Three_Item_List_With_Correct_DeltaWorkDays(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 25})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 3, len(taceItems))
	assert.Equal(t, 125, taceItems[0].DeltaWorkDays)
	assert.Equal(t, 100, taceItems[1].DeltaWorkDays)
	assert.Equal(t, 25, taceItems[2].DeltaWorkDays)
}

func Test_WorkDaysPartSplitter_Three_Peoples_Different_Working_Days_Should_Return_A_Three_Item_List_With_Correct_NbItems(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 25})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 3, len(taceItems))
	assert.Equal(t, 1, taceItems[0].NbItems)
	assert.Equal(t, 2, taceItems[1].NbItems)
	assert.Equal(t, 3, taceItems[2].NbItems)
}

func Test_WorkDaysPartSplitter_Four_Peoples_Different_Working_250_125_125_25_Days_Should_Return_A_Three_Item_List_With_Correct_NbItems(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 25})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	assert.Equal(t, 3, len(taceItems))
	assert.Equal(t, 1, taceItems[0].NbItems)
	assert.Equal(t, 3, taceItems[1].NbItems)
	assert.Equal(t, 4, taceItems[2].NbItems)
}

func Test_WorkDaysPartSplitter_DeltaWorkDays_Accumulation_Should_Be_Equal_To_First_Item_TotalWorkDays(t *testing.T) {
	workDaysPartSplitter := new(WorkDaysPartSplitter)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 250})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 125})
	peoplesInTribe = append(peoplesInTribe, domain.PeopleInTribe{PeriodWorkDays: 25})
	taceItems := workDaysPartSplitter.Split(peoplesInTribe)

	deltaWorkDaysAccumulation := 0
	for _, taceItem := range taceItems {
		deltaWorkDaysAccumulation += taceItem.DeltaWorkDays
	}
	assert.Equal(t, taceItems[0].TotalWorkDays, deltaWorkDaysAccumulation)
}
