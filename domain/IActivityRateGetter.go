package domain

type IActivityRateGetter interface {
	Get(acessToken string, peopleId string, beginPeriod string, endPeriod string) (*ActivityRate, error)
}
