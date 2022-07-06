package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var OCTOPOD_ROOT_URL string = "https://octopod.octo.com/api/v0"

func TimeInputGetter(acessToken string, peopleId string, beginPeriod string, endPeriod string, resultPerPage uint) (*TimeInput, error) {
	if acessToken == "" {
		return nil, errors.New("access token can't be empty")
	}

	var totalTimeInput TimeInput
	var timeInput TimeInput

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	nbLinesLoaded := 0
	fullLoad := false

	for page := 1; !fullLoad; page++ {

		urlApi := fmt.Sprintf("%s/people/%s/time_input?from_date=%s&to_date=%s&page=%d&per_page=%d", OCTOPOD_ROOT_URL, peopleId, beginPeriod, endPeriod, page, resultPerPage)

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

		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code not 200 in tieInputGetter %d", response.StatusCode)
		}

		totalAvaillableLinesCount, err := strconv.Atoi(response.Header.Get("Total"))
		if err != nil {
			return nil, err
		}

		fmt.Printf("Total time Inputs : %d\n", totalAvaillableLinesCount)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &timeInput)
		if err != nil {
			return nil, err
		}
		nbTimes := len(timeInput)
		nbLinesLoaded += nbTimes
		fmt.Printf("Count of time Inputs : %d\n", nbTimes)
		fullLoad = nbLinesLoaded >= totalAvaillableLinesCount
		if totalTimeInput == nil {
			totalTimeInput = timeInput
			continue
		}
		totalTimeInput.Concat(timeInput)
	}

	return &totalTimeInput, nil
}
