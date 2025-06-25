package main

import "fmt"

/*

	Конструкция Switch - Case

*/

func main() {

	name := "Anna"

	switch name {
	case "Oleg":
		fmt.Println("Hello" + name)
	case "Denis":
		fmt.Println("Hello " + name)

	default:
		fmt.Println("Your name not in Switch - Case constuction")

	}

	number := 14

	switch {
	case number > 10:
		fmt.Printf("Your number is %d\n", number)
		fallthrough

	case number < 15:
		fmt.Printf("Yes your number is %d\n", number)
	}

}
