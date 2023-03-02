package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TargetTace_For_Confirmed_ShouldBe_85(t *testing.T) {
	targetTaces := domain.TargetTaces{Loader: usecases.MockTargetTaceLoader}

	tace, _ := targetTaces.GetTargetTaceForJobId(50)

	assert.Equal(t, 85, tace, "Target Tace should be 85 for a confirmed consultant, job id 50")
}

func Test_TargetTace_State_For_Confirmed_ShouldBe_True(t *testing.T) {
	targetTaces := domain.TargetTaces{Loader: usecases.MockTargetTaceLoader}

	_, state := targetTaces.GetTargetTaceForJobId(50)

	assert.True(t, state, "Target Tace state should be true for a confirmed consultant, job id 50")
}

func Test_TargetTace_State_For_Consultant_ShouldBe_False(t *testing.T) {
	targetTaces := domain.TargetTaces{Loader: usecases.MockTargetTaceLoader}

	_, state := targetTaces.GetTargetTaceForJobId(49)

	assert.False(t, state, "Target Tace state should be false for a consultant, job id 49")
}
