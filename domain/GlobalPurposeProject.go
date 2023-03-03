package domain

type IGlobalPurposeProject interface {
	IsGlobalPurpose(string) bool
	Init()
}

type GlobalPurposeProject struct {
	ProjectID string `json:"ProjectID"`
	Title     string `json:"Title"`
}
