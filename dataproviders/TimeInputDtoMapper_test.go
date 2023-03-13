package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TimeInputDto_No_TimeInputElement_ToTimeInput_Should_Return_Nil(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	timeInputDto := new(TimeInputDto)

	timeInput := timeInputDto.ToTimeInput(&globalPurposeProjectsManager)

	assert.Nil(t, timeInput, "TimeInputDto without timeInputElement should return nil")
}

func Test_TimeInputDto_With_One_TimeInputElement_ToTimeInput_Should_Return_Len_1(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	timeInputDto := new(TimeInputDto)
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", 1.0)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = 123456
	timeInputElement.Activity.Kind = domain.KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = "Intercontrat"

	timeInputDto.Add(timeInputElement)

	timeInput := timeInputDto.ToTimeInput(&globalPurposeProjectsManager)

	assert.Equal(t, 1, len(*timeInput), fmt.Sprintf("TimeInputDto with one timeInputElement should return len 1 but was %d", len(*timeInput)))
}

func Test_TimeInputDto_With_GlobalPurposeProject_ToTimeInput_Should_Return_TimeInput_With_Activity_GlobalPurpose_True(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	timeInputDto := new(TimeInputDto)
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", 1.0)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = 123456
	timeInputElement.Activity.Kind = domain.KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = "Intercontrat"
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.ID = 852
	timeInputElement.Activity.Project.Reference = "5678-0002"
	timeInputElement.Activity.Project.Name = ""
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = "OCTO"

	timeInputDto.Add(timeInputElement)

	timeInput := timeInputDto.ToTimeInput(&globalPurposeProjectsManager)

	assert.True(t, (*timeInput)[0].Activity.GlobalPurpose, "TimeInputDto with project reference in Global purpose list should return time input with activity global purpose true")
}

func Test_TimeInputDto_Without_GlobalPurposeProject_ToTimeInput_Should_Return_TimeInput_With_Activity_GlobalPurpose_False(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	timeInputDto := new(TimeInputDto)
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", 1.0)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = 123456
	timeInputElement.Activity.Kind = domain.KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = "Intercontrat"
	timeInputElement.Activity.Project = new(Project)
	timeInputElement.Activity.Project.ID = 852
	timeInputElement.Activity.Project.Reference = "1234"
	timeInputElement.Activity.Project.Name = ""
	timeInputElement.Activity.Project.Customer = new(Customer)
	timeInputElement.Activity.Project.Customer.Name = "OCTO"

	timeInputDto.Add(timeInputElement)

	timeInput := timeInputDto.ToTimeInput(&globalPurposeProjectsManager)

	assert.False(t, (*timeInput)[0].Activity.GlobalPurpose, "TimeInputDto with project reference in Global purpose list should return time input with activity global purpose true")
}
