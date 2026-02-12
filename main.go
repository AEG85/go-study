package main

import (
	"fmt"
	"sync"
	"time"
)

// Используем простую int переменную
var voteNumber int = 0

// Используем atomic int64
// var voteNumber atomic.Int64

func main() {

	timeStart := time.Now()
	var wg = &sync.WaitGroup{}
	var mtx = &sync.Mutex{}

	wg.Add(1)
	go Vote(wg, mtx)

	wg.Add(1)
	go Vote(wg, mtx)

	wg.Add(1)
	go VoteRead(wg, mtx)

	wg.Wait()
	// Используем обычный int
	mtx.Lock()
	fmt.Println("Колличество голосов: ", voteNumber)
	mtx.Unlock()

	// Используем atomic
	// fmt.Println("Колличество голосов: ", voteNumber.Load())

	fmt.Println("Время выполнения: ", time.Since(timeStart))
}

func Vote(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := range 100_000 {
		_ = i
		// Исеользуем обычный int
		mtx.Lock()
		voteNumber++
		mtx.Unlock()

		// Используем atomic
		// voteNumber.Add(1)
	}
}

func VoteRead(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := range 100_000 {
		_ = i
		// Исеользуем обычный int
		mtx.Lock()
		_ = voteNumber
		mtx.Unlock()

		// Используем atomic
		// voteNumber.Add(1)
	}
}
