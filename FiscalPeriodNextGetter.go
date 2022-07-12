package main

func (periodFiscal *Period) Next() {
	periodFiscal.Start = periodFiscal.Start.AddDate(1, 0, 0)
	periodFiscal.End = periodFiscal.End.AddDate(1, 0, 0)
}
