package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go RandomInt(chan1)
	go RandomInt(chan2)
	go RandomInt(chan3)

	fmt.Println("Пишем из первого канала: ", <-chan1)
	fmt.Println("Пишем из второго канала: ", <-chan2)
	fmt.Println("Пишем из третьего канала: ", <-chan3)

}

func RandomInt(ch chan int) {
	random := rand.Int()
	ch <- random
}
