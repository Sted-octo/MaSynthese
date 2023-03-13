package dataproviders

import (
	"Octoptimist/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TimeInputDto_No_TimeInputElement_ToTimeInput_Should_Return_Nil(t *testing.T) {
	timeInputDto := new(TimeInputDto)

	timeInput := timeInputDto.ToTimeInput()

	assert.Nil(t, timeInput, "TimeInputDto without timeInputElement should return nil")
}

func Test_TimeInputDto_With_One_TimeInputElement_ToTimeInput_Should_Return_Len_1(t *testing.T) {
	timeInputDto := new(TimeInputDto)
	var timeInputElement *TimeInputElement = new(TimeInputElement)
	timeInputElement.TimeInDays = fmt.Sprintf("%f", 1.0)
	timeInputElement.Day = "1973-10-07"
	timeInputElement.Activity = *new(Activity)
	timeInputElement.Activity.ID = 123456
	timeInputElement.Activity.Kind = domain.KIND_NOT_BILLABLE
	timeInputElement.Activity.Title = "Intercontrat"

	timeInputDto.Add(timeInputElement)

	timeInput := timeInputDto.ToTimeInput()

	assert.Equal(t, 1, len(*timeInput), fmt.Sprintf("TimeInputDto with one timeInputElement should return len 1 but was %d", len(*timeInput)))
}
