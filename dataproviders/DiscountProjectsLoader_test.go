package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsDiscount_For_Missing_Reference_Should_return_false(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: usecases.MockDiscountProjectsLoader}

	isDiscount := DiscountProjects.IsDiscount("0000")

	assert.False(t, isDiscount, "Project not in list should return discount false")
}

func Test_IsDiscount_For_Existing_Reference_Should_return_true(t *testing.T) {
	DiscountProjects := domain.DiscountProjects{Loader: usecases.MockDiscountProjectsLoader}

	isDiscount := DiscountProjects.IsDiscount("2024-0005I")

	assert.True(t, isDiscount, "Project in list should return dicount true")
}
