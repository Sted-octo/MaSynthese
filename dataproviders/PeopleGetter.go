package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PeopleGetter(accessToken string) (*domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := "https://octopod.octo.com/api/v0/people/me"

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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var people domain.People

	err = json.Unmarshal(body, &people)
	if err != nil {
		return nil, err
	}

	return &people, nil
}

func PeopleByIdGetter(accessToken string, peopleId string) (*domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("https://octopod.octo.com/api/v0/people/%s", peopleId)

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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var people domain.People

	err = json.Unmarshal(body, &people)
	if err != nil {
		return nil, err
	}

	return &people, nil
}
