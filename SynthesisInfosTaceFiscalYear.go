package main

import "fmt"

func (infos *SynthesisInfos) manageTaceFiscalYear(periodFiscal *Period) *ActivityRate {
	infos.Datas.FiscalYear = periodFiscal.End.Format("06")

	if infos.Datas.StartDate == periodFiscal.Start.Format("2006-01-02") &&
		infos.Datas.EndDate == periodFiscal.End.Format("2006-01-02") {
		infos.CssClass.TacePeriod = ""
	}

	activityRateFiscalYear, err := ActivityRateGetter(infos.AccessToken, infos.Datas.Id, periodFiscal.Start.Format("2006-01-02"), periodFiscal.End.Format("2006-01-02"))
	if err == nil {
		infos.Datas.TaceFiscalYear = fmt.Sprintf("%.2f", activityRateFiscalYear.Value*100.0)
		infos.CssClass.TaceFiscalYear = "bigText"
	}
	return activityRateFiscalYear
}
