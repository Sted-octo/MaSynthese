package domain

import "Octoptimist/tools"

func (timeInput *TimeInput) toMapConverter() *map[string]TimeInput {
	dictionnary := make(map[string]TimeInput)

	for _, currentTimeInput := range *timeInput {
		dictionnary[tools.DateToString(currentTimeInput.Day)] = append(dictionnary[tools.DateToString(currentTimeInput.Day)], currentTimeInput)
	}
	return &dictionnary
}
