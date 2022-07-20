package main

type ITargetTace interface {
	GetTargetTace(jobId int) int
	Init()
}

type TargetTace struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	TargetTace int    `json:"TargetTace"`
}
