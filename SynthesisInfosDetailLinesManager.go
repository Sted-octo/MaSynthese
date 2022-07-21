package main

import (
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines() error {

	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 400)
	if err != nil {
		return err
	}

	synthesisLines := timeInput.timeInputAggregator(time.Now())

	sort.Sort(ByAssending(synthesisLines))

	sl := SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
	return nil
}
