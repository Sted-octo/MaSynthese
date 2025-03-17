package tools

import (
	"log"
	"os"
)

func Debug(message string) {
	if os.Getenv("DEBUG") == "true" {
		log.Println(message)
	}
}
