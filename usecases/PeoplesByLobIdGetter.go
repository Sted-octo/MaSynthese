package usecases

import (
	"Octoptimist/domain"
	"time"
)

type PeoplesByLobIdGetter struct{}

func (p PeoplesByLobIdGetter) Get(accessToken string, lobId int64, period *domain.Period, lobGetter domain.ILobGetter, peoplesByLeagueIdGetter domain.IPeoplesByLeagueIdGetter) ([]domain.People, error) {

	lob, err := lobGetter.Get(accessToken, lobId)

	if err != nil {
		return nil, err
	}

	var leagueId int64 = lob.LeagueId

	peoplesInLeague, err := peoplesByLeagueIdGetter.Get(accessToken, leagueId)
	if err != nil {
		return nil, err
	}

	peoples := make([]domain.People, 0)
	for _, people := range peoplesInLeague {
		if people.LobId == lobId {
			if validEntryDate, err := time.Parse("2006-01-02", people.EntryDate); err == nil {
				if validEntryDate.After(period.End) {
					continue
				}
			}
			if validLeavingDate, err := time.Parse("2006-01-02", people.LeavingDate); err == nil {
				if validLeavingDate.Before(period.Start) {
					continue
				}
			}
			peoples = append(peoples, people)
		}
	}

	return peoples, nil
}
