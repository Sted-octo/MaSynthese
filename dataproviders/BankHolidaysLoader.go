package dataproviders

import (
	"Octoptimist/domain"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func BankHolidaysLoader() (map[int][]domain.BankHoliday, error) {

	var err error = nil
	file, err := os.Open("./private/jours_feries_metropole.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var dayBreaks map[int][]domain.BankHoliday = make(map[int][]domain.BankHoliday)

	firstLine := true

	const dateLayout = "2006-01-02"

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
		year, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalln("conversion string to int error in bankHolidaysLoader")
		}

		holidate, _ := time.Parse(dateLayout, record[0])
		newDayBreak := domain.BankHoliday{DayDate: holidate, Year: year, Zone: record[2], Name: record[3]}

		dayBreaks[newDayBreak.Year] = append(dayBreaks[newDayBreak.Year], newDayBreak)
	}
	if err != nil {
		log.Fatalln("error loading csv file")
	}
	return dayBreaks, nil
}
