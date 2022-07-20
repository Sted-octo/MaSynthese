package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TargetTace_For_Confirmed_ShouldBe_85(t *testing.T) {
	targetTaces := TargetTaces{Loader: mockTargetTaceLoader}

	tace, _ := targetTaces.GetTargetTaceForJobId(50)

	assert.Equal(t, 85, tace, "Target Tace should be 85 for a confirmed consultant, job id 50")
}

func Test_TargetTace_State_For_Confirmed_ShouldBe_True(t *testing.T) {
	targetTaces := TargetTaces{Loader: mockTargetTaceLoader}

	_, state := targetTaces.GetTargetTaceForJobId(50)

	assert.True(t, state, "Target Tace state should be true for a confirmed consultant, job id 50")
}

func Test_TargetTace_State_For_Consultant_ShouldBe_False(t *testing.T) {
	targetTaces := TargetTaces{Loader: mockTargetTaceLoader}

	_, state := targetTaces.GetTargetTaceForJobId(49)

	assert.False(t, state, "Target Tace state should be false for a consultant, job id 49")
}

func mockTargetTaceLoader() (map[int]TargetTace, error) {
	targetTaceMap := make(map[int]TargetTace)
	targetTaceMap[50] = TargetTace{ID: 50, Name: "Confirmé", TargetTace: 85}
	targetTaceMap[51] = TargetTace{ID: 51, Name: "Sénior", TargetTace: 80}
	return targetTaceMap, nil
}
