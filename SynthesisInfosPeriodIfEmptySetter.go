package main

func (infos *SynthesisInfos) setPeriodIfEmpty(fiscalPeriod *Period) {

	if infos.Datas.StartDate == "" && infos.Datas.EndDate == "" {
		infos.Datas.StartDate = fiscalPeriod.Start.Format("2006-01-02")
		infos.Datas.EndDate = fiscalPeriod.End.Format("2006-01-02")
	}
}
