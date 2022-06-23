package main

import "fmt"

func timeInputElementNotBillable(activityId int64, name string, timeInDay float64) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = "permanent"
	timeInputElement.Activity.Title = name

	return timeInputElement
}

func timeInputElementBillable(activityId int64, name string, timeInDay float64, clientName string, reference string) *TimeInputElement {
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", timeInDay)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = activityId
	timeInputElement.Activity.Kind = "permanent"
	timeInputElement.Activity.Title = name
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.Reference = reference
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = clientName

	return timeInputElement
}

func timeInputOneJsonGetter() string {
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

func timeInputTwoJsonGetter() string {
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