package domain

func (timeInput *TimeInput) Clone() *TimeInput {
	if len(*timeInput) == 0 {
		return nil
	}
	cloned := make(TimeInput, len(*timeInput))
	for i, element := range *timeInput {
		clonedElement := TimeInputElement{
			Day:        element.Day,
			TimeInDays: element.TimeInDays,
			Activity: Activity{
				ID:            element.Activity.ID,
				Title:         element.Activity.Title,
				Kind:          element.Activity.Kind,
				GlobalPurpose: element.Activity.GlobalPurpose,
				Discount:      element.Activity.Discount,
			},
		}

		// Clone Project if it exists
		if element.Activity.Project != nil {
			clonedElement.Activity.Project = &Project{
				ID:        element.Activity.Project.ID,
				Name:      element.Activity.Project.Name,
				Reference: element.Activity.Project.Reference,
			}

			// Clone Customer if it exists
			if element.Activity.Project.Customer != nil {
				clonedElement.Activity.Project.Customer = &Customer{
					ID:   element.Activity.Project.Customer.ID,
					Name: element.Activity.Project.Customer.Name,
				}
			}
		}

		cloned[i] = clonedElement
	}
	return &cloned
}
