package domain

type PeopleInTribe struct {
	Actif                    bool
	Person                   People
	OctopodFyTace            ActivityRate
	CustomFYTace             ActivityRate
	PeriodTace               ActivityRate
	PeriodWithDiscountTace   ActivityRate
	OptimistTace             ActivityRate
	OptimistWithDiscountTace ActivityRate
	PeriodWorkDays           int
}
