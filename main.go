package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"study/models"
	"study/pgx/connection"
	"study/pgx/requests"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

const (
	WrongHttpMethod    = "Неправильный http метод в запросе"
	CantWriteResponse  = "Не получилось записать ответ"
	WrongJson          = "Неверно сформирован json в запросе"
	InvalidRequestBody = "Не получилось прочитать тело запроса"
	InvalidJson        = "Не получилось преобразовать данные в json формат"
)

func main() {
	ctx := context.Background()
	conn, err := connection.CreateConnection(ctx)
	if err != nil {
		panic("Не удалось подключиться к БД!")
	}

	if err := requests.CreateTable(ctx, conn); err != nil {
		panic("Не получилось создать таблицу сотрудников!")
	}

	r := chi.NewRouter()
	r.Route("/employee", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			getEmployeesListHandler(w, r, ctx, conn)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			addEmployeeHandler(w, r, ctx, conn)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			getEmployeeByID(w, r, ctx, conn)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			deleteEmployeeHandler(w, r, ctx, conn)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		getEmployeesListHandler(w, r, ctx, conn)
	})

	http.ListenAndServe(":9091", r)
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request, ctx context.Context, conn *pgx.Conn) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID должен быть числом"))
		return
	}

	employee, err := requests.SelectRow(ctx, conn, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Сотрудник не найден"))
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func addEmployeeHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, conn *pgx.Conn) {
	employee := models.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(WrongJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}

	_, err := requests.InsertRow(ctx, conn, employee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось добавить сотрудника"))
		fmt.Println(err)
		return
	}
	if _, err := w.Write([]byte("Сотрудник добавлен успешно!")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}
}

func getEmployeesListHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, conn *pgx.Conn) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := requests.SelectRows(ctx, conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Не получилось достать сотрудников из базы")
	}

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		w.WriteHeader(http.StatusNoContent)
		if _, err := w.Write([]byte(InvalidJson)); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		return
	}
}

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, conn *pgx.Conn) {
	idStr := chi.URLParam(r, "id")
	employeeID, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID должен быть числом"))
		return
	}

	employee, err := requests.SelectRow(ctx, conn, employeeID)
	if err != nil {
		if _, err := w.Write([]byte("Токого сотрудника не существует")); err != nil {
			fmt.Println(CantWriteResponse)
			return
		}
		fmt.Println("Такой книги не существует")
		return
	}

	employeeSlice := []int{employee.ID}
	_, err = requests.DeleteRow(ctx, conn, employeeSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Не удалось удалить сотрудника"))
		fmt.Println(err)
		return
	}

	if _, err := w.Write([]byte("Сотрудник с ID " + idStr + " удален ")); err != nil {
		fmt.Println(CantWriteResponse)
		return
	}

}
