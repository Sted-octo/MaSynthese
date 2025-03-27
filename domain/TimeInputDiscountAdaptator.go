package domain

func (timeInput *TimeInput) TimeInputDiscountAdaptator(withDiscount bool, discountProjectsManager *DiscountProjects) *TimeInput {
	if !withDiscount {
		return timeInput
	}

	for i := range *timeInput {
		if (*timeInput)[i].Activity.Project != nil &&
			discountProjectsManager.IsDiscount((*timeInput)[i].Activity.Project.Reference) {
			(*timeInput)[i].Activity.Kind = KIND_BILLABLE
			(*timeInput)[i].Activity.Discount = true
		}
	}

	return timeInput
}
