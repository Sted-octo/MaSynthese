package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

var bankHolidays *BankHolidays

type BankHolidays struct {
	Loader          func() (map[int][]BankHoliday, error)
	BankHolidaysMap map[int][]BankHoliday
}

func CreateBankHolydays() {
	bankHolidays = &BankHolidays{Loader: bankHolidayLoader}
}

func GetBankHolidayInstance() *BankHolidays {
	if bankHolidays == nil {
		CreateBankHolydays()
	}
	bankHolidays.Init()
	return bankHolidays
}

func (d *BankHolidays) IsHoliday(dateToTest time.Time) bool {
	d.Init()

	for _, holidate := range d.BankHolidaysMap[dateToTest.Year()] {
		if datesEquals(dateToTest, holidate.DayDate) {
			return true
		}
	}

	return false
}

func (d *BankHolidays) Init() {
	if d.BankHolidaysMap == nil {
		lock.Lock()
		defer lock.Unlock()
		if d.BankHolidaysMap == nil {
			var err error
			d.BankHolidaysMap, err = d.Loader()
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func bankHolidayLoader() (map[int][]BankHoliday, error) {

	var err error = nil
	file, err := os.Open("./private/jours_feries_metropole.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)

	var dayBreaks map[int][]BankHoliday = make(map[int][]BankHoliday)

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
			log.Fatalln("conversion string to int error in bankHolidayLoader")
		}

		holidate, _ := time.Parse(dateLayout, record[0])
		newDayBreak := BankHoliday{DayDate: holidate, Year: year, Zone: record[2], Name: record[3]}

		dayBreaks[newDayBreak.Year] = append(dayBreaks[newDayBreak.Year], newDayBreak)
	}
	if err != nil {
		log.Fatalln("error loading csv file")
	}
	return dayBreaks, nil
}
