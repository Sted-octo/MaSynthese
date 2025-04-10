package dataproviders

import (
	"Octoptimist/tools"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_PeoplesByLeagueIdGetter_Should_Return_err_Nil(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^`+tools.OctopodUrlApiGetter()+`/people\?all=true&league_id=(\d+)\z`,
		httpmock.NewStringResponder(200, PeoplesByLeagueJsonGetter()))

	accessToken := "123"
	var leagueId int64 = 44

	peopleByLeagueIdGetter := PeoplesByLeagueIdGetter{}

	_, err := peopleByLeagueIdGetter.Get(accessToken, leagueId)

	assert.Nil(t, err, "PeoplesByLeagueIdGetter should return err nil")
}

func Test_PeoplesByLeagueIdGetter_Should_Return_7_Peoples_When_Mock_Is_Correct(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^`+tools.OctopodUrlApiGetter()+`/people\?all=true&league_id=(\d+)\z`,
		httpmock.NewStringResponder(200, PeoplesByLeagueJsonGetter()))

	accessToken := "123"
	var leagueId int64 = 44

	peopleByLeagueIdGetter := PeoplesByLeagueIdGetter{}

	peoples, _ := peopleByLeagueIdGetter.Get(accessToken, leagueId)

	assert.Equal(t, 7, len(peoples), "PeoplesByLeagueIdGetter should return 7 peoples when mock is correct")
}
