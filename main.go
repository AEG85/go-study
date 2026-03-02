package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string  `json:"name"`
	Adress   string  `json:"adress"`
	Age      int     `json:"age"`
	Marreged bool    `json:"marreged"`
	Height   float64 `json:"height"`
}

func NewUser() User {
	return User{}
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/get", getHandler)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Не удалось запустить api сервер")
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	user := NewUser()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("Не удалось записать ответ")
			return
		}
		return
	}
	fmt.Println("Передан пользователь", user)
	if _, err := w.Write([]byte("Данные пользователя переданы")); err != nil {
		fmt.Println("Не получилоьс записать ответ!")
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	user := NewUser()
	user.Name = "Евгеинй"
	user.Adress = "Омск, Поселкова"
	user.Age = 41
	user.Height = 175.5
	user.Marreged = true
	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Не удалось получить пользователя!")
	}
	if _, err := w.Write(userJson); err != nil {
		fmt.Println("Не удалось записать ответ!")
	}
}
