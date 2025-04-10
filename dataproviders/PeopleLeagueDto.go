package dataproviders

import "time"

type PeopleLeagueDto struct {
	ID                     int64         `json:"id"`
	LastName               string        `json:"last_name"`
	FirstName              string        `json:"first_name"`
	Nickname               string        `json:"nickname"`
	RegistrationNumber     string        `json:"registration_number"`
	PhotoURL               interface{}   `json:"photo_url"`
	URL                    string        `json:"url"`
	Email                  string        `json:"email"`
	PhoneNumber            interface{}   `json:"phone_number"`
	JobAlias               string        `json:"job_alias"`
	Job                    Job           `json:"job"`
	Lob                    Lob           `json:"lob"`
	WorkSchedule           WorkSchedule  `json:"work_schedule"`
	IncludedInActivityRate bool          `json:"included_in_activity_rate"`
	LCR                    float64       `json:"lcr"`
	ManagerID              int64         `json:"manager_id"`
	Manager                Manager       `json:"manager"`
	InIDF                  bool          `json:"in_idf"`
	CreatedAt              time.Time     `json:"created_at"`
	UpdatedAt              time.Time     `json:"updated_at"`
	EntryDate              string        `json:"entry_date"`
	LeavingDate            string        `json:"leaving_date"`
	Description            string        `json:"description"`
	ResumeLink             interface{}   `json:"resume_link"`
	ResumeLinkUpdatedAt    interface{}   `json:"resume_link_updated_at"`
	MiniBioLink            interface{}   `json:"mini_bio_link"`
	MiniBioLinkUpdatedAt   interface{}   `json:"mini_bio_link_updated_at"`
	Links                  []interface{} `json:"links"`
	StaffingInfo           interface{}   `json:"staffing_info"`
	Certifications         []interface{} `json:"certifications"`
	PersonRoles            []interface{} `json:"person_roles"`
}

type Manager struct {
	ID                 int64  `json:"id"`
	LastName           string `json:"last_name"`
	FirstName          string `json:"first_name"`
	Nickname           string `json:"nickname"`
	RegistrationNumber string `json:"registration_number"`
	PhotoURL           string `json:"photo_url"`
	URL                string `json:"url"`
}
