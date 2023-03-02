package main

import (
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"strconv"
	"time"
)

func (infos *SynthesisInfos) manageTotalWorkDay() {
	startPeriod, _ := time.Parse("2006-01-02", infos.Datas.StartDate)
	endPeriod, _ := time.Parse("2006-01-02", infos.Datas.EndDate)
	period := domain.NewPeriod(startPeriod, endPeriod, infrastructure.BankHolidaysSingletonGetter())
	totalWorkDays, err := period.TotalWorkDaysGetter()
	if err == nil {
		infos.Datas.TotalWorkDays = strconv.Itoa(totalWorkDays)
		infos.CssClass.TotalWorkDays = "bigText"
	}
}
