package domain

type IGlobalPurposeProject interface {
	IsGlobalPurpose(string) bool
	Init()
}

type GlobalPurposeProject struct {
	Reference string `json:"Reference"`
	Title     string `json:"Title"`
}
