package dataproviders

type TimeInputDto []TimeInputElement

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

func (timeInput *TimeInputDto) Add(timeInputToAdd *TimeInputElement) {
	*timeInput = append(*timeInput, *timeInputToAdd)
}
