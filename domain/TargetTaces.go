package domain

import (
	"log"
	"sync"
)

var lockTargetTace = &sync.Mutex{}

type TargetTaces struct {
	Loader         func() (map[int]TargetTace, error)
	TargetTacesMap map[int]TargetTace
}

func (targetTace *TargetTaces) GetTargetTaceForJobId(jobId int) (int, bool) {

	targetTace.Init()

	if value, ok := targetTace.TargetTacesMap[jobId]; ok {
		return value.TargetTace, true
	}

	return 0, false
}

func (targetTace *TargetTaces) Init() {
	if targetTace.TargetTacesMap == nil {
		lockTargetTace.Lock()
		defer lockTargetTace.Unlock()
		if targetTace.TargetTacesMap == nil {
			var err error
			targetTace.TargetTacesMap, err = targetTace.Loader()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
