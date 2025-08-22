package usecases

import (
	"Octoptimist/domain"
)

type TribeTaceCalculator struct{}

func (tribeTaceCalculator *TribeTaceCalculator) Calculate(peoplesInTribe []domain.PeopleInTribe, period domain.Period) *domain.AllActivityRates {
	activityRates := domain.AllActivityRates{}
	activityRates.RecalculatedPeriodActivityRate = domain.ActivityRate{Value: 0.0}
	activityRates.OctopodFiscalYearActivityRate = domain.ActivityRate{Value: 0.0}
	activityRates.RecalculatedFiscalYearActivityRate = domain.ActivityRate{Value: 0.0}
	activityRates.RecalculatedPeriodWithDiscountActivityRate = domain.ActivityRate{Value: 0.0}
	activityRates.OptimistActivityRate = domain.ActivityRate{Value: 0.0}
	activityRates.OptimistWithDiscountActivityRate = domain.ActivityRate{Value: 0.0}

	if len(peoplesInTribe) == 0 {
		return &activityRates
	}

	activePeopleInTribe := make([]domain.PeopleInTribe, 0)

	for _, peopleInTribe := range peoplesInTribe {
		if !peopleInTribe.Actif {
			continue
		}
		if peopleInTribe.TargetTace == 0 {
			continue
		}
		activePeopleInTribe = append(activePeopleInTribe, peopleInTribe)
	}
	if len(activePeopleInTribe) == 0 {
		return &activityRates
	}

	var tacedAccumulation float64 = 0.0
	var tacedWithDiscountAccumulation float64 = 0.0
	var taceAbleAccumulation float64 = 0.0

	for _, peopleInTribe := range activePeopleInTribe {
		taceAbleAccumulation += peopleInTribe.TaceAble
		tacedAccumulation += peopleInTribe.Taced
		tacedWithDiscountAccumulation += peopleInTribe.TacedWithDiscount
	}
	if taceAbleAccumulation == 0.0 {
		return &activityRates
	}

	activityRates.RecalculatedPeriodActivityRate.Value = tacedAccumulation / taceAbleAccumulation
	activityRates.RecalculatedPeriodWithDiscountActivityRate.Value = tacedWithDiscountAccumulation / taceAbleAccumulation

	return &activityRates
}
