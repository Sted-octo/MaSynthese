package domain

type IDiscountProject interface {
	IsDiscount(string) bool
	Init()
}

type DiscountProject struct {
	Reference string `json:"Reference"`
	Title     string `json:"Title"`
}
