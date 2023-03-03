package presenters

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/tools"
	"fmt"
)

func (infos *SynthesisInfos) manageTaceFiscalYear(periodFiscal *domain.Period) *domain.ActivityRate {
	infos.Datas.FiscalYear = periodFiscal.End.Format("06")

	if infos.Datas.StartDate == tools.DateToString(periodFiscal.Start) &&
		infos.Datas.EndDate == tools.DateToString(periodFiscal.End) {
		infos.CssClass.TacePeriod = ""
	}

	activityRateFiscalYear, err := dataproviders.ActivityRateGetter(infos.AccessToken, infos.Datas.Id, tools.DateToString(periodFiscal.Start), tools.DateToString(periodFiscal.End))
	if err == nil {
		infos.Datas.TaceFiscalYear = fmt.Sprintf("%.2f", activityRateFiscalYear.Value*100.0)
		infos.CssClass.TaceFiscalYear = "bigText"
	}
	return activityRateFiscalYear
}