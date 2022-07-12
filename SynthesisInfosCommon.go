package main

func (infos *SynthesisInfos) synthesisCommon(periodFiscal *Period) {
	infos.manageInfosPeople()

	infos.manageSynthesisDetailLines()

	infos.manageTotalWorkDay()

	infos.manageTacePeriod()

	infos.manageTaceFiscalYear(periodFiscal)

	infos.manageTaceOptimist(periodFiscal)
}
