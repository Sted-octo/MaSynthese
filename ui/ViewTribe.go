package ui

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"Octoptimist/presenters"
	"Octoptimist/tools"
	"Octoptimist/usecases"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func Tribe(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tribeGET(w, r)
	}
	if r.Method == "POST" {
		tribePOST(w, r)
	}
}

func tribeGET(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("ui/tribe.html"))
	infos := presenters.TribeInfos{}
	apiToken, err := tools.GetAPITokenFromJWT(r)
	if err == nil {
		infos.AccessToken = apiToken
	}

	if r.URL.Query().Get("tribe") != "" {
		infos.TribeName = r.URL.Query().Get("tribe")
	}
	if r.URL.Query().Get("id") != "" {
		infos.TribeId = r.URL.Query().Get("id")
	}

	fiscalPeriod := infos.InitFiscalPeriod()

	infos.SetPeriodIfEmpty(fiscalPeriod)

	returnInfos := manageTribeInfos(&infos, fiscalPeriod, fiscalPeriod.Start, fiscalPeriod.End)

	t.Execute(w, returnInfos)
}

func manageTribeInfos(infos *presenters.TribeInfos, fiscalPeriod *domain.Period, startDay time.Time, endDay time.Time) *presenters.TribeInfos {
	var peoplesInTribe []domain.PeopleInTribe

	lobGetter := dataproviders.LobGetter{}
	peoplesByLeagueIdGetter := dataproviders.PeoplesByLeagueIdGetter{}

	lobId, _ := strconv.ParseInt(infos.TribeId, 10, 64)

	peoplesByLobIdGetter := usecases.PeoplesByLobIdGetter{}
	peoples, err := peoplesByLobIdGetter.Get(infos.AccessToken, lobId, fiscalPeriod, &lobGetter, &peoplesByLeagueIdGetter)

	if err == nil {
		tribeManager := usecases.TribeManager{}

		peoplesInTribe = tribeManager.Manage(
			peoples,
			startDay,
			endDay,
			fiscalPeriod,
			infrastructure.BankHolidaysSingletonGetter(),
			infrastructure.GlobalPurposeProjectsSingletonGetter(),
			infrastructure.DiscountProjectsSingletonGetter(),
			infrastructure.TargetTacesSingletonGetter(),
			&dataproviders.TimeInputGetter{},
			&dataproviders.ActivityRateGetter{},
			infos.AccessToken,
			strings.ToUpper(infos.BypassNicknames))

		for _, peopleInTribe := range peoplesInTribe {
			tribeMember := presenters.TribeMember{
				FirstName: peopleInTribe.Person.FirstName,
				Name:      peopleInTribe.Person.LastName,
				Nickname:  peopleInTribe.Person.Nickname,
			}
			tribeMember.StillInTribe = true
			if peopleInTribe.Person.LeavingDate != "" {
				tribeMember.StillInTribe = false
				if leavingDate, err := time.Parse("2006-01-02", peopleInTribe.Person.LeavingDate); err == nil {
					if startDay.Before(leavingDate) && leavingDate.Before(endDay) {
						tribeMember.LeavingDate = peopleInTribe.Person.LeavingDate
					}
				}
			}

			tribeMember.JobName = peopleInTribe.Person.JobName
			if peopleInTribe.TargetTace != 0 {
				tribeMember.TargetTace = strconv.Itoa(peopleInTribe.TargetTace)
			}
			tribeMember.TotalWorkDays = strconv.Itoa(peopleInTribe.PeriodWorkDays)
			tribeMember.TacePeriod = fmt.Sprintf("%.2f", peopleInTribe.ActivityRates.RecalculatedPeriodActivityRate.Value*100.0)
			tribeMember.TaceFiscalYear = fmt.Sprintf("%.2f", peopleInTribe.ActivityRates.OctopodFiscalYearActivityRate.Value*100.0)
			tribeMember.TaceOptimist = fmt.Sprintf("%.2f", peopleInTribe.ActivityRates.OptimistActivityRate.Value*100.0)
			tribeMember.TacePeriodWithDiscount = fmt.Sprintf("%.2f", peopleInTribe.ActivityRates.RecalculatedPeriodWithDiscountActivityRate.Value*100.0)
			tribeMember.TaceOptimistWithDiscount = fmt.Sprintf("%.2f", peopleInTribe.ActivityRates.OptimistWithDiscountActivityRate.Value*100.0)

			infos.Members = append(infos.Members, tribeMember)
		}

		tribeTaceCalculator := new(usecases.TribeTaceCalculator)

		activityRates := tribeTaceCalculator.Calculate(peoplesInTribe)

		infos.Datas.TaceFiscalYear = fmt.Sprintf("%.2f", activityRates.OctopodFiscalYearActivityRate.Value*100.0)
		infos.CssClass.TaceFiscalYear = "bigText"

		infos.Datas.TaceOptimist = fmt.Sprintf("%.2f", activityRates.OptimistActivityRate.Value*100.0)
		infos.CssClass.TaceOptimist = "bigText"

		infos.Datas.TacePeriodWithDiscount = fmt.Sprintf("%.2f", activityRates.RecalculatedPeriodWithDiscountActivityRate.Value*100.0)
		infos.CssClass.TacePeriodWithDiscount = "bigText"

		infos.Datas.TacePeriod = fmt.Sprintf("%.2f", activityRates.RecalculatedPeriodActivityRate.Value*100.0)
		infos.CssClass.TacePeriod = "bigText"

		infos.Datas.TaceOptimistWithDiscount = fmt.Sprintf("%.2f", activityRates.OptimistWithDiscountActivityRate.Value*100.0)
		infos.CssClass.TaceOptimistWithDiscount = "bigText"

		infos.Datas.FiscalYear = fiscalPeriod.End.Format("06")

		sort.SliceStable(infos.Members, func(i, j int) bool {
			return infos.Members[i].JobName < infos.Members[j].JobName
		})

	}

	infos.MembersCount = strconv.Itoa(len(peoplesInTribe))

	return infos
}

func tribePOST(w http.ResponseWriter, r *http.Request) {

	infos, areParametersValid := validateTribeParameters(r)
	var returnInfos *presenters.TribeInfos
	if len(r.Form["btnExit"]) > 0 {
		err := dataproviders.TokenRevoker(infos.AccessToken)
		if err != nil {
			log.Printf("error revoking token : %s\n", err)
		}
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}

	if areParametersValid {

		fiscalPeriod := infos.InitFiscalPeriod()

		infos.SetPeriodIfEmpty(fiscalPeriod)

		if len(r.Form["btnFYPrev"]) > 0 {
			fiscalPeriod.Previous()
			infos.Datas.StartDate = tools.DateToString(fiscalPeriod.Start)
			infos.Datas.EndDate = tools.DateToString(fiscalPeriod.End)
		}

		if len(r.Form["btnFYNext"]) > 0 {
			fiscalPeriod.Next()
			infos.Datas.StartDate = tools.DateToString(fiscalPeriod.Start)
			infos.Datas.EndDate = tools.DateToString(fiscalPeriod.End)
		}

		if startDay, err := time.Parse("2006-01-02", infos.Datas.StartDate); err == nil {
			if endDay, err := time.Parse("2006-01-02", infos.Datas.EndDate); err == nil {
				returnInfos = manageTribeInfos(&infos, fiscalPeriod, startDay, endDay)
			}
		}

	}

	t, _ := template.ParseFiles("ui/tribe.html")

	t.Execute(w, returnInfos)
}

func validateTribeParameters(r *http.Request) (presenters.TribeInfos, bool) {
	r.ParseForm()
	state := true
	infos := presenters.TribeInfos{}

	apiToken, err := tools.GetAPITokenFromJWT(r)
	if err == nil {
		infos.AccessToken = apiToken
	}

	if len(r.Form["tribeName"]) > 0 {
		infos.TribeName = r.Form["tribeName"][0]
	}
	if len(r.Form["tribeId"]) > 0 {
		infos.TribeId = r.Form["tribeId"][0]
	}

	if len(r.Form["startdate"]) > 0 {
		infos.Datas.StartDate = r.Form["startdate"][0]
	}
	if len(r.Form["enddate"]) > 0 {
		infos.Datas.EndDate = r.Form["enddate"][0]
	}

	if len(r.Form["bypass"]) > 0 {
		infos.BypassNicknames = r.Form["bypass"][0]
	}

	if infos.Datas.StartDate == "" {
		infos.CssClass.StartDate = "error"
		state = false
	}
	if infos.Datas.EndDate == "" {
		infos.CssClass.EndDate = "error"
		state = false
	}

	if infos.Datas.StartDate != "" && infos.Datas.EndDate != "" {
		if startDay, err := time.Parse("2006-01-02", infos.Datas.StartDate); err == nil {
			if endDay, err := time.Parse("2006-01-02", infos.Datas.EndDate); err == nil {
				if startDay.After(endDay) {
					infos.CssClass.EndDate = "error"
					state = false
				}
			}
		}
	}

	return infos, state
}
