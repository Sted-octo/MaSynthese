package domain

import (
	"Octoptimist/tools"
	"strconv"
	"time"
)

func (timeInput *TimeInput) TimeInputAggregator(pivot time.Time) []SynthesisLine {
	nbTimes := len(*timeInput)
	if nbTimes == 0 {
		return nil
	}

	var synthesisLines []SynthesisLine
	activityMap := make(map[int64]*SynthesisLine)

	for _, currentTimeInput := range *timeInput {
		day, _ := time.Parse("2006-01-02", currentTimeInput.Day)

		if existingLine, exist := activityMap[currentTimeInput.Activity.ID]; exist {
			if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
				existingLine.TimeSum += decimal
				if day.Before(pivot) || tools.DatesEquals(day, pivot) {
					existingLine.TimeSumDone += decimal
				} else {
					existingLine.TimeSumTodo += decimal
				}
			}
			continue
		}

		newLine := &SynthesisLine{
			ActivityID: currentTimeInput.Activity.ID,
			Title:      currentTimeInput.Activity.Title,
			Kind:       currentTimeInput.Activity.Kind,
		}
		if currentTimeInput.Activity.IsDayBreak() {
			newLine.Kind = KIND_ABSENCE
		}

		if decimal, err := strconv.ParseFloat(currentTimeInput.TimeInDays, 64); err == nil {
			newLine.TimeSum = decimal
			if day.Before(pivot) || tools.DatesEquals(day, pivot) {
				newLine.TimeSumDone = decimal
			} else {
				newLine.TimeSumTodo = decimal
			}
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
