package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var storage = make(map[int]Message, 100)
var mu = sync.Mutex{}

type Message struct {
	Title     string `json:"title"`
	PostIndex int    `json:"postIndex"`
	Text      string `json:"text"`
	IsExpress bool   `json:"isExpress"`
}

func NewMessage(title string, postIndex int, text string, isExpress bool) Message {

	if title == "" || postIndex == 0 || text == "" {
		return Message{}
	}
	return Message{
		Title:     title,
		PostIndex: postIndex,
		Text:      text,
		IsExpress: isExpress,
	}

}

func main() {
	http.HandleFunc("/send", addMessageHandler)
	http.HandleFunc("/delete/", deleteMessageHandler)
	http.HandleFunc("/storage", storageHandler)
	http.HandleFunc("/message/", messageHandler)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка сервера: ", err)
	}
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Не верный метод запроса!")); err != nil {
			fmt.Println("Не удалось записать тело ответа")
			return
		}
		return
	}
	message := NewMessage("", 0, "", false)
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("Не получилось записать ответ")
			return
		}
		return
	}

	mu.Lock()
	randIndex := getRandIndex()
	storage[randIndex] = message
	fmt.Println("Добавлено сообщеине:", message)
	printStorage()
	mu.Unlock()
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Не верный метод запроса!")); err != nil {
			fmt.Println("Не удалось записать тело ответа")
		}
	}
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Не верный метод запроса!")); err != nil {
			fmt.Println("Не удалось записать тело ответа")
		}
	}
	mu.Lock()
	w.WriteHeader(http.StatusOK)
	writeResponse(w, storage)
	mu.Unlock()
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte("Не верный метод запроса!")); err != nil {
			fmt.Println("Не удалось записать тело ответа")
		}
	}
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
		if err := writeResponse(w, msg); err != nil {
			fmt.Println(err)
		}
		return
	}
	mu.Lock()
	msg, ok := storage[index]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		msg := "Не найден элемент с таким ID!"
		fmt.Println(msg, err)
		if err := writeResponse(w, msg); err != nil {
			fmt.Println(err)
		}
		mu.Unlock()
		return
	}
	mu.Unlock()
	w.WriteHeader(http.StatusAccepted)
	writeResponse(w, msg)
}

func writeResponse(w http.ResponseWriter, msg any) error {
	message, err := json.Marshal(msg)
	if err != nil {
		return errors.New("Не получилось разобрать JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(message); err != nil {
		return errors.New("е получилось записать ответ!")
	}

	return nil
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
