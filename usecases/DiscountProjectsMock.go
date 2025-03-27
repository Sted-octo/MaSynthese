package usecases

import "Octoptimist/domain"

func MockDiscountProjectsLoader() (map[string]domain.DiscountProject, error) {
	discountProjectMap := make(map[string]domain.DiscountProject)
	discountProjectMap["2024-0005I"] = domain.DiscountProject{Reference: "2024-0005I", Title: "Négo commerciale et autres"}
	return discountProjectMap, nil
}
