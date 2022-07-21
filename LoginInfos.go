package main

type FormLoginInfo struct {
	AuthCode string
	Id       string
}

type LoginInfos struct {
	CssClass    FormLoginInfo
	Datas       FormLoginInfo
	AccessToken string
	Debug       string
	Error       string
}
