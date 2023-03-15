package tools

import (
	"fmt"
	"os"
)

func Debug(message string) {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println(message)
	}
}
