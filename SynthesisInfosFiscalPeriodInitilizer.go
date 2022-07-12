package main

import "time"

func (infos *SynthesisInfos) initFiscalPeriod() *Period {
	day := time.Now()
	if !(infos.Datas.StartDate == "" && infos.Datas.EndDate == "") {
		if convertedDay, err := time.Parse("2006-01-02", infos.Datas.StartDate); err == nil {
			day = convertedDay
		} else {
			if convertedDay, err := time.Parse("2006-01-02", infos.Datas.EndDate); err == nil {
				day = convertedDay
			}
		}
	}

	fiscalPeriod := FiscalPeriodGetter(day, GetBankHolidayInstance())

	return fiscalPeriod
}
