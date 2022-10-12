package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	maintenant := time.Now()

	f, err := os.Create(fmt.Sprintf("%s-rpp%d-%d-%d-%d-%d.log", peopleId, resultPerPage, maintenant.Hour(), maintenant.Minute(), maintenant.Second(), maintenant.Nanosecond()))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	f.WriteString(header())

	urlApi1 := fmt.Sprintf("%s/people/%s/time_input?from_date=%s&to_date=%s&page=1&per_page=1", OCTOPOD_ROOT_URL, peopleId, beginPeriod, endPeriod)

	fmt.Println(urlApi1)

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

	urlApi2 := fmt.Sprintf("%s/people/%s/time_input?from_date=%s&to_date=%s&page=1&per_page=%d", OCTOPOD_ROOT_URL, peopleId, beginPeriod, endPeriod, totalAvaillableLinesCount)

	fmt.Println(urlApi2)

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

	fmt.Printf("Total time Inputs : %d\n", totalAvaillableLinesCount)
	body, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &timeInput)
	if err != nil {
		return nil, err
	}
	nbTimes := len(timeInput)

	for _, ti := range timeInput {
		f.WriteString(line(ti))
	}

	nbLinesLoaded += nbTimes
	fmt.Printf("Count of time Inputs : %d\n", nbTimes)
	totalTimeInput = timeInput

	return &totalTimeInput, nil
}

func header() string {
	return fmt.Sprintln("Day;TimeInDays;ActivityID;ActivityTitle;ActivityNbDays;AverageDailyRate;ActivityKind;ActivityStaffingNeededFrom;ActivityExpertise;ProjectID;ProjectURL;ProjectName;ProjectKind;ProjectReference;ProjectStatus;CustomerID;CustomerName;CustomerAccountManagerID;CustomerAccountManagerEmail")
}

func line(ti TimeInputElement) string {
	returnLine := fmt.Sprintf("%s;%s;%d;%s;%v;%v;%v;%v;%v;",
		ti.Day,
		ti.TimeInDays,
		ti.Activity.ID,
		ti.Activity.Title,
		ti.Activity.NbDays,
		ti.Activity.AverageDailyRate,
		ti.Activity.Kind,
		ti.Activity.StaffingNeededFrom,
		ti.Activity.Expertise)
	if ti.Activity.Project == nil {
		returnLine += "\n"
	} else {
		returnLine += fmt.Sprintf("%d;%s;%s;%s;%s;%v;",
			ti.Activity.Project.ID,
			ti.Activity.Project.URL,
			ti.Activity.Project.Name,
			ti.Activity.Project.Kind,
			ti.Activity.Project.Reference,
			ti.Activity.Project.Status)
		if ti.Activity.Project.Customer == nil {
			returnLine += "\n"
		} else {
			returnLine += fmt.Sprintf("%d;%s;%d;%s\n",
				ti.Activity.Project.Customer.ID,
				ti.Activity.Project.Customer.Name,
				ti.Activity.Project.Customer.AccountManagerID,
				ti.Activity.Project.Customer.AccountManagerEmail)
		}
	}
	return returnLine
}
