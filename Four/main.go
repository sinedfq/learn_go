package main

import "fmt"

/*

	Циклы

*/

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println("Number is", i)
	}

	fmt.Println("========")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		if i%2 != 0 {
			continue // Пропуск итерации цикла
		}

		if i == 50 {
			fmt.Println("Stop")
			break
		}
	}

	fmt.Println("========")

	// Цикл While в Go нет, но можно сделать подобие

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("========")

	nums := []int{1, 2, 3, 4, 5} // Срез

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println("========")

	for index, element := range nums {
		fmt.Printf("Index %d Elemnet %d\n", index, element)
	}

	for _, element := range nums {
		fmt.Printf("Elemnet %d\n", element)
	}

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		fmt.Printf("\n")
	}

}
