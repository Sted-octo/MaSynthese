package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func TokenGetter(clientId string, clientSecret string, redirectUrl string, authCode string) (*domain.Token, error) {

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	baseURL := tools.OctopodDomainGetter() + "/api/oauth/token"
	values := url.Values{}
	if authCode == "" {
		values.Add("grant_type", "client_credentials")
	} else {
		values.Add("grant_type", "authorization_code")
		values.Add("code", authCode)
		values.Add("redirect_uri", redirectUrl)
		values.Add("scope", "public")
	}

	values.Add("client_id", clientId)
	values.Add("client_secret", clientSecret)

	// Construction de l'URL complète
	urlApi := fmt.Sprintf("%s?%s", baseURL, values.Encode())

	tools.Debug(urlApi)

	request, err := http.NewRequest("POST", urlApi, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json")

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
