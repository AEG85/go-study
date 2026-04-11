package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 5
	for i := range counter {
		fmt.Println("Hello from Docker!: ", i)
		time.Sleep(time.Second * 1)
	}
}
