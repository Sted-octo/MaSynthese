package domain

func PeopleActiveWithTargetTace(Taced float64, TacedWithDiscount float64, TaceAble float64) *PeopleInTribe {
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
				Value: Taced / TaceAble,
			},
			RecalculatedFiscalYearActivityRate: ActivityRate{
				Value: Taced / TaceAble,
			},
			RecalculatedPeriodActivityRate: ActivityRate{
				Value: Taced / TaceAble,
			},
			RecalculatedPeriodWithDiscountActivityRate: ActivityRate{
				Value: TacedWithDiscount / TaceAble,
			},
			OptimistActivityRate: ActivityRate{
				Value: Taced / TaceAble,
			},
			OptimistWithDiscountActivityRate: ActivityRate{
				Value: TacedWithDiscount / TaceAble,
			},
		},
		PeriodWorkDays:    250,
		TargetTace:        80,
		TaceAble:          TaceAble,
		Taced:             Taced,
		TacedWithDiscount: TacedWithDiscount,
	}
}

func PeopleInactiveWithTargetTace(Taced float64, TacedWithDiscount float64, TaceAble float64) *PeopleInTribe {
	peopleInTribe := PeopleActiveWithTargetTace(Taced, TacedWithDiscount, TaceAble)
	peopleInTribe.Actif = false

	return peopleInTribe
}

func PeopleActiveWithoutTargetTace(Taced float64, TacedWithDiscount float64, TaceAble float64) *PeopleInTribe {
	peopleInTribe := PeopleActiveWithTargetTace(Taced, TacedWithDiscount, TaceAble)
	peopleInTribe.TargetTace = 0.0
	peopleInTribe.TaceAble = 0.0

	return peopleInTribe

}
