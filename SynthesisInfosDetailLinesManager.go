package main

import (
	"Octoptimist/domain"
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines(periodFiscal *domain.Period) (*domain.TimeInput, error) {

	pivotDate := time.Now()
	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 50)
	if err != nil {
		return nil, err
	}
	timeInput = timeInput.TimeInputEnricher(periodFiscal, pivotDate)

	synthesisLines := timeInput.TimeInputAggregator(pivotDate)

	sort.Sort(ByAssending(synthesisLines))

	sl := SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
	return timeInput, nil
}
