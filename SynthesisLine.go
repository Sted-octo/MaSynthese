package main

type SynthesisLine struct {
	ActivityID   int64   `json:"activityId"`
	Title        string  `json:"title"`
	CustomerName string  `json:"customerName"`
	TimeSum      float64 `json:"timeSum"`
	Kind         string  `json:"kind"`
	Reference    string  `json:"reference"`
	ProjectName  string  `json:"projectName"`
}
