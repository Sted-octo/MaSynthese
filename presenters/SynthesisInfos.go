package presenters

import "Octoptimist/domain"

type FormInfos struct {
	Id                          string
	StartDate                   string
	EndDate                     string
	TotalWorkDays               string
	TacePeriod                  string
	FiscalYear                  string
	TaceFiscalYear              string
	TaceOptimist                string
	TaceInternal                string
	TacePeriodWithDiscount      string
	TaceOptimistWithDiscount    string
	NGram                       string
	Human                       PeopleInfos
	IncludeDiscountInTacePeriod bool
}

type PeopleInfos struct {
	Quadri     string
	FirstName  string
	LastName   string
	EntryDate  string
	ID         string
	Team       string
	TeamId     int64
	JobName    string
	TargetTace string
}

type SynthesisInfos struct {
	CssClass      FormInfos
	Datas         FormInfos
	AccessToken   string
	Lines         []domain.SynthesisLine
	ModeConnexion string
}
