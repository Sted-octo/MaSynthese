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

	var err error
	if timeInput == nil {
		timeInput, err = TimeInputGetter(infos.AccessToken, infos.Datas.Id, periodFiscal.Start.Format("2006-01-02"), periodFiscal.End.Format("2006-01-02"), 50)
		if err != nil {
			return err
		}
	}
	totalWorkDays, err := periodFiscal.TotalWorkDaysGetter()
	if err != nil {
		return err
	}

	activityOptimistRateFiscalYear, err := timeInput.ActivityRateCalculator(time.Now(), totalWorkDays)
	if err == nil {
		infos.Datas.TaceOptimist = fmt.Sprintf("%.2f", activityOptimistRateFiscalYear.Value*100.0)
		infos.CssClass.TaceOptimist = "bigText"
	}

	activityInternalRateFiscalYear, err := timeInput.ActivityRateInternalCalculator(time.Now(), totalWorkDays)
	if err == nil && activityInternalRateFiscalYear.Value != activityRateFY {
		infos.Datas.TaceInternal = fmt.Sprintf("%.2f", activityInternalRateFiscalYear.Value*100.0)
		infos.CssClass.TaceInternal = "bigText"
	}
	return nil
}
