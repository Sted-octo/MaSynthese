package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Kind_Billable_ShouldBe_Facturable(t *testing.T) {
	kind := KIND_BILLABLE
	expexted := "Facturable"

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}

func Test_Kind_NotBillable_ShouldBe_NonFacturable(t *testing.T) {
	kind := KIND_NOT_BILLABLE
	expexted := "Non facturable"

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}

func Test_Kind_Unknown_ShouldBe_TheSame(t *testing.T) {
	kind := "Unkown"
	expexted := kind

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}

func Test_Kind_Permanent_ShouldBe_ActiviteeParmanente(t *testing.T) {
	kind := KIND_PERMANENT
	expexted := "Activitées permanentes"

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}

func Test_Kind_Absence_ShouldBe_Absences(t *testing.T) {
	kind := KIND_ABSENCE
	expexted := "Absences"

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}

func Test_Kind_Internal_ShouldBe_Interne(t *testing.T) {
	kind := KIND_INTERNAL
	expexted := "Interne"

	translated := KindTranslator(kind)

	assert.Equal(t, expexted, translated, fmt.Sprintf("%s should be translated as %s, but was %s", kind, expexted, translated))
}
