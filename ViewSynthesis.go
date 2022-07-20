package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

func Synthesis(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		synthesisGET(w, r)
	}
	if r.Method == "POST" {
		synthesisPOST(w, r)
	}
}

func synthesisGET(w http.ResponseWriter, r *http.Request) {
	log.Println("synthesisGET")
	t := template.Must(template.ParseFiles("synthesis.html"))
	infos := SynthesisInfos{}
	cookie, err := r.Cookie("AccessToken")
	if err == nil {
		infos.AccessToken = cookie.Value
	}

	if r.URL.Query().Get("id") != "" {
		infos.Datas.Id = r.URL.Query().Get("id")
	}

	if r.URL.Query().Get("mode") != "" {
		infos.ModeConnexion = r.URL.Query().Get("mode")
	}

	fiscalPeriod := infos.initFiscalPeriod()

	infos.setPeriodIfEmpty(fiscalPeriod)

	infos.synthesisCommon(fiscalPeriod)

	t.Execute(w, infos)
}

func synthesisPOST(w http.ResponseWriter, r *http.Request) {
	log.Println("synthesisPOST")

	infos, state := validateSynthesisParameters(r)

	infos.manageExit(r, w)

	if state {

		fiscalPeriod := infos.initFiscalPeriod()

		infos.setPeriodIfEmpty(fiscalPeriod)

		if len(r.Form["btnFYPrev"]) > 0 {
			fiscalPeriod.Previous()
			infos.Datas.StartDate = fiscalPeriod.Start.Format("2006-01-02")
			infos.Datas.EndDate = fiscalPeriod.End.Format("2006-01-02")
		}

		if len(r.Form["btnFYNext"]) > 0 {
			fiscalPeriod.Next()
			infos.Datas.StartDate = fiscalPeriod.Start.Format("2006-01-02")
			infos.Datas.EndDate = fiscalPeriod.End.Format("2006-01-02")
		}

		infos.synthesisCommon(fiscalPeriod)
	}

	t, _ := template.ParseFiles("synthesis.html")

	t.Execute(w, infos)
}

func validateSynthesisParameters(r *http.Request) (SynthesisInfos, bool) {
	r.ParseForm()
	state := true
	infos := SynthesisInfos{}

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