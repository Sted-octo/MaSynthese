package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func TimeInputGetter(acessToken string) (*TimeInput, error) {

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := "https://octopod.octo.com/api/v0/people/2142666213/time_input?from_date=2022-03-01&to_date=2022-06-07&page=1&per_page=100"

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

	var timeInput TimeInput

	err = json.Unmarshal(body, &timeInput)
	if err != nil {
		return nil, err
	}

	return &timeInput, nil
}
