package pointers

import "fmt"

func Demo2(numbers []int) {
	numbers[0] = 100
	fmt.Println("Number in the Demo", numbers[0])
}