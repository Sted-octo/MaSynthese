package main

import (
	"Octoptimist/infrastructure"
	"Octoptimist/ui"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error
	go infrastructure.BankHolidaysSingletonGetter().Init()
	go infrastructure.TargetTacesSingletonGetter().Init()
	go infrastructure.GlobalPurposeProjectsSingletonGetter().Init()
	go infrastructure.DiscountProjectsSingletonGetter().Init()
	go infrastructure.PeoplesGlobalMapSingletonGetter()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", ui.Login)
	http.HandleFunc("/synthesis", ui.Synthesis)
	http.HandleFunc("/tribe", ui.Tribe)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
