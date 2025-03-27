package domain

import (
	"Octoptimist/tools"
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

		if existingLine, exist := activityMap[currentTimeInput.Activity.ID]; exist {
			existingLine.TimeSum += currentTimeInput.TimeInDays
			if currentTimeInput.Day.Before(pivot) || tools.DatesEquals(currentTimeInput.Day, pivot) {
				existingLine.TimeSumDone += currentTimeInput.TimeInDays
			} else {
				existingLine.TimeSumTodo += currentTimeInput.TimeInDays
			}
			existingLine.IsGlobalPurpose = existingLine.IsGlobalPurpose || currentTimeInput.Activity.GlobalPurpose
			existingLine.IsDiscount = existingLine.IsDiscount || currentTimeInput.Activity.Discount
			continue
		}

		newLine := &SynthesisLine{
			ActivityID:      currentTimeInput.Activity.ID,
			Title:           currentTimeInput.Activity.Title,
			Kind:            currentTimeInput.Activity.Kind,
			IsGlobalPurpose: currentTimeInput.Activity.GlobalPurpose,
			IsDiscount:      currentTimeInput.Activity.Discount,
		}
		if currentTimeInput.Activity.IsDayBreak() {
			newLine.Kind = KIND_ABSENCE
		}

		newLine.TimeSum = currentTimeInput.TimeInDays
		if currentTimeInput.Day.Before(pivot) || tools.DatesEquals(currentTimeInput.Day, pivot) {
			newLine.TimeSumDone = currentTimeInput.TimeInDays
		} else {
			newLine.TimeSumTodo = currentTimeInput.TimeInDays
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
