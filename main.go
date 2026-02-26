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

	pressureChanel := make(chan string, 8)
	seismoChanel := make(chan string, 8)
	wetChanel := make(chan string, 8)

	wg := &sync.WaitGroup{}

	pressureContex, presssureCancel := context.WithCancel(context.Background())
	seismoContex, seismoCancel := context.WithCancel(context.Background())
	wetContex, wetCancel := context.WithCancel(context.Background())

	meteoData := []string{}
	meteMu := sync.Mutex{}

	defer func() {
		meteMu.Lock()
		defer meteMu.Unlock()
		for _, val := range meteoData {
			fmt.Println(val)
		}
	}()

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
	go func() {
		defer wg.Done()
		meteMu.Lock()
		meteoData = append(meteoData, "Начали сбор данных с датчика давления!")
		meteMu.Unlock()
		pressure.PressureSensor(pressureContex, pressureChanel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		meteMu.Lock()
		meteoData = append(meteoData, "Начали сбор данных с сейсмо датчика!")
		meteMu.Unlock()
		pressure.SeismoSensor(seismoContex, seismoChanel)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		meteMu.Lock()
		meteoData = append(meteoData, "Начали сбор данных с датчика влажности!")
		meteMu.Unlock()
		pressure.WetSensor(wetContex, wetChanel)
	}()

	go func() {
		wg.Wait()
		close(pressureChanel)
		close(seismoChanel)
		close(wetChanel)
	}()

	// Для общего канала
	// for meteoData := range meteoChanel {
	// 	fmt.Println(meteoData)
	// }

	for pressureChanel != nil || seismoChanel != nil || wetChanel != nil {
		select {
		case presureData, ok := <-pressureChanel:
			if !ok {
				pressureChanel = nil
				continue
			}
			meteMu.Lock()
			meteoData = append(meteoData, "Данные датчика давнеия получены: "+presureData)
			meteMu.Unlock()
		case seismoData, ok := <-seismoChanel:
			if !ok {
				seismoChanel = nil
				continue
			}
			meteMu.Lock()
			meteoData = append(meteoData, "Данные сейсмо датчика получены: "+seismoData)
			meteMu.Unlock()
		case wetData, ok := <-wetChanel:
			if !ok {
				wetChanel = nil
				continue
			}
			meteMu.Lock()
			meteoData = append(meteoData, "Данные датчика влажности получены: "+wetData)
			meteMu.Unlock()
		}
	}
}
