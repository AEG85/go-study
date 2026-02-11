package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		go Gardening(wg, i+1)
	}
	wg.Wait()

}

func Gardening(wg *sync.WaitGroup, gardenerNumber int) {
	defer wg.Done()
	timeIntervals := []int{500, 600, 700, 800, 900, 1000}

	randIndex := rand.Intn(6)
	randTimeDuration := timeIntervals[randIndex]

	fmt.Println("Я садовник ", gardenerNumber, ", я поливаю! ", randTimeDuration, "мсек")

	time.Sleep(time.Duration(randTimeDuration) * time.Millisecond)
	fmt.Println("Я садовник ", gardenerNumber, ", я удобряю!", randTimeDuration, "мсек")

}
