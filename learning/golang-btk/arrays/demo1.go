package arrays

import "fmt"

func Demo1() {
	var numbers [5]int

	numbers[2] = 7
	fmt.Println(numbers)
	fmt.Println(numbers[2])
}