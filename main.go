package main

import (
	"context"
	"fmt"
	"study/pressure"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Все датчики отработали сеанс завершен!")
	meteoChanel := make(chan string)
	wg := &sync.WaitGroup{}
	pressureContex, presssureCancel := context.WithCancel(context.Background())
	seismoContex, seismoCancel := context.WithCancel(context.Background())
	wetContex, wetCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		presssureCancel()
	}()

	go func() {
		time.Sleep(4 * time.Second)
		seismoCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		wetCancel()
	}()

	wg.Add(1)
	go pressure.PressureSensor(pressureContex, wg, meteoChanel)
	wg.Add(1)
	go pressure.SeismoSensor(seismoContex, wg, meteoChanel)
	wg.Add(1)
	go pressure.WetSensor(wetContex, wg, meteoChanel)

	go func() {
		wg.Wait()
		close(meteoChanel)
	}()

	for meteoData := range meteoChanel {
		fmt.Println(meteoData)
	}

	// Альтернатива с select
	// for {
	// 	metoData, ok := <-meteoChanel
	// 	if !ok {
	// 		return
	// 	} else {
	// 		fmt.Println(metoData)
	// 	}
	// }

}
