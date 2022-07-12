package main

import (
	"log"
	"strconv"
)

func (infos *SynthesisInfos) manageInfosPeople() {
	var people *People
	var err error
	if infos.ModeConnexion == MODE_CONNEXION_AUTH {
		people, err = PeopleGetter(infos.AccessToken)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if infos.ModeConnexion == MODE_CONNEXION_ID {
		people, err = PeopleByIdGetter(infos.AccessToken, infos.Datas.Id)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if people != nil && people.ID != 0 {
		infos.Datas.Id = strconv.Itoa(int(people.ID))
		infos.Datas.Human.ID = strconv.Itoa(int(people.ID))
		infos.Datas.Human.Quadri = people.Nickname
		infos.Datas.Human.FirstName = people.FirstName
		infos.Datas.Human.LastName = people.LastName
		infos.Datas.Human.Team = people.Lob.Abbreviation
		infos.CssClass.Human.Quadri = "bigText"
		infos.CssClass.Human.Team = "bigText secondaryColor"
		infos.CssClass.AuthCode = "hidden"
		infos.CssClass.Human.ID = "smallText"
		infos.CssClass.Human.EntryDate = "smallText"
		infos.Datas.Human.EntryDate = people.EntryDate
	}
}
