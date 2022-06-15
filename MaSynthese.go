package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	timeInput, err = TimeInputGetter(token.AccessToken, "2142666213", "2022-03-01", "2022-06-15", 100)
	if err != nil {
		log.Fatalln(err)
	}

	synthesisLines := timeInput.timeInputAggregator()

	sort.Sort(ByAssending(synthesisLines))

	var cumul float64 = 0
	currentKind := ""

	for _, synthesisLine := range synthesisLines {
		if currentKind != "" && currentKind != synthesisLine.Kind {
			fmt.Printf("Total %s : %f\n", currentKind, cumul)
			cumul = 0.0
		}
		fmt.Printf("%s %s %s %s %f\n", synthesisLine.Kind, synthesisLine.Title, synthesisLine.Reference, synthesisLine.CustomerName, synthesisLine.TimeSum)
		cumul += synthesisLine.TimeSum
		currentKind = synthesisLine.Kind
	}
}
