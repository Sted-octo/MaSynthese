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
	infos := LoginInfos{}
	if r.URL.Query().Get("code") != "" {
		infos.Datas.AuthCode = r.URL.Query().Get("code")
		manageToken(&infos)
		if infos.Datas.AuthCode != "" {
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s", MODE_CONNEXION_AUTH, infos.AccessToken), http.StatusTemporaryRedirect)
			return
		}
	}

	t.Execute(w, infos)
}

func loginPOST(w http.ResponseWriter, r *http.Request) {

	infos, state := validateLoginParameters(r)
	if state {
		manageToken(&infos)
		if len(r.Form["btnAuth"]) > 0 {
			if infos.Datas.AuthCode != "" {
				http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s", MODE_CONNEXION_AUTH, infos.AccessToken), http.StatusTemporaryRedirect)
				return
			}
		}
		if len(r.Form["btnId"]) > 0 {
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&code=%s&id=%s", MODE_CONNEXION_ID, infos.AccessToken, infos.Datas.Id), http.StatusTemporaryRedirect)
			return
		}
		if len(r.Form["btnGoogle"]) > 0 {
			http.Redirect(w, r, fmt.Sprintf("https://octopod.octo.com/api/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URL")), http.StatusTemporaryRedirect)
			return
		}
	}
	t, _ := template.ParseFiles("login.html")

	t.Execute(w, infos)
}

func manageToken(infos *LoginInfos) {
	if infos.AccessToken == "" {
		token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REDIRECT_URL"), infos.Datas.AuthCode)
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
