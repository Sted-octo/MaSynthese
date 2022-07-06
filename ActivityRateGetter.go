package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ActivityRateGetter(acessToken string, peopleId string, beginPeriod string, endPeriod string) (*ActivityRate, error) {
	if acessToken == "" {
		return nil, errors.New("access token can't be empty")
	}

	var activityRate ActivityRate

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("%s/people/%s/activity_rate?from_date=%s&to_date=%s&include_pipe=false", OCTOPOD_ROOT_URL, peopleId, beginPeriod, endPeriod)

	fmt.Println(urlApi)

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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &activityRate)
	if err != nil {
		return nil, err
	}

	return &activityRate, nil
}