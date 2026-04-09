package main

import (
	"fmt"
	"os"
)

func main() {
	myAge := os.Getenv("my_age")

	if myAge != "" {
		fmt.Println("Мой возраст:", myAge)
	} else {
		fmt.Println("Возраст не задан!")
	}
}
