package infrastructure

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
)

var DiscountProjects *domain.DiscountProjects

func createDiscountProjects() {
	DiscountProjects = &domain.DiscountProjects{Loader: dataproviders.DiscountProjectsLoader}
}

func DiscountProjectsSingletonGetter() *domain.DiscountProjects {
	if DiscountProjects == nil {
		createDiscountProjects()
	}
	DiscountProjects.Init()
	return DiscountProjects
}
