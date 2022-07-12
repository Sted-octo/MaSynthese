package main

import "fmt"

func (infos *SynthesisInfos) manageTacePeriod() {
	activityRate, err := ActivityRateGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate)
	if err == nil {
		infos.Datas.TacePeriod = fmt.Sprintf("%.2f", activityRate.Value*100.0)
		infos.CssClass.TacePeriod = "bigText"
	}
}
