package dataproviders

import (
	"Octoptimist/domain"
	"encoding/csv"
	"io"
	"os"
)

func DiscountProjectsLoader() (map[string]domain.DiscountProject, error) {

	var err error = nil
	file, err := os.Open("./private/projets_nego_commerciale.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var DiscountProjectMap map[string]domain.DiscountProject = make(map[string]domain.DiscountProject)

	firstLine := true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if firstLine {
			firstLine = !firstLine
			continue
		}

		newDiscountProject := domain.DiscountProject{Reference: record[0], Title: record[1]}

		DiscountProjectMap[newDiscountProject.Reference] = newDiscountProject
	}

	return DiscountProjectMap, nil
}
