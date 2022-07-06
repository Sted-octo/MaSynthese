package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
}

type IndexInfos struct {
	CssClass    FormInfos
	Datas       FormInfos
	AccessToken string
	Lines       []SynthesisLine
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
	r.ParseForm()

	infos, state := validateParameters(r)

	if state {
		manageToken(&infos)

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

func manageToken(infos *IndexInfos) {
	if infos.AccessToken == "" {
		token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), infos.Datas.AuthCode)
		if err != nil {
			log.Fatalln(err)
		}
		infos.AccessToken = token.AccessToken
	}
}

func manageInfosPeople(infos *IndexInfos) {
	var people *People
	var err error
	if infos.Datas.AuthCode != "" {
		people, err = PeopleGetter(infos.AccessToken)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
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
		infos.CssClass.Human.Quadri = "bigText"
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
	infos.Datas.FiscalYear = time.Now().Format("06")

	periodFiscal := FiscalPeriodGetter()

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

func validateParameters(r *http.Request) (IndexInfos, bool) {
	r.ParseForm()
	state := true
	infos := IndexInfos{}
	if len(r.Form["idOctoUser"]) > 0 {
		infos.Datas.Id = r.Form["idOctoUser"][0]
	}
	if len(r.Form["startdate"]) > 0 {
		infos.Datas.StartDate = r.Form["startdate"][0]
	}
	if len(r.Form["enddate"]) > 0 {
		infos.Datas.EndDate = r.Form["enddate"][0]
	}

	if len(r.Form["idOctoUser"]) > 0 {
		infos.Datas.AuthCode = r.Form["authCode"][0]
	}

	if len(r.Form["accessToken"]) > 0 {
		infos.AccessToken = r.Form["accessToken"][0]
	}

	if infos.Datas.Id == "" && infos.Datas.AuthCode == "" {
		infos.CssClass.Id = "error"
		infos.CssClass.AuthCode = "error"
		state = false
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
