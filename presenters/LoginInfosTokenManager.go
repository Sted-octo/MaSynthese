package presenters

import (
	"Octoptimist/dataproviders"
	"errors"
	"os"
)

func (infos *LoginInfos) ManageToken() error {
	if infos.AccessToken == "" {
		token, err := dataproviders.TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REDIRECT_URL"), infos.Datas.AuthCode)
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
