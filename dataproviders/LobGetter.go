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

type LobGetter struct{}

func (l *LobGetter) Get(accessToken string, lobId int64) (*domain.Lob, error) {
	if accessToken == "" {
		return nil, errors.New("access token can't be empty")
	}

	var lob *domain.Lob

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := fmt.Sprintf("%s/lobs/%d", tools.OctopodUrlApiGetter(), lobId)

	tools.Debug(urlApi)

	request, err := http.NewRequest("GET", urlApi, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", "Bearer "+accessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var lobDto LobDto

	err = json.Unmarshal(body, &lobDto)
	if err != nil {
		return nil, err
	}

	lob = &domain.Lob{
		ID:       lobDto.ID,
		Name:     lobDto.Name,
		LeagueId: lobDto.League.ID,
	}

	return lob, nil
}
