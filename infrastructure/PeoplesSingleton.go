package infrastructure

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"os"
)

var peoplesGlobalMap *domain.PeoplesGlobalMap

func createPeoplesGlobalMap(accessToken string) {
	peoplesGlobalMap = &domain.PeoplesGlobalMap{Loader: dataproviders.PeoplesGetter}
}

func PeoplesGlobalMapSingletonGetter() *domain.PeoplesGlobalMap {
	var token *domain.Token = nil
	var err error
	if peoplesGlobalMap == nil {
		token, err = dataproviders.TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
		if err == nil {
			createPeoplesGlobalMap(token.AccessToken)
		}
	}
	if token == nil {
		token, err = dataproviders.TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
		if err != nil {
			return nil
		}
	}

	peoplesGlobalMap.Init(token.AccessToken)
	return peoplesGlobalMap
}
