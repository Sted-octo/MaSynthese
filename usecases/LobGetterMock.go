package usecases

import "Octoptimist/domain"

type LobGetterMock struct{}

func (l LobGetterMock) Get(accessToken string, lobId int64) (*domain.Lob, error) {
	lob := &domain.Lob{
		ID:       44,
		Name:     "LA HAUT",
		LeagueId: 117,
	}
	return lob, nil
}
