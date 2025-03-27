package domain

import (
	"log"
	"sync"
)

var lockDiscountProjects = &sync.Mutex{}

type DiscountProjects struct {
	Loader              func() (map[string]DiscountProject, error)
	DiscountProjectsMap map[string]DiscountProject
}

func (d *DiscountProjects) IsDiscount(reference string) bool {
	d.Init()

	_, projectIdFound := d.DiscountProjectsMap[reference]

	return projectIdFound
}

func (d *DiscountProjects) Init() {
	if d.DiscountProjectsMap == nil {
		lockDiscountProjects.Lock()
		defer lockDiscountProjects.Unlock()
		if d.DiscountProjectsMap == nil {
			var err error
			d.DiscountProjectsMap, err = d.Loader()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
