package main

import (
	"fmt"
	"math"
)

/*

	Математические операторы

*/

func main() {
	a := 3
	b := 5

	sum := a + b
	fmt.Println(sum)
	sum_minus := a - b
	fmt.Println(sum_minus)
	sum_multiply := a * b
	fmt.Println(sum_multiply)

	var fA float64 = 3
	var fB float64 = 4
	sum_division := fA / fB
	fmt.Println(sum_division)

	iA := 20
	iB := 4
	sum_divisionInt := iA / iB
	fmt.Println(sum_divisionInt)

	modA := 20
	modB := 6
	sum_divisionMod := modA % modB
	fmt.Println(sum_divisionMod)

	var mA float64 = 144
	result := math.Sqrt(mA)
	fmt.Println(result)

	var mB float64 = 25.1241241
	resultCeil := math.Ceil(mB)
	/*
		Ceil округляет к большему значениею, т.е 25.0000021 => 26
		Floar округляет к меньшему значению, т.е 25.6 => 25
		Round округляет нормально 25.6 => 26  || 25.4 => 25
	*/
	fmt.Println(resultCeil)

}
