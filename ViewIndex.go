package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"
	"time"
)

type FormInfos struct {
	Id            string
	StartDate     string
	EndDate       string
	TotalWorkDays string
}

type IndexInfos struct {
	CssClass FormInfos
	Datas    FormInfos
	Lines    []SynthesisLine
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
		var timeInput *TimeInput
		timeInput, err := TimeInputGetter(token.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 200)
		if err != nil {
			log.Fatalln(err)
		}

		synthesisLines := timeInput.timeInputAggregator()

		sort.Sort(ByAssending(synthesisLines))

		sl := SynthesisLines(synthesisLines)

		synthesisLines = sl.Accumulate()

		infos.Lines = synthesisLines

		startPeriod, _ := time.Parse("2006-01-02", infos.Datas.StartDate)
		endPeriod, _ := time.Parse("2006-01-02", infos.Datas.EndDate)
		period := NewPeriod(startPeriod, endPeriod, GetBankHolidayInstance())
		totalWorkDays, err := period.TotalWorkDaysGetter()
		if err == nil {
			infos.Datas.TotalWorkDays = strconv.Itoa(totalWorkDays)
			infos.CssClass.TotalWorkDays = "bigNumber"
		}
	}

	t, _ := template.ParseFiles("index.html")

	t.Execute(w, infos)
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

	if infos.Datas.Id == "" {
		infos.CssClass.Id = "error"
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
