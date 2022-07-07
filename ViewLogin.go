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

type FormLoginInfo struct {
	AuthCode string
	Id       string
}

type LoginInfos struct {
	CssClass    FormLoginInfo
	Datas       FormLoginInfo
	AccessToken string
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginGET(w, r)
	}
	if r.Method == "POST" {
		loginPOST(w, r)
	}
}

func loginGET(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("login.html"))
	infos := IndexInfos{}
	t.Execute(w, infos)
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	infos, state := validateLoginParameters(r)
	if state {
		manageToken(&infos)
		if infos.Datas.AuthCode != "" {
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s", MODE_CONNEXION_AUTH, infos.AccessToken), http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s&id=%s", MODE_CONNEXION_ID, infos.AccessToken, infos.Datas.Id), http.StatusTemporaryRedirect)
		return
	}
	t, _ := template.ParseFiles("login.html")

	t.Execute(w, infos)
}

func manageToken(infos *LoginInfos) {
	if infos.AccessToken == "" {
		token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), infos.Datas.AuthCode)
		if err != nil {
			log.Fatal(err)
		}
		infos.AccessToken = token.AccessToken
	}
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

	if infos.Datas.Id == "" && infos.Datas.AuthCode == "" {
		infos.CssClass.Id = "error"
		infos.CssClass.AuthCode = "error"
		state = false
	}

	return infos, state
}
