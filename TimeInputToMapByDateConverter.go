package main

func (timeInput *TimeInput) toMapConverter() *map[string]TimeInput {
	dictionnary := make(map[string]TimeInput)

	for _, currentTimeInput := range *timeInput {
		dictionnary[currentTimeInput.Day] = append(dictionnary[currentTimeInput.Day], currentTimeInput)
	}
	return &dictionnary
}
