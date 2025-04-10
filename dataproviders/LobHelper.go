package dataproviders

func LobJsonGetter() string {
	return `
	{
  "id": 117,
  "name": "LOB_FAKE",
  "abbreviation": "LOFA",
  "url": "https://octopod.octo.tools/api/v0/lobs/117",
  "photo_url": "",
  "active": true,
  "leaders": [
    {
      "id": 123456789,
      "last_name": "GOLDMAN",
      "first_name": "Jean-Jacques",
      "nickname": "JJG",
      "registration_number": "13671363",
      "photo_url": "",
      "url": "https://octopod.octo.tools/api/v0/people/123456789"
    },
    {
      "id": 987654321,
      "last_name": "JONES",
      "first_name": "Mickael",
      "nickname": "MIJO",
      "registration_number": "11627088",
      "photo_url": "",
      "url": "https://octopod.octo.tools/api/v0/people/987654321"
    }
  ],
  "description": "",
  "turnover_type": "from_internal_team",
  "in_activity_rate": true,
  "timesheet_email_alert_enabled": true,
  "created_at": "2018-01-08T19:09:48Z",
  "updated_at": "2023-10-18T12:30:24Z",
  "email": "tribu@octo-fake.com",
  "subsidiary": {
    "id": 1,
    "name": "OCTO France",
    "region": "fr",
    "chrono_prefix": "FR",
    "locale": "fr",
    "timezone": "Europe/Paris",
    "active": true,
    "url": "https://octopod.octo.tools/api/v0/subsidiaries/1",
    "currency": {
      "symbol": "€"
    }
  },
  "league": {
    "id": 44,
    "name": "LEAGUE_FAKE",
    "description": null,
    "photo_url": null,
    "email": null,
    "active": true
  }
}
	`
}
