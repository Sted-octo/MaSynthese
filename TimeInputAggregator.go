package main

import (
	"strconv"
	"time"
)

var ACTIVITY_ID_RTT int64 = 2140298843
var ACTIVITY_ID_CONGES_PAYE int64 = 2140309429
var ACTIVITY_ID_SICK_DAY int64 = 2140312911
var ACTIVITY_ID_PART_TIME_BREAK int64 = 2140316822
var ACTIVITY_ID_PARENTAL_BREAK int64 = 3000050819
var ACTIVITY_ID_NO_SALARY_BREAK int64 = 3000030459
var ACTIVITY_ID_MEDICAL_CARE int64 = 3000030462
var ACTIVITY_ID_FAMILY_DAY int64 = 3000030457
var ACTIVITY_ID_PARENT_DAY int64 = 3000030458
var ACTIVITY_ID_MEDICAL_PART_TIME_BREAK int64 = 3000050818
var ACTIVITY_ID_AUTORIZED_BREAK int64 = 3000050820
var ACTIVITY_ID_NO_EXCUSE_BREAK int64 = 3000050821
var ACTIVITY_ID_PERSONAL_CARE int64 = 3000065641

func (timeInput *TimeInput) timeInputAggregator(pivot time.Time) []SynthesisLine {
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
				if day.Before(pivot) || datesEquals(day, pivot) {
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
			if day.Before(pivot) || datesEquals(day, pivot) {
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
