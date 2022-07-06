package main

import (
	"log"
	"net/http"
)

var token *Token

func main() {
	var err error
	go GetBankHolidayInstance().Init()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Index)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
