package domain

import (
	"Octoptimist/tools"
	"fmt"
	"time"
)

func TimeInputElementNotBillable(activityId int64, name string, timeInDay float64) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = name

	return timeInputElement
}

func TimeInputElementPermanent(activityId int64, name string, timeInDay float64) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_PERMANENT
	timeInputElement.Activity.Title = name

	return timeInputElement
}

func TimeInputElementPermanentAt(activityId int64, name string, timeInDay float64, day time.Time) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%g", timeInDay)
	timeInputElement.Day = tools.DateToString(day)
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_PERMANENT
	timeInputElement.Activity.Title = name

	return timeInputElement
}

func TimeInputElementNotBillableAt(activityId int64, name string, timeInDay float64, day time.Time) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = tools.DateToString(day)
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = name

	return timeInputElement
}

func TimeInputElementBillable(activityId int64, name string, timeInDay float64, clientName string, reference string) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_BILLABLE
	timeInputElement.Activity.Title = name
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.Reference = reference
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = clientName

	return timeInputElement
}

func TimeInputElementBillableAtDay(day time.Time, activityId int64, name string, timeInDay float64, clientName string, reference string) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = tools.DateToString(day)
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_BILLABLE
	timeInputElement.Activity.Title = name
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.Reference = reference
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = clientName

	return timeInputElement
}

func TimeInputElementAbsence(activityId int64, timeInDay float64) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_PERMANENT
	timeInputElement.Activity.Title = KIND_ABSENCE

	return timeInputElement
}

func TimeInputElementInternal(activityId int64, name string, timeInDay float64, clientName string, reference string) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = KIND_INTERNAL
	timeInputElement.Activity.Title = name
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.Reference = reference
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = clientName

	return timeInputElement
}

func TimeInputOneJsonGetter() string {
	return `[
		{
			"day": "2022-03-01",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		}]`
}

func TimeInputTwoJsonGetter() string {
	return `[
		{
			"day": "2022-03-01",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		},
		{
			"day": "2022-03-02",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		}]`
}

func TimeInputThreeJsonGetter() string {
	return `[
		{
			"day": "2022-03-01",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		},
		{
			"day": "2022-03-02",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		},
		{
			"day": "2022-03-03",
			"time_in_days": "1.0",
			"activity": {
				"id": 2140318361,
				"title": "Intercontrat ",
				"role": null,
				"nb_days": null,
				"average_daily_rate": null,
				"kind": "permanent",
				"staffing_needed_from": null,
				"expertise": null,
				"project": null
			}
		}]`
}
