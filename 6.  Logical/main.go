package main

import "fmt"

/*

	Логические операторы

*/

func main() {

	a := 1
	b := 4

	if a >= 1 {
		fmt.Println("Super")
	}

	if a >= 1 && b > 3 { // Логическое " AND "
		fmt.Println("A >= 1 and B > 3")
	}

	if a >= 1 || b > 2 { // Логическое " OR "
		fmt.Println("A >= 1 or B > 2")
	}

	if a != 2 { // Логическая " NOT "
		fmt.Print("A != 2")
	}

}
