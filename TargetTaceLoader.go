package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

var lockTargetTace = &sync.Mutex{}

var targetTaces *TargetTaces

type TargetTaces struct {
	Loader         func() (map[int]TargetTace, error)
	TargetTacesMap map[int]TargetTace
}

func CreateTargetTaces() {
	targetTaces = &TargetTaces{Loader: targetTaceLoader}
}

func GetTargetTacesInstance() *TargetTaces {
	if targetTaces == nil {
		CreateTargetTaces()
	}
	targetTaces.Init()
	return targetTaces
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

func targetTaceLoader() (map[int]TargetTace, error) {

	var err error = nil
	file, err := os.Open("./private/tace_objectif.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var targetTace map[int]TargetTace = make(map[int]TargetTace)

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

		targetTace[id] = TargetTace{ID: id, Name: record[1], TargetTace: tace}

	}
	if err != nil {
		log.Fatalln("error loading csv file tace objectif")
	}
	return targetTace, nil
}
