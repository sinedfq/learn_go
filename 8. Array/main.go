package main

import (
	"fmt"
	"math"
	"sort"
)

/*

	Работа с масивами

*/

func main() {

	// Первый способ объявления массива (не популярен)
	var names [3]string
	names[0] = "Denis"
	names[1] = "Oleg"
	names[2] = "Olga"

	// Второй способ объявления массива
	arNames := [3]string{"Denis", "Oleg", "Nataha"}
	fmt.Println(arNames, len(arNames))

	for i := 0; i < len(arNames); i++ {
		fmt.Println(arNames[i])
	}

	// Среднее арифметическое
	var result float64 = 0
	marks := [5]float64{5, 4, 5, 3, 5}

	for i := 0; i < len(marks); i++ {
		result += marks[i]
	}
	fmt.Println("Result: " + fmt.Sprint(math.Round(result/float64(len(marks)))))

	/*

		Многомерные массивы

	*/

	multArray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(multArray[1][2])

	/*

		Работа со срезами

	*/

	slice := []int{3, 4, 6, 1, 2, 7}
	slice = append(slice, 10)
	slice[0] = 11
	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)

	strSlice := []string{"a", "b", "g", "c", "e", "d"}
	fmt.Println(strSlice)
	sort.Strings(strSlice)
	fmt.Println(strSlice)

	rSlice := []int{1, 2, 3, 4, 5, 6, 7}

	for _, el := range rSlice {
		fmt.Printf("El: %d\n", el)
	}

}
