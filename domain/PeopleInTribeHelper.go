package domain

func PeopleActiveWithTargetTaceFullTime(activityRate ActivityRate) *PeopleInTribe {
	return &PeopleInTribe{
		Actif: true,
		Person: People{
			ID:          1,
			EntryDate:   "2023-01-01",
			LeavingDate: "",
			FirstName:   "J.J",
			LastName:    "Goldman",
			Nickname:    "JJG",
			JobName:     "Singer",
		},
		ActivityRates: AllActivityRates{
			OctopodFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
		},
		PeriodWorkDays: 250,
		TargetTace:     80,
	}
}
func PeopleInactiveWithTargetTaceFullTime(activityRate ActivityRate) *PeopleInTribe {
	return &PeopleInTribe{
		Actif: false,
		Person: People{
			ID:          2,
			EntryDate:   "2023-01-01",
			LeavingDate: "",
			FirstName:   "J.J",
			LastName:    "Goldman",
			Nickname:    "JJG",
			JobName:     "Singer",
		},
		ActivityRates: AllActivityRates{
			OctopodFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
		},
		PeriodWorkDays: 250,
		TargetTace:     50,
	}
}
func PeopleActiveWithoutTargetTaceFullTime(activityRate ActivityRate) *PeopleInTribe {
	return &PeopleInTribe{
		Actif: true,
		Person: People{
			ID:          3,
			EntryDate:   "2023-01-01",
			LeavingDate: "",
			FirstName:   "J.J",
			LastName:    "Goldman",
			Nickname:    "JJG",
			JobName:     "Singer",
		},
		ActivityRates: AllActivityRates{
			OctopodFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
		},
		PeriodWorkDays: 250,
		TargetTace:     0,
	}
}

func PeopleActiveWithTargetTaceHalfTime(activityRate ActivityRate) *PeopleInTribe {
	return &PeopleInTribe{
		Actif: true,
		Person: People{
			ID:          1,
			EntryDate:   "2023-01-01",
			LeavingDate: "2024-05-31",
			FirstName:   "J.J",
			LastName:    "Goldman",
			Nickname:    "JJG",
			JobName:     "Singer",
		},
		ActivityRates: AllActivityRates{
			OctopodFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedFiscalYearActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			RecalculatedPeriodWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
			OptimistWithDiscountActivityRate: ActivityRate{
				Value: activityRate.Value,
			},
		},
		PeriodWorkDays: 125,
		TargetTace:     80,
	}
}
