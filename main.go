package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

type Fun struct {
	Name string
}

func main() {

	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	mailBox := make([]string, 0, 100_000)

	Vasya := Fun{Name: "Vacya"}
	Petya := Fun{Name: "Petya"}

	wg.Add(2)
	go Vasya.SendLetter(&mailBox, wg, mtx)
	go Petya.SendLetter(&mailBox, wg, mtx)

	wg.Wait()

	mtx.Lock()
	pp.Println(len(mailBox))
	mtx.Unlock()

}

func (fun *Fun) SendLetter(mailBox *[]string, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	message := "Я тебя обожаю!"
	for i := range 1000 {
		_ = i
		mtx.Lock()
		*mailBox = append(*mailBox, message)
		mtx.Unlock()
	}
}
