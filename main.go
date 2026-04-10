package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"study/feature_postgres/models"
	simpleconnetion "study/feature_postgres/simple_connetion"
	simplesql "study/feature_postgres/simple_sql"

	"github.com/k0kubun/pp/v3"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnetion.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connection success!")

	if err := simplesql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("Table created succesfull!")

	createUser := os.Getenv("NEW_USER")
	if createUser == "" {
		fmt.Println("Не задана переменная окружения NEW_USER")
		return
	}

	if createUser == "YES" {
		scanner := bufio.NewScanner(os.Stdin)
		var name string
		for {
			fmt.Print("Введите полное имя пользоватлея: ")
			scanner.Scan()
			name = scanner.Text()
			name = strings.TrimSpace(name)
			if name == "" {
				fmt.Println("Вы не ввели полное имя")
				continue
			}
			words := strings.Fields(name)
			if len(words) < 2 {
				fmt.Println("Нужно указать минимум два слова")
				continue
			}
			break
		}

		var phoneNember string

		fmt.Print("Введите номер телефона: ")
		scanner.Scan()
		phoneNember = scanner.Text()
		phoneNember = strings.TrimSpace(phoneNember)

		// Добавление пользователя
		user := models.User{
			FullName:    name,
			PhoneNumber: phoneNember,
		}
		commandTag, err := simplesql.InsertRow(ctx, conn, user)
		if err != nil {
			panic(err)
		}
		fmt.Println(commandTag)
		fmt.Println("User added!")
	}

	if createUser == "NO" {
		// Получение всех пользователей
		users, err := simplesql.SelectRows(ctx, conn)
		if err != nil {
			panic(err)
		}
		for i := range users {
			pp.Println(users[i])
		}
		fmt.Println("Users selected!")
	}
}
