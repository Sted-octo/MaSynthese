package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsGlobalPurpose_For_Missing_ProjectID__Should_return_false(t *testing.T) {
	globalPurposeProjects := domain.GlobalPurposeProjects{Loader: usecases.MockGloablPurposeProjectLoader}

	isGlobalPurpose := globalPurposeProjects.IsGlobalPurpose("0000")

	assert.False(t, isGlobalPurpose, "Project not in list should return global purpose false")
}

func Test_IsGlobalPurpose_For_Existing_ProjectID__Should_return_true(t *testing.T) {
	globalPurposeProjects := domain.GlobalPurposeProjects{Loader: usecases.MockGloablPurposeProjectLoader}

	isGlobalPurpose := globalPurposeProjects.IsGlobalPurpose("1234-0001")

	assert.True(t, isGlobalPurpose, "Project in list should return global purpose trus")
}
