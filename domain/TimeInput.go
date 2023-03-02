package domain

type TimeInput []TimeInputElement

type TimeInputElement struct {
	Day        string   `json:"day"`
	TimeInDays string   `json:"time_in_days"`
	Activity   Activity `json:"activity"`
}

type Activity struct {
	ID                 int64    `json:"id"`
	Title              string   `json:"title"`
	Role               *Role    `json:"role"`
	NbDays             *string  `json:"nb_days"`
	AverageDailyRate   *string  `json:"average_daily_rate"`
	Kind               string   `json:"kind"`
	StaffingNeededFrom *string  `json:"staffing_needed_from"`
	Expertise          *string  `json:"expertise"`
	Project            *Project `json:"project"`
}

type Project struct {
	ID        int64     `json:"id"`
	URL       string    `json:"url"`
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	Reference string    `json:"reference"`
	Status    *string   `json:"status"`
	Customer  *Customer `json:"customer"`
}

type Customer struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	AccountManagerID    int64  `json:"account_manager_id"`
	AccountManagerEmail string `json:"account_manager_email"`
}

type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (timeInput *TimeInput) Add(timeInputToAdd *TimeInputElement) {
	*timeInput = append(*timeInput, *timeInputToAdd)
}

func (timeInput *TimeInput) Concat(timeInputToAdd *TimeInput) TimeInput {
	la := len(*timeInput)
	c := make(TimeInput, la+len(*timeInputToAdd))
	_ = copy(c, *timeInput)
	_ = copy(c[la:], *timeInputToAdd)
	return c
}

var ACTIVITY_ID_RTT int64 = 2140298843
var ACTIVITY_ID_CONGES_PAYE int64 = 2140309429
var ACTIVITY_ID_SICK_DAY int64 = 2140312911
var ACTIVITY_ID_PART_TIME_BREAK int64 = 2140316822
var ACTIVITY_ID_PARENTAL_BREAK int64 = 3000050819
var ACTIVITY_ID_NO_SALARY_BREAK int64 = 3000030459
var ACTIVITY_ID_MEDICAL_CARE int64 = 3000030462
var ACTIVITY_ID_FAMILY_DAY int64 = 3000030457
var ACTIVITY_ID_PARENT_DAY int64 = 3000030458
var ACTIVITY_ID_MEDICAL_PART_TIME_BREAK int64 = 3000050818
var ACTIVITY_ID_AUTORIZED_BREAK int64 = 3000050820
var ACTIVITY_ID_NO_EXCUSE_BREAK int64 = 3000050821
var ACTIVITY_ID_PERSONAL_CARE int64 = 3000065641
var ACTIVITY_ID_INTERCONTRAT int64 = 2140318361

func (activity *Activity) IsDayBreak() bool {
	return activity.ID == ACTIVITY_ID_RTT ||
		activity.ID == ACTIVITY_ID_CONGES_PAYE ||
		activity.ID == ACTIVITY_ID_SICK_DAY ||
		activity.ID == ACTIVITY_ID_PART_TIME_BREAK ||
		activity.ID == ACTIVITY_ID_PARENTAL_BREAK ||
		activity.ID == ACTIVITY_ID_NO_SALARY_BREAK ||
		activity.ID == ACTIVITY_ID_MEDICAL_CARE ||
		activity.ID == ACTIVITY_ID_FAMILY_DAY ||
		activity.ID == ACTIVITY_ID_PARENT_DAY ||
		activity.ID == ACTIVITY_ID_MEDICAL_PART_TIME_BREAK ||
		activity.ID == ACTIVITY_ID_AUTORIZED_BREAK ||
		activity.ID == ACTIVITY_ID_NO_EXCUSE_BREAK ||
		activity.ID == ACTIVITY_ID_PERSONAL_CARE
}
