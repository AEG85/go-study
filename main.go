package main

import (
	"errors"
	"fmt"
	"math/rand"
	"study/banking"
	"time"
)

func main() {
	currency, err := banking.GetCurrency("rub")
	if err != nil {
		fmt.Println(err.Error())
	}
	bill, err := banking.NewBill(currency, 100)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for {
			fmt.Println("Текущий баланс: ", bill.GetBalance())

			balance, err := bill.Pay(10)
			if err != nil {
				fmt.Println("❌Ошибка проведения операции: ", err.Error())
				break
			} else {
				fmt.Println("Новый баланс счета: ", balance)
			}
			time.Sleep(1 * time.Second)

			balance, err = bill.TakeCach(5)
			if err != nil {
				fmt.Println("❌Ошибка проведения операции: ", err.Error())
				break
			} else {
				fmt.Println("Новый баланс счета: ", balance)
			}
			time.Sleep(1 * time.Second)

			if rand.Intn(100) < 30 {
				err := errors.New("❌Ошибка проведения операции")
				fmt.Println(err.Error())
				break
			}
			time.Sleep(1 * time.Second)
		}
	}

}
