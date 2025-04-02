package ui

import (
	"Octoptimist/presenters"
	"Octoptimist/tools"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginGET(w, r)
	}
	if r.Method == "POST" {
		loginPOST(w, r)
	}
}

func loginGET(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("ui/login.html"))
	infos := presenters.LoginInfos{Debug: os.Getenv("DEBUG")}
	if r.URL.Query().Get("code") != "" {
		infos.Datas.AuthCode = r.URL.Query().Get("code")
		err := infos.ManageToken()
		if err != nil {
			http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
			return
		}
		if infos.Datas.AuthCode != "" {
			tools.StoreAPITokenInJWT(w, infos.AccessToken)
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s", presenters.MODE_CONNEXION_AUTH), http.StatusTemporaryRedirect)
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
		case "sc":
			infos.Error = "Error during synthesis process, need to reconnect"
		}

		state = false
	}
	if state {

		if len(r.Form["btnAuth"]) > 0 {

			err := infos.ManageToken()
			if err != nil {
				http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
				return
			}
			if infos.Datas.AuthCode != "" {
				tools.StoreAPITokenInJWT(w, infos.AccessToken)
				http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s", presenters.MODE_CONNEXION_AUTH), http.StatusFound)
				return
			}
		}
		if len(r.Form["btnId"]) > 0 {
			err := infos.ManageToken()
			if err != nil {
				http.Redirect(w, r, "/loginform?err=tk", http.StatusTemporaryRedirect)
				return
			}
			tools.StoreAPITokenInJWT(w, infos.AccessToken)
			http.Redirect(w, r, fmt.Sprintf("/synthesis?mode=%s&id=%s", presenters.MODE_CONNEXION_ID, infos.Datas.Id), http.StatusTemporaryRedirect)
			return
		}
		if len(r.Form["btnGoogle"]) > 0 {
			http.Redirect(w, r, fmt.Sprintf(tools.OctopodDomainGetter()+"/api/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", os.Getenv("CLIENT_ID"), os.Getenv("REDIRECT_URL")), http.StatusFound)
			return
		}
	}
	t, _ := template.ParseFiles("ui/login.html")

	t.Execute(w, infos)
}

func validateLoginParameters(r *http.Request) (presenters.LoginInfos, bool) {
	r.ParseForm()
	state := true
	infos := presenters.LoginInfos{Debug: os.Getenv("DEBUG")}

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
