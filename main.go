package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var storage = make(map[int]string, 100)
var mu = sync.Mutex{}

func main() {
	http.HandleFunc("/send", addMessageHandler)
	http.HandleFunc("/delete/", deleteMessageHandler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка сервера: ", err)
	}
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении: ", err)
		return
	}
	mu.Lock()
	randIndex := getRandIndex()
	storage[randIndex] = string(body)
	fmt.Println("Добавлено сообщеине:", string(body))
	printStorage()
	mu.Unlock()
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении: ", err)
		return
	}

	mu.Lock()
	index, err := strconv.Atoi(string(body))
	if err != nil {
		fmt.Println("ID элемента должно быть числом! ", string(body))
		printStorage()
		mu.Unlock()
		return
	}
	if _, ok := storage[index]; !ok {
		fmt.Println("Нет элемента с таким ID:", string(body))
		printStorage()
		mu.Unlock()
		return
	}
	fmt.Println("Удаляем сообщение:", storage[index])
	delete(storage, index)
	printStorage()
	mu.Unlock()
}

func getRandIndex() int {
	for {
		storageLen := len(storage)
		if storageLen < 1 {
			storageLen = 1
		}
		storageLen *= 2
		index := rand.Intn(storageLen)
		if _, ok := storage[index]; ok {
			continue
		}
		return index
	}

}

func printStorage() {
	for key, val := range storage {
		fmt.Println(key, val)
	}
}
