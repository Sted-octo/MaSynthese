package main

import (
	"log"
	"net/http"
)

func (infos SynthesisInfos) manageExit(r *http.Request, w http.ResponseWriter) {
	if len(r.Form["btnExit"]) > 0 {
		err := TokenRevoker(infos.AccessToken)
		if err != nil {
			log.Printf("error revoking token : %s\n", err)
		}
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}
