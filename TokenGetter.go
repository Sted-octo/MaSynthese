package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func TokenGetter(clientId string, clientSecret string, authCode string) (*Token, error) {

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := "https://octopod.octo.com/api/oauth/token?grant_type=client_credentials&client_id=" + clientId + "&client_secret=" + clientSecret + "&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob"

	if authCode != "" {
		urlApi = "https://octopod.octo.com/api/oauth/token?grant_type=authorization_code&code=" + authCode + "&client_id=" + clientId + "&client_secret=" + clientSecret + "&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob"
	}
	fmt.Println(urlApi)

	request, err := http.NewRequest("POST", urlApi, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "x-www-form-urlencoded")
	request.Header.Add("accept", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var token Token

	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
