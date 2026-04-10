package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	outPut := os.Getenv("OUTPUT_COUNT")

	if outPut == "" {
		fmt.Println("Переменная окружения OUTPUT_COUNT должна быть задана")
		return
	}

	outPutCount, err := strconv.Atoi(outPut)

	if err != nil {
		fmt.Println("Ошибка преобразования OUTPUT_COUNT в целое число")
		return
	}

	for index := range outPutCount {
		fmt.Println("Работа цикла:", index+1)
	}
}
