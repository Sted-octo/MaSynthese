package domain

type People struct {
	ID              int64  `json:"id"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	Nickname        string `json:"nickname"`
	EntryDate       string `json:"entry_date"`
	LeavingDate     string `json:"leaving_date"`
	JobId           int64  `json:"jobid"`
	JobName         string `json:"jobname"`
	LobId           int64  `json:"lobid"`
	LobAbbreviation string `json:"lobabbreviation"`
	LeagueId        int64  `json:"leagueid"`
	LeagueName      string `json:"leaguename"`
}
