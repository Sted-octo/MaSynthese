package usecases

import "Octoptimist/domain"

func MockGlobalPurposeProjectsLoader() (map[string]domain.GlobalPurposeProject, error) {
	globalPurposeProjectMap := make(map[string]domain.GlobalPurposeProject)
	globalPurposeProjectMap["1234-0001"] = domain.GlobalPurposeProject{Reference: "1234-0001", Title: "Projet d'intéret Général"}
	globalPurposeProjectMap["5678-0002"] = domain.GlobalPurposeProject{Reference: "5678-0002", Title: "Mécénat et Solidarité"}
	return globalPurposeProjectMap, nil
}
