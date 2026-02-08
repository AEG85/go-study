package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	chan3 := make(chan float64)

	go IntegerGenerator(chan1)
	go StringGenerator(chan2)
	go FloatGenerator(chan3)

	for {
		select {
		case intValue := <-chan1:
			fmt.Println("Читаем int: ", intValue)
		case strValue := <-chan2:
			fmt.Println("Читаем string: ", strValue)
		case flatValue := <-chan3:
			fmt.Println("Читаем flat: ", flatValue)
		}
	}
}

func IntegerGenerator(ch chan int) {
	i := 1
	for {
		time.Sleep(500 * time.Millisecond)
		ch <- i
		i++
	}
}

func StringGenerator(ch chan string) {
	i := 1
	for {
		time.Sleep(1 * time.Second)
		ch <- "Строка " + strconv.Itoa(i)
		i++
	}
}

func FloatGenerator(ch chan float64) {
	for {
		time.Sleep(5 * time.Second)
		ch <- rand.Float64()
	}
}
