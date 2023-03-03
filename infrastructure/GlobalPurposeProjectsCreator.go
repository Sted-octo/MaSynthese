package infrastructure

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
)

var globalPurposeProjects *domain.GlobalPurposeProjects

func createglobalPurposeProjects() {
	globalPurposeProjects = &domain.GlobalPurposeProjects{Loader: dataproviders.GlobalPurposeProjectsLoader}
}

func GlobalPurposeProjectsSingletonGetter() *domain.GlobalPurposeProjects {
	if globalPurposeProjects == nil {
		createglobalPurposeProjects()
	}
	globalPurposeProjects.Init()
	return globalPurposeProjects
}
