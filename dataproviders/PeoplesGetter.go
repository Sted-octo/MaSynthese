package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func PeoplesGetter(accessToken string) (map[string]domain.People, map[string][]domain.People, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := tools.OctopodUrlApiGetter() + "/people?order_by=nickname&order=asc"

	tools.Debug(urlApi)

	request, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer "+accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	var peoples []domain.People

	err = json.Unmarshal(body, &peoples)
	if err != nil {
		return nil, nil, err
	}

	peoplesMap := make(map[string]domain.People)
	peopleByTribeMap := make(map[string][]domain.People)

	for _, people := range peoples {
		peoplesMap[people.Nickname] = people
		peopleByTribeMap[strconv.FormatInt(people.LobId, 10)] = append(peopleByTribeMap[strconv.FormatInt(people.LobId, 10)], people)
	}

	return peoplesMap, peopleByTribeMap, nil
}
