package main

import "strconv"

func (timeInput *TimeInput) timeInputAggregator() []SynthesisLine {
	nbTimes := len(*timeInput)
	if nbTimes == 0 {
		return nil
	}

	var synthesisLines []SynthesisLine
	activityMap := make(map[int64]*SynthesisLine)

	for indx := range *timeInput {
		var currentTimeInput *TimeInputElement = &(*timeInput)[indx]

		if existingLine, exist := activityMap[currentTimeInput.Activity.ID]; exist {

			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				existingLine.TimeSum += decimal
			}
			continue
		}

		newLine := &SynthesisLine{
			ActivityID: currentTimeInput.Activity.ID,
			Title:      currentTimeInput.Activity.Title,
			Kind:       currentTimeInput.Activity.Kind,
			TypeLine:   LineNormal,
		}

		if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
			newLine.TimeSum = decimal
		}
		if currentTimeInput.Activity.Project != nil {
			newLine.Reference = currentTimeInput.Activity.Project.Reference
			newLine.ProjectName = currentTimeInput.Activity.Project.Name

			if currentTimeInput.Activity.Project.Customer != nil {
				newLine.CustomerName = currentTimeInput.Activity.Project.Customer.Name
			}
		}

		activityMap[currentTimeInput.Activity.ID] = newLine
	}

	for key := range activityMap {
		synthesisLines = append(synthesisLines, *activityMap[key])
	}

	return synthesisLines
}
