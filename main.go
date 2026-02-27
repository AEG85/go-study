package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerMain)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
	}
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	topic := []string{
		"Хорошего дня!",
		"Добрейшего добра!",
		"Мира вам, миряне!",
		"Вечер в хату!",
	}
	index := rand.IntN(len(topic))
	randPhrase := topic[index]

	w.Write([]byte(randPhrase))
}
