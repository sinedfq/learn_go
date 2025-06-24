package main

import (
	"fmt"
	"math"
)

/*

	Условные операторы

*/

func main() {

	num := 3
	if num > 0 {
		fmt.Println("Number is greater then 0")
	} else if num < 0 {
		fmt.Println("Number is under 0")
	} else if num == 3 {
		fmt.Println("Number equals 3")
	}

	var a float64
	var b float64
	var c float64

	fmt.Println("Введни первое число из квадратного уравнения: ")
	fmt.Scan(&a)
	fmt.Println("Введни второе число из квадратного уравнения: ")
	fmt.Scan(&b)
	fmt.Println("Введни третье число из квадратного уравнения: ")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x1 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Ваше уравнение имеет следующие параметры: \n D = " +
			fmt.Sprint(D) +
			"\nx1 = " + fmt.Sprint(x1) +
			"\nx2 = " + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)

		fmt.Println("Ваше выражение имеет 1 корень: \n D =" + fmt.Sprint(D) + "\n x = " + fmt.Sprint(x))
	} else {
		fmt.Println("Ваше выражение не имеет так как D < 0, D = " + fmt.Sprint(D))
	}

	fmt.Scan(&a)

	/*
		Чтобы оптимизировать размер .exe файла нужно указать флаги:
		-ldflags "-s -w"
		Полная команла будет выглядеть так:
		go build -ldflags "-s -w" main.go
		Данные флаги убирает информацию о дебаге
	*/
}
