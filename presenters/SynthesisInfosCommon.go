package presenters

import (
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"Octoptimist/tools"
	"fmt"
	"time"
)

func (infos *SynthesisInfos) SynthesisCommon(periodFiscal *domain.Period, fromGetVerb bool, withDiscount bool) error {
	entryDate, err := infos.manageInfosPeople()
	if err != nil {
		return err
	}
	if fromGetVerb {
		fmt.Println(infos.Datas.Human.Quadri)
	}

	var startDay time.Time
	var endDay time.Time
	if convertedDay, err := time.Parse("2006-01-02", infos.Datas.StartDate); err == nil {
		startDay = convertedDay
	}
	if convertedDay, err := time.Parse("2006-01-02", infos.Datas.EndDate); err == nil {
		endDay = convertedDay
	}
	if startDay.Before(entryDate) && entryDate.Before(endDay) {
		startDay = entryDate
	}

	if periodFiscal.Start.Before(entryDate) && entryDate.Before(periodFiscal.End) {
		periodFiscal.Start = entryDate
	}

	periodAnalysis := domain.NewPeriod(startDay, endDay, infrastructure.BankHolidaysSingletonGetter())

	timeInputPeriod, err := infos.manageSynthesisDetailLines(periodAnalysis, withDiscount)
	if err != nil {
		return err
	}

	totalWorkDays, err := infos.manageTotalWorkDay(periodAnalysis)
	if err != nil {
		return err
	}

	infos.manageTacePeriod(timeInputPeriod, totalWorkDays)

	activityRateFY := infos.manageTaceFiscalYear(periodFiscal)

	if infos.Datas.TacePeriod == infos.Datas.TaceFiscalYear {
		infos.CssClass.TacePeriod = ""
	}

	if !tools.DatesEquals(periodAnalysis.Start, periodFiscal.Start) || !tools.DatesEquals(periodAnalysis.End, periodFiscal.End) {
		timeInputPeriod = nil
	}

	err = infos.manageTaceCustom(periodFiscal, activityRateFY.Value, timeInputPeriod)
	if err != nil {
		return err
	}
	return nil
}
