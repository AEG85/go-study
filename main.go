package main

import (
	"fmt"
	"study/calculator"
)

func main() {

	calc := calculator.NewCalculator(10.1212, 0)

	fmt.Println(calculator.MakeStringArgs(calc.X, calc.Y))

	res, err := calc.Devide()
	if err == nil {
		fmt.Println("Результат деления: ", res)
	} else {
		fmt.Println("При делении возникла ошибка: ", err.Error())
	}

	res, err = calc.Summ()
	if err == nil {
		fmt.Println("Результат сложения: ", res)
	} else {
		fmt.Println("При сложении возникла ошибка: ", err.Error())
	}

	res, err = calc.Subtract()
	if err == nil {
		fmt.Println("Результат вычитания: ", res)
	} else {
		fmt.Println("При вычитании возникла ошибка: ", err.Error())
	}

	res, err = calc.Multyply()
	if err == nil {
		fmt.Println("Результат умножении: ", res)
	} else {
		fmt.Println("При умножении возникла ошибка: ", err.Error())
	}

}
