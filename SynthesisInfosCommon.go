package main

import "Octoptimist/domain"

func (infos *SynthesisInfos) synthesisCommon(periodFiscal *domain.Period) error {
	err := infos.manageInfosPeople()
	if err != nil {
		return err
	}

	timeInput, err := infos.manageSynthesisDetailLines(periodFiscal)
	if err != nil {
		return err
	}

	infos.manageTotalWorkDay()

	infos.manageTacePeriod()

	activityRateFY := infos.manageTaceFiscalYear(periodFiscal)

	err = infos.manageTaceCustom(periodFiscal, activityRateFY.Value, timeInput)
	if err != nil {
		return err
	}
	return nil
}
