package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var token *Token
	var timeInput *TimeInput

	token, err := TokenGetter(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("AccessToken:")
	fmt.Println(token.AccessToken)

	timeInput, err = TimeInputGetter(token.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result:")
	fmt.Println(timeInput)
}
