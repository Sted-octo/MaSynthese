package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ActivityRateGetter struct{}

func (a *ActivityRateGetter) Get(acessToken string, peopleId string, beginPeriod string, endPeriod string) (*domain.ActivityRate, error) {
	if acessToken == "" {
		return nil, errors.New("access token can't be empty")
	}

	var activityRate domain.ActivityRate

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("%s/people/%s/activity_rate?from_date=%s&to_date=%s&include_pipe=false", tools.OctopodUrlApiGetter(), peopleId, beginPeriod, endPeriod)

	tools.Debug(urlApi)

	request, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer "+acessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &activityRate)
	if err != nil {
		return nil, err
	}

	return &activityRate, nil
}
