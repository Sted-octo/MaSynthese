package domain

import (
	"log"
	"sync"
)

var lockPeoples = &sync.Mutex{}

type PeoplesGlobalMap struct {
	Loader           func(string) (map[string]People, map[string][]People, error)
	PeopleMap        map[string]People
	PeopleByTribeMap map[string][]People
}

func (peoples *PeoplesGlobalMap) Init(accessToken string) {
	if peoples.PeopleMap == nil {
		lockPeoples.Lock()
		defer lockPeoples.Unlock()
		if peoples.PeopleMap == nil {
			var err error
			peoples.PeopleMap, peoples.PeopleByTribeMap, err = peoples.Loader(accessToken)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
