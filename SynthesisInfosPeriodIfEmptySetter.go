package main

func (infos *SynthesisInfos) setPeriodIfEmpty(fiscalPeriod *Period) {

	if infos.Datas.StartDate == "" && infos.Datas.EndDate == "" {
		infos.Datas.StartDate = DateToString(fiscalPeriod.Start)
		infos.Datas.EndDate = DateToString(fiscalPeriod.End)
	}
}
