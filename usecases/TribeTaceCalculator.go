package usecases

import "Octoptimist/domain"

type TribeTaceCalculator struct{}

func (tribeTaceCalculator *TribeTaceCalculator) Calculate(peoplesInTribe []domain.PeopleInTribe) *domain.AllActivityRates {
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

	workDaysPartSplitter := new(WorkDaysPartSplitter)
	tribeTaceItems := workDaysPartSplitter.Split(activePeopleInTribe)

	maxTotalWorkDays := tribeTaceItems[0].TotalWorkDays

	for _, peopleInTribe := range activePeopleInTribe {
		for i, taceItem := range tribeTaceItems {
			if taceItem.TotalWorkDays <= peopleInTribe.PeriodWorkDays {
				peopleTaceItem := domain.TaceItem{
					TotalWorkDays: taceItem.TotalWorkDays,
					DeltaWorkDays: taceItem.DeltaWorkDays,
					NbItems:       taceItem.NbItems,
				}
				peopleTaceItem.ActivityRates.RecalculatedPeriodActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.RecalculatedPeriodActivityRate.Value / float64(maxTotalWorkDays)
				peopleTaceItem.ActivityRates.OctopodFiscalYearActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.OctopodFiscalYearActivityRate.Value / float64(maxTotalWorkDays)
				peopleTaceItem.ActivityRates.RecalculatedFiscalYearActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.RecalculatedFiscalYearActivityRate.Value / float64(maxTotalWorkDays)
				peopleTaceItem.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value / float64(maxTotalWorkDays)
				peopleTaceItem.ActivityRates.OptimistActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.OptimistActivityRate.Value / float64(maxTotalWorkDays)
				peopleTaceItem.ActivityRates.OptimistWithDiscountActivityRate.Value = float64(taceItem.DeltaWorkDays) * peopleInTribe.ActivityRates.OptimistWithDiscountActivityRate.Value / float64(maxTotalWorkDays)

				peopleInTribe.TaceItems = append(peopleInTribe.TaceItems, peopleTaceItem)

				tribeTaceItems[i].ActivityRates.RecalculatedPeriodActivityRate.Value += peopleTaceItem.ActivityRates.RecalculatedPeriodActivityRate.Value
				tribeTaceItems[i].ActivityRates.OctopodFiscalYearActivityRate.Value += peopleTaceItem.ActivityRates.OctopodFiscalYearActivityRate.Value
				tribeTaceItems[i].ActivityRates.RecalculatedFiscalYearActivityRate.Value += peopleTaceItem.ActivityRates.RecalculatedFiscalYearActivityRate.Value
				tribeTaceItems[i].ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value += peopleTaceItem.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value
				tribeTaceItems[i].ActivityRates.OptimistActivityRate.Value += peopleTaceItem.ActivityRates.OptimistActivityRate.Value
				tribeTaceItems[i].ActivityRates.OptimistWithDiscountActivityRate.Value += peopleTaceItem.ActivityRates.OptimistWithDiscountActivityRate.Value
			}
		}
	}

	for _, taceItem := range tribeTaceItems {
		taceItem.ActivityRates.RecalculatedPeriodActivityRate.Value = taceItem.ActivityRates.RecalculatedPeriodActivityRate.Value / float64(taceItem.NbItems)
		taceItem.ActivityRates.OctopodFiscalYearActivityRate.Value = taceItem.ActivityRates.OctopodFiscalYearActivityRate.Value / float64(taceItem.NbItems)
		taceItem.ActivityRates.RecalculatedFiscalYearActivityRate.Value = taceItem.ActivityRates.RecalculatedFiscalYearActivityRate.Value / float64(taceItem.NbItems)
		taceItem.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value = taceItem.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value / float64(taceItem.NbItems)
		taceItem.ActivityRates.OptimistActivityRate.Value = taceItem.ActivityRates.OptimistActivityRate.Value / float64(taceItem.NbItems)
		taceItem.ActivityRates.OptimistWithDiscountActivityRate.Value = taceItem.ActivityRates.OptimistWithDiscountActivityRate.Value / float64(taceItem.NbItems)

		activityRates.RecalculatedPeriodActivityRate.Value += taceItem.ActivityRates.RecalculatedPeriodActivityRate.Value
		activityRates.OctopodFiscalYearActivityRate.Value += taceItem.ActivityRates.OctopodFiscalYearActivityRate.Value
		activityRates.RecalculatedFiscalYearActivityRate.Value += taceItem.ActivityRates.RecalculatedFiscalYearActivityRate.Value
		activityRates.RecalculatedPeriodWithDiscountActivityRate.Value += taceItem.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value
		activityRates.OptimistActivityRate.Value += taceItem.ActivityRates.OptimistActivityRate.Value
		activityRates.OptimistWithDiscountActivityRate.Value += taceItem.ActivityRates.OptimistWithDiscountActivityRate.Value
	}

	return &activityRates
}
