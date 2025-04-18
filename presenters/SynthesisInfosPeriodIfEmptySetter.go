package presenters

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
)

func (infos *SynthesisInfos) SetPeriodIfEmpty(fiscalPeriod *domain.Period) {

	if infos.Datas.StartDate == "" && infos.Datas.EndDate == "" {
		infos.Datas.StartDate = tools.DateToString(fiscalPeriod.Start)
		infos.Datas.EndDate = tools.DateToString(fiscalPeriod.End)
	}
}

func (infos *TribeInfos) SetPeriodIfEmpty(fiscalPeriod *domain.Period) {

	if infos.Datas.StartDate == "" && infos.Datas.EndDate == "" {
		infos.Datas.StartDate = tools.DateToString(fiscalPeriod.Start)
		infos.Datas.EndDate = tools.DateToString(fiscalPeriod.End)
	}
	infos.Datas.FiscalYear = fiscalPeriod.End.Format("06")
}
