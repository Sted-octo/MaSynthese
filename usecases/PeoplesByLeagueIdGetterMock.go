package usecases

import "Octoptimist/domain"

type PeoplesByLeagueIdGetterMockOnePeopleLob1 struct{}

func (p PeoplesByLeagueIdGetterMockOnePeopleLob1) Get(accessToken string, leagueId int64) ([]domain.People, error) {
	peoples := make([]domain.People, 0)
	people := domain.People{
		ID:              1,
		FirstName:       "John",
		LastName:        "Doe",
		Nickname:        "JODO",
		LeagueId:        117,
		LeagueName:      "CAMPS",
		LobId:           1,
		LobAbbreviation: "CHOCOLATINE",
		JobId:           1,
		JobName:         "Manager",
		EntryDate:       "2023-01-01",
		LeavingDate:     "2023-12-31",
	}

	peoples = append(peoples, people)
	return peoples, nil
}

type PeoplesByLeagueIdGetterMockOnePeopleLob44 struct{}

func (p PeoplesByLeagueIdGetterMockOnePeopleLob44) Get(accessToken string, leagueId int64) ([]domain.People, error) {
	peoples := make([]domain.People, 0)
	people := domain.People{
		ID:              1,
		FirstName:       "John",
		LastName:        "Doe",
		Nickname:        "JODO",
		LeagueId:        117,
		LeagueName:      "CAMPS",
		LobId:           44,
		LobAbbreviation: "PAIN_AU_CHOCOLAT",
		JobId:           1,
		JobName:         "Manager",
		EntryDate:       "2023-01-01",
		LeavingDate:     "2023-12-31",
	}

	peoples = append(peoples, people)
	return peoples, nil
}
