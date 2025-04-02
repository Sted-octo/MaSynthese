package domain

type ITimeInputGetter interface {
	Get(acessToken string, peopleId string, beginPeriod string, endPeriod string, resultPerPage uint, globalPurposeProjectsManager *GlobalPurposeProjects) (*TimeInput, error)
}
