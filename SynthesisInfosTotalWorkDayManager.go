package main

import (
	"strconv"
	"time"
)

func (infos *SynthesisInfos) manageTotalWorkDay() {
	startPeriod, _ := time.Parse("2006-01-02", infos.Datas.StartDate)
	endPeriod, _ := time.Parse("2006-01-02", infos.Datas.EndDate)
	period := NewPeriod(startPeriod, endPeriod, GetBankHolidayInstance())
	totalWorkDays, err := period.TotalWorkDaysGetter()
	if err == nil {
		infos.Datas.TotalWorkDays = strconv.Itoa(totalWorkDays)
		infos.CssClass.TotalWorkDays = "bigText"
	}
}
