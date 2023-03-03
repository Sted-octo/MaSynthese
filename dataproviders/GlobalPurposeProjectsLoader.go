package dataproviders

import (
	"Octoptimist/domain"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func GlobalPurposeProjectsLoader() (map[string]domain.GlobalPurposeProject, error) {

	var err error = nil
	file, err := os.Open("./private/projets_interet_general.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var globalPurposeProjectMap map[string]domain.GlobalPurposeProject = make(map[string]domain.GlobalPurposeProject)

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

		newGlobalPurposeProject := domain.GlobalPurposeProject{ProjectID: record[0], Title: record[1]}

		globalPurposeProjectMap[newGlobalPurposeProject.ProjectID] = newGlobalPurposeProject
	}
	if err != nil {
		log.Fatalln("error loading csv file")
	}
	return globalPurposeProjectMap, nil
}
