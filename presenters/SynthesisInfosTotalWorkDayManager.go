package presenters

import (
	"Octoptimist/domain"
	"strconv"
)

func (infos *SynthesisInfos) manageTotalWorkDay(period *domain.Period) (int, error) {

	totalWorkDays, err := period.TotalWorkDaysGetter()
	if err != nil {
		return 0, err
	}
	infos.Datas.TotalWorkDays = strconv.Itoa(totalWorkDays)
	infos.CssClass.TotalWorkDays = "bigText"
	return totalWorkDays, nil
}
