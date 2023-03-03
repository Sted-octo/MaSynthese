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
	if peoplesGlobalMap == nil {
		token, _ = dataproviders.TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
		createPeoplesGlobalMap(token.AccessToken)
	}
	if token == nil {
		token, _ = dataproviders.TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
	}

	peoplesGlobalMap.Init(token.AccessToken)
	return peoplesGlobalMap
}
