package main

import (
	"log"
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines() {

	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 400)
	if err != nil {
		log.Fatalln(err)
	}

	synthesisLines := timeInput.timeInputAggregator(time.Now())

	sort.Sort(ByAssending(synthesisLines))

	sl := SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
}
