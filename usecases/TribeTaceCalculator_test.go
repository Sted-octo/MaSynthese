package usecases

import (
	"Octoptimist/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TribeTaceCalculator_Shoul_Return_0_When_Input_List_Is_Empty(t *testing.T) {
	tribeTaceCalculator := new(TribeTaceCalculator)
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate([]domain.PeopleInTribe{})

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_Should_Return_0_When_Input_List_Is_Nil(t *testing.T) {
	tribeTaceCalculator := new(TribeTaceCalculator)
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate(nil)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_FullTime_In_Tribe_With_TargetTace_0_Should_Return_0(t *testing.T) {
	activityRate := domain.ActivityRate{Value: 10.0}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithoutTargetTaceFullTime(activityRate))
	var expected float64 = 0.0

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_OnePeople_FullTime_In_Tribe_With_TargetTace_Should_Return_People_ActivityRate(t *testing.T) {
	activityRate := domain.ActivityRate{Value: 10.0}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceFullTime(activityRate))

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

	assert.Equal(t, activityRate.Value, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, activityRate.Value, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_FullTime_In_Tribe_One_Inactive_Should_Return_Active_People_ActivityRate(t *testing.T) {
	activityRatePeopleActive := domain.ActivityRate{Value: 10.0}
	activityRatePeopleInactive := domain.ActivityRate{Value: 20.0}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceFullTime(activityRatePeopleActive))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleInactiveWithTargetTaceFullTime(activityRatePeopleInactive))

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

	assert.Equal(t, activityRatePeopleActive.Value, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, activityRatePeopleActive.Value, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_FullTime_In_Tribe_Actives_Should_Return_Average_ActivityRate(t *testing.T) {
	activityRatePeopleActive1 := domain.ActivityRate{Value: 10.0}
	activityRatePeopleActive2 := domain.ActivityRate{Value: 20.0}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceFullTime(activityRatePeopleActive1))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceFullTime(activityRatePeopleActive2))
	var expected float64 = 15.0

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistWithDiscountActivityRate.Value)
}

func Test_TribeTaceCalculator_TwoPeoples_Actives_In_Tribe_On_Fulltime_One_Haltime_Should_Return_Ponderated_Average_ActivityRate(t *testing.T) {
	activityRatePeopleActiveFulltime := domain.ActivityRate{Value: 100.0}
	activityRatePeopleActiveHalftime := domain.ActivityRate{Value: 50.0}
	tribeTaceCalculator := new(TribeTaceCalculator)
	peoplesInTribe := make([]domain.PeopleInTribe, 0)
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceFullTime(activityRatePeopleActiveFulltime))
	peoplesInTribe = append(peoplesInTribe, *domain.PeopleActiveWithTargetTaceHalfTime(activityRatePeopleActiveHalftime))

	var expected float64 = 87.5

	activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

	assert.Equal(t, expected, activityRates.RecalculatedPeriodActivityRate.Value)
	assert.Equal(t, expected, activityRates.OctopodFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedFiscalYearActivityRate.Value)
	assert.Equal(t, expected, activityRates.RecalculatedPeriodWithDiscountActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistActivityRate.Value)
	assert.Equal(t, expected, activityRates.OptimistWithDiscountActivityRate.Value)
}
