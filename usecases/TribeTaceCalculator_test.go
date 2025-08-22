package usecases

import (
	"Octoptimist/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TribeTaceCalculator_Shoul_Return_0_When_Input_List_Is_Empty(t *testing.T) {
	period := Period250WorkDaysGetter()
	tribeTaceCalculator := new(TribeTaceCalculator)
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate([]domain.PeopleInTribe{}, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_Should_Return_0_When_Input_List_Is_Nil(t *testing.T) {
	period := Period250WorkDaysGetter()
	tribeTaceCalculator := new(TribeTaceCalculator)
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate(nil, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_FullTime_In_Tribe_With_TargetTace_0_Should_Return_0(t *testing.T) {
	period := Period250WorkDaysGetter()
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithoutTargetTace(20.0, 20.0, 200.0))
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_FullTime_In_Tribe_With_TargetTace_Should_Return_People_ActivityRate(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced float64 = 20.0
	var tacedWithDiscount float64 = taced
	var taceAble float64 = 200.0
	activityRate := domain.ActivityRate{Value: taced / taceAble}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced, tacedWithDiscount, taceAble))

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, activityRate.Value, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_FullTime_In_Tribe_One_Inactive_Should_Return_Active_People_ActivityRate(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced float64 = 20.0
	var tacedWithDiscount float64 = taced
	var taceAble float64 = 200.0
	activityRatePeopleActive := domain.ActivityRate{Value: taced / taceAble}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced, tacedWithDiscount, taceAble))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleInactiveWithTargetTace(taced, tacedWithDiscount, taceAble))

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, activityRatePeopleActive.Value, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_FullTime_In_Tribe_Actives_Should_Return_Average_ActivityRate(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced1 float64 = 20.0
	var taced2 float64 = 40.0
	var tacedWithDiscount1 float64 = taced1
	var tacedWithDiscount2 float64 = taced2
	var taceAble float64 = 200.0
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced1, tacedWithDiscount1, taceAble))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced2, tacedWithDiscount2, taceAble))
	var expected float64 = (taced1 + taced2) / (taceAble * 2)

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_Actives_In_Tribe_On_Fulltime_One_Haltime_Should_Return_Ponderated_Average_ActivityRate(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced1 float64 = 200.0
	var taced2 float64 = 50.0
	var tacedWithDiscount1 float64 = taced1
	var tacedWithDiscount2 float64 = taced2
	var taceAble1 float64 = 200.0
	var taceAble2 float64 = 100.0

	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced1, tacedWithDiscount1, taceAble1))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced2, tacedWithDiscount2, taceAble2))

	var expected float64 = (taced1 + taced2) / (taceAble1 + taceAble2)

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_Without_Taceable_Day_Should_Return_ActivityRate_0_Without_Divide_By_Zero_Error(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced float64 = 20.0
	var tacedWithDiscount float64 = taced
	var taceAble float64 = 0.0
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced, tacedWithDiscount, taceAble))
	var expected float64 = 0

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_With_Disount_Day_Should_Return_ActivityRate_Different_Than_ActivityRateWithDiscount(t *testing.T) {
	period := Period250WorkDaysGetter()
	var taced float64 = 100.0
	var tacedWithDiscount float64 = 150.0
	var taceAble float64 = 200.0
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTace(taced, tacedWithDiscount, taceAble))
	var expectedWithoutDiscount float64 = taced / taceAble
	var expectedWithDiscount float64 = tacedWithDiscount / taceAble

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe, period)

	assert.Equal(t, expectedWithoutDiscount, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expectedWithDiscount, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
}
