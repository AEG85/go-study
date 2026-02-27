package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", handlerDog)
	http.HandleFunc("/cat", handlerCat)
	http.HandleFunc("/cow", handlerCow)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Главная страница!"))
		if err != nil {
			fmt.Println("Произошла ошибка: ", err.Error())
		}
	})

	err := http.ListenAndServe("localhost:9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка: ", err.Error())
	}
}

func handlerDog(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я собака, я говорю 'ГАВ'"))
	if err != nil {
		fmt.Println("Произошла ошибка: ", err.Error())
	}
}

func handlerCat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я кошка, я говорю 'Мяу'"))
	if err != nil {
		fmt.Println("Произошла ошибка: ", err.Error())
	}
}

func handlerCow(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я корова я говорю 'Мууу'"))
	if err != nil {
		fmt.Println("Произошла ошибка: ", err.Error())
	}
}
