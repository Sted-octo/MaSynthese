package dataproviders

import (
	"Octoptimist/tools"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_PeopleGetter_Should_Return_Ok(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", tools.OctopodUrlApiGetter()+"/people/me",
		httpmock.NewStringResponder(200, PeopleJsonGetter()))

	accessToken := "123"

	people, _ := PeopleGetter(accessToken)

	assert.NotNil(t, people, "PeopleGetter should return a not nil objet")
}

func Test_PeopleByIdGetter_Should_Return_Ok(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^`+tools.OctopodUrlApiGetter()+`/people/(\d+)\z`,
		httpmock.NewStringResponder(200, PeopleJsonGetter()))

	accessToken := "123"
	id := "123"

	people, _ := PeopleByIdGetter(accessToken, id)

	assert.NotNil(t, people, "PeopleGetter should return a not nil objet")
}
