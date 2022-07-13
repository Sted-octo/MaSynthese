package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var MODE_CONNEXION_AUTH string = "A"
var MODE_CONNEXION_ID string = "I"
var MODE_CONNEXION_GOOGLE string = "G"

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginGET(w, r)
	}
	if r.Method == "POST" {
		loginPOST(w, r)
	}
}

func loginGET(w http.ResponseWriter, r *http.Request) {
	log.Println("loginGET")
	t := template.Must(template.ParseFiles("login.html"))
	infos := LoginInfos{}
	if r.URL.Query().Get("code") != "" {
		infos.Datas.AuthCode = r.URL.Query().Get("code")
		infos.manageToken()
		if infos.Datas.AuthCode != "" {
			log.Println("loginGET with parameter code")
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s", MODE_CONNEXION_AUTH, infos.AccessToken), http.StatusTemporaryRedirect)
			return
		}
	}

	t.Execute(w, infos)
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	log.Println("loginPOST")
	infos, state := validateLoginParameters(r)
	if state {

		if len(r.Form["btnAuth"]) > 0 {
			infos.manageToken()
			if infos.Datas.AuthCode != "" {
				log.Println("loginPOST with parameter AuthCode")
				http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s", MODE_CONNEXION_AUTH, infos.AccessToken), http.StatusFound)
				return
			}
		}
		if len(r.Form["btnId"]) > 0 {
			infos.manageToken()
			log.Println("loginPOST with parameter ID")
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s&id=%s", MODE_CONNEXION_ID, infos.AccessToken, infos.Datas.Id), http.StatusFound)
			return
		}
		if len(r.Form["btnGoogle"]) > 0 {
			log.Println("loginPOST redirect to /oauth/authorize")
			http.Redirect(w, r, fmt.Sprintf("https://octopod.octo.com/api/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URL")), http.StatusFound)
			return
		}
	}
	t, _ := template.ParseFiles("login.html")

	t.Execute(w, infos)
}

func validateLoginParameters(r *http.Request) (LoginInfos, bool) {
	r.ParseForm()
	state := true
	infos := LoginInfos{}

	if r.URL.Query().Get("code") != "" {
		infos.Datas.AuthCode = r.URL.Query().Get("code")
		return infos, state
	}

	if len(r.Form["idOctoUser"]) > 0 {
		infos.Datas.Id = r.Form["idOctoUser"][0]
	}

	if len(r.Form["authCode"]) > 0 {
		infos.Datas.AuthCode = r.Form["authCode"][0]
	}

	if len(r.Form["btnAuth"]) > 0 && infos.Datas.AuthCode == "" {
		infos.CssClass.Id = ""
		infos.CssClass.AuthCode = "error"
		state = false
	}

	if len(r.Form["btnId"]) > 0 && infos.Datas.Id == "" {
		infos.CssClass.Id = "error"
		infos.CssClass.AuthCode = ""
		state = false
	}

	if len(r.Form["btnGoogle"]) > 0 {
		infos.CssClass.Id = ""
		infos.CssClass.AuthCode = ""
	}

	return infos, state
}
