package usecases

import (
	"Octoptimist/domain"
	"time"
)

func MockBankHolidayLoader() (map[int][]domain.BankHoliday, error) {
	bankHolidayMap := make(map[int][]domain.BankHoliday)
	bankHolidayMap[2021] = append(bankHolidayMap[2021], domain.BankHoliday{DayDate: time.Date(2021, time.May, 13, 0, 0, 0, 0, timeZoneGetter("Europe/Paris")),
		Year: 2021, Zone: "Métropole", Name: "Ascension"})
	bankHolidayMap[2022] = append(bankHolidayMap[2022], domain.BankHoliday{DayDate: time.Date(2022, time.May, 26, 0, 0, 0, 0, timeZoneGetter("Europe/Paris")),
		Year: 2022, Zone: "Métropole", Name: "Ascension"})
	bankHolidayMap[2023] = append(bankHolidayMap[2023], domain.BankHoliday{DayDate: time.Date(2023, time.May, 18, 0, 0, 0, 0, timeZoneGetter("Europe/Paris")),
		Year: 2023, Zone: "Métropole", Name: "Ascension"})
	return bankHolidayMap, nil
}

func timeZoneGetter(zone string) *time.Location {
	t, _ := time.LoadLocation(zone)
	return t
}
