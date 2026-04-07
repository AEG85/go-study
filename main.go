package main

import (
	"context"
	"fmt"
	simpleconnetion "study/feature_postgres/simple_connetion"
	simplesql "study/feature_postgres/simple_sql"
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

	simplesql.ListPages(ctx, conn, 10)

}
