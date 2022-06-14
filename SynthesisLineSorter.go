package main

type ByAssending []SynthesisLine

func (a ByAssending) Len() int {
	return len(a)
}

func (a ByAssending) Less(i, j int) bool {
	if a[i].Kind != "" && a[j].Kind != "" {
		if a[i].Kind == a[j].Kind {
			return a[i].Title < a[j].Title
		}
		return a[i].Kind < a[j].Kind
	}
	if a[i].Reference == "" && a[j].Reference == "" {
		if a[i].Kind != "" && a[j].Kind != "" {
			return a[i].Kind < a[j].Kind
		}
		return true
	}
	if a[i].Reference != "" && a[j].Reference != "" {
		return a[i].Reference < a[j].Reference
	}
	if a[j].Reference != "" && a[i].Reference == "" {
		return true
	}
	return false
}

func (a ByAssending) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
