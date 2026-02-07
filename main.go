package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Возникла паника: ", p)
		}
	}()
	// a := 0
	// b := 1 / a
	// fmt.Println(b, a)

	// var mapTest map[int]int
	// mapTest[2] = 123

	// slice := make([]int, 2, 2)
	// c := slice[4]
	// fmt.Println(c)

	panic("💥 Это моя собственная паника! Я сам её вызвал!")

}
