package domain

import (
	"log"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type BankHolidays struct {
	Loader          func() (map[int][]BankHoliday, error)
	BankHolidaysMap map[int][]BankHoliday
}

func (d *BankHolidays) IsHoliday(dateToTest time.Time) bool {
	d.Init()
	year, month, day := dateToTest.Date()
	for _, holidate := range d.BankHolidaysMap[dateToTest.Year()] {

		holiYear, holiMonth, holiDay := holidate.DayDate.Date()
		if year == holiYear && month == holiMonth && day == holiDay {
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
