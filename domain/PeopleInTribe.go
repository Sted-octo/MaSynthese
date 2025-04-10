package domain

type PeopleInTribe struct {
	Actif          bool
	Person         People
	ActivityRates  AllActivityRates
	PeriodWorkDays int
	TargetTace     int
	TaceItems      []TaceItem
}
