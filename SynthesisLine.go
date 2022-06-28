package main

type SynthesisLine struct {
	ActivityID   int64   `json:"activityId"`
	Title        string  `json:"title"`
	CustomerName string  `json:"customerName"`
	TimeSum      float64 `json:"timeSum"`
	Kind         string  `json:"kind"`
	Reference    string  `json:"reference"`
	ProjectName  string  `json:"projectName"`
	TypeLine     int8    `json:"TypeLine"`
	RowCount     int64   `json:"RowCount"`
}

const (
	LineNormal   int8 = 0
	LineSubTotal int8 = 1
	LineTotal    int8 = 2
)
