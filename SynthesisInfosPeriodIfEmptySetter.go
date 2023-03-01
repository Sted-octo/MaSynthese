package main

import "Octoptimist/tools"

func (infos *SynthesisInfos) setPeriodIfEmpty(fiscalPeriod *Period) {

	if infos.Datas.StartDate == "" && infos.Datas.EndDate == "" {
		infos.Datas.StartDate = tools.DateToString(fiscalPeriod.Start)
		infos.Datas.EndDate = tools.DateToString(fiscalPeriod.End)
	}
}
