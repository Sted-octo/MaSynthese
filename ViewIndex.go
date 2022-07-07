package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"
	"time"
)

type FormInfos struct {
	Id             string
	StartDate      string
	EndDate        string
	TotalWorkDays  string
	TacePeriod     string
	FiscalYear     string
	TaceFiscalYear string
	TaceOptimist   string
	AuthCode       string
	Human          PeopleInfos
}

type PeopleInfos struct {
	Quadri    string
	FirstName string
	LastName  string
	EntryDate string
	ID        string
	Team      string
}

type IndexInfos struct {
	CssClass      FormInfos
	Datas         FormInfos
	AccessToken   string
	Lines         []SynthesisLine
	ModeConnexion string
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		indexGET(w, r)
	}
	if r.Method == "POST" {
		indexPOST(w, r)
	}
}

func indexGET(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	infos := IndexInfos{}
	t.Execute(w, infos)
}

func indexPOST(w http.ResponseWriter, r *http.Request) {

	infos, state := validateIndexParameters(r)

	if state {

		manageInfosPeople(&infos)

		manageSynthesisDetailLines(&infos)

		manageTotalWorkDay(&infos)

		manageTacePeriod(&infos)

		manageTaceFiscalYear(&infos)

		manageTaceOptimist(&infos)
	}

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, infos)
}

func manageInfosPeople(infos *IndexInfos) {
	var people *People
	var err error
	if infos.ModeConnexion == MODE_CONNEXION_AUTH {
		people, err = PeopleGetter(infos.AccessToken)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if infos.ModeConnexion == MODE_CONNEXION_ID {
		people, err = PeopleByIdGetter(infos.AccessToken, infos.Datas.Id)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if people != nil && people.ID != 0 {
		infos.Datas.Id = strconv.Itoa(int(people.ID))
		infos.Datas.Human.ID = strconv.Itoa(int(people.ID))
		infos.Datas.Human.Quadri = people.Nickname
		infos.Datas.Human.FirstName = people.FirstName
		infos.Datas.Human.LastName = people.LastName
		infos.Datas.Human.Team = people.Lob.Abbreviation
		infos.CssClass.Human.Quadri = "bigText"
		infos.CssClass.Human.Team = "bigText secondaryColor"
		infos.CssClass.AuthCode = "hidden"
		infos.CssClass.Human.ID = "smallText"
		infos.CssClass.Human.EntryDate = "smallText"
		infos.Datas.Human.EntryDate = people.EntryDate
	}

}

func manageTaceOptimist(infos *IndexInfos) {
	periodFiscal := FiscalPeriodGetter()

	if infos.Datas.Human.EntryDate != "" {
		if startDay, err := time.Parse("2006-01-02", infos.Datas.Human.EntryDate); err == nil {
			if startDay.After(periodFiscal.Start) && startDay.Before(periodFiscal.End) {
				periodFiscal.Start = startDay
			}
		}
	}

	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, periodFiscal.Start.Format("2006-01-02"), periodFiscal.End.Format("2006-01-02"), 400)
	if err != nil {
		log.Fatalln(err)
	}
	totalWorkDays, err := periodFiscal.TotalWorkDaysGetter()
	if err != nil {
		log.Fatalln(err)
	}

	activityOptimistRateFiscalYear, err := timeInput.ActivityRateCalculator(time.Now(), totalWorkDays)
	if err == nil {
		infos.Datas.TaceOptimist = fmt.Sprintf("%.2f", activityOptimistRateFiscalYear.Value*100.0)
		infos.CssClass.TaceOptimist = "bigText"
	}
}

func manageTaceFiscalYear(infos *IndexInfos) {
	periodFiscal := FiscalPeriodGetter()
	infos.Datas.FiscalYear = periodFiscal.End.Format("06")

	if infos.Datas.StartDate == periodFiscal.Start.Format("2006-01-02") &&
		infos.Datas.EndDate == periodFiscal.End.Format("2006-01-02") {
		infos.CssClass.TacePeriod = "hidden"
	}

	activityRateFiscalYear, err := ActivityRateGetter(infos.AccessToken, infos.Datas.Id, periodFiscal.Start.Format("2006-01-02"), periodFiscal.End.Format("2006-01-02"))
	if err == nil {
		infos.Datas.TaceFiscalYear = fmt.Sprintf("%.2f", activityRateFiscalYear.Value*100.0)
		infos.CssClass.TaceFiscalYear = "bigText"
	}
}

func manageTacePeriod(infos *IndexInfos) {
	activityRate, err := ActivityRateGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate)
	if err == nil {
		infos.Datas.TacePeriod = fmt.Sprintf("%.2f", activityRate.Value*100.0)
		infos.CssClass.TacePeriod = "bigText"
	}
}

func manageTotalWorkDay(infos *IndexInfos) {
	startPeriod, _ := time.Parse("2006-01-02", infos.Datas.StartDate)
	endPeriod, _ := time.Parse("2006-01-02", infos.Datas.EndDate)
	period := NewPeriod(startPeriod, endPeriod, GetBankHolidayInstance())
	totalWorkDays, err := period.TotalWorkDaysGetter()
	if err == nil {
		infos.Datas.TotalWorkDays = strconv.Itoa(totalWorkDays)
		infos.CssClass.TotalWorkDays = "bigText"
	}
}

func manageSynthesisDetailLines(infos *IndexInfos) {

	timeInput, err := TimeInputGetter(infos.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 400)
	if err != nil {
		log.Fatalln(err)
	}

	synthesisLines := timeInput.timeInputAggregator(time.Now())

	sort.Sort(ByAssending(synthesisLines))

	sl := SynthesisLines(synthesisLines)

	synthesisLines = sl.Accumulate()

	infos.Lines = synthesisLines
}

func validateIndexParameters(r *http.Request) (IndexInfos, bool) {
	r.ParseForm()
	state := true
	infos := IndexInfos{}

	if r.URL.Query().Get("code") != "" {
		infos.AccessToken = r.URL.Query().Get("code")
	}

	if r.URL.Query().Get("id") != "" {
		infos.Datas.Id = r.URL.Query().Get("id")
	}

	if r.URL.Query().Get("mode") != "" {
		infos.ModeConnexion = r.URL.Query().Get("mode")
	}

	if infos.Datas.Id == "" && len(r.Form["idOctoUser"]) > 0 {
		infos.Datas.Id = r.Form["idOctoUser"][0]
	}
	if len(r.Form["startdate"]) > 0 {
		infos.Datas.StartDate = r.Form["startdate"][0]
	}
	if len(r.Form["enddate"]) > 0 {
		infos.Datas.EndDate = r.Form["enddate"][0]
	}

	if infos.AccessToken == "" && len(r.Form["accessToken"]) > 0 {
		infos.AccessToken = r.Form["accessToken"][0]
	}

	if infos.ModeConnexion == "" && len(r.Form["mode"]) > 0 {
		infos.ModeConnexion = r.Form["mode"][0]
	}

	if infos.Datas.StartDate == "" {
		infos.CssClass.StartDate = "error"
		state = false
	}
	if infos.Datas.EndDate == "" {
		infos.CssClass.EndDate = "error"
		state = false
	}

	return infos, state
}
