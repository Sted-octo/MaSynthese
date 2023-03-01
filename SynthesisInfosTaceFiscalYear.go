package main

import "fmt"

func (infos *SynthesisInfos) manageTaceFiscalYear(periodFiscal *Period) *ActivityRate {
	infos.Datas.FiscalYear = periodFiscal.End.Format("06")

	if infos.Datas.StartDate == DateToString(periodFiscal.Start) &&
		infos.Datas.EndDate == DateToString(periodFiscal.End) {
		infos.CssClass.TacePeriod = ""
	}

	activityRateFiscalYear, err := ActivityRateGetter(infos.AccessToken, infos.Datas.Id, DateToString(periodFiscal.Start), DateToString(periodFiscal.End))
	if err == nil {
		infos.Datas.TaceFiscalYear = fmt.Sprintf("%.2f", activityRateFiscalYear.Value*100.0)
		infos.CssClass.TaceFiscalYear = "bigText"
	}
	return activityRateFiscalYear
}
