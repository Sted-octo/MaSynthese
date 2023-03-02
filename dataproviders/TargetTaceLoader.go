package dataproviders

import (
	"Octoptimist/domain"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func TargetTaceLoader() (map[int]domain.TargetTace, error) {

	var err error = nil
	file, err := os.Open("./private/tace_objectif.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var targetTace map[int]domain.TargetTace = make(map[int]domain.TargetTace)

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
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalln("conversion id string to int error in targetTaceLoader")
		}

		tace, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalln("conversion tace string to int error in targetTaceLoader")
		}

		targetTace[id] = domain.TargetTace{ID: id, Name: record[1], TargetTace: tace}

	}
	if err != nil {
		log.Fatalln("error loading csv file tace objectif")
	}
	return targetTace, nil
}
