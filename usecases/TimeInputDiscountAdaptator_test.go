package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"
)

func Test_Without_Discount_One_TimeInput_Should_Not_Return_nil(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: MockDiscountProjectsLoader}
	timeInputs = new(domain.TimeInput)
	timeInputs.Add(domain.TimeInputElementBillableAtDay(tools.DateSimple(2023, time.January, 12), 123, "Audit", 1.0, "OCTO", "123456"))

	modifiedTimeInputs := timeInputs.TimeInputDiscountAdaptator(false, &DiscountProjects)

	if modifiedTimeInputs == nil {
		t.Errorf("TimeInputDiscountAdaptator without discount should not return nil")
	}
}

func Test_With_Discount_One_NotBillable_Discount_TimeInput_Should_Change_Kind_As_Billable(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: MockDiscountProjectsLoader}
	timeInputs = new(domain.TimeInput)
	timeInputs.Add(domain.TimeInputElementDiscountAtDay(tools.DateSimple(2023, time.January, 12), 123, "Négo", 1, "RETAILLER", "2024-0005I"))

	modifiedTimeInputs := timeInputs.TimeInputDiscountAdaptator(true, &DiscountProjects)

	if (*modifiedTimeInputs)[0].Activity.Kind != domain.KIND_BILLABLE {
		t.Errorf("TimeInputDiscountAdaptator with discount for a discount timeInput should have activity king Billable")
	}
}

func Test_With_Discount_One_NotBillable_No_Discount_TimeInput_Should_Stay_Not_Billable(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: MockDiscountProjectsLoader}
	timeInputs = new(domain.TimeInput)
	timeInputs.Add(domain.TimeInputElementNotBillableAt(123, "Négo", 1, tools.DateSimple(2023, time.January, 12)))

	modifiedTimeInputs := timeInputs.TimeInputDiscountAdaptator(true, &DiscountProjects)

	if (*modifiedTimeInputs)[0].Activity.Kind != domain.KIND_NOT_BILLABLE {
		t.Errorf("TimeInputDiscountAdaptator with discount for a non discount timeInput should stay not Billable")
	}
}

func Test_With_Discount_One_NotBillable_Discount_TimeInput_Should_Change_Discount_Bool_True(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: MockDiscountProjectsLoader}
	timeInputs = new(domain.TimeInput)
	timeInputs.Add(domain.TimeInputElementDiscountAtDay(tools.DateSimple(2023, time.January, 12), 123, "Négo", 1, "RETAILLER", "2024-0005I"))

	modifiedTimeInputs := timeInputs.TimeInputDiscountAdaptator(true, &DiscountProjects)

	if (*modifiedTimeInputs)[0].Activity.Discount != true {
		t.Errorf("TimeInputDiscountAdaptator with discount for a discount timeInput should change Discount bool to true")
	}
}
