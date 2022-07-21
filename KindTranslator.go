package main

func KindTranslator(kind string) string {
	switch kind {
	case KIND_BILLABLE:
		return "Facturable"
	case KIND_NOT_BILLABLE:
		return "Non facturable"
	case KIND_PERMANENT:
		return "Activitées permanentes"
	case KIND_ABSENCE:
		return "Absences"
	}

	return kind
}
