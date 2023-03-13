package dataproviders

import (
	"Octoptimist/domain"
	"strconv"
	"time"
)

func (timeInputDto *TimeInputDto) ToTimeInput() *domain.TimeInput {
	if len(*timeInputDto) == 0 {
		return nil
	}
	var timeInput *domain.TimeInput = new(domain.TimeInput)

	for _, ti := range *timeInputDto {
		var timeInputElement *domain.TimeInputElement = new(domain.TimeInputElement)
		if validDate, err := time.Parse("2006-01-02", ti.Day); err == nil {
			timeInputElement.Day = validDate
		}
		if decimal, err := strconv.ParseFloat(ti.TimeInDays, 64); err == nil {
			timeInputElement.TimeInDays = decimal
		}
		timeInputElement.Activity = *new(domain.Activity)
		timeInputElement.Activity.ID = ti.Activity.ID
		timeInputElement.Activity.Title = ti.Activity.Title
		timeInputElement.Activity.Kind = ti.Activity.Kind
		if ti.Activity.Project != nil {
			timeInputElement.Activity.Project = new(domain.Project)
			timeInputElement.Activity.Project.ID = ti.Activity.Project.ID
			timeInputElement.Activity.Project.Name = ti.Activity.Project.Name
			timeInputElement.Activity.Project.Reference = ti.Activity.Project.Reference
			if ti.Activity.Project.Customer != nil {
				timeInputElement.Activity.Project.Customer = new(domain.Customer)
				timeInputElement.Activity.Project.Customer.ID = ti.Activity.Project.Customer.ID
				timeInputElement.Activity.Project.Customer.Name = ti.Activity.Project.Customer.Name
			}
		}

		timeInput.Add(timeInputElement)
	}

	return timeInput
}
