package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var storage = []string{}
var mu = sync.RWMutex{}

func main() {
	http.HandleFunc("/send", mainHandler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка сервера: ", err)
	}

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении: ", err)
		return
	}
	mu.Lock()
	storage = append(storage, string(body))
	mu.Unlock()
	mu.RLock()
	for _, val := range storage {
		fmt.Println(val)
	}
	mu.RUnlock()
}
