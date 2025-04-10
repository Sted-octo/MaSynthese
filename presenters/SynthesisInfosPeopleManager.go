package presenters

import (
	"Octoptimist/dataproviders"
	"Octoptimist/domain"
	"Octoptimist/infrastructure"
	"strconv"
	"time"
)

func (infos *SynthesisInfos) manageInfosPeople() (time.Time, error) {
	var people *domain.People
	var err error
	if infos.ModeConnexion == MODE_CONNEXION_AUTH {
		people, err = dataproviders.PeopleGetter(infos.AccessToken)
		if err != nil {
			return time.Now(), err
		}
	}
	if infos.ModeConnexion == MODE_CONNEXION_ID {
		people, err = dataproviders.PeopleByIdGetter(infos.AccessToken, infos.Datas.Id)
		if err != nil {
			return time.Now(), err
		}
	}

	if people != nil && people.ID != 0 {
		infos.Datas.Id = strconv.Itoa(int(people.ID))
		infos.Datas.Human.ID = strconv.Itoa(int(people.ID))
		infos.Datas.Human.Quadri = people.Nickname
		infos.Datas.Human.FirstName = people.FirstName
		infos.Datas.Human.LastName = people.LastName
		infos.Datas.Human.Team = people.LobAbbreviation
		infos.Datas.Human.TeamId = people.LobId
		infos.CssClass.Human.Quadri = "bigText"
		infos.CssClass.Human.Team = "bigText secondaryColor"
		infos.CssClass.Human.ID = "smallText"
		infos.CssClass.Human.EntryDate = "smallText"
		infos.Datas.Human.EntryDate = people.EntryDate

		if targetTace, ok := infrastructure.TargetTacesSingletonGetter().GetTargetTaceForJobId(int(people.JobId)); ok {
			infos.Datas.Human.JobName = people.JobName
			infos.Datas.Human.TargetTace = strconv.Itoa(targetTace)
			infos.CssClass.Human.TargetTace = "bigText"
		}
		if convertedDay, err := time.Parse("2006-01-02", people.EntryDate); err == nil {
			return convertedDay, nil
		}
	}

	return time.Now(), nil
}
