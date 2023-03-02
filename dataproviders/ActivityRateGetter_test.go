package dataproviders

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_ActivityRateGetter_No_AccessToken_Should_Return_Nil(t *testing.T) {
	accessToken := ""

	_, err := ActivityRateGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10")

	assert.Error(t, err, "ActivityRateGetter error should not be nil")
}

func Test_ActivityRateGetter_Should_Return_Ok(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFunc := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, activityRateJsonGetter())
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/activity_rate\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&include_pipe=false\z`,
		responderFunc)

	accessToken := "123"

	activityRate, _ := ActivityRateGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10")

	assert.NotNil(t, activityRate, "ActivityRateGetter should return a not nil object")
}

func Test_ActivityRateGetter_Return_Value_Should_be_Dot31(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFunc := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, activityRateJsonGetter())
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/activity_rate\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&include_pipe=false\z`,
		responderFunc)

	accessToken := "123"

	activityRate, _ := ActivityRateGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10")

	assert.Equal(t, 0.31, activityRate.Value, "ActivityRate value should be 0.31")
}
