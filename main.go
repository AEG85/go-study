package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Не получилось запустить http сервер!")
	}
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("В запросе использован неверный метод!")
		if _, err := w.Write([]byte("В запросе использован неверный метод!")); err != nil {
			fmt.Println("Не получилось записать ответ!")
			return
		}
		return
	}

	msg := "Успешный запрос!"
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println("Не получилось записать ответ!")
		return
	}
	fmt.Println(msg)
}
