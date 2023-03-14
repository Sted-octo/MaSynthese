package presenters

import "Octoptimist/domain"

func (infos *SynthesisInfos) SynthesisCommon(periodFiscal *domain.Period) error {
	err := infos.manageInfosPeople()
	if err != nil {
		return err
	}

	timeInput, err := infos.manageSynthesisDetailLines(periodFiscal)
	if err != nil {
		return err
	}

	infos.manageTotalWorkDay()

	activityRateFY := infos.manageTaceFiscalYear(periodFiscal)

	err = infos.manageTaceCustom(periodFiscal, activityRateFY.Value, timeInput)
	if err != nil {
		return err
	}
	return nil
}
