package presenters

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines(periodFiscal *domain.Period) (*domain.TimeInput, error) {

	pivotDate := time.Now()
	timeInput, err := dataproviders.TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 50, infrastructure.GlobalPurposeProjectsSingletonGetter())
	if err != nil {
		return nil, err
	}
	timeInput = timeInput.TimeInputEnricher(periodFiscal, pivotDate)

	synthesisLines := timeInput.TimeInputAggregator(pivotDate)

	sort.Sort(domain.ByAssending(synthesisLines))

	sl := domain.SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
	return timeInput, nil
}
