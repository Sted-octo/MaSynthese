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

func PeopleGetter(accessToken string) (*domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := tools.OctopodUrlApiGetter() + "/people/me"

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

	var peopleDto PeopleDto

	err = json.Unmarshal(body, &peopleDto)
	if err != nil {
		return nil, err
	}

	people := domain.People{
		ID:              peopleDto.ID,
		LastName:        peopleDto.LastName,
		FirstName:       peopleDto.FirstName,
		Nickname:        peopleDto.Nickname,
		EntryDate:       peopleDto.EntryDate,
		JobId:           peopleDto.Job.ID,
		JobName:         peopleDto.Job.Name,
		LobId:           peopleDto.Lob.ID,
		LobAbbreviation: peopleDto.Lob.Abbreviation,
		LeagueId:        peopleDto.Lob.League.ID,
		LeagueName:      peopleDto.Lob.League.Name,
	}
	if peopleDto.LeavingDate != nil {
		people.LeavingDate = peopleDto.LeavingDate.(string)
	}

	return &people, nil
}

func PeopleByIdGetter(accessToken string, peopleId string) (*domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf(tools.OctopodUrlApiGetter()+"/people/%s", peopleId)

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

	var peopleDto PeopleDto

	err = json.Unmarshal(body, &peopleDto)
	if err != nil {
		return nil, err
	}
	people := domain.People{
		ID:              peopleDto.ID,
		LastName:        peopleDto.LastName,
		FirstName:       peopleDto.FirstName,
		Nickname:        peopleDto.Nickname,
		EntryDate:       peopleDto.EntryDate,
		JobId:           peopleDto.Job.ID,
		JobName:         peopleDto.Job.Name,
		LobId:           peopleDto.Lob.ID,
		LobAbbreviation: peopleDto.Lob.Abbreviation,
		LeagueId:        peopleDto.Lob.League.ID,
		LeagueName:      peopleDto.Lob.League.Name,
	}

	if peopleDto.LeavingDate != nil {
		people.LeavingDate = peopleDto.LeavingDate.(string)
	}

	return &people, nil
}
