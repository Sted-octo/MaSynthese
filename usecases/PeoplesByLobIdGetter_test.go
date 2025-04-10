package usecases

import (
	"Octoptimist/domain"
	"Octoptimist/tools"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PeoplesByLobIdGetter_Should_Empty_List_When_No_People_Are_In_League_And_Lob(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	lobGetter := &LobGetterMock{}
	peoplesByLeagueIdGetter := &PeoplesByLeagueIdGetterMockOnePeopleLob1{}

	start := tools.DateSimple(2022, time.March, 18)
	end := tools.DateSimple(2022, time.March, 17)
	period := domain.NewPeriod(start, end, &bankHolidays)

	var lobId int64 = 44
	accessToken := "123"

	peoplesByLobIdGetter := PeoplesByLobIdGetter{}
	peoples, _ := peoplesByLobIdGetter.Get(accessToken, lobId, period, lobGetter, peoplesByLeagueIdGetter)

	assert.Equal(t, 0, len(peoples), "People list count should be 0 when no people are in league and lob")
}

func Test_PeoplesByLobIdGetter_Should_One_People_When_One_People_Are_In_League_And_Lob(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	lobGetter := &LobGetterMock{}
	peoplesByLeagueIdGetter := &PeoplesByLeagueIdGetterMockOnePeopleLob44{}

	start := tools.DateSimple(2023, time.January, 1)
	end := tools.DateSimple(2023, time.December, 31)
	period := domain.NewPeriod(start, end, &bankHolidays)

	var lobId int64 = 44
	accessToken := "123"

	peoplesByLobIdGetter := PeoplesByLobIdGetter{}
	peoples, _ := peoplesByLobIdGetter.Get(accessToken, lobId, period, lobGetter, peoplesByLeagueIdGetter)

	assert.Equal(t, 1, len(peoples), "People list count should be 1 when one people is in league and lob")
}

func Test_PeoplesByLobIdGetter_Should_Empty_List_When_One_People_Are_In_League_And_Lob_With_EntryDate_After_End_Period(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	lobGetter := &LobGetterMock{}
	peoplesByLeagueIdGetter := &PeoplesByLeagueIdGetterMockOnePeopleLob44{}

	start := tools.DateSimple(2022, time.January, 1)
	end := tools.DateSimple(2022, time.December, 31)
	period := domain.NewPeriod(start, end, &bankHolidays)

	var lobId int64 = 44
	accessToken := "123"

	peoplesByLobIdGetter := PeoplesByLobIdGetter{}
	peoples, _ := peoplesByLobIdGetter.Get(accessToken, lobId, period, lobGetter, peoplesByLeagueIdGetter)

	assert.Equal(t, 0, len(peoples), "People list count should be 0 when one people is in league and lob but entry date after end period")
}

func Test_PeoplesByLobIdGetter_Should_Empty_List_When_One_People_Are_In_League_And_Lob_With_LeavingDate_Beforer_Start_Period(t *testing.T) {
	bankHolidays := domain.BankHolidays{Loader: MockBankHolidaysLoader}
	lobGetter := &LobGetterMock{}
	peoplesByLeagueIdGetter := &PeoplesByLeagueIdGetterMockOnePeopleLob44{}

	start := tools.DateSimple(2024, time.January, 1)
	end := tools.DateSimple(2024, time.December, 31)
	period := domain.NewPeriod(start, end, &bankHolidays)

	var lobId int64 = 44
	accessToken := "123"

	peoplesByLobIdGetter := PeoplesByLobIdGetter{}
	peoples, _ := peoplesByLobIdGetter.Get(accessToken, lobId, period, lobGetter, peoplesByLeagueIdGetter)

	assert.Equal(t, 0, len(peoples), "People list count should be 0 when one people is in league and lob but leaving date beafore start period")
}
