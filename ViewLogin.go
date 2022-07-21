package main

import (
	"fmt"
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
	t := template.Must(template.ParseFiles("login.html"))
	infos := LoginInfos{Debug: os.Getenv("DEBUG")}
	if r.URL.Query().Get("code") != "" {
		infos.Datas.AuthCode = r.URL.Query().Get("code")
		err := infos.manageToken()
		if err != nil {
			http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
			return
		}
		if infos.Datas.AuthCode != "" {
			cookie := http.Cookie{
				Name:   "AccessToken",
				Value:  infos.AccessToken,
				Path:   "/",
				MaxAge: 1,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s", MODE_CONNEXION_AUTH), http.StatusTemporaryRedirect)
			return
		}
	}

	t.Execute(w, infos)
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	infos, state := validateLoginParameters(r)
	if r.URL.Query().Get("err") != "" {
		errCode := r.URL.Query().Get("err")
		switch errCode {
		case "tk":
			infos.Error = "Invalid Token, need to reconnect"
			break
		case "sc":
			infos.Error = "Error during synthesis process, need to reconnect"
			break
		}

		state = false
	}
	if state {

		if len(r.Form["btnAuth"]) > 0 {

			err := infos.manageToken()
			if err != nil {
				http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
				return
			}
			if infos.Datas.AuthCode != "" {
				cookie := http.Cookie{
					Name:   "AccessToken",
					Value:  infos.AccessToken,
					Path:   "/",
					MaxAge: 1,
				}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s", MODE_CONNEXION_AUTH), http.StatusFound)
				return
			}
		}
		if len(r.Form["btnId"]) > 0 {
			err := infos.manageToken()
			if err != nil {
				http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
				return
			}
			cookie := http.Cookie{
				Name:   "AccessToken",
				Value:  infos.AccessToken,
				Path:   "/",
				MaxAge: 1,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&id=%s", MODE_CONNEXION_ID, infos.Datas.Id), http.StatusTemporaryRedirect)
			return
		}
		if len(r.Form["btnGoogle"]) > 0 {
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
	infos := LoginInfos{Debug: os.Getenv("DEBUG")}

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
