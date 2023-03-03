package usecases

import "Octoptimist/domain"

func MockGloablPurposeProjectLoader() (map[string]domain.GlobalPurposeProject, error) {
	globalPurposeProjectMap := make(map[string]domain.GlobalPurposeProject)
	globalPurposeProjectMap["1234-0001"] = domain.GlobalPurposeProject{ProjectID: "1234-0001", Title: "Projet d'intéret Général"}
	globalPurposeProjectMap["5678-0002"] = domain.GlobalPurposeProject{ProjectID: "5678-0002", Title: "Mécénat et Solidarité"}
	return globalPurposeProjectMap, nil
}
