package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PeoplesByLeagueIdGetter struct{}

func (p *PeoplesByLeagueIdGetter) Get(accessToken string, leagueId int64) ([]domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("%s/people?all=true&league_id=%d", tools.OctopodUrlApiGetter(), leagueId)

	tools.Debug(urlApi)

	request, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer "+accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var peopleLeagues []PeopleLeagueDto

	err = json.Unmarshal(body, &peopleLeagues)
	if err != nil {
		return nil, err
	}

	peoples := make([]domain.People, 0)

	for _, peopleLeague := range peopleLeagues {
		people := domain.People{
			ID:              peopleLeague.ID,
			LastName:        peopleLeague.LastName,
			FirstName:       peopleLeague.FirstName,
			Nickname:        peopleLeague.Nickname,
			JobId:           peopleLeague.Job.ID,
			JobName:         peopleLeague.Job.Name,
			LobId:           peopleLeague.Lob.ID,
			LobAbbreviation: peopleLeague.Lob.Abbreviation,
			EntryDate:       peopleLeague.EntryDate,
			LeavingDate:     peopleLeague.LeavingDate,
			LeagueId:        peopleLeague.Lob.League.ID,
			LeagueName:      peopleLeague.Lob.League.Name,
		}

		peoples = append(peoples, people)
	}

	return peoples, nil
}
