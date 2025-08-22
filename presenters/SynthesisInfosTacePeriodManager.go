package presenters

import (
	"Octoptimist/domain"
	"fmt"
	"time"
)

func (infos *SynthesisInfos) manageTacePeriod(timeInput *domain.TimeInput, totalWorkDays int) {
	pivotDate := time.Now()
	activityRatePeriod, _, _ := timeInput.ActivityRateCalculator(pivotDate, totalWorkDays)

	infos.Datas.TacePeriod = fmt.Sprintf("%.2f", activityRatePeriod.Value*100.0)
	infos.CssClass.TacePeriod = "bigText"

}
