package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"text/template"
)

type FormInfos struct {
	Id        string
	StartDate string
	EndDate   string
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
		timeInput, err := TimeInputGetter(token.AccessToken, infos.Datas.Id, infos.Datas.StartDate, infos.Datas.EndDate, 100)
		if err != nil {
			log.Fatalln(err)
		}

		synthesisLines := timeInput.timeInputAggregator()

		sort.Sort(ByAssending(synthesisLines))
		infos.Lines = synthesisLines

		var cumul float64 = 0
		currentKind := ""

		for _, synthesisLine := range synthesisLines {
			if currentKind != "" && currentKind != synthesisLine.Kind {
				fmt.Printf("Total %s : %f\n", currentKind, cumul)
				cumul = 0.0
			}
			fmt.Printf("%s %s %s %s %f\n", synthesisLine.Kind, synthesisLine.Title, synthesisLine.Reference, synthesisLine.CustomerName, synthesisLine.TimeSum)
			cumul += synthesisLine.TimeSum
			currentKind = synthesisLine.Kind
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
