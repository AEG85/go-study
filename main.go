package main

import (
	"context"
	"fmt"
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
	// Добавление книги
	// book := models.Book{
	// 	Title:           "Колобок2",
	// 	Author:          "Нород",
	// 	Review:          "Очень хорошая книжка",
	// 	PublicationYear: 1010,
	// }
	// commandTag, err := simplesql.InsertRow(ctx, conn, book)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(commandTag)
	// fmt.Println("Book added!")

	// Обновление книги
	// timeReaded := time.Now()
	// book := models.Book{
	// 	ID:              1,
	// 	Title:           "Ватсон",
	// 	Author:          "Артур Конендоил",
	// 	Review:          "Очень очень хорошая книжка",
	// 	IsRead:          true,
	// 	DateReaded:      &timeReaded,
	// 	PublicationYear: 1979,
	// }
	// commandTag, err := simplesql.UpdateRow(ctx, conn, book)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(commandTag)
	// fmt.Println("Book updated!")

	// Удаление книг
	// booksIds := []int{1, 5}

	// commandTag, err := simplesql.DeleteRow(ctx, conn, booksIds)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(commandTag)
	// fmt.Println("Books deleted!")

	// Получение всех книг
	books, err := simplesql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}
	for i := range books {
		pp.Println(books[i])
	}
	fmt.Println("Books selected!")

}
