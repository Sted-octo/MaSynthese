package main

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"strconv"
)

func (infos *SynthesisInfos) manageInfosPeople() error {
	var people *domain.People
	var err error
	if infos.ModeConnexion == MODE_CONNEXION_AUTH {
		people, err = dataproviders.PeopleGetter(infos.AccessToken)
		if err != nil {
			return err
		}
	}
	if infos.ModeConnexion == MODE_CONNEXION_ID {
		people, err = dataproviders.PeopleByIdGetter(infos.AccessToken, infos.Datas.Id)
		if err != nil {
			return err
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

		if targetTace, ok := infrastructure.TargetTacesSingletonGetter().GetTargetTaceForJobId(int(people.Job.ID)); ok {
			infos.Datas.Human.JobName = people.Job.Name
			infos.Datas.Human.TargetTace = strconv.Itoa(targetTace)
			infos.CssClass.Human.TargetTace = "bigText"
		}
	}
	return nil
}
