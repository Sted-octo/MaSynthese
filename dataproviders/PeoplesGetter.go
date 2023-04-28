package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func PeoplesGetter(accessToken string) (map[string]domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := "https://octopod.octo.com/api/v0/people?order_by=nickname&order=asc"

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

	var peoples []domain.People

	err = json.Unmarshal(body, &peoples)
	if err != nil {
		return nil, err
	}

	peoplesMap := make(map[string]domain.People)

	for _, people := range peoples {
		peoplesMap[people.Nickname] = people
	}

	return peoplesMap, nil
}
