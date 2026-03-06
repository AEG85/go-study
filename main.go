package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 1. Название
// 2. Автора
// 3. Количество страниц
// 4. Информацию о том, прочитана эта книга нами, либо же нет
// 5. Время добавления в библиотеку
// 6. Время окончательного прочтения (когда книга была дочитана нами до конца)

// 1. Добавлять новые книги в нашу личную библиотеку
// 2. Отмечать отдельные книги как прочитанные
// 3. Получать информацию о какой-то конкретной книге
// 4. Получать список всех книг, с учётом возможной фильтрации по: автору, прочитано/не прочитано
// 5. Удалять книги из нашей библиотеки

type Book struct {
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	PageCount  int        `json:"pageCount"`
	IsReaded   bool       `json:"isReaded"`
	TimeInsert time.Time  `json:"timeInsert"`
	TimeReaded *time.Time `json:"timeReaded"`
}

var Books = make(map[string]Book)
var mu sync.RWMutex = sync.RWMutex{}

func main() {
	http.HandleFunc("/add", addBookHandler)
	http.HandleFunc("/readed", readedBookHandler)
	http.HandleFunc("/get", getBookHandler)
	http.HandleFunc("/list", getBooksListHandler)
	http.HandleFunc("/delete", deleteBookHandler)
	http.HandleFunc("/", getBooksListHandler)

	http.ListenAndServe(":9091", nil)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := validateHttpMethod(&w, r, "POST"); !ok {
		fmt.Println(err.Error())
		return
	}
	book := Book{
		TimeInsert: time.Now(),
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(WrongJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
	mu.Lock()
	defer mu.Unlock()
	Books[book.Title] = book
	if _, err := w.Write([]byte("Элемент добавлен успешно!")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}

}

func readedBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := validateHttpMethod(&w, r, "PATCH"); !ok {
		fmt.Println(err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(InvalidRequestBody)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
	mu.Lock()
	defer mu.Unlock()
	booksKey := string(body)
	book, ok := Books[booksKey]
	if !ok {
		if _, err := w.Write([]byte("Такой книги нет в библиотеке")); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		fmt.Println("Такой книги не существует")
		return
	}

	timeReaded := time.Now()
	book.TimeReaded = &timeReaded
	book.IsReaded = true
	Books[booksKey] = book
	if _, err := w.Write([]byte("Книга " + booksKey + " прочитана")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := validateHttpMethod(&w, r, "GET"); !ok {
		fmt.Println(err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(InvalidRequestBody)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
	mu.RLock()
	defer mu.RUnlock()
	booksKey := string(body)
	book, ok := Books[booksKey]
	if !ok {
		if _, err := w.Write([]byte("Такой книги нет в библиотеке")); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		fmt.Println("Такой книги не существует")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(http.StatusNoContent)
		if _, err := w.Write([]byte(InvalidJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}

}

func getBooksListHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := validateHttpMethod(&w, r, "GET"); !ok {
		fmt.Println(err.Error())
		return
	}

	authorFilter := r.URL.Query().Get("author")
	isReadedFilter := r.URL.Query().Get("isReaded")
	var isReadedBool bool
	var err error
	if isReadedFilter != "" {
		isReadedBool, err = strconv.ParseBool(isReadedFilter)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Неверно передан фильтр по Прочтению")); err != nil {
				fmt.Println(CantWriteResponse)
				return
			}
			return
		}
	}

	mu.RLock()
	defer mu.RUnlock()
	filteredBooks := make(map[string]Book)
	for title, book := range Books {
		// Проверяем фильтр по автору
		if authorFilter != "" && book.Author != authorFilter {
			continue
		}

		// Проверяем фильтр по прочтению
		if isReadedFilter != "" && book.IsReaded != isReadedBool {
			continue
		}

		// Если все фильтры пройдены, добавляем книгу
		filteredBooks[title] = book
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredBooks); err != nil {
		w.WriteHeader(http.StatusNoContent)
		if _, err := w.Write([]byte(InvalidJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if ok, err := validateHttpMethod(&w, r, "DELETE"); !ok {
		fmt.Println(err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(InvalidRequestBody)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
	mu.Lock()
	defer mu.Unlock()
	booksKey := string(body)
	_, ok := Books[booksKey]
	if !ok {
		if _, err := w.Write([]byte("Такой книги нет в библиотеке")); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		fmt.Println("Такой книги не существует")
		return
	}
	delete(Books, booksKey)
	if _, err := w.Write([]byte("Книга " + booksKey + " удалена")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}

}

func validateHttpMethod(w *http.ResponseWriter, r *http.Request, method string) (bool, error) {
	methodName := ""
	switch method {
	case "POST":
		methodName = http.MethodPost
	case "GET":
		methodName = http.MethodGet
	case "DELETE":
		methodName = http.MethodDelete
	case "PATCH":
		methodName = http.MethodPatch
	case "PUT":
		methodName = http.MethodPut
	}
	if r.Method != methodName {
		(*w).WriteHeader(http.StatusBadRequest)
		if _, err := (*w).Write([]byte(WrongHttpMethod)); err != nil {
			return false, errors.New(CantWriteResponse)
		}
		return false, errors.New(WrongHttpMethod)
	}
	return true, nil
}
