package main

import (
	"log"
	"os"
	"sync"
)

var lockPeoples = &sync.Mutex{}

var peoplesGlobalMap *PeoplesGlobalMap

type PeoplesGlobalMap struct {
	Loader    func(string) (map[string]People, error)
	PeopleMap map[string]People
}

func CreatePeoplesGlobalMap(accessToken string) {
	peoplesGlobalMap = &PeoplesGlobalMap{Loader: PeoplesGetter}
}

func GetPeoplesGlobalMapInstance() *PeoplesGlobalMap {
	var token *Token = nil
	if peoplesGlobalMap == nil {
		token, _ = TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
		CreatePeoplesGlobalMap(token.AccessToken)
	}
	if token == nil {
		token, _ = TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), "", "")
	}

	peoplesGlobalMap.Init(token.AccessToken)
	return peoplesGlobalMap
}

func (peoples *PeoplesGlobalMap) Init(accessToken string) {
	if peoples.PeopleMap == nil {
		lockPeoples.Lock()
		defer lockPeoples.Unlock()
		if peoples.PeopleMap == nil {
			var err error
			peoples.PeopleMap, err = peoples.Loader(accessToken)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
