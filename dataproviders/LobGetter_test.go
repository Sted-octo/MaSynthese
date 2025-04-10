package dataproviders

import (
	"Octoptimist/tools"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_LobGetter_Should_Return_not_nil_when_mock_is_correct(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^`+tools.OctopodUrlApiGetter()+`/lobs/(\d+)\z`,
		httpmock.NewStringResponder(200, LobJsonGetter()))

	accessToken := "123"
	var lobId int64 = 117

	lobGetter := LobGetter{}
	lob, _ := lobGetter.Get(accessToken, lobId)

	assert.NotNil(t, lob, "LobGetter should return a not nil objet")
}

func Test_LobGetter_Should_Return_Lob_Object_With_Valid_Informations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~^`+tools.OctopodUrlApiGetter()+`/lobs/(\d+)\z`,
		httpmock.NewStringResponder(200, LobJsonGetter()))

	accessToken := "123"
	var lobId int64 = 117
	lobName := "LOB_FAKE"
	var leagueId int64 = 44

	lobGetter := LobGetter{}
	lob, _ := lobGetter.Get(accessToken, lobId)

	assert.Equal(t, lob.ID, lobId, "LobGetter should return an object with a valid ID")
	assert.Equal(t, lob.Name, lobName, "LobGetter should return an object with a valid name")
	assert.Equal(t, lob.LeagueId, leagueId, "LobGetter should return an object with a valid LeagueId")
}
