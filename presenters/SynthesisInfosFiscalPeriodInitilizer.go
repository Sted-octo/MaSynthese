package presenters

import (
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"Octoptimist/usecases"
	"time"
)

func (infos *SynthesisInfos) InitFiscalPeriod() *domain.Period {
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

	fiscalPeriod := usecases.FiscalPeriodGetter(day, infrastructure.BankHolidaysSingletonGetter())

	return fiscalPeriod
}

func (infos *TribeInfos) InitFiscalPeriod() *domain.Period {
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

	fiscalPeriod := usecases.FiscalPeriodGetter(day, infrastructure.BankHolidaysSingletonGetter())

	return fiscalPeriod
}
