package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", mainHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка, api сервер не запустился:", err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Вы не передали name в заголовке!"))
		if err != nil {
			fmt.Println("Не получилось записать ответ!")
			return
		}
		return
	}
	fmt.Println("Передано имя:", name)

	var headersString strings.Builder
	for key, header := range r.Header {
		headerString := strings.Join(header, ", ")
		headersString.WriteString(key)
		headersString.WriteString(" : ")
		headersString.WriteString(headerString)
		headersString.WriteString("\n")
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(headersString.String()))
	if err != nil {
		msg := "Не получилось отправить ответ!"
		fmt.Println(msg)
	}
}
