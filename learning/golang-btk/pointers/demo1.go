package pointers

import "fmt"

// pointers
func Demo1(number *int) {
	*number = *number + 1
	
	fmt.Println("Pointers method's number variable = ", *number)
}