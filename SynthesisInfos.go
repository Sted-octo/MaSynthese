package main

type FormInfos struct {
	Id             string
	StartDate      string
	EndDate        string
	TotalWorkDays  string
	TacePeriod     string
	FiscalYear     string
	TaceFiscalYear string
	TaceOptimist   string
	TaceInternal   string
	AuthCode       string
	Human          PeopleInfos
}

type PeopleInfos struct {
	Quadri     string
	FirstName  string
	LastName   string
	EntryDate  string
	ID         string
	Team       string
	JobName    string
	TargetTace string
}

type SynthesisInfos struct {
	CssClass      FormInfos
	Datas         FormInfos
	AccessToken   string
	Lines         []SynthesisLine
	ModeConnexion string
}
