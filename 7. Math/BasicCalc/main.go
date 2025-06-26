package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	var action string

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Println("Ввведите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Введите операцию (+, -, /, *): ")
	fmt.Scan(&action)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел = " + fmt.Sprint(a+b))
	case action == "-":
		fmt.Println("Вычитание чисел = " + fmt.Sprint(a-b))
	case action == "/":
		fmt.Println("Деление чисел = " + fmt.Sprint(a/b))
	case action == "*":
		fmt.Println("Умножение чисел = " + fmt.Sprint(a*b))
	}
}
