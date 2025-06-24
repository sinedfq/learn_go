package main

import "fmt"

/*

	Работа с переменными

*/

func main() {
	// Обычный способ объявление переменных
	var age int8 = 23
	var number float32 = 254.213

	/*
		int  - могут быть больше или меньше 0
		uint - могуть быть только больше 0
	*/

	fmt.Println(age)
	fmt.Println(number)
	fmt.Println("========")

	// Популярный способ объявление переменных
	newAge := 16
	newNumber := 421.231
	name := "Denis"
	fmt.Println(newAge)
	fmt.Println(newNumber)
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println("========")

	// Scan
	var userName string
	var userAge uint16
	fmt.Println("What is Your Name?")
	fmt.Scan(&userName)
	fmt.Println("Hello " + userName + "!")
	fmt.Println("What is Your Age?")
	fmt.Scan(&userAge)
	fmt.Println("Your age is - " + fmt.Sprint(userAge))

	var h int8 = 5
	var g int64 = int64(h)
	fmt.Println(g)

	/*
		Нельзя присваивать одну переменную другой если у них разные типы
		Пример 1:
			var h int8 = 5
			var g int64 = h
		Так делать нельзя

		Пример 2:
			var h int8 = 5
			var g int64 = int64(h)
		Так делать можно
	*/

}
