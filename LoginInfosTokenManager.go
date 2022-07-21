package main

import (
	"errors"
	"os"
)

func (infos *LoginInfos) manageToken() error {
	if infos.AccessToken == "" {
		token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REDIRECT_URL"), infos.Datas.AuthCode)
		if err != nil {
			return err
		}
		if token.AccessToken == "" {
			return errors.New("access token is empty")
		}
		infos.AccessToken = token.AccessToken
	}
	return nil
}
