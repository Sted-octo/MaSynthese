package main

import (
	"log"
	"os"
)

func (infos *LoginInfos) manageToken() {
	if infos.AccessToken == "" {
		if infos.Datas.AuthCode != "" {
			log.Println("TokenGetter with authCode")
		} else {
			log.Println("TokenGetter without authCode")
		}

		token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REDIRECT_URL"), infos.Datas.AuthCode)
		if err != nil {
			log.Fatal(err)
		}
		infos.AccessToken = token.AccessToken
	}
}
