package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func TokenGetter(clientId string, clientSecret string, redirectUrl string, authCode string) (*domain.Token, error) {

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("https://octopod.octo.com/api/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s&redirect_uri=%s", clientId, clientSecret, redirectUrl)

	if authCode != "" {
		urlApi = fmt.Sprintf("https://octopod.octo.com/api/oauth/token?grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s", authCode, clientId, clientSecret, redirectUrl)
	}
	tools.Debug(urlApi)

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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var token domain.Token

	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
