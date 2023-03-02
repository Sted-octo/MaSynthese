package usecases

import "Octoptimist/domain"

func MockTargetTaceLoader() (map[int]domain.TargetTace, error) {
	targetTaceMap := make(map[int]domain.TargetTace)
	targetTaceMap[50] = domain.TargetTace{ID: 50, Name: "Confirmé", TargetTace: 85}
	targetTaceMap[51] = domain.TargetTace{ID: 51, Name: "Sénior", TargetTace: 80}
	return targetTaceMap, nil
}
