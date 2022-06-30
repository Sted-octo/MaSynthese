package main

import "time"

type IBankHolyday interface {
	IsHoliday(time.Time) bool
	Init()
}

type BankHoliday struct {
	DayDate time.Time `json:"DayDate"`
	Year    int       `json:"Year"`
	Zone    string    `json:"Zone"`
	Name    string    `json:"Name"`
}
