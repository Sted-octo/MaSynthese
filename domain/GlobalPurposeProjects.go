package domain

import (
	"log"
	"sync"
)

var lockGlobalPurposeProjects = &sync.Mutex{}

type GlobalPurposeProjects struct {
	Loader                   func() (map[string]GlobalPurposeProject, error)
	GlobalPurposeProjectsMap map[string]GlobalPurposeProject
}

func (d *GlobalPurposeProjects) IsGlobalPurpose(projectID string) bool {
	d.Init()

	_, projectIdFound := d.GlobalPurposeProjectsMap[projectID]

	return projectIdFound
}

func (d *GlobalPurposeProjects) Init() {
	if d.GlobalPurposeProjectsMap == nil {
		lockGlobalPurposeProjects.Lock()
		defer lockGlobalPurposeProjects.Unlock()
		if d.GlobalPurposeProjectsMap == nil {
			var err error
			d.GlobalPurposeProjectsMap, err = d.Loader()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
