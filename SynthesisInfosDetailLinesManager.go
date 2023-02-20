package main

import (
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines(periodFiscal *Period) (*TimeInput, error) {

	pivotDate := time.Now()
	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 50)
	if err != nil {
		return nil, err
	}
	timeInput = timeInput.timeInputEnricher(periodFiscal, pivotDate)

	synthesisLines := timeInput.timeInputAggregator(pivotDate)

	sort.Sort(ByAssending(synthesisLines))

	sl := SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
	return timeInput, nil
}
