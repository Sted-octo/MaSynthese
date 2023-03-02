package domain

type People struct {
	ID                              int64          `json:"id"`
	LastName                        string         `json:"last_name"`
	FirstName                       string         `json:"first_name"`
	Nickname                        string         `json:"nickname"`
	URL                             string         `json:"url"`
	Email                           string         `json:"email"`
	Job                             Job            `json:"job"`
	Lob                             Lob            `json:"lob"`
	WorkSchedule                    WorkSchedule   `json:"work_schedule"`
	IncludedInActivityRate          bool           `json:"included_in_activity_rate"`
	CreatedAt                       string         `json:"created_at"`
	UpdatedAt                       string         `json:"updated_at"`
	EntryDate                       string         `json:"entry_date"`
	LeavingDate                     interface{}    `json:"leaving_date"`
	ActiveRoles                     []WorkSchedule `json:"active_roles"`
	LCR                             float64        `json:"lcr"`
	StaffingInfo                    string         `json:"staffing_info"`
	MandatoryMainPaidLeaveFulfilled bool           `json:"mandatory_main_paid_leave_fulfilled"`
}

type WorkSchedule struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Job struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Active       bool   `json:"active"`
	Appointable  bool   `json:"appointable"`
	NaturalOrder int64  `json:"natural_order"`
}

type Lob struct {
	ID                         int64        `json:"id"`
	Abbreviation               string       `json:"abbreviation"`
	URL                        string       `json:"url"`
	Name                       string       `json:"name"`
	TurnoverType               string       `json:"turnover_type"`
	Active                     bool         `json:"active"`
	InActivityRate             bool         `json:"in_activity_rate"`
	TimesheetEmailAlertEnabled bool         `json:"timesheet_email_alert_enabled"`
	CreatedAt                  string       `json:"created_at"`
	UpdatedAt                  string       `json:"updated_at"`
	Subsidiary                 Subsidiary   `json:"subsidiary"`
	League                     WorkSchedule `json:"league"`
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
