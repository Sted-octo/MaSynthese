package dataproviders

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

func PeoplesByLeagueJsonGetter() string {
	return `
[
    {
        "id": 2142666999,
        "last_name": "RUISE",
        "first_name": "Olivia",
        "nickname": "OLRU",
        "registration_number": "123456987",
        "photo_url": "",
        "url": "https://octopod.octo.tools/api/v0/people/2142666999",
        "email": "olivia.ruise@octo-fake.com",
        "phone_number": null,
        "job_alias": null,
        "job": {
            "id": 59,
            "name": "Senior Manager Biz",
            "url": "https://octopod.octo.tools/api/v0/jobs/59",
            "created_at": "2015-02-20T19:36:31Z",
            "updated_at": "2024-12-17T13:26:04Z",
            "active": true,
            "appointable": true,
            "natural_order": 9
        },
        "lob": {
            "id": 152,
            "name": "OCTO CHOCOLATINE",
            "abbreviation": "ATIN",
            "url": "https://octopod.octo.tools/api/v0/lobs/152",
            "photo_url": null,
            "active": true,
            "leaders": [],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-03-27T15:46:56Z",
            "updated_at": "2021-03-12T11:05:04Z",
            "email": "OCTO.chocolatine@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": true,
        "lcr": 1195.92,
        "manager_id": 2142664162,
        "manager": {
            "id": 2142664162,
            "last_name": "DUPONT",
            "first_name": "JEAN",
            "nickname": "JDUP",
            "registration_number": "11629592",
            "photo_url": "",
            "url": "https://octopod.octo.tools/api/v0/people/2142664162"
        },
        "in_idf": false,
        "created_at": "2025-02-13T08:09:10Z",
        "updated_at": "2025-03-25T08:15:50Z",
        "entry_date": "2025-03-01",
        "leaving_date": null,
        "description": null,
        "resume_link": null,
        "resume_link_updated_at": null,
        "mini_bio_link": null,
        "mini_bio_link_updated_at": null,
        "links": [],
        "staffing_info": null,
        "certifications": [],
        "person_roles": [
            {
                "id": 7681,
                "person_id": 2142666999,
                "main": false,
                "role": {
                    "id": 12,
                    "name": "Directeur de Mission",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2024-10-25T11:57:31Z"
                }
            }
        ]
    },
    {
        "id": 2142666920,
        "last_name": "JAMES",
        "first_name": "Morganne",
        "nickname": "MOJA",
        "registration_number": "123456789",
        "photo_url": "",
        "url": "https://octopod.octo.tools/api/v0/people/2142666920",
        "email": "morganne.james@octo-fake.com",
        "phone_number": "",
        "job_alias": null,
        "job": {
            "id": 49,
            "name": "Consultant(e)",
            "url": "https://octopod.octo.tools/api/v0/jobs/49",
            "created_at": "2015-02-20T19:34:35Z",
            "updated_at": "2024-12-17T13:35:23Z",
            "active": true,
            "appointable": true,
            "natural_order": 1
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT 🍺",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": null,
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux.chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 12,
            "name": "FR - Forfait 40 Heures"
        },
        "included_in_activity_rate": true,
        "lcr": 446.32,
        "manager_id": 2142664933,
        "manager": {
            "id": 2142664933,
            "last_name": "VANDENBROUCKE",
            "first_name": "Rémi",
            "nickname": "REVA",
            "registration_number": "11626983",
            "photo_url": null,
            "url": "https://octopod.octo.tools/api/v0/people/2142664933"
        },
        "in_idf": false,
        "created_at": "2024-08-09T10:13:29Z",
        "updated_at": "2025-03-05T13:07:04Z",
        "entry_date": "2024-10-08",
        "leaving_date": null,
        "description": null,
        "resume_link": "https://docs.google.com/presentation/d/1ORQ9L3kuKMcoR32VDb0SL0JQTXH_NWpXeaivN4UkAuM/edit?usp=drivesdk",
        "resume_link_updated_at": "2025-01-07T00:00:00Z",
        "mini_bio_link": "https://docs.google.com/presentation/d/1NTLDt3YzoBFYFpW_9Md8btH37wP_TkfQrzw4MKAGjC8/edit?usp=drivesdk",
        "mini_bio_link_updated_at": "2025-01-07T00:00:00Z",
        "links": [],
        "staffing_info": null,
        "certifications": [],
        "person_roles": [
            {
                "id": 6788,
                "person_id": 2142666920,
                "main": false,
                "role": {
                    "id": 8,
                    "name": "Dev Back",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            },
            {
                "id": 6789,
                "person_id": 2142666920,
                "main": false,
                "role": {
                    "id": 10,
                    "name": "Dev FullStack",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            },
            {
                "id": 7043,
                "person_id": 2142666920,
                "main": true,
                "role": {
                    "id": 27,
                    "name": "Développeur",
                    "primary": true,
                    "secondary": false,
                    "active": true,
                    "created_at": "2024-10-25T11:57:58Z",
                    "updated_at": "2025-01-23T09:32:30Z"
                }
            }
        ]
    },
    {
        "id": 2142666715,
        "last_name": "NEWHOME",
        "first_name": "otto",
        "nickname": "OTNE",
        "registration_number": "13671363",
        "photo_url": null,
        "url": "https://octopod.octo.tools/api/v0/people/2142666715",
        "email": "otto.newhome@octo-fake.com",
        "phone_number": null,
        "job_alias": "Manager Business Développeur Confirmé",
        "job": {
            "id": 68,
            "name": "Manager Biz Confirmé",
            "url": "https://octopod.octo.tools/api/v0/jobs/68",
            "created_at": "2019-02-19T15:31:28Z",
            "updated_at": "2024-12-17T13:26:24Z",
            "active": true,
            "appointable": true,
            "natural_order": 8
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT 🍺",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": null,
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux.chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": false,
        "lcr": 1153.2,
        "manager_id": 2142665051,
        "manager": {
            "id": 2142665051,
            "last_name": "GAULOIS",
            "first_name": "Henry",
            "nickname": "GAHE",
            "registration_number": "963658268",
            "photo_url": "",
            "url": "https://octopod.octo.tools/api/v0/people/2142665051"
        },
        "in_idf": false,
        "created_at": "2023-04-20T14:32:44Z",
        "updated_at": "2025-03-25T11:28:46Z",
        "entry_date": "2023-06-20",
        "leaving_date": null,
        "description": "",
        "resume_link": null,
        "resume_link_updated_at": null,
        "mini_bio_link": "",
        "mini_bio_link_updated_at": "2024-10-09T00:00:00Z",
        "links": [],
        "staffing_info": "",
        "certifications": [
            {
                "id": 624,
                "title": "ITIL®Continual Service Improvement Certificate",
                "certifier": {
                    "id": 11,
                    "name": "Autre"
                },
                "issue_date": "2015-09-18",
                "expiration_date": null
            },
            {
                "id": 625,
                "title": "ITIL®Foundation Certificate in IT Service Management",
                "certifier": {
                    "id": 11,
                    "name": "Autre"
                },
                "issue_date": "2014-09-18",
                "expiration_date": null
            },
            {
                "id": 626,
                "title": "ITIL® Service Strategy Certificate (ITILSS)",
                "certifier": {
                    "id": 11,
                    "name": "Autre"
                },
                "issue_date": "2015-11-18",
                "expiration_date": null
            }
        ],
        "person_roles": [
            {
                "id": 5857,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 24,
                    "name": "Product Manager",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2022-03-24T14:00:50Z",
                    "updated_at": "2025-01-23T13:30:59Z"
                }
            },
            {
                "id": 5858,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 7,
                    "name": "Delivery Manager",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2025-01-24T09:32:59Z"
                }
            },
            {
                "id": 5859,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 12,
                    "name": "Directeur de Mission",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2024-10-25T11:57:31Z"
                }
            },
            {
                "id": 5880,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 25,
                    "name": "Coach Produit",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2022-04-01T10:36:17Z",
                    "updated_at": "2022-04-01T10:36:17Z"
                }
            },
            {
                "id": 5881,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 17,
                    "name": "Product Owner / Co-PO",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2025-01-23T09:32:49Z"
                }
            },
            {
                "id": 5882,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 26,
                    "name": "Scrum Master",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2022-04-27T15:28:25Z",
                    "updated_at": "2022-04-27T15:28:25Z"
                }
            },
            {
                "id": 5883,
                "person_id": 2142666715,
                "main": false,
                "role": {
                    "id": 3,
                    "name": "Coach Agile",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            },
            {
                "id": 7348,
                "person_id": 2142666715,
                "main": true,
                "role": {
                    "id": 38,
                    "name": "Biz Dev",
                    "primary": true,
                    "secondary": false,
                    "active": true,
                    "created_at": "2024-10-25T13:08:10Z",
                    "updated_at": "2024-10-25T13:08:10Z"
                }
            }
        ]
    },
    {
        "id": 2142666622,
        "last_name": "BERNARD",
        "first_name": "Sarah",
        "nickname": "SABE",
        "registration_number": "13623474",
        "photo_url": "",
        "url": "https://octopod.octo.tools/api/v0/people/2142666622",
        "email": "sarah.bernard@octo-fake.com",
        "phone_number": "",
        "job_alias": null,
        "job": {
            "id": 51,
            "name": "Consultant(e) Senior",
            "url": "https://octopod.octo.tools/api/v0/jobs/51",
            "created_at": "2015-02-20T19:34:35Z",
            "updated_at": "2024-12-17T13:35:10Z",
            "active": true,
            "appointable": true,
            "natural_order": 3
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": "",
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux.chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": true,
        "lcr": 702.1168,
        "manager_id": 2142664834,
        "manager": {
            "id": 2142664834,
            "last_name": "RON",
            "first_name": "Pat",
            "nickname": "PARO",
            "registration_number": "741258654",
            "photo_url": "",
            "url": "https://octopod.octo.tools/api/v0/people/2142664834"
        },
        "in_idf": false,
        "created_at": "2022-12-29T10:35:05Z",
        "updated_at": "2025-01-15T12:03:31Z",
        "entry_date": "2023-01-03",
        "leaving_date": null,
        "description": "",
        "resume_link": "",
        "resume_link_updated_at": "2022-07-12T00:00:00Z",
        "mini_bio_link": "",
        "mini_bio_link_updated_at": "2024-10-31T00:00:00Z",
        "links": [],
        "staffing_info": "",
        "certifications": [],
        "person_roles": [
            {
                "id": 5453,
                "person_id": 2142666622,
                "main": false,
                "role": {
                    "id": 2,
                    "name": "Animateur/Facilitateur",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            },
            {
                "id": 5454,
                "person_id": 2142666622,
                "main": false,
                "role": {
                    "id": 26,
                    "name": "Scrum Master",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2022-04-27T15:28:25Z",
                    "updated_at": "2022-04-27T15:28:25Z"
                }
            },
            {
                "id": 5455,
                "person_id": 2142666622,
                "main": false,
                "role": {
                    "id": 3,
                    "name": "Coach Agile",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            },
            {
                "id": 7152,
                "person_id": 2142666622,
                "main": true,
                "role": {
                    "id": 17,
                    "name": "Product Owner / Co-PO",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2025-01-23T09:32:49Z"
                }
            }
        ]
    },
    {
        "id": 2142666525,
        "last_name": "TESLA",
        "first_name": "Nicolas",
        "nickname": "NITE",
        "registration_number": "963654741",
        "photo_url": null,
        "url": "https://octopod.octo.tools/api/v0/people/2142666525",
        "email": "nicolas.tesla@octo-.com",
        "phone_number": null,
        "job_alias": "Consultant(e) Senior",
        "job": {
            "id": 51,
            "name": "Consultant(e) Senior",
            "url": "https://octopod.octo.tools/api/v0/jobs/51",
            "created_at": "2015-02-20T19:34:35Z",
            "updated_at": "2024-12-17T13:35:10Z",
            "active": true,
            "appointable": true,
            "natural_order": 3
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT 🍺",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": "",
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux_chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": true,
        "lcr": 823.36,
        "manager_id": 2142664834,
        "manager": {
            "id": 2142664834,
            "last_name": "RON",
            "first_name": "Pat",
            "nickname": "PARO",
            "registration_number": "13456741",
            "photo_url": "",
            "url": "https://octopod.octo.tools/api/v0/people/2142664834"
        },
        "in_idf": false,
        "created_at": "2022-08-23T12:20:53Z",
        "updated_at": "2023-10-16T12:38:43Z",
        "entry_date": "2022-09-07",
        "leaving_date": "2022-12-15",
        "description": null,
        "resume_link": null,
        "resume_link_updated_at": null,
        "mini_bio_link": null,
        "mini_bio_link_updated_at": null,
        "links": [],
        "staffing_info": "",
        "certifications": [],
        "person_roles": [
            {
                "id": 5048,
                "person_id": 2142666525,
                "main": false,
                "role": {
                    "id": 3,
                    "name": "Coach Agile",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            }
        ]
    },
    {
        "id": 2142665570,
        "last_name": "WAVE",
        "first_name": "Paul",
        "nickname": "PAWA",
        "registration_number": "789123456",
        "photo_url": null,
        "url": "https://octopod.octo.tools/api/v0/people/2142665570",
        "email": "paul.wave@octo-fake.com",
        "phone_number": null,
        "job_alias": "Consultant(e) Senior",
        "job": {
            "id": 51,
            "name": "Consultant(e) Senior",
            "url": "https://octopod.octo.tools/api/v0/jobs/51",
            "created_at": "2015-02-20T19:34:35Z",
            "updated_at": "2024-12-17T13:35:10Z",
            "active": true,
            "appointable": true,
            "natural_order": 3
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT 🍺",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": "",
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux.chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": true,
        "lcr": 757.92,
        "manager_id": 2142664834,
        "manager": {
            "id": 2142664834,
            "last_name": "RON",
            "first_name": "Pat",
            "nickname": "PARO",
            "registration_number": "741258789",
            "photo_url": null,
            "url": "https://octopod.octo.tools/api/v0/people/2142664834"
        },
        "in_idf": null,
        "created_at": "2020-02-27T13:51:35Z",
        "updated_at": "2023-10-16T12:37:23Z",
        "entry_date": "2020-03-01",
        "leaving_date": "2022-06-30",
        "description": null,
        "resume_link": null,
        "resume_link_updated_at": null,
        "mini_bio_link": null,
        "mini_bio_link_updated_at": null,
        "links": [],
        "staffing_info": "",
        "certifications": [],
        "person_roles": [
            {
                "id": 2440,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 5,
                    "name": "Data Engineer",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2025-01-23T13:45:40Z"
                }
            },
            {
                "id": 2441,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 6,
                    "name": "manager",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2025-01-24T09:32:25Z"
                }
            },
            {
                "id": 2442,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 18,
                    "name": "Tech Lead",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2024-10-25T11:58:12Z"
                }
            },
            {
                "id": 2893,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 12,
                    "name": "Directeur de Mission",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2024-10-25T11:57:31Z"
                }
            },
            {
                "id": 3379,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 21,
                    "name": "Consultant",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-02-25T16:20:06Z",
                    "updated_at": "2025-01-23T13:30:34Z"
                }
            },
            {
                "id": 3377,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 1,
                    "name": "Architecte",
                    "primary": true,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2024-10-25T11:57:45Z"
                }
            },
            {
                "id": 3861,
                "person_id": 2142665570,
                "main": false,
                "role": {
                    "id": 13,
                    "name": "Expert",
                    "primary": false,
                    "secondary": true,
                    "active": true,
                    "created_at": "2020-01-14T09:37:38Z",
                    "updated_at": "2020-01-14T09:37:38Z"
                }
            }
        ]
    },
    {
        "id": 2142665327,
        "last_name": "REY",
        "first_name": "daisy",
        "nickname": "DARE",
        "registration_number": "11798522",
        "photo_url": null,
        "url": "https://octopod.octo.tools/api/v0/people/2142665327",
        "email": "daisy.rey@octo-fake.com",
        "phone_number": null,
        "job_alias": "Chargée RH et communication",
        "job": {
            "id": 41,
            "name": "Admin",
            "url": "https://octopod.octo.tools/api/v0/jobs/41",
            "created_at": "2013-03-26T10:02:51Z",
            "updated_at": "2023-09-26T08:36:02Z",
            "active": true,
            "appointable": true,
            "natural_order": 10041
        },
        "lob": {
            "id": 117,
            "name": "OCTO LA-HAUT 🍺",
            "abbreviation": "LAHO",
            "url": "https://octopod.octo.tools/api/v0/lobs/117",
            "photo_url": null,
            "active": true,
            "leaders": [
                {
                    "id": 2142664834,
                    "last_name": "RON",
                    "first_name": "Pat",
                    "nickname": "PARO",
                    "registration_number": "79654369",
                    "photo_url": "",
                    "url": "https://octopod.octo.tools/api/v0/people/2142664834"
                }
            ],
            "description": "",
            "turnover_type": "from_internal_team",
            "in_activity_rate": true,
            "timesheet_email_alert_enabled": true,
            "created_at": "2018-01-08T19:09:48Z",
            "updated_at": "2023-10-18T12:30:24Z",
            "email": "octo.pain.aux.chocolats@accenture-fake.com",
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
                "name": "CAMPEMENT",
                "description": null,
                "photo_url": null,
                "email": null,
                "active": true
            }
        },
        "work_schedule": {
            "id": 1,
            "name": "FR - Forfait Jours"
        },
        "included_in_activity_rate": false,
        "lcr": 676.56,
        "manager_id": 2142665051,
        "manager": {
            "id": 2142665051,
            "last_name": "GAULOIS",
            "first_name": "Henry",
            "nickname": "GAHE",
            "registration_number": "456321987",
            "photo_url": "",
            "url": "https://octopod.octo.tools/api/v0/people/2142665051"
        },
        "in_idf": false,
        "created_at": "2019-08-12T14:44:57Z",
        "updated_at": "2025-01-07T04:02:28Z",
        "entry_date": "2019-09-16",
        "leaving_date": "2025-01-31",
        "description": "",
        "resume_link": null,
        "resume_link_updated_at": null,
        "mini_bio_link": null,
        "mini_bio_link_updated_at": null,
        "links": [],
        "staffing_info": null,
        "certifications": [],
        "person_roles": []
    }
]`
}
