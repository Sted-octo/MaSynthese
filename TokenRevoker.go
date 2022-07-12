package main

import (
	"fmt"
	"net/http"
	"time"
)

func TokenRevoker(accessToken string) error {

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("https://octopod.octo.com/api/oauth/revoke?token=%s", accessToken)

	fmt.Println(urlApi)

	request, err := http.NewRequest("POST", urlApi, nil)

	if err != nil {
		return err
	}

	request.Header.Add("content-type", "x-www-form-urlencoded")
	request.Header.Add("accept", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("revoke token return response code : %d", response.StatusCode)
	}

	return nil
}
