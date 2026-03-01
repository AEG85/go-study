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
	http.HandleFunc("/status/", testStatusHandler)
	http.HandleFunc("/storage", storageHandler)
	http.HandleFunc("/message/", messageHandler)
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

func storageHandler(w http.ResponseWriter, r *http.Request) {
	storageList := ""
	mu.Lock()
	for _, val := range storage {
		storageList += val + "\n"
	}
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
	writeResponse(w, storageList)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка при чтении: ", err)
		return
	}
	index, err := strconv.Atoi(string(body))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Вы передали не число!"
		fmt.Println(msg, err)
		writeResponse(w, msg)
		return
	}
	mu.Lock()
	msg, ok := storage[index]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		msg := "Не найден элемент с таким ID!"
		fmt.Println(msg, err)
		writeResponse(w, msg)
		mu.Unlock()
		return
	}
	mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
	writeResponse(w, msg)
}

func testStatusHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Не передано тело запроса!"
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		writeResponse(w, msg)
		return
	}

	statusCode, err := strconv.Atoi(string(body))
	if err != nil {
		msg := "Статус код должен быть числом! Вы передали:"
		fmt.Println(msg, string(body))
		writeResponse(w, msg)
		return
	}
	switch statusCode {
	case 200:
		w.WriteHeader(http.StatusOK)
		msg := "Получен ответ!"
		writeResponse(w, msg)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
		msg := "Ошибка в запросе!"
		writeResponse(w, msg)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Ошибка сервера!"
		writeResponse(w, msg)
	case 404:
		w.WriteHeader(http.StatusNotFound)
		msg := "Запрашиваемый ресурс не найден!"
		writeResponse(w, msg)
	case 403:
		w.WriteHeader(http.StatusForbidden)
		msg := "Недостаточно прав!"
		writeResponse(w, msg)
	case 201:
		w.WriteHeader(http.StatusCreated)
		msg := "Элемент добавлен!"
		writeResponse(w, msg)
	case 409:
		w.WriteHeader(http.StatusConflict)
		msg := "Нельзя обновить элемент с такими данными, конфликт!"
		writeResponse(w, msg)
	case 301:
		w.WriteHeader(http.StatusMovedPermanently)
		msg := "Ресурс перенесен постоянно!"
		writeResponse(w, msg)
	case 302:
		w.WriteHeader(http.StatusFound)
		msg := "Ресурс верменно перенесен!"
		writeResponse(w, msg)
	}
}

func writeResponse(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Println("Не получилось записать ответ!")
	}
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
