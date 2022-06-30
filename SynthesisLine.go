package main

type SynthesisLine struct {
	ActivityID     int64   `json:"activityId"`
	Title          string  `json:"title"`
	CustomerName   string  `json:"customerName"`
	TimeSum        float64 `json:"timeSum"`
	Kind           string  `json:"kind"`
	Reference      string  `json:"reference"`
	ProjectName    string  `json:"projectName"`
	IsLineSubTotal bool    `json:"IsLineSubTotal"`
	IsLineTotal    bool    `json:"IsLineTotal"`
	RowCount       int64   `json:"RowCount"`
}
