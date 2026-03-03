package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Не удалось запустить сервер!")
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Не правильный метод запроса")); err != nil {
			fmt.Println("Не получилось записать ответ")
			return
		}
		return
	}
	name := r.URL.Query().Get("name")
	fmt.Println("Имя: ", name)
	lastName := r.URL.Query().Get("lastName")
	fmt.Println("Фамилия: ", lastName)
	params := r.URL.Query()
	for k, v := range params {
		fmt.Printf("Название: %s, Значение: %s \n", k, v[0])
	}
}
