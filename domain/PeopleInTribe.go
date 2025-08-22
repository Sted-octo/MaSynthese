package domain

type PeopleInTribe struct {
	Actif                     bool
	Person                    People
	ActivityRates             AllActivityRates
	PeriodWorkDays            int
	TargetTace                int
	TaceItems                 []TaceItem
	TaceAble                  float64
	Taced                     float64
	TacedWithDiscount         float64
	TacedOptimist             float64
	TacedOptimistWithDiscount float64
}
