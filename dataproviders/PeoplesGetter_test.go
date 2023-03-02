package dataproviders

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_PeoplesGetter_Should_Return_Ok(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://octopod.octo.com/api/v0/people?order_by=nickname&order=asc",
		httpmock.NewStringResponder(200, PeoplesJsonGetter()))

	accessToken := "123"

	peoples, _ := PeoplesGetter(accessToken)

	assert.NotNil(t, peoples, "PeoplesGetter should return a not nil objet")
}

func Test_PeoplesGetter_Map_Count_Should_Return_2(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://octopod.octo.com/api/v0/people?order_by=nickname&order=asc",
		httpmock.NewStringResponder(200, PeoplesJsonGetter()))

	accessToken := "123"

	peoples, _ := PeoplesGetter(accessToken)

	assert.Equal(t, 2, len(peoples), "PeoplesGetter return map length should be 2")
}

func Test_PeoplesGetter_For_Nickname_MODE_Should_Return_Firstname_Dexter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://octopod.octo.com/api/v0/people?order_by=nickname&order=asc",
		httpmock.NewStringResponder(200, PeoplesJsonGetter()))

	accessToken := "123"

	peoples, _ := PeoplesGetter(accessToken)

	assert.Equal(t, "Dexter", peoples["MODE"].FirstName, "PeoplesGetter for nickname MODE return firstname Dexter")
}

func Test_PeoplesGetter_For_Nickname_DEMO_Should_Return_Lastname_MORGAN(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://octopod.octo.com/api/v0/people?order_by=nickname&order=asc",
		httpmock.NewStringResponder(200, PeoplesJsonGetter()))

	accessToken := "123"

	peoples, _ := PeoplesGetter(accessToken)

	assert.Equal(t, "MORGAN", peoples["DEMO"].LastName, "PeoplesGetter for nickname DEMO return lastname MORGAN")
}
