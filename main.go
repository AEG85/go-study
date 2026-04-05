package main

import (
	"context"
	"fmt"
	simpleconnetion "study/feature_postgres/simple_connetion"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnetion.CheckConnection(ctx)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connection success!")
}
