package domain

type IPeoplesByLeagueIdGetter interface {
	Get(accessToken string, leagueId int64) ([]People, error)
}
