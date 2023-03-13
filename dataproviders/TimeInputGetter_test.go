package dataproviders

import (
	"Octoptimist/domain"
	"Octoptimist/usecases"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_TimeInputGetter_No_AccessToken_Should_Return_Nil(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	accessToken := ""

	_, err := TimeInputGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10", 2, &globalPurposeProjectsManager)

	if err == nil {
		t.Errorf("TimeInputGetter error should not be nil")
	}
}

func Test_TimeInputGetter_Should_Return_Ok(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFunc := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, domain.TimeInputOneJsonGetter())
		resp.Header.Add("Total", "1")
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=(\d+)&per_page=(\d+)\z`,
		responderFunc)

	accessToken := "123"

	timeInput, _ := TimeInputGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10", 2, &globalPurposeProjectsManager)

	if timeInput == nil {
		t.Errorf("TimeInputGetter error should return a new timeInput object")
	}
}

func Test_TimeInputGetter_One_Page_One_TimeInput_Count_Shouldbe_1(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFunc := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, domain.TimeInputOneJsonGetter())
		resp.Header.Add("Total", "1")
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=1\z`,
		responderFunc)
	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=2\z`,
		responderFunc)

	accessToken := "123"

	timeInput, _ := TimeInputGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10", 2, &globalPurposeProjectsManager)

	expected := 1

	if len(*timeInput) != expected {
		t.Errorf("TimeInputGetter error should return a list of %d object", expected)
	}
}

func Test_TimeInputGetter_One_Page_Two_TimeInputs_Count_Shouldbe_2(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFunc := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, domain.TimeInputTwoJsonGetter())
		resp.Header.Add("Total", "2")
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=2\z`,
		responderFunc)
	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=1\z`,
		responderFunc)

	accessToken := "123"

	timeInput, _ := TimeInputGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10", 2, &globalPurposeProjectsManager)

	expected := 2

	if len(*timeInput) != expected {
		t.Errorf("TimeInputGetter error should return a list of %d object but was %d", expected, len(*timeInput))
	}
}

func Test_TimeInputGetter_Two_Pages_Three_TimeInputs_Count_Shouldbe_3(t *testing.T) {
	globalPurposeProjectsManager := domain.GlobalPurposeProjects{Loader: usecases.MockGlobalPurposeProjectsLoader}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	responderFuncPage1 := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, domain.TimeInputOneJsonGetter())
		resp.Header.Add("Total", "3")
		return resp, nil
	}

	responderFuncPage2 := func(*http.Request) (*http.Response, error) {
		resp := httpmock.NewStringResponse(200, domain.TimeInputThreeJsonGetter())
		resp.Header.Add("Total", "3")
		return resp, nil
	}

	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=1\z`,
		responderFuncPage1)
	httpmock.RegisterResponder("GET", `=~^https://octopod\.octo\.com/api/v0/people/(\d+)/time_input\?from_date=(\d{4}-\d{2}-\d{2})&to_date=(\d{4}-\d{2}-\d{2})&page=1&per_page=3\z`,
		responderFuncPage2)

	accessToken := "123"

	timeInput, _ := TimeInputGetter(accessToken, "2142666213", "2022-03-01", "2022-03-10", 2, &globalPurposeProjectsManager)

	expected := 3

	if len(*timeInput) != expected {
		t.Errorf("TimeInputGetter error should return a list of %d object but was %d", expected, len(*timeInput))
	}
}
