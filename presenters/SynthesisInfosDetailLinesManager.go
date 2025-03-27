package presenters

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"Octoptimist/tools"
	"sort"
	"time"
)

func (infos *SynthesisInfos) manageSynthesisDetailLines(period *domain.Period, withDiscount bool) (*domain.TimeInput, error) {

	pivotDate := time.Now()
	timeInput, err := dataproviders.TimeInputGetter(infos.AccessToken, infos.Datas.Id, tools.DateToString(period.Start), tools.DateToString(period.End), 50, infrastructure.GlobalPurposeProjectsSingletonGetter())
	if err != nil {
		return nil, err
	}

	timeInput = timeInput.TimeInputDiscountAdaptator(withDiscount, infrastructure.DiscountProjectsSingletonGetter())

	timeInput = timeInput.TimeInputEnricher(period, pivotDate)

	synthesisLines := timeInput.TimeInputAggregator(pivotDate)

	sort.Sort(domain.ByAssending(synthesisLines))

	sl := domain.SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
	return timeInput, nil
}
