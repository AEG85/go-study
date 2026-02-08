package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go func() {
		for i := range 3 {
			fmt.Println("Я 1 горутина мой текст: " + strconv.Itoa(i))
		}
	}()

	go func() {
		for i := range 3 {
			time.Sleep(1 * time.Second)
			fmt.Println("Я 2 горутина мой текст: " + strconv.Itoa(i))
		}
	}()

	go func() {
		for i := range 3 {
			fmt.Println("Я 3 горутина мой текст: " + strconv.Itoa(i))
		}
	}()
	time.Sleep(4 * time.Second)
}
