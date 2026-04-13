package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Position string `json:"position"`
}

var Employees = make(map[int]Employee)
var mu sync.RWMutex = sync.RWMutex{}

const (
	WrongHttpMethod    = "Неправильный http метод в запросе"
	CantWriteResponse  = "Не получилось записать ответ"
	WrongJson          = "Неверно сформирован json в запросе"
	InvalidRequestBody = "Не получилось прочитать тело запроса"
	InvalidJson        = "Не получилось преобразовать данные в json формат"
)

func main() {
	r := chi.NewRouter()
	r.Route("/employee", func(r chi.Router) {
		r.Get("/", getEmployeesListHandler)
		r.Post("/", addEmployeeHandler)
		r.Get("/{id}", getEmployeeByID)
		r.Delete("/{id}", deleteEmployeeHandler)
	})
	r.Get("/", getEmployeesListHandler)

	http.ListenAndServe(":9091", r)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID должен быть числом"))
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	employee, exists := Employees[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Сотрудник не найден"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func addEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(WrongJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
	mu.Lock()
	defer mu.Unlock()
	Employees[employee.ID] = employee
	if _, err := w.Write([]byte("Сотрудник добавлен успешно!")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}
}

func getEmployeesListHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Employees); err != nil {
		w.WriteHeader(http.StatusNoContent)
		if _, err := w.Write([]byte(InvalidJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	employeeID, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID должен быть числом"))
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, ok := Employees[employeeID]
	if !ok {
		if _, err := w.Write([]byte("Токого сотрудника не существует")); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		fmt.Println("Такой книги не существует")
		return
	}
	delete(Employees, employeeID)
	if _, err := w.Write([]byte("Сотрудник с ID " + idStr + " удален")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}
}
