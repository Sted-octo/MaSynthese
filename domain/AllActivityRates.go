package domain

type AllActivityRates struct {
	OctopodFiscalYearActivityRate              ActivityRate
	RecalculatedFiscalYearActivityRate         ActivityRate
	RecalculatedPeriodActivityRate             ActivityRate
	RecalculatedPeriodWithDiscountActivityRate ActivityRate
	OptimistActivityRate                       ActivityRate
	OptimistWithDiscountActivityRate           ActivityRate
}
