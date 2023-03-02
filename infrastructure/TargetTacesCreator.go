package infrastructure

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
)

var targetTaces *domain.TargetTaces

func createTargetTaces() {
	targetTaces = &domain.TargetTaces{Loader: dataproviders.TargetTaceLoader}
}

func TargetTacesSingletonGetter() *domain.TargetTaces {
	if targetTaces == nil {
		createTargetTaces()
	}
	targetTaces.Init()
	return targetTaces
}
