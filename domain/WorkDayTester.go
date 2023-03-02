package domain

import "time"

func IsWorkDay(date time.Time) bool {
	day := date.Weekday()
	return day == time.Monday || day == time.Tuesday || day == time.Wednesday || day == time.Thursday || day == time.Friday
}
