package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan bool, 3)
	go TestGorutina(ch, 1)
	go TestGorutina(ch, 2)
	go TestGorutina(ch, 3)

	fmt.Println("Ок", <-ch)
	fmt.Println("Ок", <-ch)
	fmt.Println("Ок", <-ch)
	fmt.Println("Время выполнения: ", time.Since(start))
}

func TestGorutina(chanal chan bool, number int) {
	for i := range 2 {
		time.Sleep(1 * time.Second)
		fmt.Println("Я горутина " + strconv.Itoa(number) + ", делаю вывод на экран в " + strconv.Itoa(i+1) + " раз")
	}
	chanal <- true
}
