package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func TimeInputGetter(acessToken string, peopleId string, beginPeriod string, endPeriod string, resultPerPage uint, globalPurposeProjectsManager *domain.GlobalPurposeProjects) (*domain.TimeInput, error) {
	if acessToken == "" {
		return nil, errors.New("access token can't be empty")
	}

	var totalTimeInput domain.TimeInput
	var timeInputDto TimeInputDto

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	nbLinesLoaded := 0

	urlApi1 := fmt.Sprintf("%s/people/%s/time_input?from_date=%s&to_date=%s&page=1&per_page=1", tools.OctopodUrlApiGetter(), peopleId, beginPeriod, endPeriod)

	tools.Debug(urlApi1)

	request1, err := http.NewRequest("GET", urlApi1, nil)

	if err != nil {
		return nil, err
	}

	request1.Header.Add("content-type", "application/json")
	request1.Header.Add("authorization", "Bearer "+acessToken)

	response1, err := httpClient.Do(request1)
	if err != nil {
		return nil, err
	}
	defer response1.Body.Close()

	if response1.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not 200 in tieInputGetter %d", response1.StatusCode)
	}

	totalAvaillableLinesCount, err := strconv.Atoi(response1.Header.Get("Total"))
	if err != nil {
		return nil, err
	}
	if totalAvaillableLinesCount == 0 {
		return new(domain.TimeInput), nil
	}

	urlApi2 := fmt.Sprintf("%s/people/%s/time_input?from_date=%s&to_date=%s&page=1&per_page=%d", tools.OctopodUrlApiGetter(), peopleId, beginPeriod, endPeriod, totalAvaillableLinesCount)

	tools.Debug(urlApi2)

	request2, err := http.NewRequest("GET", urlApi2, nil)

	if err != nil {
		return nil, err
	}

	request2.Header.Add("content-type", "application/json")
	request2.Header.Add("authorization", "Bearer "+acessToken)

	response2, err := httpClient.Do(request2)
	if err != nil {
		return nil, err
	}
	defer response2.Body.Close()

	if response2.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code not 200 in tieInputGetter %d", response2.StatusCode)
	}

	msg := fmt.Sprintf("Total time Inputs : %d\n", totalAvaillableLinesCount)
	tools.Debug(msg)

	body, err := io.ReadAll(response2.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &timeInputDto)
	if err != nil {
		return nil, err
	}
	nbTimes := len(timeInputDto)

	nbLinesLoaded += nbTimes
	msg = fmt.Sprintf("Count of time Inputs : %d\n", nbTimes)
	tools.Debug(msg)

	totalTimeInput = *timeInputDto.ToTimeInput(globalPurposeProjectsManager)

	return &totalTimeInput, nil
}
