package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeSlice := []int{
		100,
		200,
		300,
		400,
		500,
		600,
		700,
		800,
		900,
		1000,
	}

	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		randIndex := rand.Intn(len(timeSlice))
		time.Sleep(time.Duration(timeSlice[randIndex]) * time.Millisecond)
		fmt.Println("Задержка в 1 горутине равна: ", timeSlice[randIndex])
		chan1 <- rand.Int()
	}()

	go func() {
		randIndex := rand.Intn(len(timeSlice))
		time.Sleep(time.Duration(timeSlice[randIndex]) * time.Millisecond)
		fmt.Println("Задержка во 2 горутине равна: ", timeSlice[randIndex])
		chan2 <- rand.Int()
	}()

	time.Sleep(3000 * time.Millisecond)

	select {
	case number := <-chan1:
		fmt.Println("Пишем из первой секции: ", number)
	case number := <-chan2:
		fmt.Println("Пишем из второй секции: ", number)
	default:
		fmt.Println("Пишем из дефолтной секции")
	}

}
