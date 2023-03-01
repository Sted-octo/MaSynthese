package main

import (
	"fmt"
	"time"
)

func (infos *SynthesisInfos) manageTaceCustom(periodFiscal *Period, activityRateFY float64, timeInput *TimeInput) error {

	if infos.Datas.Human.EntryDate != "" {
		if startDay, err := time.Parse("2006-01-02", infos.Datas.Human.EntryDate); err == nil {
			if startDay.After(periodFiscal.Start) && startDay.Before(periodFiscal.End) {
				periodFiscal.Start = startDay
			}
		}
	}

	pivotDate := time.Now()
	var err error
	if timeInput == nil {
		timeInput, err = TimeInputGetter(infos.AccessToken, infos.Datas.Id, DateToString(periodFiscal.Start), DateToString(periodFiscal.End), 50)
		if err != nil {
			return err
		}
		timeInput = timeInput.timeInputEnricher(periodFiscal, pivotDate)
	}

	periodFiscalTotalWorkDays, err := periodFiscal.TotalWorkDaysGetter()
	if err != nil {
		return err
	}

	activityOptimistRateFiscalYear, err := timeInput.ActivityRateCalculator(pivotDate, periodFiscalTotalWorkDays)
	if err == nil {
		infos.Datas.TaceOptimist = fmt.Sprintf("%.2f", activityOptimistRateFiscalYear.Value*100.0)
		infos.CssClass.TaceOptimist = "bigText"
	}

	activityInternalRateFiscalYear, err := timeInput.ActivityRateInternalCalculator(pivotDate, periodFiscalTotalWorkDays)
	if err == nil && activityInternalRateFiscalYear.Value != activityRateFY {
		infos.Datas.TaceInternal = fmt.Sprintf("%.2f", activityInternalRateFiscalYear.Value*100.0)
		infos.CssClass.TaceInternal = "bigText"
	}

	return nil
}
