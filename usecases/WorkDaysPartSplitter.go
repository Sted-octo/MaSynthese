package usecases

import (
	"Octoptimist/domain"
	"sort"
)

type WorkDaysPartSplitter struct{}

func (w *WorkDaysPartSplitter) Split(peoplesInTribe []domain.PeopleInTribe) []domain.TaceItem {
	var result []domain.TaceItem
	uniqueMap := make(map[int]int)

	for _, people := range peoplesInTribe {
		uniqueMap[people.PeriodWorkDays] = uniqueMap[people.PeriodWorkDays] + 1
	}

	for value := range uniqueMap {
		result = append(result, domain.TaceItem{TotalWorkDays: value, DeltaWorkDays: value, NbItems: uniqueMap[value]})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TotalWorkDays > result[j].TotalWorkDays
	})

	if len(result) <= 1 {
		return result
	}

	precItemsCount := 0
	for i := 0; i < len(result)-1; i++ {
		result[i].DeltaWorkDays = result[i].TotalWorkDays - result[i+1].TotalWorkDays
		result[i].NbItems = result[i].NbItems + precItemsCount
		precItemsCount = result[i].NbItems
	}
	result[len(result)-1].NbItems = result[len(result)-1].NbItems + precItemsCount

	return result
}
