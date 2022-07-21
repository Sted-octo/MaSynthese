package main

func (infos *SynthesisInfos) synthesisCommon(periodFiscal *Period) error {
	err := infos.manageInfosPeople()
	if err != nil {
		return err
	}

	err = infos.manageSynthesisDetailLines()
	if err != nil {
		return err
	}

	infos.manageTotalWorkDay()

	infos.manageTacePeriod()

	infos.manageTaceFiscalYear(periodFiscal)

	err = infos.manageTaceOptimist(periodFiscal)
	if err != nil {
		return err
	}
	return nil
}
