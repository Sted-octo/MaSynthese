package dataproviders

import "time"

type LobDto struct {
	ID                         int64      `json:"id"`
	Name                       string     `json:"name"`
	Abbreviation               string     `json:"abbreviation"`
	URL                        string     `json:"url"`
	PhotoURL                   string     `json:"photo_url"`
	Active                     bool       `json:"active"`
	Leaders                    []Leader   `json:"leaders"`
	Description                string     `json:"description"`
	TurnoverType               string     `json:"turnover_type"`
	InActivityRate             bool       `json:"in_activity_rate"`
	TimesheetEmailAlertEnabled bool       `json:"timesheet_email_alert_enabled"`
	CreatedAt                  time.Time  `json:"created_at"`
	UpdatedAt                  time.Time  `json:"updated_at"`
	Email                      string     `json:"email"`
	Subsidiary                 Subsidiary `json:"subsidiary"`
	League                     League     `json:"league"`
}

type Leader struct {
	ID                 int64  `json:"id"`
	LastName           string `json:"last_name"`
	FirstName          string `json:"first_name"`
	Nickname           string `json:"nickname"`
	RegistrationNumber string `json:"registration_number"`
	PhotoURL           string `json:"photo_url"`
	URL                string `json:"url"`
}

type League struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	PhotoURL    interface{} `json:"photo_url"`
	Email       interface{} `json:"email"`
	Active      bool        `json:"active"`
}

type Subsidiary struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Region       string   `json:"region"`
	ChronoPrefix string   `json:"chrono_prefix"`
	Locale       string   `json:"locale"`
	Timezone     string   `json:"timezone"`
	Active       bool     `json:"active"`
	URL          string   `json:"url"`
	Currency     Currency `json:"currency"`
}

type Currency struct {
	Symbol string `json:"symbol"`
}
