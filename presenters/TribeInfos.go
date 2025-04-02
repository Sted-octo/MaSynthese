package presenters

type TribeMember struct {
	FirstName                string
	Name                     string
	Nickname                 string
	JobName                  string
	TargetTace               string
	TotalWorkDays            string
	TaceFiscalYear           string
	TaceOptimist             string
	TacePeriodWithDiscount   string
	TaceOptimistWithDiscount string
	TacePeriod               string
}

type TribeInfos struct {
	CssClass        FormInfos
	Datas           FormInfos
	AccessToken     string
	TribeName       string
	TribeId         string
	Members         []TribeMember
	MembersCount    string
	BypassNicknames string
}
