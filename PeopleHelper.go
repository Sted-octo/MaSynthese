package main

func PeopleJsonGetter() string {
	return `{
		"id": 123456789,
		"last_name": "GOLDMAN",
		"first_name": "Jean",
		"nickname": "JEGO",
		"url": "https://octopod.fake/api/v0/people/123456789",
		"email": "jean.goldman@octo-fake.com",
		"job": {
		  "id": 50,
		  "name": "Confirmé",
		  "url": "https://octopod.fake/api/v0/jobs/50",
		  "created_at": "2015-02-20T19:34:35Z",
		  "updated_at": "2015-03-10T22:26:05Z",
		  "active": true,
		  "appointable": true,
		  "natural_order": 2
		},
		"lob": {
		  "id": 117,
		  "abbreviation": "CHTI",
		  "url": "https://octopod.fake/api/v0/lobs/117",
		  "name": "OCTO Nord 🍺",
		  "turnover_type": "from_internal_team",
		  "active": true,
		  "in_activity_rate": true,
		  "timesheet_email_alert_enabled": true,
		  "created_at": "2018-01-08T19:09:48Z",
		  "updated_at": "2021-03-12T11:04:46Z",
		  "subsidiary": {
			"id": 1,
			"name": "OCTO France",
			"region": "fr",
			"chrono_prefix": "FR",
			"locale": "fr",
			"timezone": "Europe/Paris",
			"active": true,
			"url": "https://octopod.fake/api/v0/subsidiaries/1",
			"currency": {
			  "symbol": "€"
			}
		  },
		  "league": {
			"id": 44,
			"name": "CAMPEMENT"
		  }
		},
		"work_schedule": {
		  "id": 1,
		  "name": "FR - Forfait Jours"
		},
		"included_in_activity_rate": true,
		"created_at": "2022-02-28T11:43:04Z",
		"updated_at": "2022-06-23T07:08:43Z",
		"entry_date": "2022-03-01",
		"leaving_date": null,
		"active_roles": [
		  {
			"id": 8,
			"name": "Dev Back"
		  },
		  {
			"id": 9,
			"name": "Dev Front"
		  }
		],
		"lcr": 700.0,
		"staffing_info": "J’ai une solide expérience.",
		"mandatory_main_paid_leave_fulfilled": true
	  }`
}

func PeoplesJsonGetter() string {
	return `[
		{
			"id": 123,
			"last_name": "MORGAN",
			"first_name": "Dexter",
			"nickname": "MODE",
			"url": "https://octopod.octo.com/api/v0/people/123",
			"email": "dexter.morgan@octo.com",
		"job": {
				"id": 51,
				"name": "Senior",
				"url": "https://octopod.octo.com/api/v0/jobs/51"
			},
			"lob": {
				"id": 385,
				"abbreviation": "PRODUCT-STACK",
				"url": "https://octopod.octo.com/api/v0/lobs/385"
			},
			"work_schedule": {
				"id": 12,
				"name": "FR - Forfait 40 Heures"
			},
			"included_in_activity_rate": true,
			"created_at": "2022-03-01T14:42:12Z",
			"updated_at": "2022-07-04T12:56:20Z",
			"entry_date": "2022-04-01",
			"leaving_date": null,
			"active_roles": [
				{
					"id": 8,
					"name": "Dev Back"
				},
				{
					"id": 10,
					"name": "Dev FullStack"
				},
				{
					"id": 18,
					"name": "Tech Lead"
				}
			]
		},
		{
			"id": 456,
			"last_name": "MORGAN",
			"first_name": "Debra",
			"nickname": "DEMO",
			"url": "https://octopod.octo.com/api/v0/people/456",
			"email": "debra.morgan@octo.com",
			"job": {
				"id": 49,
				"name": "Consultant",
				"url": "https://octopod.octo.com/api/v0/jobs/49"
			},
			"lob": {
				"id": 167,
				"abbreviation": "INPL",
				"url": "https://octopod.octo.com/api/v0/lobs/167"
			},
			"work_schedule": {
				"id": 1,
				"name": "FR - Forfait Jours"
			},
			"included_in_activity_rate": true,
			"created_at": "2020-05-25T08:17:09Z",
			"updated_at": "2022-04-19T08:56:08Z",
			"entry_date": "2020-06-08",
			"leaving_date": null,
			"active_roles": [
				{
					"id": 5,
					"name": "Data Engineer"
				},
				{
					"id": 1,
					"name": "Architecte"
				},
				{
					"id": 18,
					"name": "Tech Lead"
				},
				{
					"id": 13,
					"name": "Expert"
				},
				{
					"id": 21,
					"name": "Consultant Conseil"
				}
			]
		}
	]`
}
